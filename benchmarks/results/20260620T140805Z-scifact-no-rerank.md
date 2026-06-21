# Mnemosyne BEIR scifact Results

- Generated: 2026-06-20T14:08:05Z
- Dataset: `scifact`
- Queries: 300
- Corpus documents: 5183
- Search limit: 100
- Rerank candidates: 100
- Rerank enabled: False

| Metric | Score |
| --- | ---: |
| `ndcg@10` | 0.6990 |
| `mrr@10` | 0.6688 |
| `recall@10` | 0.8238 |
| `recall@100` | 0.9500 |
| `map@100` | 0.6601 |

## Breakdown

| Detail | Count |
| --- | ---: |
| `queries` | 300 |
| `queries_with_rank_1_hit` | 177 |
| `queries_with_top_10_hit` | 251 |
| `queries_with_top_100_hit` | 285 |
| `queries_missing_at_100` | 15 |

### First Relevant Rank

| Bucket | Queries |
| --- | ---: |
| `rank_1` | 177 |
| `rank_2_3` | 42 |
| `rank_4_10` | 32 |
| `rank_11_100` | 34 |
| `missing@100` | 15 |

### Missing At 100

| Query ID | Missing Doc IDs | Query | Top 10 Doc IDs |
| --- | --- | --- | --- |
| 1 | 31715818 | 0-dimensional biomaterials show inductive properties. | 25657127, 7583104, 10982689, 20454006, 14082855, 40212412, 4346436, 58050905, 13481880, 825728 |
| 128 | 8290953 | Arterioles have a larger lumen diameter than venules. | 12549585, 34876410, 79231308, 3588621, 79447, 22972632, 4890578, 2056197, 9310407, 21380348 |
| 132 | 7975937 | Aspirin inhibits the production of PGE2. | 6923961, 20732789, 3866315, 4583180, 15648443, 2443495, 9159125, 58006489, 39903312, 27615329 |
| 312 | 6173523 | De novo assembly of sequence data has more specific contigs than unassembled sequence data. | 14464451, 1544804, 32770503, 14275671, 20375264, 15893330, 7260461, 40721190, 36082224, 13123189 |
| 421 | 11172205 | Flexible molecules experience greater steric hindrance in the tumor microenviroment than rigid molecules. | 1065627, 17388232, 12866641, 6944800, 10812605, 8702697, 4429118, 40667577, 30919024, 8774475 |
| 437 | 18399038 | Functional consequences of genomic alterations due to Myelodysplastic syndrome (MDS) are poorly understood due to the... | 2359152, 3391547, 5765455, 39851630, 8083310, 1617327, 12240507, 8385277, 5836, 7239105 |
| 502 | 13071728 | Healthcare delivery efficiency in crowded delivery centers is impaired by improving structural, logistical, and inter... | 70516463, 10577574, 17626822, 39059143, 4345757, 31019903, 79231308, 10854174, 11748341, 153755807 |
| 560 | 40096222 | Immune responses result in the development of inflammatory Th17 cells and anti-inflammatory iTregs. | 11233339, 29347970, 14644164, 19005293, 25726838, 1855679, 45096063, 39084565, 16119973, 45449835 |
| 834 | 5483793 | NOX2-independent pathways can generate peroxynitrite by reacting with nitrogen intermediates. | 23100962, 23122306, 22059387, 35828148, 6259170, 13823200, 8536018, 14848619, 1354567, 21297708 |
| 1110 | 13770184 | Suboptimal nutrition is not predictive of chronic disease | 8529693, 23245050, 9244474, 28894097, 25837950, 24273592, 23377475, 43483151, 23013317, 24396137 |
| 1199 | 16760369 | The benefits of colchicine were achieved with effective widespread use of secondary prevention strategies such as hig... | 7454794, 21557614, 4687948, 5698494, 8780599, 14260013, 1287809, 2391552, 11614737, 34268160 |
| 1213 | 14407673 | The deregulated and prolonged activation of monocytes has deleterious effects in inflammatory diseases. | 14386505, 21395936, 9334631, 36444198, 5035827, 39084565, 22621251, 44562221, 5107861, 21498497 |
| 1245 | 7662395 | The one-child policy has been successful in lowering population growth. | 23531592, 167944455, 85693958, 40500723, 3623127, 1624106, 12428497, 30292811, 3001685, 11748341 |
| 1332 | 5304891 | Tumor necrosis factor alpha (TNF-α) and interleukin-1 (IL-1) are pro-inflammatory cytokines that inhibit IL-6 and IL-10. | 2844490, 20722510, 12705056, 44935041, 11233339, 39264456, 36386637, 9334631, 6397191, 22889972 |
| 1362 | 8290953 | Venules have a larger lumen diameter than arterioles. | 34876410, 3588621, 79231308, 7764903, 2056197, 9310407, 22972632, 17741440, 12549585, 496873 |

### Lowest MRR@10

| Query ID | MRR@10 | Recall@100 | First Relevant Rank | Query |
| --- | ---: | ---: | ---: | --- |
| 1 | 0.0000 | 0.0000 | missing | 0-dimensional biomaterials show inductive properties. |
| 1110 | 0.0000 | 0.0000 | missing | Suboptimal nutrition is not predictive of chronic disease |
| 1199 | 0.0000 | 0.0000 | missing | The benefits of colchicine were achieved with effective widespread use of secondary prevention strategies such as hig... |
| 1213 | 0.0000 | 0.0000 | missing | The deregulated and prolonged activation of monocytes has deleterious effects in inflammatory diseases. |
| 1245 | 0.0000 | 0.0000 | missing | The one-child policy has been successful in lowering population growth. |
| 128 | 0.0000 | 0.0000 | missing | Arterioles have a larger lumen diameter than venules. |
| 132 | 0.0000 | 0.0000 | missing | Aspirin inhibits the production of PGE2. |
| 1332 | 0.0000 | 0.0000 | missing | Tumor necrosis factor alpha (TNF-α) and interleukin-1 (IL-1) are pro-inflammatory cytokines that inhibit IL-6 and IL-10. |
| 1362 | 0.0000 | 0.0000 | missing | Venules have a larger lumen diameter than arterioles. |
| 312 | 0.0000 | 0.0000 | missing | De novo assembly of sequence data has more specific contigs than unassembled sequence data. |
