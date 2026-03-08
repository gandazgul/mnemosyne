# Mnemosyne

A local document storage and retrieval CLI tool built in Go. Store small
documents (sentences to paragraphs) and retrieve them using hybrid search:
full-text (BM25) + vector similarity (cosine), combined with Reciprocal Rank
Fusion and local cross-encoder reranking.

All ML inference runs locally via ONNX Runtime. No cloud APIs required.

## Features (Planned)

- **Document storage** in SQLite with metadata support
- **Full-text search** via SQLite FTS5 with BM25 ranking
- **Vector search** via sqlite-vec with cosine similarity
- **Hybrid search** combining both via Reciprocal Rank Fusion (RRF)
- **Local reranking** with a cross-encoder model (ONNX Runtime)
- **Configurable models** -- swap embedding or reranker models via config
- **No cloud dependencies** -- everything runs on your machine

## Prerequisites

- **Go 1.21+** -- [Install Go](https://go.dev/dl/)
- **GCC** -- required by `mattn/go-sqlite3` (CGO). On macOS: `xcode-select --install`
- **Task** -- task runner. Install: `brew install go-task` or see [taskfile.dev](https://taskfile.dev/installation/)
- **HuggingFace CLI** (later phases) -- for model downloads: `pip install huggingface_hub`

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
```

## Available Tasks

```bash
task build            # Build the binary
task test             # Run all tests
task clean            # Remove build artifacts
task lint             # Run linter (requires golangci-lint)
task download-models  # Download ONNX models from HuggingFace
task setup            # Install deps + download models
```

## Project Structure

```
mnemosyne/
├── cmd/                      # CLI commands (Cobra)
│   ├── root.go               # Root command + welcome message
│   └── version.go            # version subcommand
├── internal/                 # Internal packages (added in later phases)
│   ├── config/               # Configuration loading
│   ├── db/                   # SQLite, FTS5, sqlite-vec
│   ├── embedding/            # ONNX embedding model
│   ├── reranker/             # ONNX cross-encoder reranker
│   └── search/               # Hybrid search + RRF
├── models/                   # ONNX model files (gitignored)
├── main.go                   # Entry point
├── Taskfile.yml              # Build/test/run tasks
├── IMPLEMENTATION_PLAN.md    # Detailed phased build plan
└── go.mod
```

## Implementation Status

- [x] **Phase 1**: Skeleton CLI + project setup
- [ ] **Phase 2**: SQLite + document storage (CRUD)
- [ ] **Phase 3**: Full-text search (FTS5 + BM25)
- [ ] **Phase 4**: Embedding model (ONNX Runtime)
- [ ] **Phase 5**: Vector storage + search (sqlite-vec)
- [ ] **Phase 6**: Hybrid search + Reciprocal Rank Fusion
- [ ] **Phase 7**: Cross-encoder reranker
- [ ] **Phase 8**: Polish and extras

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
| Embedding model  | EmbeddingGemma-300M (768-dim, configurable) |
| Reranker model   | ms-marco-MiniLM-L-6-v2 (configurable) |
| Task runner      | [Task](https://taskfile.dev/)    |

## License

MIT
