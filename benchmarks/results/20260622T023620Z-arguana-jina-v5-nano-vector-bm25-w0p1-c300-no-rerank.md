# Mnemosyne BEIR arguana Results

- Generated: 2026-06-22T02:36:20Z
- Dataset: `arguana`
- Queries: 1406
- Corpus documents: 8674
- Search limit: 100
- Source mode: `hybrid`
- Fusion: `vector-bm25`
- BM25 weight: `0.1`
- Rerank candidates: 300
- Rerank enabled: False
- Run label: `jina-v5-nano`
- Mnemosyne config: `/Users/gandazgul/Documents/web/mnemosyne/cli/configs/jina-v5-text-nano-retrieval.yaml`

| Metric | Score |
| --- | ---: |
| `ndcg@10` | 0.4718 |
| `mrr@10` | 0.3280 |
| `recall@10` | 0.9154 |
| `recall@100` | 0.9936 |
| `map@100` | 0.3327 |

## Breakdown

| Detail | Count |
| --- | ---: |
| `queries` | 1406 |
| `queries_with_rank_1_hit` | 3 |
| `queries_with_top_10_hit` | 1287 |
| `queries_with_top_100_hit` | 1397 |
| `queries_missing_at_100` | 9 |

### First Relevant Rank

| Bucket | Queries |
| --- | ---: |
| `rank_1` | 3 |
| `rank_2_3` | 829 |
| `rank_4_10` | 455 |
| `rank_11_100` | 110 |
| `missing@100` | 9 |

### Missing At 100

| Query ID | Missing Doc IDs | Query | Top 10 Doc IDs |
| --- | --- | --- | --- |
| test-free-speech-debate-yfsdfkhbwu-con03a | test-free-speech-debate-yfsdfkhbwu-con03b | Universities should exchange ideas not impose them  Of all possible institutions, for a university to suggest that it... | test-free-speech-debate-yfsdfkhbwu-con03a, test-free-speech-debate-yfsdfkhbwu-pro01a, test-free-speech-debate-yfsdfkh... |
| test-free-speech-debate-ldhwprhs-con01a | test-free-speech-debate-ldhwprhs-con01b | Regardless of the views expressed, freedom of speech means that all opinions should be heard.  Allowing politicians t... | test-free-speech-debate-ldhwprhs-con01a, training-digital-freedoms-fehwbawdh-con01a, training-law-aullgsmhwchs-con01b... |
| test-free-speech-debate-nshbcsbawc-pro04a | test-free-speech-debate-nshbcsbawc-pro04b | Freedom of expression, like any right is fairly meaningless if it’s only respected when it’s convenient.  Recognising... | test-free-speech-debate-nshbcsbawc-pro04a, training-free-speech-debate-nvhsibsv-con01a, training-free-speech-debate-n... |
| test-international-amehbuaisji-pro01a | test-international-amehbuaisji-pro01b | The ICC is a force for good, and the all states should be seen to be standing fully behind it.  The International Cri... | test-international-amehbuaisji-pro01a, validation-law-hrilhbiccfg-pro01a, test-law-hrilpgwhwr-con05b, test-law-hrilpg... |
| test-philosophy-pphbclsbs-pro01a | test-philosophy-pphbclsbs-pro01b | National security is something that must be protected even at the cost of  Terrorism is part of the modern world and... | test-philosophy-pphbclsbs-pro01a, training-law-cplghrhwrgo-con02a, validation-international-gsidfphb-con02a, validati... |
| test-education-ufsdfkhbwu-con03a | test-education-ufsdfkhbwu-con03b | Universities should exchange ideas not impose them  Of all possible institutions, for a university to suggest that it... | test-free-speech-debate-yfsdfkhbwu-con03a, test-free-speech-debate-yfsdfkhbwu-pro01a, test-free-speech-debate-yfsdfkh... |
| test-politics-dhwem-pro06a | test-politics-dhwem-pro06b | PMCs can be made much more legitimate by regulation  Currently mercenary work as a profession is not regulated by law... | test-politics-dhwem-pro06a, training-international-apdwhbpa-pro01a, test-politics-dhwem-pro01b, test-politics-dhwem-c... |
| test-science-sghwbdgmo-con03a | test-science-sghwbdgmo-con03b | Genetically modified organisms will prevent starvation due to global climate changes.  The temperature of the earth i... | test-science-sghwbdgmo-con03a, test-science-sghwbdgmo-con02a, test-science-sghwbdgmo-con02b, test-science-sghwbdgmo-p... |
| test-society-asfhwapg-con04a | test-society-asfhwapg-con04b | A liability regime not patents.  There are alternatives to the kind of blanket patenting that stifles innovation and... | test-society-asfhwapg-con04a, test-society-asfhwapg-pro02a, test-society-asfhwapg-con02b, test-society-asfhwapg-con02... |

### Lowest MRR@10

| Query ID | MRR@10 | Recall@100 | First Relevant Rank | Query |
| --- | ---: | ---: | ---: | --- |
| test-education-ufsdfkhbwu-con03a | 0.0000 | 0.0000 | missing | Universities should exchange ideas not impose them  Of all possible institutions, for a university to suggest that it... |
| test-free-speech-debate-ldhwprhs-con01a | 0.0000 | 0.0000 | missing | Regardless of the views expressed, freedom of speech means that all opinions should be heard.  Allowing politicians t... |
| test-free-speech-debate-nshbcsbawc-pro04a | 0.0000 | 0.0000 | missing | Freedom of expression, like any right is fairly meaningless if it’s only respected when it’s convenient.  Recognising... |
| test-free-speech-debate-yfsdfkhbwu-con03a | 0.0000 | 0.0000 | missing | Universities should exchange ideas not impose them  Of all possible institutions, for a university to suggest that it... |
| test-international-amehbuaisji-pro01a | 0.0000 | 0.0000 | missing | The ICC is a force for good, and the all states should be seen to be standing fully behind it.  The International Cri... |
| test-philosophy-pphbclsbs-pro01a | 0.0000 | 0.0000 | missing | National security is something that must be protected even at the cost of  Terrorism is part of the modern world and... |
| test-politics-dhwem-pro06a | 0.0000 | 0.0000 | missing | PMCs can be made much more legitimate by regulation  Currently mercenary work as a profession is not regulated by law... |
| test-science-sghwbdgmo-con03a | 0.0000 | 0.0000 | missing | Genetically modified organisms will prevent starvation due to global climate changes.  The temperature of the earth i... |
| test-society-asfhwapg-con04a | 0.0000 | 0.0000 | missing | A liability regime not patents.  There are alternatives to the kind of blanket patenting that stifles innovation and... |
| test-culture-cgeeghwmeo-pro03a | 0.0000 | 1.0000 | 21 | Avoids self-segregation  In a time when the US has begun to overcome racial segregation, and legal discrimination in... |
