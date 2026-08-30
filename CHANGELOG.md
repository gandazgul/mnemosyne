# Changelog

## Unreleased

### Changed

- Rename active documentation, installer examples, configuration paths, benchmark
  reproduction commands, and integration links to Mnemoteca.
- Add an authoritative migration guide for moving from Mnemosyne to Mnemoteca,
  including macOS/Linux installer prompts, manual export/import, Windows ZIP
  migration with no compatibility shim, partial-import recovery, cleanup safety,
  and integration replacement order.
- Rename release artifacts and installer defaults to Mnemoteca. The POSIX
  installer now installs `mnemoteca` and owns the guided Mnemosyne export/import
  transition.

## v0.2.3 (2026-06-23)

Historical entries below refer to the product name, paths, and repository links
that existed at the time of release.

### New Features

- **YAML config loading and JSON search output** — Load config from `~/.config/mnemosyne/config.yaml` with env override support. Add `-f json` output format for search. Support custom `onnx_file` in reranker config and `task_id` ONNX input for embedding models. ([c9b75a3](https://github.com/gandazgul/mnemosyne/commit/c9b75a3))
- **Vector-first BM25 fusion search** — Default search now retrieves vector candidates first, then applies in-memory BM25 reranking with blended scores. Add `--fusion`, `--bm25-weight`, `--fts-only`, `--vector-only` CLI flags. ([dc960dc](https://github.com/gandazgul/mnemosyne/commit/dc960dc))

### Bug Fixes and Improvements

- Fix unchecked `fmt.Fprintf` return values in `version.go`. ([d1ee2d7](https://github.com/gandazgul/mnemosyne/commit/d1ee2d7))
- Add logo to README header and ignore `.history/` directory. ([7298683](https://github.com/gandazgul/mnemosyne/commit/7298683))
- Publish LongMemEval full 500-question results and prune stale smoke test runs. ([80c3273](https://github.com/gandazgul/mnemosyne/commit/80c3273))
