# [Mnemosyne](https://en.wikipedia.org/wiki/Mnemosyne)

A local document storage and retrieval CLI tool built in Go. Store small
documents (sentences to paragraphs) and retrieve them using hybrid search:
full-text (BM25) + vector similarity (cosine), combined with Reciprocal Rank
Fusion and local cross-encoder reranking.

All ML inference runs locally via ONNX Runtime. No cloud APIs required.

## Features

- **Document storage** in SQLite with metadata support
- **Full-text search** via SQLite FTS5 with BM25 ranking
- **Vector search** via sqlite-vec with cosine similarity
- **Hybrid search** combining both via Reciprocal Rank Fusion (RRF)
- **Local reranking** with a cross-encoder model (ONNX Runtime) *(coming soon)*
- **Automatic setup** -- downloads ONNX Runtime and ML models on first use (~500 MB one-time)
- **Configurable models** -- swap embedding or reranker models via config
- **No cloud dependencies** -- everything runs on your machine

## Prerequisites

- **Go 1.21+** -- [Install Go](https://go.dev/dl/)
- **GCC** -- required by `mattn/go-sqlite3` (CGO). On macOS: `xcode-select --install`
- **Task** -- task runner. Install: `brew install go-task` or see [taskfile.dev](https://taskfile.dev/installation/)

## Quick Start

```bash
# Clone the repo
git clone https://github.com/gandazgul/mnemosyne.git
cd mnemosyne

# Build
task build

# Run
./mnemosyne

# See available commands
./mnemosyne --help

# Check version
./mnemosyne version

# Download ONNX Runtime and ML models (~500 MB one-time)
# This also happens automatically on first 'add' or 'search'.
./mnemosyne setup

# Initialize a collection (uses current directory name by default)
./mnemosyne init

# Add documents (triggers model download on first use if not already set up)
./mnemosyne add "Go is a statically typed programming language"
./mnemosyne add "Rust focuses on memory safety and zero-cost abstractions"
./mnemosyne add --file notes.txt

# Search documents (hybrid: FTS5 + vector, fused with RRF)
./mnemosyne search "programming language"
./mnemosyne search --limit 5 "systems programming"

# List documents
./mnemosyne list

# Delete a document by ID
./mnemosyne delete 1

# Use a named collection
./mnemosyne init --name myproject
./mnemosyne add --name myproject "some text"
./mnemosyne search --name myproject "some query"

# Delete an entire collection
./mnemosyne forget myproject
```

## Available Tasks

```bash
task build            # Build the binary
task test             # Run all tests
task clean            # Remove build artifacts
task lint             # Run linter (requires golangci-lint)
task download-models  # Download ONNX models from HuggingFace (dev workflow)
```

## Setup

Mnemosyne requires ONNX Runtime and two ML models to generate embeddings:

| Component | Size | Source |
|-----------|------|--------|
| ONNX Runtime | ~38 MB | GitHub Releases |
| snowflake-arctic-embed-m-v1.5 (embedding) | ~420 MB | HuggingFace |
| ms-marco-MiniLM-L-6-v2 (reranker) | ~80 MB | HuggingFace |

**Automatic**: On first use of `add` or `search`, Mnemosyne detects missing
components and downloads them automatically to `~/.local/share/mnemosyne/`.

**Manual**: Run `mnemosyne setup` to download everything upfront:

```bash
./mnemosyne setup
```

The command is idempotent -- it skips files that are already downloaded. No
HuggingFace account or API token is required (all models are openly licensed).

## Project Structure

```
mnemosyne/
├── cmd/                      # CLI commands (Cobra)
│   ├── root.go               # Root command + welcome message
│   ├── version.go            # version subcommand
│   ├── init.go               # Initialize a collection
│   ├── add.go                # Add a document (embeds + stores vector)
│   ├── list.go               # List documents
│   ├── delete.go             # Delete a document by ID
│   ├── forget.go             # Delete an entire collection
│   ├── search.go             # Search (hybrid: FTS5 + vector + RRF)
│   ├── setup.go              # Download ONNX Runtime + ML models
│   └── helpers.go            # Shared helpers (resolve collection, open DB/embedder)
├── internal/
│   ├── config/
│   │   └── config.go         # Configuration loading + defaults
│   ├── db/
│   │   ├── sqlite.go         # DB init, migrations, connection
│   │   ├── collections.go    # CRUD for collections table
│   │   ├── documents.go      # CRUD for documents table
│   │   ├── fts.go            # FTS5 full-text search queries
│   │   └── vectors.go        # sqlite-vec vector insert/query (KNN)
│   ├── embedding/            # ONNX embedding (tokenizer + embedder)
│   │   ├── embedder.go       # Embedder interface + ONNX implementation
│   │   ├── tokenizer.go      # HuggingFace tokenizer wrapper
│   │   └── *_test.go         # Unit + integration tests
│   ├── setup/                # Auto-download of runtime + models
│   │   ├── platform.go       # Platform detection, URL construction
│   │   ├── download.go       # HTTP download with resume + checksum
│   │   └── setup.go          # Orchestration (Check, Run, EnsureReady)
│   ├── reranker/             # ONNX cross-encoder reranker (Phase 7)
│   └── search/               # Hybrid search + RRF
│       ├── hybrid.go         # Search engine (orchestrates FTS + vector + RRF)
│       └── rrf.go            # Reciprocal Rank Fusion algorithm
├── models/                   # ONNX model files (gitignored)
├── lib/                      # Native libraries (gitignored)
├── main.go                   # Entry point
├── Taskfile.yml              # Build/test/run tasks
├── IMPLEMENTATION_PLAN.md    # Detailed phased build plan
└── go.mod
```

## Implementation Status

- [x] **Phase 1**: Skeleton CLI + project setup
- [x] **Phase 2**: SQLite + document storage (CRUD)
- [x] **Phase 3**: Full-text search (FTS5 + BM25)
- [x] **Phase 4**: Embedding model (ONNX Runtime)
- [x] **Phase 5**: Vector storage + search (sqlite-vec)
- [x] **Phase 6**: Hybrid search + Reciprocal Rank Fusion
- [ ] **Phase 7**: Cross-encoder reranker
- [ ] **Phase 8**: Polish and extras
- [ ] **Phase 9**: GitHub CI/CD + versioned releases

See [IMPLEMENTATION_PLAN.md](IMPLEMENTATION_PLAN.md) for the full plan with architecture
diagrams, database schema, search pipeline details, and Go concepts covered per phase.

## Technology Stack

| Component        | Library                          |
| ---------------- | -------------------------------- |
| CLI              | [Cobra](https://github.com/spf13/cobra) |
| SQLite driver    | [go-sqlite3](https://github.com/mattn/go-sqlite3) (CGO) |
| Vector search    | [sqlite-vec](https://github.com/asg017/sqlite-vec) |
| Full-text search | SQLite FTS5 (built-in)           |
| ML inference     | [ONNX Runtime](https://github.com/yalue/onnxruntime_go) |
| Tokenizer        | [daulet/tokenizers](https://github.com/daulet/tokenizers) (HuggingFace) |
| Embedding model  | snowflake-arctic-embed-m-v1.5 (256-dim, Apache 2.0) |
| Reranker model   | ms-marco-MiniLM-L-6-v2 (cross-encoder) |
| Task runner      | [Task](https://taskfile.dev/)    |

## Acknowledgements

Built with the help of [OpenCode](https://opencode.ai/) using Claude Opus and Gemini.

## License

MIT
