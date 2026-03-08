# Mnemosyne - Implementation Plan

A Go CLI tool for storing and retrieving small documents (sentences to doc chunks) using
**hybrid search**: full-text search (FTS5) + vector similarity search (sqlite-vec), combined
with **Reciprocal Rank Fusion (RRF)**, and refined by a local **cross-encoder reranker** model.
All inference runs locally via ONNX Runtime. Models are configurable.

**Documents belong to collections.** A collection is created via `mnemosyne init` and all
operations (`add`, `search`, `list`) are scoped to a collection. The collection is specified
via `--name <collection>`, or defaults to the base name of the current working directory
(e.g. running from `/home/user/myproject` uses collection `myproject`).

## Architecture

```
┌─────────────────────────────────────────────────────┐
│                   CLI (Cobra)                       │
│  mnemosyne init | add | search | list | delete       │
└─────────────┬───────────────────────────────────────┘
              │
┌─────────────▼───────────────────────────────────────┐
│               Core Application Layer                │
│  ┌──────────┐  ┌──────────┐  ┌───────────────────┐  │
│  │ Ingest   │  │ Search   │  │ Rank/Fuse/Rerank  │  │
│  │ Service  │  │ Service  │  │ Service            │  │
│  └────┬─────┘  └────┬─────┘  └────────┬──────────┘  │
└───────┼──────────────┼─────────────────┼────────────┘
        │              │                 │
┌───────▼──────────────▼─────────────────▼────────────┐
│                  Storage Layer                       │
│  SQLite + FTS5 (full-text) + sqlite-vec (vectors)   │
│  ┌─────────────┐ ┌────────────┐ ┌────────────────┐  │
│  │ collections │ │ documents  │ │ docs_fts       │  │
│  │ (names)     │ │ (content)  │ │ (FTS5 index)   │  │
│  └─────────────┘ └────────────┘ └────────────────┘  │
│                  ┌──────────────────┐                │
│                  │ docs_vec (vec0)  │                │
│                  │ (vector index)   │                │
│                  └──────────────────┘                │
└─────────────────────────────────────────────────────┘
        │                                 │
┌───────▼─────────────────────────────────▼───────────┐
│              ML Inference Layer                      │
│  ONNX Runtime (yalue/onnxruntime_go)                │
│  ┌─────────────────────┐ ┌────────────────────────┐ │
│  │ Embedding Model     │ │ Reranker Model         │ │
│  │ EmbeddingGemma-300M │ │ ms-marco-MiniLM-L-6   │ │
│  │ (768-dim, config.)  │ │ (cross-encoder)        │ │
│  └─────────────────────┘ └────────────────────────┘ │
└─────────────────────────────────────────────────────┘
```

## Technology Stack

