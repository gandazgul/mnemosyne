# [Mnemosyne](https://en.wikipedia.org/wiki/Mnemosyne)

<p align="center">
  <img src="logo.png" alt="Mnemosyne logo" width="200"/>
</p>

A local document storage and retrieval CLI tool built in Go. Store small
documents (sentences to paragraphs) and retrieve them using hybrid search:
vector similarity (cosine) with a small BM25 lexical boost, plus optional local
cross-encoder reranking.

All ML inference runs locally via ONNX Runtime. No cloud APIs required.

## Features

- **Document storage** in SQLite with metadata support
- **Semantic Markdown Chunking** using `yuin/goldmark` AST to intelligently
  split and preserve heading context when adding `.md` files
- **Full-text search** via SQLite FTS5 with BM25 ranking
- **Vector search** via sqlite-vec with cosine similarity
- **Hybrid search** using vector-first retrieval with BM25 lexical reranking
- **Legacy RRF fusion** for explicit full-text + vector candidate union
- **Local reranking** with a cross-encoder model (ONNX Runtime) _(coming soon)_
- **Automatic setup** -- downloads ONNX Runtime and ML models on first use
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

## Installation

### Install the CLI

Install the latest release to `~/.local/bin`:

```bash
curl -fsSL https://raw.githubusercontent.com/gandazgul/mnemosyne/main/install.sh | sh
```

Make sure `~/.local/bin` is on your `PATH`, then run:

```bash
mnemosyne --help
mnemosyne setup
```

`mnemosyne setup` downloads ONNX Runtime and the local embedding/reranker
models. Setup is idempotent and also runs automatically on first `add` or
`search`. No HuggingFace account or API token is required for the built-in
models.

### Connect Mnemosyne to your agent

Install the CLI once, then add the integration for each assistant you use. The
integrations share the same local Mnemosyne collections, so memories can follow
you across agents.

