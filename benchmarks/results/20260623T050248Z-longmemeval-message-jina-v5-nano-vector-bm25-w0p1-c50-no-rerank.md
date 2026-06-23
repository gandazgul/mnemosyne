# Mnemosyne LongMemEval Results

- Generated: 2026-06-23T05:02:48Z
- Dataset: `longmemeval_s_cleaned`
- Questions: 500
- Document mode: `message`
- Search limit: 50
- Rerank candidates: 50
- Fusion: `vector-bm25`
- BM25 weight: `0.1`
- Rerank enabled: False
- Run label: `jina-v5-nano`
- Mnemosyne config: `/Users/gandazgul/Documents/web/mnemosyne/cli/configs/jina-v5-text-nano-retrieval.yaml`

| Metric | Score |
| --- | ---: |
| `recall_any@5` | 0.9880 |
| `recall_any@10` | 0.9980 |
| `recall_all@5` | 0.9300 |
| `recall_all@10` | 0.9820 |
| `mrr@10` | 0.9586 |
| `ndcg@10` | 0.9575 |

## First Relevant Rank

| Bucket | Queries |
| --- | ---: |
| `rank_1` | 467 |
| `rank_2_3` | 22 |
| `rank_4_5` | 5 |
| `rank_6_10` | 5 |
| `missing@10` | 1 |

## Per Type

| Question Type | Count | R@5 | R@10 | nDCG@10 |
| --- | ---: | ---: | ---: | ---: |
| `knowledge-update` | 78 | 1.0000 | 1.0000 | 0.9871 |
| `multi-session` | 133 | 0.9925 | 1.0000 | 0.9496 |
| `single-session-assistant` | 56 | 1.0000 | 1.0000 | 1.0000 |
| `single-session-preference` | 30 | 0.9333 | 1.0000 | 0.8975 |
| `single-session-user` | 70 | 1.0000 | 1.0000 | 0.9866 |
| `temporal-reasoning` | 133 | 0.9774 | 0.9925 | 0.9284 |

## Missing At 5

| Query ID | Type | Question | Answer Sessions | Top 10 Sessions |
| --- | --- | --- | --- | --- |
| 60bf93ed_abs | multi-session | How many days did it take for my iPad case to arrive after I bought it? | answer_e0956e0a_abs_2, answer_e0956e0a_abs_1 | 1e91cdf0, 841da171_2, 36d5bbde_1, cdf068b1_3, c1e170f0_1, answer_e0956e0a_abs_2, answer_e0956e0a_abs_1, 9ef698bc_2, u... |
| d6233ab6 | single-session-preference | I've been feeling nostalgic lately. Do you think it would be a good idea to attend my high school reunion? | answer_b0fac439 | 94bc18df_3, 0e726047, 32f28c7b_1, ultrachat_125013, f916c63a_2, ecfd2047_1, ultrachat_329160, c927ffbb, e6c3a50a, ans... |
| 1c0ddc50 | single-session-preference | Can you suggest some activities I can do during my commute to work? | answer_8da8c7e0 | a0aa5035, 2aa70c9c_1, b33e89b5_1, c7ca6dff, ultrachat_494933, answer_8da8c7e0, 2566382f_2, 99883f38_1, 2f0c1f4e, bdf3... |
| gpt4_e061b84f | temporal-reasoning | What is the order of the three sports events I participated in during the past month, from earliest to latest? | answer_8c64ce25_2, answer_8c64ce25_1, answer_8c64ce25_3 | ultrachat_194928, 0a6bf5e4_1, 9345f7dc_4, 1446b088_2, ultrachat_129606, answer_8c64ce25_2, 78d28576_2, 3275acf9, 539c... |
| gpt4_59149c78 | temporal-reasoning | I mentioned that I participated in an art-related event two weeks ago. Where was that event held at? | answer_d00ba6d1_1, answer_d00ba6d1_2 | 23754665, 765ce8a7_2, b2341a22, a8ac3d1d_1, ultrachat_58926, answer_d00ba6d1_2, answer_d00ba6d1_1, ultrachat_463998 |
| gpt4_8279ba03 | temporal-reasoning | What kitchen appliance did I buy 10 days ago? | answer_56521e66_1 | 50d66391_4, d1a1b9ea_1, 518d26d3_4, 570fe405, bb107057_2, 652c0717_3, b357fb8b_2, ultrachat_129760, a8ac3d1d_3, 89749... |
