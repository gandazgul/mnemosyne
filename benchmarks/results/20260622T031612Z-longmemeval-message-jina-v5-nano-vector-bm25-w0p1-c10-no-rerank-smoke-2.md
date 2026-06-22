# Mnemosyne LongMemEval Results

- Generated: 2026-06-22T03:16:12Z
- Dataset: `longmemeval_s_cleaned`
- Questions: 2
- Document mode: `message`
- Search limit: 10
- Rerank candidates: 10
- Fusion: `vector-bm25`
- BM25 weight: `0.1`
- Rerank enabled: False
- Max sessions per question: 8 plus answer sessions
- Run label: `jina-v5-nano`
- Mnemosyne config: `/Users/gandazgul/Documents/web/mnemosyne/cli/configs/jina-v5-text-nano-retrieval.yaml`

| Metric | Score |
| --- | ---: |
| `recall_any@5` | 1.0000 |
| `recall_any@10` | 1.0000 |
| `recall_all@5` | 1.0000 |
| `recall_all@10` | 1.0000 |
| `mrr@10` | 1.0000 |
| `ndcg@10` | 1.0000 |

## First Relevant Rank

| Bucket | Queries |
| --- | ---: |
| `rank_1` | 2 |
| `rank_2_3` | 0 |
| `rank_4_5` | 0 |
| `rank_6_10` | 0 |
| `missing@10` | 0 |

## Per Type

| Question Type | Count | R@5 | R@10 | nDCG@10 |
| --- | ---: | ---: | ---: | ---: |
| `single-session-user` | 2 | 1.0000 | 1.0000 | 1.0000 |