| Component          | Library / Tool                                            | Purpose                                   |
| ------------------ | --------------------------------------------------------- | ----------------------------------------- |
| CLI framework      | `spf13/cobra`                                             | Subcommand routing, flags, help           |
| SQLite driver      | `mattn/go-sqlite3` (CGO)                                  | SQLite access with extension support      |
| Vector search      | `asg017/sqlite-vec` + `sqlite-vec-go-bindings/cgo`        | Store embeddings, cosine similarity KNN   |
| Full-text search   | SQLite FTS5 (build tag `sqlite_fts5`)                     | BM25 keyword search                       |
| ONNX inference     | `yalue/onnxruntime_go`                                    | Run embedding + reranker models locally   |
| Tokenizer          | HuggingFace `tokenizer.json` + Go tokenizer lib           | Tokenize text for ONNX models             |
| Rank fusion        | Custom implementation                                     | Reciprocal Rank Fusion (RRF)              |
| Task runner        | [Taskfile](https://taskfile.dev/) (`Taskfile.yml`)        | Build, test, model download               |

## Configuration

Models and settings are configurable via `~/.config/mnemosyne/config.yaml`, environment
variables, or CLI flags. Example:

```yaml
db_path: "~/.local/share/mnemosyne/mnemosyne.db"

embedding:
  model_path: "~/.local/share/mnemosyne/models/embeddinggemma-300m"
  dimensions: 768            # supports MRL truncation: 768, 512, 256, 128
  max_seq_length: 2048
  query_prefix: "task: search result | query: "
  document_prefix: "title: none | text: "

reranker:
  model_path: "~/.local/share/mnemosyne/models/ms-marco-MiniLM-L-6-v2"
  max_seq_length: 512
  enabled: true

search:
  rrf_k: 60                 # RRF constant
  top_k: 10                 # results to return
  rerank_candidates: 50     # candidates to pass to reranker

huggingface:
  token: ""                  # or set HF_TOKEN env var
```

## Database Schema

```sql
-- Collections table
CREATE TABLE IF NOT EXISTS collections (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Main documents table (belongs to a collection)
CREATE TABLE IF NOT EXISTS documents (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    collection_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    metadata TEXT,            -- JSON blob for tags, source, etc.
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (collection_id) REFERENCES collections(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_documents_collection ON documents(collection_id);

-- FTS5 virtual table for full-text search
CREATE VIRTUAL TABLE IF NOT EXISTS docs_fts USING fts5(
    content,
    content=documents,
    content_rowid=id
);

-- Triggers to keep FTS in sync with documents table
CREATE TRIGGER IF NOT EXISTS docs_fts_ai AFTER INSERT ON documents BEGIN
    INSERT INTO docs_fts(rowid, content) VALUES (new.id, new.content);
END;
CREATE TRIGGER IF NOT EXISTS docs_fts_ad AFTER DELETE ON documents BEGIN
    INSERT INTO docs_fts(docs_fts, rowid, content) VALUES('delete', old.id, old.content);
END;
CREATE TRIGGER IF NOT EXISTS docs_fts_au AFTER UPDATE ON documents BEGIN
    INSERT INTO docs_fts(docs_fts, rowid, content) VALUES('delete', old.id, old.content);
    INSERT INTO docs_fts(rowid, content) VALUES (new.id, new.content);
END;

-- sqlite-vec virtual table for vector search (dimension from config)
CREATE VIRTUAL TABLE IF NOT EXISTS docs_vec USING vec0(
    document_id INTEGER PRIMARY KEY,
    embedding float[768] distance_metric=cosine
);
```

**Note**: FTS5 and sqlite-vec tables are global (not per-collection). Queries filter
by collection by joining through the `documents` table on `collection_id`.

## Search Pipeline

```
User Query + Collection Name (--name or cwd)
    |
    +--> Resolve collection_id from name
    |
    +-->  FTS5 Search (BM25) filtered by collection_id  -->  Ranked List A
    |
    +-->  Embed query -->  sqlite-vec KNN filtered by collection_id  -->  Ranked List B
    |
    v
Reciprocal Rank Fusion
    RRF_score(d) = SUM( 1/(k + rank_i(d)) )   [k=60 typically]
    |
    v
Top-N Candidates
    |
    v
Cross-Encoder Reranker (ONNX)
    Score each (query, document) pair
    |
    v
Final Ranked Results
```

## Key Interfaces

Models are swappable via interfaces -- changing models is a config change, not a code change:

```go
// Embedder generates vector embeddings from text.
type Embedder interface {
    Embed(text string) ([]float32, error)
    EmbedBatch(texts []string) ([][]float32, error)
    Dimensions() int
}

// Reranker scores query-document relevance.
type Reranker interface {
    Score(query string, documents []string) ([]float32, error)
}
```

## ONNX Models

### Embedding: EmbeddingGemma-300M (default, configurable)

- HuggingFace: `google/embeddinggemma-300m`
- ONNX: `onnx-community/embeddinggemma-300m-ONNX`
- Dimensions: 768 (MRL truncatable to 512, 256, 128)
- Max sequence length: 2048 tokens
- Requires HuggingFace token for download (gated model)
- Uses task-specific prefixes for queries vs documents

### Reranker: ms-marco-MiniLM-L-6-v2 (default, configurable)

- HuggingFace: `cross-encoder/ms-marco-MiniLM-L-6-v2`
- Input: tokenized (query + document) pair
- Output: relevance score (float32)
- Max sequence length: 512 tokens

## Project Structure

```
mnemosyne/
├── cmd/                      # Cobra commands
│   ├── root.go               # Root command, global flags
│   ├── init.go               # mnemosyne init --name <collection>
│   ├── add.go                # mnemosyne add --name <collection> "text"
│   ├── search.go             # mnemosyne search --name <collection> "query"
│   ├── list.go               # mnemosyne list --name <collection>
│   ├── delete.go             # mnemosyne delete <id>
│   └── version.go            # mnemosyne version
├── internal/
│   ├── config/
│   │   └── config.go         # YAML config loading, model paths, dimensions
│   ├── db/
│   │   ├── sqlite.go         # DB init, migrations, connection
│   │   ├── collections.go    # CRUD for collections table
│   │   ├── documents.go      # CRUD for documents table (collection-scoped)
│   │   ├── fts.go            # FTS5 search queries (collection-scoped)
│   │   └── vectors.go        # sqlite-vec insert/query (collection-scoped)
│   ├── embedding/
│   │   ├── embedder.go       # Embedder interface + ONNX implementation
│   │   └── tokenizer.go      # Text tokenization
│   ├── reranker/
│   │   └── reranker.go       # Reranker interface + ONNX implementation
│   └── search/
│       ├── hybrid.go         # Orchestrates FTS + vector search
│       └── rrf.go            # Reciprocal Rank Fusion
├── models/                   # ONNX model files (gitignored)
├── main.go                   # Entry point
├── go.mod
├── go.sum
├── Taskfile.yml              # Build, test, model download tasks
├── config.example.yaml       # Example config for users
├── .gitignore
└── README.md
```

---

## Phase 1: Skeleton CLI + Project Setup

**Goal**: A working Go CLI that builds, runs, and prints a hello message. Includes the
`init` command for creating collections.

### Tasks

- [ ] Initialize Go module (`github.com/gandazgul/mnemosyne`)
- [ ] Install Cobra dependency
- [ ] Create `main.go` entry point
- [ ] Create `cmd/root.go` with root command (prints welcome message)
- [ ] Create `cmd/version.go` subcommand
- [ ] Create `cmd/init.go` - `mnemosyne init [--name <collection>]`
  - If `--name` is not provided, use the base name of the current working directory
  - Creates a new collection in the database (or confirms it already exists)
  - Prints the collection name on success
- [ ] Create `Taskfile.yml` with `build`, `test`, and `clean` tasks
- [ ] Create `.gitignore`
- [ ] Verify: `task build` succeeds, `./mnemosyne` prints hello, `./mnemosyne version` works
- [ ] Verify: `./mnemosyne init` creates a collection named after the cwd
- [ ] Verify: `./mnemosyne init --name myproject` creates a collection named "myproject"

### Go Concepts Introduced

- Modules and `go.mod`
- Packages and imports
- `main()` function as entry point
- Exported vs unexported identifiers
- Third-party dependencies (Cobra)
- `os.Getwd()` and `filepath.Base()` for deriving default collection name
- Build tags (placeholder for Phase 2)

---

## Phase 2: SQLite + Document Storage

**Goal**: Store and retrieve documents in SQLite, scoped to collections.

### Tasks

- [ ] Add `mattn/go-sqlite3` dependency
- [ ] Create `internal/config/config.go` with DB path and model settings
- [ ] Create `internal/db/sqlite.go` - DB initialization, connection, schema migration
  - Schema includes `collections` and `documents` tables with foreign key
- [ ] Create `internal/db/collections.go` - Create, GetByName, List, Delete collections
- [ ] Create `internal/db/documents.go` - Insert, List, Delete, GetByID (all scoped to collection_id)
- [ ] Implement `cmd/add.go` - `mnemosyne add [--name <collection>] "some text"` and `--file path`
  - Resolves collection by `--name` flag or cwd base name
  - Errors if the collection does not exist (must `init` first)
- [ ] Implement `cmd/list.go` - `mnemosyne list [--name <collection>]` with optional `--limit`
  - Resolves collection by `--name` flag or cwd base name
- [ ] Implement `cmd/delete.go` - `mnemosyne delete <id>`
- [ ] Update `Taskfile.yml` build to include `-tags "sqlite_fts5"` (prep for Phase 3)
- [ ] Write tests for collections CRUD and document CRUD operations
- [ ] Verify: init, add, list, delete all work end-to-end within a collection

### Go Concepts Introduced

- `database/sql` interface
- CGO basics (needed for go-sqlite3)
- Error handling patterns (`if err != nil`)
- Structs and methods
- Testing with `go test`
- Build tags

---

## Phase 3: Full-Text Search (FTS5)

**Goal**: Search documents by keyword using SQLite FTS5 with BM25 ranking, scoped to a collection.

### Tasks

- [ ] Create FTS5 virtual table + sync triggers in schema migration
- [ ] Create `internal/db/fts.go` - FTS5 search query returning ranked results
  - Join with `documents` table to filter by `collection_id`
- [ ] Implement `cmd/search.go` - `mnemosyne search [--name <collection>] "query"` (FTS5 only for now)
  - Resolves collection by `--name` flag or cwd base name
  - Errors if the collection does not exist
- [ ] Display results with BM25 scores and document snippets
- [ ] Handle edge cases: empty results, special characters in queries
- [ ] Write tests for FTS5 search (including collection scoping)
- [ ] Verify: search returns relevant results ranked by BM25 within the correct collection

### Go Concepts Introduced

- Virtual tables in SQLite
- SQL query building
- Formatting output (tables/columns in terminal)
- Slice operations and sorting

---

## Phase 4: Embedding Model (ONNX)

**Goal**: Generate vector embeddings from text using a local ONNX model.

### Tasks

- [ ] Add `yalue/onnxruntime_go` dependency
- [ ] Add tokenizer library dependency (e.g. `daulet/tokenizers`)
- [ ] Define `Embedder` interface in `internal/embedding/embedder.go`
- [ ] Implement ONNX-based embedder for EmbeddingGemma-300M
- [ ] Implement tokenizer wrapper in `internal/embedding/tokenizer.go`
- [ ] Handle query vs document prefixes from config
- [ ] Add `download-models` task to `Taskfile.yml` (uses `hf` CLI + HF_TOKEN)
- [ ] Add `download-onnxruntime` task for the shared library
- [ ] Write tests for embedding generation (deterministic output check)
- [ ] Verify: generate embeddings, check dimensions match config

### Go Concepts Introduced

- CGO / FFI (ONNX Runtime native library)
- Interfaces and implementations
- `[]float32` slices, binary serialization
- Environment variables (`HF_TOKEN`)
- Config-driven behavior

---

## Phase 5: Vector Storage + Search (sqlite-vec)

**Goal**: Store document embeddings and perform KNN vector search.

### Tasks

- [ ] Add `asg017/sqlite-vec-go-bindings/cgo` dependency
- [ ] Create `internal/db/vectors.go` - insert/query vectors via sqlite-vec
- [ ] Create `docs_vec` virtual table with dimension from config + cosine distance
- [ ] On `mnemosyne add`: embed document content and store vector alongside document
- [ ] On `mnemosyne search`: embed query, run KNN search via `MATCH` operator, filter by collection
- [ ] Display vector search results with cosine distances
- [ ] Handle re-embedding when model/dimensions change (migration strategy)
- [ ] Write tests for vector insert and KNN query
- [ ] Verify: vector search returns semantically similar results

### Go Concepts Introduced

- Binary data in SQLite (BLOB serialization with `SerializeFloat32`)
- Type assertions
- Multiple return values
- Parallel data paths (document + embedding stored together)

---

## Phase 6: Hybrid Search + Reciprocal Rank Fusion

**Goal**: Combine FTS5 and vector search results using RRF.

### Tasks

- [ ] Implement RRF algorithm in `internal/search/rrf.go`
- [ ] Create `internal/search/hybrid.go` - orchestrates both searches, fuses results (collection-scoped)
- [ ] Update `cmd/search.go` to use hybrid search by default
- [ ] Add flags: `--mode fts|vector|hybrid` to choose search mode
- [ ] Add `--rrf-k` flag (default from config)
- [ ] Display combined scores and which sources contributed
- [ ] Write tests for RRF (known inputs -> expected ranking)
- [ ] Verify: hybrid search outperforms either method alone on test data

### RRF Algorithm

```
RRF_score(d) = SUM over all rankings R of: 1 / (k + rank_R(d))

Where:
  - k = 60 (constant, configurable)
  - rank_R(d) = position of document d in ranking R (1-indexed)
  - Documents not present in a ranking are ignored for that ranking
```

### Go Concepts Introduced

- Algorithm implementation
- Maps (`map[int]float64` for score accumulation)
- Sorting by custom criteria (`sort.Slice`)
- Enums/constants for search modes
- Flag handling with Cobra

---

## Phase 7: Reranker (ONNX Cross-Encoder)

**Goal**: Re-score top candidates using a cross-encoder model for final ranking.

### Tasks

- [ ] Define `Reranker` interface in `internal/reranker/reranker.go`
- [ ] Implement ONNX-based cross-encoder reranker
- [ ] Tokenize (query, document) pairs for cross-encoder input
- [ ] After RRF, pass top-N candidates through reranker
- [ ] Display final reranked results with reranker scores
- [ ] Add `--no-rerank` flag to skip reranking
- [ ] Add `--rerank-candidates` flag (default from config)
- [ ] Update `download-models` task to include reranker model
- [ ] Write tests for reranker scoring
- [ ] Verify: reranking improves result quality on test queries

### Go Concepts Introduced

- Struct composition
- More ONNX Runtime usage (different model architecture)
- Goroutines (optional: parallel scoring)
- Benchmarking with `go test -bench`

---

## Phase 8: Polish and Extras

**Goal**: Production-ready CLI with good UX and documentation.

### Tasks

- [ ] Create `config.example.yaml` with all options documented
- [ ] Support `--config` global flag to specify config file path
- [ ] `mnemosyne collections` command - list all collections with document counts
- [ ] Metadata support: `mnemosyne add --tag "golang" --source "book.pdf" "text"`
- [ ] Batch import: `mnemosyne add --file document.txt` (splits into chunks)
- [ ] Batch import from stdin: `cat doc.txt | mnemosyne add --stdin`
- [ ] Pretty terminal output (table formatting, colors)
- [ ] `mnemosyne stats` command (document count, DB size, model info)
- [ ] Improve error messages and help text
- [ ] Add shell completion support (Cobra built-in)
- [ ] Performance profiling and optimization
- [ ] Write integration tests (full pipeline: add -> search -> verify)
- [ ] Update README with usage examples and screenshots

### Go Concepts Introduced

- File I/O and streaming (`os.Stdin`, `bufio.Scanner`)
- JSON marshal/unmarshal for metadata
- `text/tabwriter` for formatted output
- Context and cancellation patterns
- Integration testing patterns

---

## Go Concepts Learned Per Phase (Summary)

| Phase | Key Go Concepts                                                    |
| ----- | ------------------------------------------------------------------ |
| 1     | Modules, packages, imports, `main()`, basic types, Cobra          |
| 2     | `database/sql`, CGO, error handling, structs, methods, testing     |
| 3     | SQL queries, formatting output, slices, sorting                    |
| 4     | CGO/FFI (ONNX), interfaces, float32 arrays, env vars, config      |
| 5     | Binary data, type assertions, serialization                        |
| 6     | Algorithms, maps, custom sorting, enums, flags                     |
| 7     | Struct composition, goroutines (optional), benchmarking            |
| 8     | File I/O, JSON, streaming, context, integration testing            |
