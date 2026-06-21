# [Mnemosyne](https://en.wikipedia.org/wiki/Mnemosyne)

<p align="center">
  <img src="logo.png" alt="Mnemosyne logo" width="200"/>
</p>

A local document storage and retrieval CLI tool built in Go. Store small
documents (sentences to paragraphs) and retrieve them using hybrid search:
full-text (BM25) + vector similarity (cosine), combined with Reciprocal Rank
Fusion and local cross-encoder reranking.

All ML inference runs locally via ONNX Runtime. No cloud APIs required.

## Features

- **Document storage** in SQLite with metadata support
- **Semantic Markdown Chunking** using `yuin/goldmark` AST to intelligently
  split and preserve heading context when adding `.md` files
- **Full-text search** via SQLite FTS5 with BM25 ranking
- **Vector search** via sqlite-vec with cosine similarity
- **Hybrid search** combining both via Reciprocal Rank Fusion (RRF)
- **Local reranking** with a cross-encoder model (ONNX Runtime) _(coming soon)_
- **Automatic setup** -- downloads ONNX Runtime and ML models on first use (~500
  MB one-time)
- **Configurable models** -- swap embedding or reranker models via config
- **No cloud dependencies** -- everything runs on your machine

## Prerequisites

For the one-command install below, you only need `sh`, `curl` or `wget`, and
`tar`.

For building from source:

