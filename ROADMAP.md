# Mnemosyne Roadmap

This document outlines the planned features, ideas for the future, and explicitly out-of-scope concepts for Mnemosyne.

## 🎯 Planned Features (Priority)

* **Export/Import (Backup & Restore)**: Add commands to easily dump collections and restore them on other machines.
  * **Format**: JSONL (JSON Lines) to handle streaming and large collections.
  * **Data Included**: Include raw vector data to make imports fast and model-independent.
  * **Scope Options**: `mnemosyne collection -n name --export` (single collection) or `mnemosyne --export` (full DB, with size confirmation warning).
* **Re-indexing Tool**: Add a command (e.g., `mnemosyne collection re-index`) to safely regenerate FTS and Vector data if the user changes their embedding model or vector dimensions in the config.
  * **Process**: In-place rebuild (drop/recreate `docs_vec` virtual table with new dimensions, re-embed everything with a progress bar).
  * **Safety**: Wrapped in a SQLite transaction for safe rollback on failure or cancellation.

## ✅ Completed

* **Semantic Chunking & Markdown Ingestion**: Improve ingestion to intelligently chunk markdown files based on semantic boundaries. This allows entire project `.md` files to be ingested as contextual memories.
  * **Strategy**: Parse Markdown AST and chunk by headings.
  * **Context Preservation**: Prepend heading hierarchy path to chunk text for better LLM context.
  * **Fallback**: Use paragraph/sentence splitting if a section is too large.
* **Memory Classification & Metadata**: Support adding metadata/tags during ingestion to classify memories (e.g., "always-load" vs. "contextual"). This allows consuming tools and agents to know which memories must be read entirely versus which should be queried dynamically.
* **Short Name Flag**: Added `-n` as a short flag for `--name` across all relevant commands.
* **Init Safety**: Running `mnemosyne init` in a directory that matches an existing collection's name now errors out to prevent accidental linking.
* **Global Collection Flag**: Added `-g` or `--global` flag as a dedicated shortcut for `--name global` to streamline global memory access and prevent typos.

## 💡 Ideas for Later

* **Daemon/Server Mode (Batch Ingest)**: Keep ONNX models loaded in memory via a background process. While its utility for regular CLI usage is limited, it could serve well later for a dedicated batch ingest command.
* **Linux Releases**: Provide pre-compiled binaries for popular Linux distros to avoid the need for users to set up a Go environment or compile from source.
* **Interactive TUI**: An interactive terminal UI (using something like `bubbletea`) to visually explore collections, scroll through document chunks, and live-preview search results.

## 🚫 Out of Scope

* **Local LLM Integration (`ask` command)**: Mnemosyne is a specialized storage/retrieval tool meant to be used by agents and scripts, not a direct QA chat interface.
* **Rich Document Ingestion**: Parsing complex formats (PDFs, Word documents, etc.) is outside the scope. Ingestion will remain tightly focused on short snippets, plain text, markdown, and basic HTML.