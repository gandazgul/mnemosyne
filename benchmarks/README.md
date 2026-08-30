# Benchmarks

Mnemoteca includes benchmark harnesses for retrieval quality. The current
published memory-system result was produced before the rename, when the product
was named Mnemosyne. Current reproduction commands use Mnemoteca.

## Headline Results

These runs used local inference only: no cloud API and no LLM reranker.

| Benchmark | Document Mode | Model Config | Fusion | Recall@5 | Recall@10 | MRR@10 | nDCG@10 |
| --- | --- | --- | --- | ---: | ---: | ---: | ---: |
| LongMemEval | `message` | `jina-v5-nano` | `vector-bm25`, weight `0.10` | 0.9880 | 0.9980 | 0.9586 | 0.9575 |
| LongMemEval | `session` | `jina-v5-nano` | `vector-bm25`, weight `0.10` | 0.9920 | 0.9940 | 0.9491 | 0.9486 |

The `message` run is the more representative Mnemosyne result: each
user/assistant message was stored as a short document, and retrieved documents
were scored by whether they mapped back to the correct LongMemEval answer
session. The `session` run is included because it is the closest shape to
MemPalace's raw LongMemEval setup: one document per session.

Full generated result files:

- [`20260623T050248Z-longmemeval-message-jina-v5-nano-vector-bm25-w0p1-c50-no-rerank.md`](results/20260623T050248Z-longmemeval-message-jina-v5-nano-vector-bm25-w0p1-c50-no-rerank.md)
- [`20260623T050248Z-longmemeval-message-jina-v5-nano-vector-bm25-w0p1-c50-no-rerank.json`](results/20260623T050248Z-longmemeval-message-jina-v5-nano-vector-bm25-w0p1-c50-no-rerank.json)
- [`20260622T043321Z-longmemeval-session-jina-v5-nano-vector-bm25-w0p1-c50-no-rerank.md`](results/20260622T043321Z-longmemeval-session-jina-v5-nano-vector-bm25-w0p1-c50-no-rerank.md)
- [`20260622T043321Z-longmemeval-session-jina-v5-nano-vector-bm25-w0p1-c50-no-rerank.json`](results/20260622T043321Z-longmemeval-session-jina-v5-nano-vector-bm25-w0p1-c50-no-rerank.json)

## Reproduce LongMemEval

Build Mnemoteca first:

```bash
task build
```

Run a small mechanics check:

```bash
python3 benchmarks/longmemeval/run.py \
  --config ./configs/jina-v5-text-nano-retrieval.yaml \
  --run-label jina-v5-nano \
  --doc-mode message \
  --max-queries 2 \
  --max-sessions 8 \
  --no-rerank
```

Run the published short-document result shape:

```bash
python3 benchmarks/longmemeval/run.py \
  --config ./configs/jina-v5-text-nano-retrieval.yaml \
  --run-label jina-v5-nano \
  --doc-mode message \
  --no-rerank \
  --limit 50 \
  --rerank-candidates 50
```

Run the session-document comparison shape:

```bash
python3 benchmarks/longmemeval/run.py \
  --config ./configs/jina-v5-text-nano-retrieval.yaml \
  --run-label jina-v5-nano \
  --doc-mode session \
  --no-rerank \
  --limit 50 \
  --rerank-candidates 50
```

The harness downloads `longmemeval_s_cleaned.json` into
`benchmarks/data/longmemeval/` unless the file already exists or
`--skip-download` is used. Scratch databases are written under
`benchmarks/work/longmemeval/`. Result summaries are written to
`benchmarks/results/`.

The full message-mode run imported 246,738 short documents across 500 isolated
per-question databases and took about 7.1 hours on the original local run. The
session-mode run imported 23,796 documents and took about 1.3 hours.

Use `--reuse-db` only when the per-question databases for the same document
mode, model config, and collection already exist.

## Metrics

For LongMemEval, each retrieved document is mapped back to its
`longmemeval_session_id` metadata value before scoring.

- `recall_any@5` / `recall_any@10`: at least one correct answer session appears
  in the top 5 or top 10.
- `recall_all@5` / `recall_all@10`: all labelled answer sessions appear in the
  top 5 or top 10. This is stricter for multi-session questions.
- `MRR@10`: reciprocal rank of the first correct answer session in the top 10.
- `nDCG@10`: ranking quality over the top 10, giving more credit when relevant
  sessions are near the top.

The generated Markdown files also include first-relevant-rank buckets,
per-question-type scores, and misses at 5.

## Comparison Notes

These numbers are retrieval metrics, not end-to-end question-answering accuracy.
They measure whether Mnemosyne retrieved the right memory session in the dated
pre-rename run. They do not measure whether an LLM can answer correctly after
retrieval.

For a rough local-only comparison, MemPalace reports LongMemEval raw ChromaDB at
96.6% R@5 / 98.2% R@10 / 0.889 nDCG@10 and hybrid v2 at 98.4% R@5 / 99.0%
R@10 / 0.934 nDCG@10. Mnemosyne's message-mode result was 98.8% Recall@5,
99.8% Recall@10, and 0.9575 nDCG@10 without an LLM reranker.

The comparison is not perfectly apples-to-apples:

- Mnemosyne reported `recall_any@K`, plus the stricter `recall_all@K` for
  questions with multiple answer sessions.
- MemPalace also reports higher LLM-reranked numbers. Those add a reader model
  after retrieval, while the Mnemosyne results above did not.
- MemPalace's own benchmark documentation flags its 100% LongMemEval result as
  partly tuned on known failures, and separately publishes a clean held-out
  no-LLM hybrid result of 98.4% R@5.

## BEIR SciFact

The BEIR harness is useful for standard document-retrieval checks on SciFact:

```bash
task bench:scifact-smoke
task bench:scifact
```

Useful flags:

```bash
task bench:scifact -- --no-rerank
task bench:scifact -- --reuse-db
task bench:scifact -- --fts-only --no-rerank
task bench:scifact -- --vector-only --no-rerank
task bench:scifact -- --fusion vector-bm25 --bm25-weight 0.10 --rerank-candidates 300 --no-rerank
```

When `--config` is provided, the harness passes it to Mnemoteca as
`MNEMOTECA_CONFIG`. When `--run-label` is provided, the benchmark uses a
separate work database under `benchmarks/work/` and includes the label in result
filenames. This prevents runs with different embedding dimensions from sharing a
SQLite vector table by accident. Use `--reuse-db` for `--fts-only` comparisons
when possible, because a fresh benchmark import still embeds vectorless corpus
files through the normal import path.

Current harness labels and scratch paths use Mnemoteca names for new runs.
Dated generated result files keep the names produced at the time of the run.
