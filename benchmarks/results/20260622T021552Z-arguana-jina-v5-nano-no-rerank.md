# Mnemosyne BEIR arguana Results

- Generated: 2026-06-22T02:15:52Z
- Dataset: `arguana`
- Queries: 1406
- Corpus documents: 8674
- Search limit: 100
- Source mode: `hybrid`
- Fusion: `rrf`
- Rerank candidates: 100
- Rerank enabled: False
- Run label: `jina-v5-nano`
- Mnemosyne config: `/Users/gandazgul/Documents/web/mnemosyne/cli/configs/jina-v5-text-nano-retrieval.yaml`

| Metric | Score |
| --- | ---: |
| `ndcg@10` | 0.4314 |
| `mrr@10` | 0.2939 |
| `recall@10` | 0.8606 |
| `recall@100` | 0.9943 |
| `map@100` | 0.3018 |

## Breakdown

| Detail | Count |
| --- | ---: |
| `queries` | 1406 |
| `queries_with_rank_1_hit` | 2 |
| `queries_with_top_10_hit` | 1210 |
| `queries_with_top_100_hit` | 1398 |
| `queries_missing_at_100` | 8 |

### First Relevant Rank

| Bucket | Queries |
| --- | ---: |
| `rank_1` | 2 |
| `rank_2_3` | 724 |
| `rank_4_10` | 484 |
| `rank_11_100` | 188 |
| `missing@100` | 8 |

### Missing At 100

| Query ID | Missing Doc IDs | Query | Top 10 Doc IDs |
| --- | --- | --- | --- |
| test-free-speech-debate-yfsdfkhbwu-con03a | test-free-speech-debate-yfsdfkhbwu-con03b | Universities should exchange ideas not impose them  Of all possible institutions, for a university to suggest that it... | test-free-speech-debate-yfsdfkhbwu-con03a, test-free-speech-debate-yfsdfkhbwu-pro01a, test-free-speech-debate-yfsdfkh... |
| test-free-speech-debate-ldhwprhs-con01a | test-free-speech-debate-ldhwprhs-con01b | Regardless of the views expressed, freedom of speech means that all opinions should be heard.  Allowing politicians t... | test-free-speech-debate-ldhwprhs-con01a, training-digital-freedoms-fehwbawdh-con01a, training-digital-freedoms-fehwba... |
| test-free-speech-debate-nshbcsbawc-pro04a | test-free-speech-debate-nshbcsbawc-pro04b | Freedom of expression, like any right is fairly meaningless if it’s only respected when it’s convenient.  Recognising... | test-free-speech-debate-nshbcsbawc-pro04a, test-free-speech-debate-nshbcsbawc-pro01a, test-law-sdfclhrppph-con02a, tr... |
| test-international-amehbuaisji-pro01a | test-international-amehbuaisji-pro01b | The ICC is a force for good, and the all states should be seen to be standing fully behind it.  The International Cri... | test-international-amehbuaisji-pro01a, test-law-hrilpgwhwr-pro03a, validation-law-hrilhbiccfg-pro01a, test-law-hrilpg... |
| test-education-ufsdfkhbwu-con03a | test-education-ufsdfkhbwu-con03b | Universities should exchange ideas not impose them  Of all possible institutions, for a university to suggest that it... | test-free-speech-debate-yfsdfkhbwu-con03a, test-free-speech-debate-yfsdfkhbwu-pro01a, test-free-speech-debate-yfsdfkh... |
| test-politics-dhwem-pro06a | test-politics-dhwem-pro06b | PMCs can be made much more legitimate by regulation  Currently mercenary work as a profession is not regulated by law... | test-politics-dhwem-pro06a, training-international-apdwhbpa-pro01a, test-politics-dhwem-pro01b, test-politics-dhwem-c... |
| test-science-sghwbdgmo-con03a | test-science-sghwbdgmo-con03b | Genetically modified organisms will prevent starvation due to global climate changes.  The temperature of the earth i... | test-science-sghwbdgmo-con03a, test-science-sghwbdgmo-pro02a, test-science-sghwbdgmo-pro02b, test-science-sghwbdgmo-c... |
| test-society-asfhwapg-con04a | test-society-asfhwapg-con04b | A liability regime not patents.  There are alternatives to the kind of blanket patenting that stifles innovation and... | test-society-asfhwapg-con04a, test-society-asfhwapg-pro02a, test-society-asfhwapg-pro02b, test-society-asfhwapg-con02... |

### Lowest MRR@10

| Query ID | MRR@10 | Recall@100 | First Relevant Rank | Query |
| --- | ---: | ---: | ---: | --- |
| test-education-ufsdfkhbwu-con03a | 0.0000 | 0.0000 | missing | Universities should exchange ideas not impose them  Of all possible institutions, for a university to suggest that it... |
| test-free-speech-debate-ldhwprhs-con01a | 0.0000 | 0.0000 | missing | Regardless of the views expressed, freedom of speech means that all opinions should be heard.  Allowing politicians t... |
| test-free-speech-debate-nshbcsbawc-pro04a | 0.0000 | 0.0000 | missing | Freedom of expression, like any right is fairly meaningless if it’s only respected when it’s convenient.  Recognising... |
| test-free-speech-debate-yfsdfkhbwu-con03a | 0.0000 | 0.0000 | missing | Universities should exchange ideas not impose them  Of all possible institutions, for a university to suggest that it... |
| test-international-amehbuaisji-pro01a | 0.0000 | 0.0000 | missing | The ICC is a force for good, and the all states should be seen to be standing fully behind it.  The International Cri... |
| test-politics-dhwem-pro06a | 0.0000 | 0.0000 | missing | PMCs can be made much more legitimate by regulation  Currently mercenary work as a profession is not regulated by law... |
| test-science-sghwbdgmo-con03a | 0.0000 | 0.0000 | missing | Genetically modified organisms will prevent starvation due to global climate changes.  The temperature of the earth i... |
| test-society-asfhwapg-con04a | 0.0000 | 0.0000 | missing | A liability regime not patents.  There are alternatives to the kind of blanket patenting that stifles innovation and... |
| test-culture-ascidfakhba-con01a | 0.0000 | 1.0000 | 28 | Artists have a fundamental property right over their creative output  Whatever the end product, be it music, film, sc... |
| test-culture-ascidfakhba-con02a | 0.0000 | 1.0000 | 16 | Artists should retain the right to control their work’s interaction with the public space even if their work is publi... |
