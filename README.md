# Mnemoteca

<p align="center">
  <img src="logo.png" alt="Mnemoteca logo" width="200"/>
</p>

**Mnemoteca — local memory for every agent.**

Mnemoteca is a local document storage and retrieval CLI built in Go. Store short
memories and retrieve them with hybrid search: vector similarity with a small
BM25 lexical boost, plus optional local cross-encoder reranking.

All ML inference runs locally with ONNX Runtime. No cloud API is required.

## Features

- SQLite document storage with metadata support.
- Semantic Markdown chunking with `yuin/goldmark` AST.
- SQLite FTS5 full-text search with BM25 ranking.
- sqlite-vec vector search with cosine similarity.
- Vector-first hybrid search with BM25 lexical reranking.
- Explicit legacy RRF fusion mode.
- Local ONNX embedding and optional reranking models.
- Automatic setup for ONNX Runtime and model downloads.
- Configurable local models.
- No cloud dependency.

## Prerequisites

For the one-command install, you need `sh`, `curl` or `wget`, and `tar`.

For building from source:

- **Go 1.21+** — [Install Go](https://go.dev/dl/).
- **GCC** — required by `mattn/go-sqlite3` (CGO). On macOS, run
  `xcode-select --install`.
- **Task** — install with `brew install go-task` or see
  [taskfile.dev](https://taskfile.dev/installation/).

## Installation

### Install the CLI

Install the latest macOS or Linux release to `~/.local/bin`:

```bash
curl -fsSL https://raw.githubusercontent.com/gandazgul/mnemoteca/main/install.sh | sh
```

The installer installs `mnemoteca`. If it finds a final Mnemosyne command, it
can guide an explicit export/import migration and offer an optional compatibility
link. If no terminal is available for prompts, it installs Mnemoteca and skips
migration, cleanup, and link creation.

Make sure `~/.local/bin` is on your `PATH`, then run:

```bash
mnemoteca --help
mnemoteca setup
```

`mnemoteca setup` downloads ONNX Runtime and the local embedding/reranker
models. Setup is idempotent. It also runs automatically on first `add` or
`search`. No HuggingFace account or API token is required for the built-in
models.

### Windows release ZIP

Download the Windows release artifact named
`mnemoteca_VERSION_windows_amd64.zip` from the Mnemoteca GitHub release, extract
it, place `mnemoteca.exe` in a directory on `PATH`, then run:

```powershell
mnemoteca.exe --help
mnemoteca.exe setup
```

The POSIX `install.sh` migration prompts and compatibility symlink do not apply
on Windows. See the migration guide before you replace an old installation.

### Migrating from Mnemosyne

Use [Migrate from Mnemosyne](docs/migrate-from-mnemosyne.md) before you remove
old data or integrations. The guide covers:

- installer-guided macOS/Linux migration;
- manual macOS/Linux migration;
- manual Windows migration with no compatibility shim;
- partial-import recovery;
- cleanup safety;
- integration replacement order.

## Connect Mnemoteca to your agent

Install the CLI once, then add the integration for each assistant you use. The
integrations share the same local Mnemoteca collections, so memories can follow
you across agents.

- [claudecode-mnemoteca](https://github.com/gandazgul/claudecode-mnemoteca) —
  Claude Code skill.
- [opencode-mnemoteca](https://github.com/gandazgul/opencode-mnemoteca) —
  OpenCode plugin with memory tools and core-memory injection.
- [pi-mnemoteca](https://github.com/gandazgul/pi-mnemoteca) — Pi extension with
  memory tools and core-memory injection.

See each integration README for install commands, agent-specific configuration,
upgrade-from-Mnemosyne steps, and usage details. The agent-facing `memory_*`
tool names stay stable.

### Other agents

For agents or IDEs that support `AGENTS.md`, project instructions, or similar
custom guidance, add this instruction block after installing the CLI:

```markdown
## Memory System

- Use `mnemoteca search -f plain [query]` and `mnemoteca search -g -f plain [query]` to search relevant memories. Use this before making decisions or taking actions.
- After significant decisions, use `mnemoteca add "memory content"` to save a concise fact. Also do this if the user explicitly asks you to remember something. Use `mnemoteca add -g "memory content"` for cross-project preferences.
- Delete contradicted memories with `mnemoteca delete [memory id]` after storing updated memories with `mnemoteca add ...` or `mnemoteca add -g ...`.
- Mark critical, always-relevant context as core with `-t core`, but use it sparingly. You can also use other tags with repeated `-t` flags, such as `mnemoteca add "database is sqlite" -t core -t tech-stack`.
- When you are done with a session, store memories that are relevant to the user and the project. This helps you recall important information in future sessions.
```

### Sleep

For conservative memory maintenance, use the [Sleep prompt](docs/sleep-prompt.md).
It preserves durable context while it removes only exact duplicates, clearly
deprecated facts, explicitly superseded memories, or lossless consolidations.

## Quick Start

```bash
# See available commands
mnemoteca --help

# Check version
mnemoteca version

# Download ONNX Runtime and ML models
# This also happens automatically on first 'add' or 'search'.
mnemoteca setup

# Initialize a collection. The default name comes from the current directory.
mnemoteca init

# Add documents
mnemoteca add "Go is a statically typed programming language"
mnemoteca add "Rust focuses on memory safety and zero-cost abstractions"
mnemoteca add --file notes.txt
mnemoteca add --file README.md # Automatically chunks by semantic headings
printf 'Piped note' | mnemoteca add --stdin
mnemoteca add "Important project preference" -t core -t preference

# Search documents. Hybrid vector-first BM25 fusion is the default.
mnemoteca search "programming language"
mnemoteca search --limit 5 "systems programming"
mnemoteca search -f json --limit 10 "programming language"
mnemoteca search --fts-only --no-rerank "programming language"
mnemoteca search --vector-only --no-rerank "programming language"
mnemoteca search --fusion rrf --no-rerank "programming language"

# List documents
mnemoteca list
mnemoteca list -f plain
mnemoteca list --limit 5
mnemoteca list -t core

# List tags used in a collection
mnemoteca tags

# Update an existing document by ID
mnemoteca update 1 "Revised memory content"
mnemoteca update 1 --file revised-note.txt
printf 'Revised piped note' | mnemoteca update 1 --stdin
mnemoteca update 1 -t reviewed
mnemoteca update 1 --replace-tags -t core
mnemoteca update 1 --replace-tags

# Delete a document by ID
mnemoteca delete 1

# Use a named collection
mnemoteca init -n myproject
mnemoteca add -n myproject "some text"
mnemoteca update -n myproject 1 "revised text"
mnemoteca search -n myproject "some query"

# Use the global collection shortcut
mnemoteca add -g "Global memory"
mnemoteca list -g
mnemoteca update -g 1 "Revised global memory"
mnemoteca search --global "global memory"

# Show collections and database stats
mnemoteca collections
mnemoteca stats
mnemoteca stats --format json

# Delete an entire collection
mnemoteca forget -n myproject
mnemoteca forget -n myproject --yes

# Export a collection to JSONL. Vectors are included for fast import.
mnemoteca export --name myproject

# Export without vectors. Import will re-embed documents.
mnemoteca export --name myproject --no-embeddings

# Export all collections
mnemoteca export --all -o ./backups/

# Import a collection or a directory of collection JSONL files
mnemoteca import myproject.jsonl
mnemoteca import myproject.jsonl --name other
mnemoteca import --dir ./backups/

# Import Claude Code memory without modifying Claude files
mnemoteca import --agent claude --dry-run
mnemoteca import --agent claude --include-user
```

## Search Output Formats

The `search` command supports three output formats:

```bash
mnemoteca search -f color "query" # default human-readable color output
mnemoteca search -f plain "query" # stable plain text output
mnemoteca search -f json "query"  # machine-readable output for scripts and benchmarks
```

JSON search output is an object with `query`, `collection`, `count`, and a
ranked `results` array. Each result includes:

- `rank`, `document_id`, `collection_id`, `content`, `created_at`;
- parsed `metadata` when document metadata contains valid JSON;
- `metadata_raw` when metadata is present but not valid JSON;
- `rrf_score`, `fts_rank`, `vec_distance`, `reranker_score`, `is_reranked`;
- `sources`, such as `["fts", "vector"]`.

This format is intended for evaluation harnesses that must map retrieved
documents back to external corpus IDs stored in metadata.

## Configuration

Mnemoteca reads `~/.config/mnemoteca/config.yaml` when it exists. You can also
point a command or benchmark run at another config file:

```bash
MNEMOTECA_CONFIG=./configs/jina-v5-text-nano-retrieval.yaml mnemoteca search "query"
MNEMOTECA_CONFIG_PATH=./configs/jina-v5-text-nano-retrieval.yaml mnemoteca search "query"
MNEMOTECA_DB_PATH=/tmp/mnemoteca.db mnemoteca add "temporary note"
```

Default data locations:

- macOS/Linux: `$XDG_DATA_HOME/mnemoteca` or `~/.local/share/mnemoteca`;
- Windows: `%LOCALAPPDATA%\mnemoteca`, then `%APPDATA%\mnemoteca`, then
  `~/AppData/Local/mnemoteca`;
- database: `<data-dir>/mnemoteca.db`.

Use [config.example.yaml](config.example.yaml) as a starting point. Important
model fields are:

- `embedding.model_path`: directory containing `tokenizer.json`, `config.json`,
  and the configured ONNX file;
- `embedding.onnx_file`: model path relative to `embedding.model_path`, usually
  `onnx/model.onnx`;
- `embedding.dimensions`: vector size stored in SQLite. Use a fresh DB or
  re-import documents when this changes;
- `embedding.pooling`: `none`, `mean`, `cls`, or `last`;
- `embedding.query_prefix` and `embedding.document_prefix`: model-specific text
  prefixes;
- `reranker.model_path`: directory containing an ONNX cross-encoder reranker;
- `reranker.enabled`: set `false` to disable reranking globally for that config.

`mnemoteca setup` auto-downloads the built-in default models. For custom model
paths, place the model files yourself and run with `MNEMOTECA_CONFIG` or copy
the config to the default location.

The `configs/` directory includes benchmark-oriented profiles for larger
retrieval models, such as Jina v5 text retrieval. These profiles use a separate
benchmark DB when passed through the harness with `--config` and `--run-label`.

Mnemoteca does not read legacy config or environment names. If you migrate from
Mnemosyne, map each custom setting to `MNEMOTECA_CONFIG`,
`MNEMOTECA_CONFIG_PATH`, `MNEMOTECA_DB_PATH`, or
`~/.config/mnemoteca/config.yaml` explicitly.

## Export and Import

Mnemoteca supports JSONL export and import for backup and transfer.

**Export** writes one JSONL file per collection. Each document includes:

- `content` and `metadata` — the original document data;
- `vector` — the raw embedding, omitted with `--no-embeddings`;
- `original_document_id` — the source database ID for inspection.

**Import** reads a JSONL file and inserts documents into the database:

- If vectors are present, import is fast and model-independent.
- If vectors are missing, Mnemoteca initializes the embedder and creates vectors.
- If vectors are missing and the embedder cannot initialize, import returns an
  error.
- Import appends. Do not import the same backup twice unless duplicate documents
  are intended.

```bash
# Full export with vectors
mnemoteca export --name myproject

# Lightweight export without vectors
mnemoteca export --name myproject --no-embeddings

# Import
mnemoteca import myproject.jsonl
```

## Benchmarks

**Headline:** Mnemosyne, the pre-rename product, retrieved the correct memory
session in the top 5 for 98.8% of questions and in the top 10 for 99.8% of
questions on the full 500-question LongMemEval cleaned set, using short
per-message documents, local Jina v5 nano embeddings, vector-BM25 fusion, and no
LLM reranker.

See [benchmarks/README.md](benchmarks/README.md) for published results,
comparison notes, and current Mnemoteca reproduction commands.

Mnemoteca includes a BEIR-compatible retrieval benchmark harness for SciFact. It
downloads the public BEIR SciFact dataset, imports the corpus into an isolated
benchmark database, runs `mnemoteca search -f json` for judged test queries, and
writes JSON plus Markdown result summaries.

```bash
# Short mechanics check over 10 judged queries and a 500-document corpus subset
task bench:scifact-smoke

# Standard full SciFact test split over the complete corpus
task bench:scifact

# Run the full benchmark in a detached screen session
task bench:scifact-bg
```

Mnemoteca also includes a LongMemEval harness for memory-system retrieval
comparisons. LongMemEval is per-question: each question has its own haystack of
conversation sessions, so the harness builds an isolated Mnemoteca DB per
question and scores whether the retrieved documents map back to the answer
session IDs.

Published no-rerank LongMemEval results from the pre-rename Mnemosyne run:

| Document Mode | Recall@5 | Recall@10 | MRR@10 | nDCG@10 |
| ------------- | -------: | --------: | -----: | ------: |
| `message`     |   0.9880 |    0.9980 | 0.9586 |  0.9575 |
| `session`     |   0.9920 |    0.9940 | 0.9491 |  0.9486 |

```bash
# Tiny mechanics check: 2 questions, trimmed haystacks
task bench:longmemeval-smoke -- --config ./configs/jina-v5-text-nano-retrieval.yaml --run-label jina-v5-nano

# Session-doc mode: one document per session, user turns joined.
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
`MAP@100`. The LongMemEval harness reports session-level `recall_any@5`,
`recall_any@10`, `recall_all@5`, `recall_all@10`, `MRR@10`, and `nDCG@10`, plus
per-question-type breakdowns. Downloaded data and scratch databases are written
under `benchmarks/data/` and `benchmarks/work/`, which are gitignored.
Published result files are written to `benchmarks/results/`.

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

When `--config` is provided, the harness passes it to Mnemoteca as
`MNEMOTECA_CONFIG`. When `--run-label` is provided, the benchmark uses a
separate work DB under `benchmarks/work/` and includes the label in result
filenames. This keeps runs with different embedding dimensions from sharing a
SQLite vector table by accident.

The background task uses `screen` so long benchmark runs survive the shell that
started them:

```bash
screen -ls
tail -f benchmarks/work/beir/scifact/full-benchmark.log
screen -r mnemoteca-scifact-benchmark
screen -S mnemoteca-scifact-benchmark -X quit
```

## Developing and Contributing

Build and contribution workflows live here. User-facing install and integration
steps are covered above. Package-specific agent setup lives in the integration
READMEs.

### Available Tasks

```bash
task build            # Build the binary
task test             # Run all tests
task clean            # Remove build artifacts
task lint             # Run linter
task download-models  # Download ONNX models from HuggingFace
task bench:scifact-smoke # Run a small BEIR SciFact smoke benchmark
task bench:scifact    # Run the full BEIR SciFact benchmark
task bench:scifact-bg # Run the full benchmark in a detached screen session
task release          # Validate local GoReleaser artifacts
```

### Project Structure

```text
mnemoteca/
├── cmd/                      # CLI commands (Cobra)
├── internal/
│   ├── backup/               # JSONL export/import of collections
│   ├── config/               # Configuration loading and defaults
│   ├── db/                   # SQLite, FTS5, and sqlite-vec storage
│   ├── embedding/            # ONNX embedding
│   ├── reranker/             # ONNX cross-encoder reranker
│   ├── search/               # Hybrid search and fusion
│   └── setup/                # Auto-download of runtime and models
├── benchmarks/               # BEIR and LongMemEval harnesses
├── configs/                  # Example and benchmark configs
├── docs/                     # User and maintainer documentation
├── main.go                   # Entry point
├── Taskfile.yml              # Build/test/run tasks
└── go.mod
```

### Technology Stack

| Component | Library |
| --- | --- |
| CLI | [Cobra](https://github.com/spf13/cobra) |
| SQLite driver | [go-sqlite3](https://github.com/mattn/go-sqlite3) (CGO) |
| Vector search | [sqlite-vec](https://github.com/asg017/sqlite-vec) |
| Full-text search | SQLite FTS5 |
| ML inference | [ONNX Runtime](https://github.com/yalue/onnxruntime_go) |
| Tokenizer | [sugarme/tokenizer](https://github.com/sugarme/tokenizer) |
| Embedding model | Jina v5 text nano retrieval, 768 dimensions |
| Reranker model | ms-marco-MiniLM-L-6-v2 cross-encoder |
| Task runner | [Task](https://taskfile.dev/) |

### Creating a Release

To create a versioned release with binaries and an automated changelog:

1. Create and push an annotated `v*` tag:
   ```bash
   git tag -a v0.3.0 -m "Release v0.3.0"
   git push origin v0.3.0
   ```
2. The Release GitHub Action workflow triggers from the tag.
3. The workflow publishes:
   - macOS `amd64` and `arm64` archives from GoReleaser on a macOS runner;
   - Linux `amd64` and `arm64` archives from native Ubuntu 22.04 runners;
   - Windows `amd64` zip archives from the native Windows/MSYS2 job.
4. The release archives contain the CLI and project files. ONNX Runtime and model
   files are installed separately by `mnemoteca setup` or first-use setup.

Use `task release` only to validate local GoReleaser artifacts. It runs a
snapshot by default. `CONFIRM=true task release` runs GoReleaser's official mode
for the current tag but does not create or push tags.

## Acknowledgements

Built with the help of [RunWield](https://github.com/gandazgul/runwield),
[OpenCode](https://opencode.ai/), [Pi](https://pi.dev), and various LLMs.

## License

MIT