- [claudecode-mnemosyne](https://github.com/gandazgul/claudecode-mnemosyne) —
  Claude Code skill, including migration commands for existing Claude Code
  memory files.
- [opencode-mnemosyne](https://github.com/gandazgul/opencode-mnemosyne) —
  OpenCode plugin with memory tools and core-memory injection.
- [pi-mnemosyne](https://github.com/gandazgul/pi-mnemosyne) — Pi extension with
  memory tools and core-memory injection.

See each integration README for install commands, agent-specific configuration,
and usage details.

#### Other agents

For agents or IDEs that support `AGENTS.md`, project instructions, or similar
custom guidance, add this instruction block after installing the CLI:

```markdown
## Memory System

- Use `mnemosyne search -f plain [query]` and `mnemosyne search -g -f plain [query]` to search relevant memories. Use this before making any decisions or taking any actions.
- After significant decisions, use `mnemosyne add "memory content"` to save a concise fact you want to remember. Also do this if the user explicitly asks you to remember something. Use `mnemosyne add -g "memory content"` for cross-project preferences.
- Delete contradicted memories with `mnemosyne delete [memory id]` after storing updated ones with `mnemosyne add ...` or `mnemosyne add -g ...`.
- Mark critical, always-relevant context as core with `-t core`, but use it sparingly. You can also use other tags with repeated `-t` flags, such as `mnemosyne add "database is sqlite" -t core -t tech-stack`.
- When you are done with a session, store any memories that you think are relevant to the user and the project. This will help you recall important information in future sessions.
```

#### Sleep

For conservative memory maintenance, use the [Sleep prompt](docs/sleep-prompt.md).
It is tuned for preserving durable context while removing only exact duplicates,
clearly deprecated facts, explicitly superseded memories, or lossless
consolidations.

## Benchmarks

**Headline:** Mnemosyne retrieves the correct memory session
in the top 5 for 98.8% of questions and in the top 10 for 99.8% of questions on
the full 500-question LongMemEval cleaned set, using short per-message
documents, local Jina v5 nano embeddings, vector-BM25 fusion, and no LLM
reranker.

See [benchmarks/README.md](benchmarks/README.md) for published results,
comparison notes, and full reproduction commands.

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

Mnemosyne also includes a LongMemEval harness for memory-system retrieval
comparisons. LongMemEval is per-question: each question has its own haystack of
conversation sessions, so the harness builds an isolated Mnemosyne DB per
question and scores whether the retrieved documents map back to the answer
session IDs.

Published no-rerank LongMemEval results:

| Document Mode | Recall@5 | Recall@10 | MRR@10 | nDCG@10 |
| ------------- | -------: | --------: | -----: | ------: |
| `message`     |   0.9880 |    0.9980 | 0.9586 |  0.9575 |
| `session`     |   0.9920 |    0.9940 | 0.9491 |  0.9486 |

```bash
# Tiny mechanics check: 2 questions, trimmed haystacks
task bench:longmemeval-smoke -- --config ./configs/jina-v5-text-nano-retrieval.yaml --run-label jina-v5-nano

# Session-doc mode: one document per session, user turns joined.
# This is the closest comparison to MemPalace's raw LongMemEval setup.
python3 benchmarks/longmemeval/run.py \
  --config ./configs/jina-v5-text-nano-retrieval.yaml \
  --run-label jina-v5-nano \
  --doc-mode session \
  --no-rerank

# Short-doc mode: one document per user/assistant message, scored by session hit.
python3 benchmarks/longmemeval/run.py \
  --config ./configs/jina-v5-text-nano-retrieval.yaml \
  --run-label jina-v5-nano \
  --doc-mode message \
  --no-rerank
```

The BEIR harness reports `nDCG@10`, `MRR@10`, `Recall@10`, `Recall@100`, and
`MAP@100`. Result JSON and Markdown also include breakdowns for first relevant
rank buckets, queries missing at 100, and the lowest `MRR@10` cases. The
LongMemEval harness reports session-level `recall_any@5`, `recall_any@10`,
`recall_all@5`, `recall_all@10`, `MRR@10`, and `nDCG@10`, plus per-question-type
breakdowns. Downloaded data and scratch databases are written under
`benchmarks/data/` and `benchmarks/work/`, which are gitignored. Published
result files are written to `benchmarks/results/`.

Useful options can be passed after `--`:

```bash
task bench:scifact -- --no-rerank
task bench:scifact -- --reuse-db
task bench:scifact -- --config ./configs/jina-v5-text-nano-retrieval.yaml --run-label jina-v5-nano
task bench:scifact -- --fts-only --no-rerank
task bench:scifact -- --vector-only --no-rerank
task bench:scifact -- --fusion rrf --no-rerank
task bench:scifact -- --fusion vector-bm25 --bm25-weight 0.10 --rerank-candidates 300 --no-rerank
task bench:scifact-smoke -- --max-queries 25 --max-docs 1000
task bench:scifact-bg -- --reuse-db
```

The default search fusion is `vector-bm25`, which retrieves vector candidates
first, computes an in-memory BM25 score over that candidate set, and blends the
scores with `--bm25-weight` (default `0.10`). Use `--fusion rrf` for the legacy
global FTS + vector Reciprocal Rank Fusion mode. Use a larger
`--rerank-candidates` value with vector-BM25 if you want lexical reranking to
improve `Recall@100`; with only 100 candidates and a 100-result limit it can
only change ordering.

When `--config` is provided, the harness passes it to Mnemosyne as
`MNEMOSYNE_CONFIG`. When `--run-label` is provided (or inferred from the config
filename), the benchmark uses a separate work DB under `benchmarks/work/` and
includes the label in result filenames. This keeps runs with different embedding
dimensions from sharing a SQLite vector table by accident. Use `--reuse-db` for
`--fts-only` comparisons when possible, since a fresh benchmark import still
embeds vectorless corpus files through the normal import path.

The background task uses `screen` so long benchmark runs survive the shell that
started them:

```bash
screen -ls
tail -f benchmarks/work/beir/scifact/full-benchmark.log
screen -r mnemosyne-scifact-benchmark
screen -S mnemosyne-scifact-benchmark -X quit
```

## Quick Start

```bash
# Run
mnemosyne

# See available commands
mnemosyne --help

# Check version
mnemosyne version

# Download ONNX Runtime and ML models
# This also happens automatically on first 'add' or 'search'.
mnemosyne setup

# Initialize a collection (uses current directory name by default)
# This is idempotent: if the collection already exists, init confirms it.
mnemosyne init

# Add documents (requires an initialized collection and triggers model download
# on first use if setup has not already run)
mnemosyne add "Go is a statically typed programming language"
mnemosyne add "Rust focuses on memory safety and zero-cost abstractions"
mnemosyne add --file notes.txt
mnemosyne add --file README.md # Automatically chunks by semantic headings
printf 'Piped note' | mnemosyne add --stdin
mnemosyne add "Important project preference" -t core -t preference

# Search documents (hybrid: vector-first BM25 fusion by default)
mnemosyne search "programming language"
mnemosyne search --limit 5 "systems programming"
mnemosyne search -f json --limit 10 "programming language"
mnemosyne search --fts-only --no-rerank "programming language"
mnemosyne search --vector-only --no-rerank "programming language"
mnemosyne search --fusion rrf --no-rerank "programming language"

# List documents
mnemosyne list

# List documents without colors, limit results, or filter by tags
mnemosyne list -f plain
mnemosyne list --limit 5
mnemosyne list -t core

# List tags used in a collection
mnemosyne tags

# Update an existing document by ID (strict: errors if the ID is missing or in another collection)
mnemosyne update 1 "Revised memory content"
mnemosyne update 1 --file revised-note.txt
printf 'Revised piped note' | mnemosyne update 1 --stdin
mnemosyne update 1 -t reviewed          # Add a tag, preserving existing tags
mnemosyne update 1 --replace-tags -t core # Replace all tags with the supplied tags
mnemosyne update 1 --replace-tags       # Clear all tags

# Delete a document by ID
mnemosyne delete 1

# Use a named collection (with --name or -n)
mnemosyne init -n myproject
mnemosyne add -n myproject "some text"
mnemosyne update -n myproject 1 "revised text"
mnemosyne search -n myproject "some query"

# Use the global collection shortcut (created automatically on first use)
mnemosyne add -g "Global memory"
mnemosyne list -g
mnemosyne update -g 1 "Revised global memory"
mnemosyne search --global "global memory"

# Show collections and database stats
mnemosyne collections
mnemosyne stats

# Delete an entire collection
mnemosyne forget -n myproject
mnemosyne forget -n myproject --yes # skip confirmation

# Export a collection to JSONL (includes vectors for fast import)
mnemosyne export --name myproject

# Export without vectors (smaller file; embeddings auto-generated on import)
mnemosyne export --name myproject --no-embeddings

# Export all collections
mnemosyne export --all

# Import a collection (auto-embeds if vectors are missing)
mnemosyne import myproject.jsonl
mnemosyne import myproject.jsonl --name other   # override collection name
mnemosyne import --dir ./backups/               # import all .jsonl files

# Import Claude Code memory without modifying Claude files
mnemosyne import --agent claude --dry-run
mnemosyne import --agent claude --include-user
```

## Search Output Formats

The `search` command supports three output formats:

```bash
mnemosyne search -f color "query" # default human-readable color output
mnemosyne search -f plain "query" # stable plain text output
mnemosyne search -f json "query"  # machine-readable output for scripts and benchmarks
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
MNEMOSYNE_CONFIG=./configs/jina-v5-text-nano-retrieval.yaml mnemosyne search "query"
MNEMOSYNE_DB_PATH=/tmp/mnemosyne.db mnemosyne add "temporary note"
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
mnemosyne export --name myproject

# Lightweight export (no vectors, ~10x smaller)
mnemosyne export --name myproject --no-embeddings

# Import (auto-embeds if vectors are missing)
mnemosyne import myproject.jsonl
```

## Developing & Contributing

Build and contribution workflows live here. User-facing install and integration
steps are covered above; package-specific agent setup lives in the integration
READMEs.

### Available Tasks

```bash
task build            # Build the binary
task test             # Run all tests
task clean            # Remove build artifacts
task lint             # Run linter (requires golangci-lint)
task download-models  # Download ONNX models from HuggingFace (dev workflow)
task bench:scifact-smoke # Run a small BEIR SciFact smoke benchmark
task bench:scifact    # Run the full BEIR SciFact benchmark
task bench:scifact-bg # Run the full benchmark in a detached screen session
task release       # Validate local GoReleaser artifacts (snapshot by default)
```

### Project Structure

```
mnemosyne/
├── cmd/                      # CLI commands (Cobra)
│   ├── root.go               # Root command + welcome message
│   ├── version.go            # version subcommand
│   ├── init.go               # Initialize a collection
│   ├── add.go                # Add a document (embeds + stores vector)
│   ├── update.go             # Update an existing document by ID
│   ├── list.go               # List documents
│   ├── delete.go             # Delete a document by ID
│   ├── forget.go             # Delete an entire collection
│   ├── search.go             # Search (hybrid: FTS5 + vector + fusion)
│   ├── setup.go              # Download ONNX Runtime + ML models
│   ├── export.go             # Export collections to JSONL
│   ├── import.go             # Import collections from JSONL (auto-embeds if needed)
│   ├── import_claude.go      # Claude Code memory import backend
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
│   └── search/               # Hybrid search + fusion
│       ├── bm25.go           # In-memory BM25 scoring for vector-first rerank
│       ├── hybrid.go         # Search engine (orchestrates FTS/vector/fusion)
│       └── rrf.go            # Result types + Reciprocal Rank Fusion algorithm
├── models/                   # ONNX model files (gitignored)
├── lib/                      # Native libraries (gitignored)
├── main.go                   # Entry point
├── Taskfile.yml              # Build/test/run tasks
├── ROADMAP.md                # Future features and out-of-scope ideas
└── go.mod
```

### Technology Stack

| Component        | Library                                                                 |
| ---------------- | ----------------------------------------------------------------------- |
| CLI              | [Cobra](https://github.com/spf13/cobra)                                 |
| SQLite driver    | [go-sqlite3](https://github.com/mattn/go-sqlite3) (CGO)                 |
| Vector search    | [sqlite-vec](https://github.com/asg017/sqlite-vec)                      |
| Full-text search | SQLite FTS5 (built-in)                                                  |
| ML inference     | [ONNX Runtime](https://github.com/yalue/onnxruntime_go)                 |
| Tokenizer        | [sugarme/tokenizer](https://github.com/sugarme/tokenizer) (HuggingFace) |
| Embedding model  | Jina v5 text nano retrieval (768-dim)                                   |
| Reranker model   | ms-marco-MiniLM-L-6-v2 (cross-encoder)                                  |
| Task runner      | [Task](https://taskfile.dev/)                                           |

### Creating a Release

To create a new versioned release with binaries and an automated changelog:

1. Create and push an annotated `v*` tag:
   ```bash
   git tag -a v0.3.0 -m "Release v0.3.0"
   git push origin v0.3.0
   ```
2. The `Release` GitHub Action workflow will automatically trigger from the tag.
3. The workflow publishes:
   - macOS `amd64` and `arm64` archives from GoReleaser on a macOS runner.
   - Linux `amd64` and `arm64` archives from native Ubuntu 22.04 runners. These
     Linux binaries use CGO and target an Ubuntu 22.04-era glibc baseline.
   - Windows `amd64` zip archives from the native Windows/MSYS2 job.
4. The release archives contain the CLI and project files. ONNX Runtime and model
   files are still installed separately by `mnemosyne setup` or first-use setup.

Use `task release` only to validate local GoReleaser artifacts. It runs a
snapshot by default; `CONFIRM=true task release` runs GoReleaser's official mode
for the current tag but does not create or push tags.

## Acknowledgements

Built with the help of [RunWield](https://github.com/gandazgul/runwield),
[OpenCode](https://opencode.ai/), [Pi](https://pi.dev), and various LLMs.

## License

MIT
