# Mnemoteca Roadmap

This document outlines planned features, later ideas, and explicitly out-of-scope
concepts for Mnemoteca.

## Planned Features

- **Re-indexing tool**: add a command, for example
  `mnemoteca collection re-index`, to safely regenerate FTS and vector data when
  a user changes the embedding model or vector dimensions in the config.
  - **Process**: rebuild in place, including `docs_vec` recreation and
    re-embedding, with progress output.
  - **Safety**: wrap the rebuild in a SQLite transaction for rollback on failure
    or cancellation.
- **Benchmarking**: use HuggingFace and other embedding and memory benchmarks to
  validate vector-search quality and embedding performance. Publish reproducible
  results linked from the README.

## Completed

Historical command examples in this section are kept only where they describe
behavior that shipped before the rename.

- **Default score thresholds**: low-scoring results are filtered by default.
  Reranker threshold `0.0` filters negative-logit results. RRF threshold `0.01`
  filters very low rank single-source results. Configure thresholds with
  `config.yaml` or `--threshold`. Use `--no-threshold` to disable.
- **Export/import backup and restore**: dump collections and restore them on
  other machines.
- **Semantic chunking and Markdown ingestion**: chunk Markdown files on semantic
  boundaries so project `.md` files can become contextual memories.
- **Memory classification and metadata**: support metadata and tags during
  ingestion to classify memories.
- **Short name flag**: add `-n` as a short flag for `--name` across relevant
  commands.
- **Init safety**: before the rename, `mnemosyne init` in a directory that
  matched an existing collection name errored to prevent accidental linking. The
  same behavior remains in `mnemoteca init`.
- **Global collection flag**: add `-g` or `--global` as a shortcut for
  `--name global`.

## Ideas for Later

- **Daemon/server mode for batch ingest**: keep ONNX models loaded in memory in a
  background process for dedicated batch import or ingest workflows.
- **Interactive TUI**: use a terminal UI, such as Bubble Tea, to explore
  collections, scroll through document chunks, and live-preview search results.

## Out of Scope

- **Local LLM integration, `ask` command**: Mnemoteca is a specialized
  storage/retrieval tool for agents and scripts, not a direct QA chat interface.
- **Rich document ingestion**: complex formats such as PDFs and Word documents
  are out of scope. Ingestion remains focused on short snippets, plain text,
  Markdown, and basic HTML.
