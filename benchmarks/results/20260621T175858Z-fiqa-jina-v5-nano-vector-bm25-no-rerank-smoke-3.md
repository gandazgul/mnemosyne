# Mnemosyne BEIR fiqa Results

- Generated: 2026-06-21T17:58:58Z
- Dataset: `fiqa`
- Queries: 3
- Corpus documents: 57638
- Search limit: 100
- Source mode: `hybrid`
- Fusion: `vector-bm25`
- BM25 weight: `0.3`
- Rerank candidates: 300
- Rerank enabled: False
- Run label: `jina-v5-nano`
- Mnemosyne config: `/Users/gandazgul/Documents/web/mnemosyne/cli/configs/jina-v5-text-nano-retrieval.yaml`

| Metric | Score |
| --- | ---: |
| `ndcg@10` | 0.2044 |
| `mrr@10` | 0.3333 |
| `recall@10` | 0.1667 |
| `recall@100` | 0.8667 |
| `map@100` | 0.2018 |

## Breakdown

| Detail | Count |
| --- | ---: |
| `queries` | 3 |
| `queries_with_rank_1_hit` | 1 |
| `queries_with_top_10_hit` | 1 |
| `queries_with_top_100_hit` | 3 |
| `queries_missing_at_100` | 0 |

### First Relevant Rank

| Bucket | Queries |
| --- | ---: |
| `rank_1` | 1 |
| `rank_2_3` | 0 |
| `rank_4_10` | 0 |
| `rank_11_100` | 2 |
| `missing@100` | 0 |

### Missing At 100

| Query ID | Missing Doc IDs | Query | Top 10 Doc IDs |
| --- | --- | --- | --- |
| 4641 | 319954, 397358 | Where should I park my rainy-day / emergency fund? | 580025, 497993, 376148, 32833, 583695, 538023, 285812, 527939, 282623, 108978 |

### Lowest MRR@10

| Query ID | MRR@10 | Recall@100 | First Relevant Rank | Query |
| --- | ---: | ---: | ---: | --- |
| 4641 | 0.0000 | 0.6000 | 11 | Where should I park my rainy-day / emergency fund? |
| 7803 | 0.0000 | 1.0000 | 26 | Can the Delta be used to calculate the option premium given a certain target? |
| 5503 | 1.0000 | 1.0000 | 1 | Tax considerations for selling a property below appraised value to family? |
