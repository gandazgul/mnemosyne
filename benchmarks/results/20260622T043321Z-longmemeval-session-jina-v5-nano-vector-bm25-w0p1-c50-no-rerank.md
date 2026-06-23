# Mnemosyne LongMemEval Results

- Generated: 2026-06-22T04:33:21Z
- Dataset: `longmemeval_s_cleaned`
- Questions: 500
- Document mode: `session`
- Search limit: 50
- Rerank candidates: 50
- Fusion: `vector-bm25`
- BM25 weight: `0.1`
- Rerank enabled: False
- Run label: `jina-v5-nano`
- Mnemosyne config: `/Users/gandazgul/Documents/web/mnemosyne/cli/configs/jina-v5-text-nano-retrieval.yaml`

| Metric | Score |
| --- | ---: |
| `recall_any@5` | 0.9920 |
| `recall_any@10` | 0.9940 |
| `recall_all@5` | 0.9360 |
| `recall_all@10` | 0.9700 |
| `mrr@10` | 0.9491 |
| `ndcg@10` | 0.9486 |

## First Relevant Rank

| Bucket | Queries |
| --- | ---: |
| `rank_1` | 461 |
| `rank_2_3` | 23 |
| `rank_4_5` | 12 |
| `rank_6_10` | 1 |
| `missing@10` | 3 |

## Per Type

| Question Type | Count | R@5 | R@10 | nDCG@10 |
| --- | ---: | ---: | ---: | ---: |
| `knowledge-update` | 78 | 1.0000 | 1.0000 | 0.9832 |
| `multi-session` | 133 | 1.0000 | 1.0000 | 0.9506 |
| `single-session-assistant` | 56 | 0.9643 | 0.9643 | 0.9533 |
| `single-session-preference` | 30 | 1.0000 | 1.0000 | 0.9170 |
| `single-session-user` | 70 | 1.0000 | 1.0000 | 0.9785 |
| `temporal-reasoning` | 133 | 0.9850 | 0.9925 | 0.9156 |

## Missing At 5

| Query ID | Type | Question | Answer Sessions | Top 10 Sessions |
| --- | --- | --- | --- | --- |
| gpt4_4929293b | temporal-reasoning | What was the the life event of one of my relatives that I participated in a week ago? | answer_add9b013_2, answer_add9b013_1 | 4090cbea, bda611f6_3, c0dbabb8, sharegpt_KFhIUCO_0, f11109b1_1, cc8252e8, answer_add9b013_1, sharegpt_81riySf_0, cae6... |
| gpt4_8279ba03 | temporal-reasoning | What kitchen appliance did I buy 10 days ago? | answer_56521e66_1 | 518d26d3_4, 570fe405, 50d66391_4, d1a1b9ea_1, bb107057_2, 89749c78_2, b357fb8b_2, 2ef55f49_3, 5ee1c179_1, dea400f8_2 |
| ceb54acb | single-session-assistant | In our previous chat, you suggested 'sexual compulsions' and a few other options for alternative terms for certain be... | answer_sharegpt_cGdjmYo_0 | de293134_1, sharegpt_iTrNGx4_13, a1cbe83a, 63b72857, f6fd00cf, sharegpt_yePPued_0, dddce60a_1, 0bbd7094_1, 197d99cc,... |
| 58470ed2 | single-session-assistant | I was going through our previous conversation about The Library of Babel, and I wanted to confirm - what did Borges s... | answer_sharegpt_U4oCSfU_7 | sharegpt_xSFBEuR_0, sharegpt_ErOTMZ3_277, sharegpt_siSd9ET_0, sharegpt_oPOTyid_141, sharegpt_6QUDIXG_104, sharegpt_e8... |