- **Go 1.21+** -- [Install Go](https://go.dev/dl/)
- **GCC** -- required by `mattn/go-sqlite3` (CGO). On macOS:
  `xcode-select --install`
- **Task** -- task runner. Install: `brew install go-task` or see
  [taskfile.dev](https://taskfile.dev/installation/)

## Quick Start

Install the latest release to `~/.local/bin`:

```bash
curl -fsSL https://raw.githubusercontent.com/gandazgul/mnemosyne/main/install.sh | sh
```

Then run:

```bash
mnemosyne --help
mnemosyne setup
```

To build from source instead:

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
# Note: If a collection with this name already exists elsewhere,
# init will error out to prevent accidental linking.
./mnemosyne init

# Add documents (triggers model download on first use if not already set up)
./mnemosyne add "Go is a statically typed programming language"
./mnemosyne add "Rust focuses on memory safety and zero-cost abstractions"
./mnemosyne add --file notes.txt
./mnemosyne add --file README.md # Automatically chunks by semantic headings

# Search documents (hybrid: FTS5 + vector, fused with RRF)
./mnemosyne search "programming language"
./mnemosyne search --limit 5 "systems programming"
./mnemosyne search -f json --limit 10 "programming language"

# List documents
./mnemosyne list

# List documents without colors
./mnemosyne list -f plain

# Delete a document by ID
./mnemosyne delete 1

# Use a named collection (with --name or -n)
./mnemosyne init -n myproject
./mnemosyne add -n myproject "some text"
./mnemosyne search -n myproject "some query"

# Use the global collection shortcut
./mnemosyne add -g "Global memory"
./mnemosyne search --global "global memory"

# Delete an entire collection
./mnemosyne forget myproject

# Export a collection to JSONL (includes vectors for fast import)
./mnemosyne export --name myproject

# Export without vectors (smaller file; embeddings auto-generated on import)
./mnemosyne export --name myproject --no-embeddings

# Export all collections
./mnemosyne export --all

# Import a collection (auto-embeds if vectors are missing)
./mnemosyne import myproject.jsonl
./mnemosyne import myproject.jsonl --name other   # override collection name
./mnemosyne import --dir ./backups/               # import all .jsonl files

# Delete an entire collection
./mnemosyne forget myproject
```

## Search Output Formats

The `search` command supports three output formats:

```bash
./mnemosyne search -f color "query" # default human-readable color output
./mnemosyne search -f plain "query" # stable plain text output
./mnemosyne search -f json "query"  # machine-readable output for scripts and benchmarks
```

JSON search output is an object with `query`, `collection`, `count`, and a
ranked `results` array. Each result includes:

- `rank`, `document_id`, `collection_id`, `content`, `created_at`
- parsed `metadata` when document metadata contains valid JSON
- `metadata_raw` when metadata is present but not valid JSON
- `rrf_score`, `fts_rank`, `vec_distance`, `reranker_score`, `is_reranked`
- `sources`, such as `["fts", "vector"]`

This format is intended for evaluation harnesses that need to map retrieved
documents back to external corpus IDs stored in metadata.

## Configuration

Mnemosyne reads `~/.config/mnemosyne/config.yaml` when it exists. You can also
point a single command or benchmark run at another config file:

```bash
MNEMOSYNE_CONFIG=./configs/jina.yaml ./mnemosyne search "query"
MNEMOSYNE_DB_PATH=/tmp/mnemosyne.db ./mnemosyne add "temporary note"
```

Use [config.example.yaml](config.example.yaml) as a starting point. The most
important model fields are:

- `embedding.model_path`: directory containing `tokenizer.json`, `config.json`,
  and the configured ONNX file
- `embedding.onnx_file`: model path relative to `embedding.model_path`, usually
  `onnx/model.onnx`
- `embedding.dimensions`: vector size stored in SQLite; use a fresh DB or
  re-import documents when this changes
- `embedding.pooling`: `none` for a pooled output, `mean` for mean pooling,
  `cls` for first-token pooling, or `last` for last-token pooling
- `embedding.query_prefix` and `embedding.document_prefix`: model-specific text
  prefixes for query/document embeddings
- `reranker.model_path`: directory containing an ONNX cross-encoder reranker
- `reranker.enabled`: set `false` to disable reranking globally for that config

`mnemosyne setup` still auto-downloads the built-in default models. For custom
model paths, place the model files yourself and run with `MNEMOSYNE_CONFIG` or
copy the config to the default location.

The `configs/` directory includes benchmark-oriented profiles for larger
retrieval models, such as Jina v5 text retrieval. These profiles use a separate
benchmark DB when passed through the BEIR harness with `--config` and
`--run-label`.

## Benchmarks

Mnemosyne includes a small BEIR-compatible retrieval benchmark harness for
SciFact. It downloads the public BEIR SciFact dataset, imports the corpus into
an isolated benchmark database, runs `mnemosyne search -f json` for the judged
test queries, and writes JSON plus Markdown result summaries.

```bash
# Short mechanics check over 10 judged queries and a 500-document corpus subset
task bench:scifact-smoke

# Standard full SciFact test split over the complete corpus
task bench:scifact

# Run the full benchmark in a detached screen session
task bench:scifact-bg
```

The harness reports `nDCG@10`, `MRR@10`, `Recall@10`, `Recall@100`, and
`MAP@100`. Result JSON and Markdown also include breakdowns for first relevant
rank buckets, queries missing at 100, and the lowest `MRR@10` cases.
Downloaded data and scratch databases are written under
`benchmarks/data/` and `benchmarks/work/`, which are gitignored. Published
result files are written to `benchmarks/results/`.

Useful options can be passed after `--`:

```bash
task bench:scifact -- --no-rerank
task bench:scifact -- --reuse-db
task bench:scifact -- --config ./configs/jina.yaml --run-label jina-v2
task bench:scifact-smoke -- --max-queries 25 --max-docs 1000
task bench:scifact-bg -- --reuse-db
```

When `--config` is provided, the harness passes it to Mnemosyne as
`MNEMOSYNE_CONFIG`. When `--run-label` is provided (or inferred from the config
filename), the benchmark uses a separate work DB under `benchmarks/work/` and
includes the label in result filenames. This keeps runs with different
embedding dimensions from sharing a SQLite vector table by accident.

The background task uses `screen` so long benchmark runs survive the shell that
started them:

```bash
screen -ls
tail -f benchmarks/work/beir/scifact/full-benchmark.log
screen -r mnemosyne-scifact-benchmark
screen -S mnemosyne-scifact-benchmark -X quit
```

## Available Tasks

```bash
task build            # Build the binary
task test             # Run all tests
task clean            # Remove build artifacts
task lint             # Run linter (requires golangci-lint)
task download-models  # Download ONNX models from HuggingFace (dev workflow)
task bench:scifact-smoke # Run a small BEIR SciFact smoke benchmark
task bench:scifact    # Run the full BEIR SciFact benchmark
task bench:scifact-bg # Run the full benchmark in a detached screen session
task release -- v0.1.0 # Create and push a new release tag
```

## Export & Import

Mnemosyne supports JSONL-based export and import for backup and transfer.

**Export** writes one JSONL file per collection. Each document includes:

- `content` and `metadata` — the original document data
- `vector` — the raw embedding (omitted with `--no-embeddings`)
- `original_document_id` — the source database ID for provenance/inspection
  (useful for memory cleanup workflows where agents summarize and prune old
  memories; ignored on import, which always assigns new IDs)

**Import** reads a JSONL file and inserts documents into the database:

- If vectors are present, import is fast and model-independent (no
  re-embedding).
- If vectors are missing (from a `--no-embeddings` export), the embedder is
  lazily initialized and vectors are auto-generated. This requires the embedding
  model to be available (auto-downloaded on first use).
- If vectors are missing and the embedder can't be initialized, a clear error is
  returned.

```bash
# Full export (includes vectors)
./mnemosyne export --name myproject

# Lightweight export (no vectors, ~10x smaller)
./mnemosyne export --name myproject --no-embeddings

# Import (auto-embeds if vectors are missing)
./mnemosyne import myproject.jsonl
```

## Creating a Release

To create a new versioned release with cross-compiled binaries and an automated
changelog:

1. Use the `release` task and pass the new semantic version tag:
   ```bash
   task release -- v0.1.0
   ```
2. The `Release` GitHub Action workflow will automatically trigger.
3. It uses GoReleaser and `zig cc` to build `linux/amd64`, `linux/arm64`,
   `darwin/amd64`, and `darwin/arm64` binaries.
4. The workflow will publish a new GitHub Release with the attached binaries and
   an auto-generated changelog based on conventional commits.

## Setup

Mnemosyne requires ONNX Runtime and two ML models to generate embeddings:

| Component                                 | Size    | Source          |
| ----------------------------------------- | ------- | --------------- |
| ONNX Runtime                              | ~38 MB  | GitHub Releases |
| snowflake-arctic-embed-m-v1.5 (embedding) | ~420 MB | HuggingFace     |
| ms-marco-MiniLM-L-6-v2 (reranker)         | ~80 MB  | HuggingFace     |

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
│   ├── export.go             # Export collections to JSONL
│   ├── import.go             # Import collections from JSONL (auto-embeds if needed)
│   ├── helpers.go            # Shared helpers (resolve collection, open DB/embedder)
│   └── format.go             # Output format validation + color helpers
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
│   ├── backup/               # JSONL export/import of collections
│   │   ├── types.go           # Header and DocRecord types
│   │   ├── export.go          # ExportCollection (streams docs to JSONL)
│   │   └── import.go          # ImportCollection (reads JSONL, auto-embeds if needed)
│   ├── reranker/             # ONNX cross-encoder reranker (Phase 7)
│   └── search/               # Hybrid search + RRF
│       ├── hybrid.go         # Search engine (orchestrates FTS + vector + RRF)
│       └── rrf.go            # Reciprocal Rank Fusion algorithm
├── models/                   # ONNX model files (gitignored)
├── lib/                      # Native libraries (gitignored)
├── main.go                   # Entry point
├── Taskfile.yml              # Build/test/run tasks
├── ROADMAP.md                # Future features and out-of-scope ideas
└── go.mod
```

## Technology Stack

| Component        | Library                                                                 |
| ---------------- | ----------------------------------------------------------------------- |
| CLI              | [Cobra](https://github.com/spf13/cobra)                                 |
| SQLite driver    | [go-sqlite3](https://github.com/mattn/go-sqlite3) (CGO)                 |
| Vector search    | [sqlite-vec](https://github.com/asg017/sqlite-vec)                      |
| Full-text search | SQLite FTS5 (built-in)                                                  |
| ML inference     | [ONNX Runtime](https://github.com/yalue/onnxruntime_go)                 |
| Tokenizer        | [sugarme/tokenizer](https://github.com/sugarme/tokenizer) (HuggingFace) |
| Embedding model  | snowflake-arctic-embed-m-v1.5 (256-dim, Apache 2.0)                     |
| Reranker model   | ms-marco-MiniLM-L-6-v2 (cross-encoder)                                  |
| Task runner      | [Task](https://taskfile.dev/)                                           |

## Integrations

### Pi & OpenCode extensions

Mnemosyne also integrates with [Pi](https://pi.dev) via the
[pi-mnemosyne](https://github.com/gandazgul/pi-mnemosyne) extension. And with
[OpenCode](https://opencode.ai/) via the
[opencode-mnemosyne](https://github.com/gandazgul/opencode-mnemosyne) plugin.

These extensions provide `memory_recall`, `memory_store`, `memory_delete` (plus
global variants) and automatically inject core memories into the system prompt.
They also re-inject core memories and tool descriptions during compactions.

See the particular extension's README files for installation and usage details.

### Generic

For other agents or IDEs that support AGENTS.md or similar add this:

```markdown
## Memory (mnemosyne)

mnemosyne is a cli memory storage and semantic retrieval tool. the memories are
kept in a project namespace named after the root folder you can access other
namespaces by using -n [namespace name].

- At the start of a session, use `mnemosyne list -t core -f plain` to get all
  core memories, these will be user preferences and project specific things you
  should know.

Then also use `mnemosyne search -f plain [query relevant to the user's prompt]`
and `mnemosyne search -g -f plain [query relevant to the user's prompt]` to
search relevant memories. `-g` searches on a global namespace for broad user
preferences.

- After significant decisions, use `mnemosyne add "new memory"` to save a
  concise fact you want to remember. Also do this if the user explicitly asks
  you to remember something.
- Delete contradicted memories with `mnemosyne delete [memory id]` before
  storing updated ones. The memory id is shown in the output in brackets e.g.
  `[123]`.
- Mark critical, always-relevant context as core (-t core) — but use sparingly.
  You can also use other tags as you see fit:
  `mnemosyne add "database is sqlite" -t tech-stack -t database`
- When you are done with a session, store any memories that you think are
  relevant to the user and the project. This will help you recall important
  information in future sessions.
```

### Sleeping

Use this prompt template to ask the model to optimize memories:

```markdown
---
description: Optimize long-term memory quality
---

You are running sleep mode to optimize long-term memory quality.

Goal:

- Improve memory signal quality for future sessions.
- Preserve high-value, durable context.
- Reduce noise, redundancy, and stale information.

Process:

1. Use \`mnemosyne export --no-embeddings\` to export all memories and core
   memories to a file ([project name].jsonl in the root directory).
2. Analyze the memories for relevance, redundancy, and importance. Optimize the
   memories by deleting irrelevant or redundant ones, and consolidating
   important but similar memories. Focus on keeping the most relevant and
   important information while minimizing noise and redundancy in the memory
   system.
3. Move memories from the core memories (tags: ['core']) to regular or vice
   versa as needed. Core memories should be reserved for the most critical and
   frequently accessed information, while regular memories can be used for less
   critical or less frequently accessed information.

Delete with \`mnemosyne delete [memory id]\` and add with \`mnemosyne add
"memory content" -t tag1 -t tag2\`.
```

## Acknowledgements

Built with the help of [OpenCode](https://opencode.ai/) & [Pi](https://pi.dev)
using Claude Opus and Gemini.

## License

MIT
