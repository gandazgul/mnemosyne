# Mnemosyne BEIR fiqa Results

- Generated: 2026-06-21T03:44:04Z
- Dataset: `fiqa`
- Queries: 3
- Corpus documents: 300
- Corpus subset limit: 300
- Search limit: 100
- Rerank candidates: 100
- Rerank enabled: False
- Run label: `jina-v5-nano-smoke`
- Mnemosyne config: `/Users/gandazgul/Documents/web/mnemosyne/cli/configs/jina-v5-text-nano-retrieval.yaml`

| Metric | Score |
| --- | ---: |
| `ndcg@10` | 1.0000 |
| `mrr@10` | 1.0000 |
| `recall@10` | 1.0000 |
| `recall@100` | 1.0000 |
| `map@100` | 1.0000 |

## Breakdown

| Detail | Count |
| --- | ---: |
| `queries` | 3 |
| `queries_with_rank_1_hit` | 3 |
| `queries_with_top_10_hit` | 3 |
| `queries_with_top_100_hit` | 3 |
| `queries_missing_at_100` | 0 |

### First Relevant Rank

| Bucket | Queries |
| --- | ---: |
| `rank_1` | 3 |
| `rank_2_3` | 0 |
| `rank_4_10` | 0 |
| `rank_11_100` | 0 |
| `missing@100` | 0 |

### Lowest MRR@10

| Query ID | MRR@10 | Recall@100 | First Relevant Rank | Query |
| --- | ---: | ---: | ---: | --- |
| 4641 | 1.0000 | 1.0000 | 1 | Where should I park my rainy-day / emergency fund? |
| 5503 | 1.0000 | 1.0000 | 1 | Tax considerations for selling a property below appraised value to family? |
| 7803 | 1.0000 | 1.0000 | 1 | Can the Delta be used to calculate the option premium given a certain target? |
