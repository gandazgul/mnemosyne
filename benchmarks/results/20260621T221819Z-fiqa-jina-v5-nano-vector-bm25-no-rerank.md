# Mnemosyne BEIR fiqa Results

- Generated: 2026-06-21T22:18:19Z
- Dataset: `fiqa`
- Queries: 648
- Corpus documents: 57638
- Search limit: 100
- Source mode: `hybrid`
- Fusion: `vector-bm25`
- BM25 weight: `0.2`
- Rerank candidates: 300
- Rerank enabled: False
- Run label: `jina-v5-nano`
- Mnemosyne config: `/Users/gandazgul/Documents/web/mnemosyne/cli/configs/jina-v5-text-nano-retrieval.yaml`

| Metric | Score |
| --- | ---: |
| `ndcg@10` | 0.4558 |
| `mrr@10` | 0.5307 |
| `recall@10` | 0.5267 |
| `recall@100` | 0.7752 |
| `map@100` | 0.3950 |

## Breakdown

| Detail | Count |
| --- | ---: |
| `queries` | 648 |
| `queries_with_rank_1_hit` | 288 |
| `queries_with_top_10_hit` | 466 |
| `queries_with_top_100_hit` | 586 |
| `queries_missing_at_100` | 62 |

### First Relevant Rank

| Bucket | Queries |
| --- | ---: |
| `rank_1` | 288 |
| `rank_2_3` | 93 |
| `rank_4_10` | 85 |
| `rank_11_100` | 120 |
| `missing@100` | 62 |

### Missing At 100

| Query ID | Missing Doc IDs | Query | Top 10 Doc IDs |
| --- | --- | --- | --- |
| 4641 | 319954, 397358 | Where should I park my rainy-day / emergency fund? | 580025, 497993, 527939, 285812, 583695, 32833, 44594, 538023, 406219, 376148 |
| 3451 | 26292, 192307, 588448, 490170 | Should you keep your stocks if you are too late to sell? | 251536, 528518, 420974, 471911, 301757, 310683, 545284, 238215, 41852, 306460 |
| 753 | 243503 | Taxes due for hobbyist Group Buy | 466718, 122185, 203791, 114418, 170632, 127004, 198090, 158738, 132780, 465447 |
| 4465 | 376575 | How to donate to charity that will make a difference? | 174033, 132881, 90591, 427322, 106786, 46381, 266342, 326167, 379353, 174543 |
| 7326 | 584295 | Do brokers execute every trade on the exchange? | 404339, 35340, 163333, 593445, 37040, 272008, 257656, 147573, 227542, 395357 |
| 5951 | 497260 | Why can't house prices be out of tune with salaries | 599860, 418034, 31663, 259777, 196405, 374480, 62702, 589470, 285525, 596834 |
| 10827 | 160786, 7748, 107554, 95282 | How much should I be contributing to my 401k given my employer's contribution? | 290105, 555377, 140330, 296405, 497561, 436930, 15841, 436071, 242556, 452592 |
| 9771 | 263955, 28740 | Is there any emprical research done on 'adding to a loser' | 137534, 130349, 18671, 490479, 38704, 541485, 356175, 375290, 137898, 468717 |
| 3724 | 508921, 279570, 497216, 552887, 199970 | Should you always max out contributions to your 401k? | 302512, 273497, 3104, 122910, 43573, 576391, 497561, 488673, 430931, 135790 |
| 5853 | 476663, 160105, 495699, 424598 | Paying Off Principal of Home vs. Investing In Mutual Fund | 284318, 64456, 182612, 387722, 473647, 473865, 186071, 301224, 494148, 336218 |
| 570 | 363591 | Employer options when setting up 401k for employees | 532839, 117845, 79375, 289064, 301616, 150883, 387876, 15841, 242529, 555377 |
| 594 | 534059 | Should a retail trader bother about reading SEC filings | 377322, 64881, 11148, 97837, 213214, 314898, 548596, 86281, 161411, 38863 |
| 10122 | 273718 | Why diversify stocks/investments? | 144261, 297100, 259084, 89084, 459970, 508540, 180855, 331008, 139368, 502495 |
| 1783 | 332314 | Freelancing Tax implication | 421924, 14609, 156063, 159709, 445298, 179359, 383172, 270426, 423625, 588211 |
| 6004 | 149555 | Put-Call parity - what is the difference between the two representations? | 374797, 122432, 13260, 216065, 278373, 345410, 247738, 232261, 118762, 37517 |
| 4837 | 531841, 20958 | When applying for a mortgage, can it also cover outstanding debts? | 97925, 245005, 204035, 32749, 516578, 225874, 294167, 372993, 249054, 257644 |
| 4415 | 67676 | How much is inflation? | 117578, 206580, 290585, 381675, 519596, 501743, 513249, 468089, 553634, 267176 |
| 8378 | 125298 | Should I wait a few days to sell ESPP Stock? | 511678, 133644, 546125, 495568, 434812, 294573, 178684, 584291, 361345, 387035 |
| 8102 | 552707, 378173, 90294 | When do I sell a stock that I hold as a long-term position? | 306460, 165970, 35752, 537212, 171819, 203638, 557877, 310683, 496209, 292045 |
| 7928 | 118633 | If I believe a stock is going to fall, what options do I have to invest on this? | 427808, 410404, 260384, 501504, 67415, 480967, 47053, 294688, 171784, 221869 |
| 4777 | 590710 | How to finance necessary repairs to our home in order to sell it? | 52351, 365342, 416382, 478413, 570318, 171631, 323449, 426678, 360872, 378384 |
| 5741 | 25943 | Learning investing and the stock market | 64168, 47973, 379546, 367845, 85558, 241423, 167088, 124219, 341148, 259081 |
| 620 | 331332 | Is it wise to have plenty of current accounts in different banks? | 543921, 535817, 467059, 164801, 572848, 489959, 380786, 61734, 104492, 48346 |
| 8271 | 415511 | Income in zero-interest environment | 376709, 249558, 380368, 52047, 119298, 137225, 136262, 65179, 83330, 283862 |
| 864 | 211364, 152072 | Why use accounting software like Quickbooks instead of Excel spreadsheets? | 30142, 472924, 24890, 566337, 329774, 157751, 2436, 78117, 222380, 402174 |
| 5888 | 540806 | Interest charges on balance transfer when purchases are involved | 263647, 125497, 336792, 543776, 579601, 213242, 358445, 545327, 490529, 429746 |
| 2568 | 388798, 127353 | How to pay with cash when car shopping? | 346042, 258247, 166314, 15696, 9146, 355310, 108739, 514238, 60261, 453301 |
| 7124 | 74615 | How come we can find stocks with a Price-to-Book ratio less than 1? | 558617, 526110, 533818, 583708, 278582, 504243, 154725, 251100, 157597, 226070 |
| 2857 | 295864 | I have around 60K $. Thinking about investing in Oil, how to proceed? | 127566, 233732, 501384, 474575, 117451, 542051, 591705, 316444, 379140, 58065 |
| 8539 | 218728, 196304, 396038 | Can the risk of investing in an asset be different for different investors? | 283074, 483123, 391861, 502495, 385881, 471817, 500863, 499166, 378403, 150475 |
| 6896 | 251704 | Selling high, pay capital gains, re-purchase later | 468047, 448659, 343219, 522319, 169240, 66834, 366560, 257625, 216827, 581780 |
| 1085 | 467737, 393710 | How do disputed debts work on credit reports? | 372039, 450031, 242013, 161422, 78328, 319276, 268777, 574122, 339365, 398258 |
| 5422 | 151973 | What are some good books for learning stocks, bonds, derivatives e.t.c for beginner with a math background? | 165294, 193555, 273906, 276786, 172587, 552371, 191688, 241423, 221319, 79517 |
| 8247 | 42521, 321114, 465313 | Tax on Stocks or ETF's | 586010, 437907, 580802, 518735, 153112, 161019, 195767, 474745, 270992, 528880 |
| 10482 | 549072 | Rollover into bond fund to do dollar cost averaging [duplicate] | 330023, 447567, 224782, 439757, 525089, 211765, 134005, 265817, 237052, 564787 |
| 8789 | 70853 | What does “profits to the shareholders jumped to 15 cents a share” mean? | 87349, 41912, 341424, 234040, 20076, 14870, 219927, 317363, 339854, 573079 |
| 2423 | 538023 | At what age should I start or stop saving money? | 529444, 417787, 396127, 234846, 337561, 235855, 272328, 553288, 104457, 468010 |
| 4523 | 594257, 119165 | What should I do with my $25k to invest as a 20 years old? | 129255, 272070, 442776, 465819, 216365, 171712, 10476, 332022, 286746, 72578 |
| 8513 | 270573 | Buy on dip when earnings fail? | 572622, 175821, 203873, 462135, 573767, 351396, 391043, 335626, 321205, 73872 |
| 5054 | 28119 | How to stress test an investment plan? | 588481, 564007, 377186, 101124, 582899, 448745, 220665, 127263, 458183, 114092 |
| 1159 | 496064 | what is the best way to do a freelancing job over the summer for a student | 55064, 347181, 460648, 156063, 163881, 112669, 132287, 449155, 439899, 271812 |
| 68 | 19183 | Intentions of Deductible Amount for Small Business | 519473, 97233, 86134, 192516, 354716, 97719, 381151, 462831, 447231, 545296 |
| 9701 | 387141, 357739 | How to bet against the London housing market? | 473883, 408865, 225682, 73283, 258048, 412502, 410404, 516214, 253359, 108399 |
| 6199 | 239214 | How can all these countries owe so much money?  Why & where did they borrow it from? | 414693, 584273, 47163, 10399, 49602, 169921, 490042, 380714, 298794, 351853 |
| 9126 | 514831 | Short an option - random assignment? | 334473, 477588, 82194, 228810, 166227, 307518, 102316, 308859, 590453, 292045 |
| 34 | 599545 | 401k Transfer After Business Closure | 458917, 15728, 144109, 424766, 551545, 490867, 483268, 226547, 387876, 154839 |
| 6395 | 166227 | Option settlement for calendar spreads | 584223, 467463, 276314, 22916, 8177, 401447, 505223, 516790, 111301, 273142 |
| 9115 | 207325 | Why does the calculation for percentage profit vary based on whether a position is short vs. long? | 419897, 158520, 422467, 245082, 331606, 314478, 428786, 154665, 232880, 410822 |
| 4767 | 280805, 568670, 224057, 22804 | New car: buy with cash or 0% financing | 420018, 584106, 429153, 306834, 166314, 56867, 451092, 256803, 9146, 59372 |
| 8351 | 472516 | What happens when a calendar spread is assigned in a non-margin account? | 102316, 141213, 45674, 273142, 527654, 596567, 23609, 228810, 272754, 516790 |
| 6131 | 381720, 170204, 416679, 2460 | Is it ever a good idea to close credit cards? | 326094, 368806, 391384, 334111, 339030, 218088, 99449, 258465, 35625, 201982 |
| 858 | 45185, 122485, 278450 | Is it bad practice to invest in stocks that fluctuate by single points throughout the day? | 293027, 567608, 214281, 146632, 30774, 208932, 573612, 139699, 433730, 55751 |
| 4019 | 6881 | How and Should I Invest (As a college 18 year old with minimal living expenses)? | 85977, 332938, 332749, 20304, 287991, 55841, 379948, 493034, 66864, 426461 |
| 6080 | 164513 | Is ScholarShare a legitimate entity for a 529 plan in California? | 22856, 201500, 236732, 277581, 468527, 83080, 2809, 233401, 115175, 535357 |
| 6959 | 205010 | What is the term for the quantity (high price minus low price) for a stock? | 428117, 468025, 303325, 169954, 577573, 229573, 373034, 412223, 304399, 599523 |
| 5402 | 491350 | Is it impossible to get a home loan with a poor credit history after a divorce? | 445163, 90579, 51728, 595029, 227485, 440063, 52250, 44105, 180214, 310790 |
| 6562 | 501157 | Cheapest way to “wire” money in an Australian bank account to a person in England, while I'm in Laos? | 473605, 60446, 282744, 582414, 183880, 182443, 549684, 308837, 462050, 385182 |
| 701 | 389446 | What are the ins/outs of writing-off part of one's rent for working at home? | 349672, 231990, 436505, 344955, 337706, 456234, 124507, 243306, 339488, 87113 |
| 7674 | 519390 | Choosing the limit when making a limit order? | 447886, 249279, 15917, 526235, 514841, 200666, 155151, 94653, 278630, 31933 |
| 5940 | 486243, 93936 | How does investment into a private company work? | 512609, 250354, 182226, 252853, 454465, 46842, 473154, 535314, 135411, 530 |
| 6612 | 205522, 322900 | If I have a lot of debt and the housing market is rising, should I rent and slowly pay off my debt or buy and roll th... | 502594, 301192, 431481, 180192, 254454, 373554, 343917, 317945, 14083, 433171 |
| 4714 | 450819 | Personal finance app where I can mark transactions as “reviewed”? | 505057, 584450, 218793, 353915, 344473, 529790, 29812, 479390, 65957, 374518 |
| 8456 | 486333 | What typically happens to unvested stock during an acquisition? | 257853, 93215, 534755, 555276, 469036, 186869, 492428, 243886, 104188, 174321 |
| 10213 | 270221, 545712 | Looking for good investment vehicle for seasonal work and savings | 195373, 38269, 446186, 100517, 272198, 112499, 263390, 386305, 96949, 80844 |
| 5196 | 172128, 114829 | I might use a credit card convenience check. What should I consider? | 565745, 402543, 85517, 85252, 2875, 289483, 456098, 393866, 481052, 581976 |
| 3006 | 269851, 568473, 328300 | Strategies for putting away money for a child's future (college, etc.)? | 512096, 127838, 258704, 332749, 8266, 372900, 290441, 303432, 538282, 211713 |
| 3909 | 312248, 404356, 193459, 245616, 353028 | How to rescue my money from negative interest? | 514003, 83330, 472837, 61586, 328499, 574011, 362730, 42475, 374400, 184838 |
| 5464 | 350399, 86691 | Resources on Buying Rental Properties | 222095, 26339, 423438, 372274, 325722, 383921, 545341, 315972, 426705, 536126 |
| 2385 | 407654 | As director, can I invoice my self-owned company? | 373059, 496064, 210889, 572690, 55440, 217472, 247760, 519321, 564453, 142645 |
| 10034 | 480749 | Tax implications of holding EWU (or other such UK ETFs) as a US citizen? | 181942, 528880, 44955, 197478, 565296, 447197, 180146, 141585, 430868, 85926 |
| 5090 | 436493 | Should I take a student loan to pursue my undergraduate studies in France? | 12988, 246286, 21913, 92430, 560681, 217831, 58005, 287507, 586289, 575421 |
| 2088 | 399875 | How would I go about selling the stock of a privately held company? | 53993, 455168, 291886, 599524, 293687, 238215, 188776, 530, 413672, 140835 |
| 9391 | 503637 | Should I replace bonds in a passive investment strategy | 535518, 248158, 107424, 136515, 577832, 545760, 171669, 155242, 494653, 283202 |
| 3148 | 178127, 438000 | Can a car company refuse to give me a copy of my contract or balance details? | 172855, 584305, 430100, 29721, 357280, 358631, 395995, 65046, 92888, 164702 |
| 4678 | 305153 | Finance, Cash or Lease? | 185405, 215225, 311748, 504918, 376016, 260095, 311446, 487678, 427884, 522532 |
| 2398 | 509391, 363810, 224654, 590489 | Frustrated Landlord | 556453, 487094, 44058, 96538, 98372, 201705, 393883, 395770, 309231, 436875 |
| 6746 | 210887 | What happens if stock purchased on margin plummets below what I have in the brokerage? | 279185, 247680, 333674, 231221, 115918, 283982, 251704, 527654, 176822, 399903 |
| 5511 | 169893, 560325, 478426, 383193, 278699, 12746 | Pay off car loan entirely or leave $1 until the end of the loan period? | 376016, 334559, 38786, 329137, 51873, 179891, 529123, 107898, 139788, 155843 |
| 8834 | 12232 | Pros/Cons of Buying Discounted Company Stock | 203139, 599156, 528827, 133644, 553331, 57387, 569303, 163396, 569224, 521095 |
| 988 | 226053, 107688 | Where should I invest my savings? | 168402, 60093, 501384, 347651, 285812, 450558, 571218, 211767, 388252, 223551 |
| 3369 | 163834, 231012, 145716, 411910, 395840 | Why should one only contribute up to the employer's match in a 401(k)? | 341493, 555377, 296405, 436071, 242556, 15841, 24231, 463892, 240373, 92370 |
| 9296 | 435746 | Why would Two ETFs tracking Identical Indexes Produce different Returns? | 148721, 206744, 159471, 368124, 408524, 428187, 285135, 492212, 410123, 209996 |
| 9245 | 194561 | Stock Options for a company bought out in cash and stock | 207253, 186869, 39345, 131488, 265111, 178497, 261487, 259560, 248393, 409818 |
| 3490 | 420529 | Tax Witholding for Stock Sale | 447651, 591157, 361482, 537371, 311782, 400730, 367742, 152960, 407602, 571124 |
| 5763 | 462019 | What is the best way to get a “rough” home appraisal prior to starting the refinance process? | 570318, 38712, 251466, 67379, 563380, 89964, 326214, 218144, 331255, 215647 |
| 4962 | 599925 | Net Cash Flows from Selling the Bond and Investing | 416839, 158363, 408661, 34949, 52149, 537603, 308276, 393838, 431386, 187110 |
| 4846 | 151104 | Is there anything comparable to/resembling CNN's Fear and Greed Index? | 98096, 3533, 335892, 538974, 415161, 183597, 320059, 270305, 317666, 489352 |
| 9403 | 6666, 328086, 345199 | Abundance of Cash - What should I do? | 410450, 159235, 570632, 14349, 551986, 273978, 70806, 499548, 215296, 420574 |
| 5993 | 367375, 272866, 230215, 55084, 5827, 352638, 426120, 63501 | Why would anyone want to pay off their debts in a way other than “highest interest” first? | 94373, 416796, 160193, 287571, 494306, 353911, 128574, 886, 156195, 431212 |
| 5710 | 232311 | Bucketing investments to track individual growths | 227364, 516267, 411856, 88417, 135765, 412830, 534323, 508610, 227733, 177036 |
| 7529 | 66607 | Does the expense ratio of a fund-of-funds include the expense ratios of its holdings? | 89297, 464337, 514529, 59249, 102904, 293626, 218261, 287537, 361013, 65587 |
| 5021 | 589285 | Is there a more flexible stock chart service, e.g. permitting choice of colours when comparing multiple stocks? | 528576, 189341, 584801, 555506, 494939, 211444, 465971, 511861, 60284, 252084 |
| 3612 | 259625 | How can I buy and sell the same stock on the same day? | 522658, 567383, 390864, 310636, 402726, 429418, 165548, 292159, 367873, 584291 |
| 4409 | 499128, 100306, 147439 | My friend wants to put my name down for a house he's buying. What risks would I be taking? | 243732, 223841, 268078, 102088, 514790, 102326, 60135, 341947, 84732, 360682 |
| 2070 | 363678 | Advantage of credit union or local community bank over larger nationwide banks such as BOA, Chase, etc.? | 550303, 587737, 469515, 578357, 30253, 597571, 590209, 249839, 38038, 408166 |
| 11039 | 53544, 249063 | Pay off credit card debt or earn employer 401(k) match? | 91183, 552383, 287876, 5203, 163287, 79363, 508534, 345895, 105557, 437706 |
| 5460 | 184337, 21174, 108514, 463885 | Paying off a loan with a loan to get a better interest rate | 77052, 344812, 343208, 106495, 470716, 327115, 529418, 243065, 576609, 555280 |
| 7925 | 318185, 402482 | Can I sell a stock immediately? | 591436, 227399, 332467, 310636, 581866, 44461, 315760, 81721, 221869, 438974 |
| 4286 | 566069 | Given advice “buy term insurance and invest the rest”, how should one “invest the rest”? | 70460, 10531, 229239, 151817, 391243, 155640, 564675, 206830, 511386, 71926 |
| 2685 | 154113, 370300, 468923 | What ways are there for us to earn a little extra side money? | 382005, 576047, 594182, 468086, 109880, 269380, 237950, 280099, 77088, 431751 |
| 1090 | 518896 | Need a formula to determine monthly payments received at time t if I'm reinvesting my returns | 393987, 446454, 16051, 19999, 179365, 520217, 296146, 573928, 281329, 209238 |
| 6122 | 44344 | Better to rent condo to daughter or put her on title? | 496166, 316794, 403515, 53840, 558251, 182039, 80269, 577658, 118246, 566184 |
| 4514 | 69485, 337764, 209804 | What intrinsic, non-monetary value does gold have as a commodity? | 156211, 471825, 426270, 146573, 317429, 240894, 408336, 99089, 80141, 277694 |
| 8507 | 370995 | When to sell a stock? | 99132, 251536, 303724, 545284, 272091, 236415, 273565, 217837, 102237, 88813 |
| 6221 | 257248, 519675, 76414, 169688, 455614, 115717 | To pay off a student loan, should I save up a lump sum payoff payment or pay extra each month? | 254245, 352363, 448791, 529551, 110081, 124705, 274108, 188384, 394474, 192602 |
| 3008 | 180192, 323406 | What are my chances at getting a mortgage with Terrible credit but High income | 102266, 285694, 407401, 231688, 47441, 574438, 2064, 78176, 251846, 455952 |
| 4007 | 521657 | What is a reasonable salary for the owner and sole member of a small S-Corp? | 556220, 260385, 370542, 170933, 205341, 334603, 388704, 543085, 508078, 315552 |
| 6644 | 175035 | How to know precisely when a SWIFT is issued by a bank? | 475527, 110198, 118396, 218761, 554518, 355870, 271596, 298587, 327623, 41383 |
| 10267 | 460398 | How should I prepare for the next financial crisis? | 178693, 143393, 569632, 305600, 36961, 326398, 369470, 182442, 96017, 87520 |
| 7622 | 253369, 378594 | Best way to pay off debt? | 220241, 457945, 373554, 480773, 353911, 345895, 388095, 157923, 416796, 508952 |
| 3767 | 153922, 392060 | What should I be doing to protect myself from identity theft? | 90632, 423809, 260580, 97686, 91986, 581889, 171510, 158285, 158008, 587778 |
| 6410 | 471723 | Will an ETF immediately reflect a reconstitution of underlying index | 454610, 71230, 214281, 330729, 87238, 295993, 200360, 87261, 313897, 47276 |
| 5030 | 215540 | Why pay for end-of-day historical prices? | 227192, 532178, 13511, 149420, 471131, 560108, 295344, 352415, 113150, 370569 |
| 6252 | 394551, 160932, 293624, 233294, 243268, 379487, 62868 | Is this mortgage advice good, or is it hooey? | 213713, 473647, 120061, 495089, 47565, 205906, 139366, 27268, 443852, 423403 |
| 885 | 337165, 409184 | How long do credit cards keep working after you disappear? | 254968, 472336, 516678, 99449, 251701, 181757, 251643, 89888, 434082, 596284 |
| 766 | 550172 | Will the ex-homeowner still owe money after a foreclosure? | 2996, 299591, 163711, 212827, 552768, 427110, 333583, 104955, 268865, 578906 |
| 8202 | 513258, 93971 | What accounted for DXJR's huge drop in stock price? | 457689, 317363, 337001, 71924, 122542, 537862, 412584, 421248, 355167, 67237 |
| 7345 | 237645 | What do these numbers mean? (futures) | 276381, 9274, 206895, 460331, 527080, 127845, 508821, 354429, 273789, 108 |
| 776 | 583640, 127263, 496899 | Can saving/investing 15% of your income starting age 25, likely make you a millionaire? | 124027, 10440, 417787, 143591, 41960, 467044, 374266, 434972, 418281, 554833 |
| 89 | 248624 | How can I deposit a check made out to my business into my personal account? | 508754, 309023, 135196, 526817, 308938, 188893, 80538, 521540, 590102, 29372 |
| 1920 | 269943 | Clarification on student expenses - To file the tax for the next year | 263485, 481114, 585356, 295562, 316482, 551187, 128980, 446889, 585023, 150271 |
| 8013 | 496159, 224231 | Frequency of investments to maximise returns (and minimise fees) | 384983, 537626, 81652, 57033, 388389, 40652, 224816, 385955, 28291, 450662 |
| 3759 | 527966, 67167, 522358 | Simplifying money management | 455457, 373772, 214248, 490065, 122378, 248663, 145812, 372743, 231412, 526159 |
| 10639 | 431799, 495774, 278453, 187039 | Short term parking of a large inheritance? | 163353, 171196, 235628, 111048, 590276, 131391, 318864, 178386, 546538, 163197 |
| 6635 | 156358 | Why don't share prices of a company rise every other Friday when the company buys shares for its own employees? | 587137, 3656, 533712, 95806, 491064, 12560, 235531, 343452, 16195, 579037 |
| 4312 | 399149 | Is it true that 90% of investors lose their money? | 282435, 222639, 167950, 285945, 497786, 116647, 532485, 170628, 431735, 300770 |
| 6525 | 181985 | Does it make sense to trade my GOOGL shares for GOOG and pocket the difference? | 106541, 550661, 98150, 156467, 362473, 105542, 498014, 378906, 147002, 53263 |
| 2590 | 589625 | Are non-residents or foreigners permitted to buy or own shares of UK companies? | 296528, 209493, 48269, 158923, 310103, 307776, 55999, 456999, 458730, 262485 |
| 5374 | 152688 | What were the main causes of the spike and drop of DRYS's stock price? | 283106, 133204, 457689, 122542, 467594, 362462, 261522, 317363, 50141, 78648 |
| 2994 | 419319 | Work on the side for my wife's company | 510317, 569145, 491844, 382005, 5840, 321743, 460721, 547793, 138428, 399882 |
| 3683 | 185909, 454501 | Can I trust the Motley Fool? | 276975, 105973, 408995, 428848, 500338, 538086, 6607, 301739, 192912, 565016 |
| 7206 | 441155, 532211, 553066 | Who Bought A Large Number Of Shares? | 34882, 351570, 573846, 65667, 358164, 558703, 350214, 444752, 552375, 327525 |
| 10246 | 512984, 77573 | Understanding the T + 3 settlement days rule | 370635, 156029, 176717, 327080, 340263, 28314, 179520, 332243, 11927, 226984 |
| 5241 | 322157, 27489 | Mortgage vs. Cash for U.S. home buy now | 344740, 281675, 213713, 111184, 438073, 390976, 266649, 42604, 426211, 309420 |
| 98 | 575929 | How can I make $250,000.00 from trading/investing/business within 5 years? | 527522, 102113, 66034, 555630, 373119, 336661, 519619, 506149, 209067, 121161 |
| 4615 | 262934 | Are solar cell panels and wind mills worth the money? | 261900, 69523, 496427, 455798, 425595, 271015, 510872, 120384, 376430, 158216 |
| 6467 | 453256, 23217, 346641, 367313 | Advice on strategy for when to sell | 88813, 240089, 217837, 109455, 203873, 368348, 99857, 130941, 83807, 504235 |
| 4289 | 24881 | Does the currency exchange rate contain any additional information at all? | 288330, 17469, 517345, 356465, 119316, 114886, 324546, 135220, 226102, 416975 |
| 4394 | 336045, 441582 | Transfer $50k to another person's account (in California, USA) | 322838, 93386, 462585, 293653, 305907, 431462, 521753, 415655, 412258, 495827 |
| 7344 | 108403 | How is the Dow divisor calculated? | 14368, 159166, 150430, 378974, 253926, 313421, 65618, 501032, 591089, 69655 |
| 6875 | 224392 | Where to find free Thailand stock recommendations and research? | 567500, 110733, 352557, 224366, 284539, 9354, 556770, 232460, 79337, 77502 |
| 10447 | 152096, 300721 | Is there an advantage to a traditional but non-deductable IRA over a taxable account? [duplicate] | 144751, 500175, 447482, 382236, 532657, 259150, 382894, 299690, 406239, 540389 |
| 5782 | 595455 | Pay off credit cards in one lump sum, or spread over a few months? | 487621, 172084, 262026, 117007, 114592, 529312, 273631, 440620, 124705, 261697 |
| 9871 | 448890, 40051, 170594 | What should I do with the 50k I have sitting in a European bank? | 367207, 73741, 293179, 76562, 175139, 292714, 433003, 212464, 74668, 219477 |
| 3625 | 414295 | What should I do with my paper financial documents? | 509617, 500751, 569812, 380263, 163168, 513248, 113830, 37582, 123366, 44204 |
| 6005 | 135415, 478457, 345895, 73310, 384626, 390689 | Why might it be advisable to keep student debt vs. paying it off quickly? | 149500, 571198, 507544, 25190, 431884, 96268, 564206, 572272, 52136, 414288 |
| 3822 | 385090, 308837 | How to change a large quantity of U.S. dollars into Euros? | 292714, 194730, 417917, 340777, 19618, 390524, 239876, 79777, 549787, 174406 |
| 7879 | 372551, 421285 | Any Tips on How to Get the Highest Returns Within 4 Months by Investing in Stocks? | 58186, 102029, 272174, 43088, 540919, 228488, 174313, 7625, 105391, 593879 |
| 3115 | 234950, 389028, 316794 | How can I live outside of the rat race of American life with 300k? | 233562, 183869, 129364, 267892, 174272, 136035, 475736, 369742, 252852, 150066 |
| 3995 | 278734, 230208 | I have more than $250,000 in a US Bank account… mistake? | 485507, 404954, 171720, 479918, 264934, 14349, 352883, 303367, 146557, 506909 |
| 8002 | 34767 | What is the tax treatment of scrip dividends in the UK? | 118786, 217006, 32600, 162454, 110983, 97842, 267067, 263312, 447197, 265159 |
| 10136 | 526115 | How to minimise the risk of a reduction in purchase power in case of Brexit for money held in a bank account? | 466950, 417740, 290930, 304007, 205685, 583903, 284305, 35511, 182195, 150543 |
| 8635 | 67107, 240215 | Is there any flaw in this investment scheme? | 493841, 46818, 151238, 510565, 447619, 203729, 303619, 480802, 365816, 202638 |
| 5206 | 563030, 28230, 117276, 300660 | Is it a good idea to get an unsecured loan to pay off a credit card that won't lower a high rate? | 298908, 595455, 287571, 516397, 225522, 69938, 153088, 340520, 504293, 2519 |
| 2713 | 388147 | Physical Checks - Mailing | 284528, 29372, 41944, 20791, 216200, 199069, 78139, 402898, 350237, 268257 |
| 9060 | 40447 | Buying puts without owning underlying | 511093, 181924, 528052, 338782, 228217, 7743, 345851, 359778, 3062, 316037 |
| 4105 | 25096 | As an investor what are side effects of Quantitative Easing in US and in EU? | 416483, 176262, 345910, 305029, 393791, 108519, 369038, 30946, 293104, 339640 |
| 2465 | 570680, 81046 | Can capital expenses for volunteer purposes be deducted from income? | 37382, 202645, 146657, 432545, 510716, 598646, 275543, 490176, 541809, 398536 |
| 4640 | 101369 | What can my relatives do to minimize their out of pocket expenses on their fathers estate | 295246, 372808, 331534, 356035, 17110, 522619, 367404, 521803, 338539, 45819 |
| 9275 | 338754 | Do I have to pay a capital gains tax if I rebuy the same stock within 30 days? | 400730, 23217, 537916, 343219, 161155, 407602, 376800, 390864, 596518, 292762 |
| 3500 | 174019 | Why invest in becoming a landlord? | 273187, 528206, 71424, 557478, 572061, 578597, 11601, 141935, 41356, 422331 |
| 1306 | 484437, 204075 | I made an investment with a company that contacted me, was it safe? | 594206, 160611, 309851, 333050, 519038, 309590, 537698, 538086, 316321, 167684 |
| 6262 | 26799 | Help required on estimating SSA benefit amounts | 34538, 118707, 2648, 498444, 83338, 529927, 320362, 430407, 390877, 15322 |
| 8632 | 213976 | Is it best to exercise options shares when they vest, or wait | 43497, 237718, 104188, 237783, 259560, 382381, 293959, 255927, 420722, 163396 |
| 6133 | 415705 | What happens to all of the options when they expire? | 7733, 575408, 11456, 581672, 242298, 132288, 358492, 177559, 428399, 116436 |
| 3771 | 488948, 198349, 217683, 49601 | Best way to buy Japanese yen for travel? | 490384, 495826, 96211, 521712, 306130, 217715, 434201, 350245, 120604, 575495 |
| 1736 | 25543, 443419 | How can people have such high credit card debts? | 399406, 569056, 562896, 372993, 437610, 174941, 517050, 298908, 475668, 99463 |
| 6814 | 340214, 223206 | Selling Stock - All or Nothing? | 513734, 590188, 400614, 67107, 66834, 154976, 279782, 276883, 337736, 178497 |
| 1322 | 399418, 64138 | Is this follow-up after a car crash a potential scam? | 226090, 44635, 114231, 332916, 567973, 397852, 283917, 524723, 98356, 91463 |
| 5185 | 210236, 317354 | Invest in low cost small cap index funds when saving towards retirement? | 196992, 376485, 241202, 580313, 262180, 524525, 503725, 268731, 106620, 523331 |
| 6909 | 127012 | Why do stocks priced above $2.00 on the ASX sometimes move in $0.005 increments? | 72633, 118232, 490584, 375019, 112946, 168080, 376399, 47217, 64943, 102026 |
| 2348 | 211867, 566573, 211622, 474234, 352271, 265874 | Why can't you just have someone invest for you and split the profits (and losses) with him? | 447619, 306430, 247486, 420544, 389004, 151412, 381757, 177194, 64410, 309851 |
| 4499 | 76996 | Is investing exlusively in a small-cap index fund a wise investment? | 52274, 196992, 501153, 228111, 517391, 235917, 335136, 372233, 312591, 376485 |
| 42 | 272709 | What are the ins/outs of writing equipment purchases off as business expenses in a home based business? | 305222, 510863, 581265, 88967, 28764, 571711, 398536, 443859, 35379, 456234 |
| 3530 | 239998 | How to exclude stock from mutual fund | 184299, 24029, 370754, 378075, 209879, 479420, 110343, 449124, 530938, 322070 |
| 9108 | 272021, 472585 | Starting an investment portfolio with Rs 5,000/- | 290757, 46967, 336218, 171189, 240351, 414116, 7748, 312821, 323067, 51848 |
| 6835 | 102243 | Are bond ETF capital gains taxed similar to stock or stock funds if held for more than 1 year? | 149305, 5710, 570546, 84238, 287950, 153112, 110343, 586010, 29502, 133093 |
| 3067 | 406156 | Should I make extra payments to my under water mortgage or increase my savings? | 477907, 476068, 3092, 560915, 131365, 423403, 90009, 468831, 83543, 336276 |
| 4125 | 344648, 72046 | Alternative means of salary for my employees | 174787, 73999, 365558, 127347, 361954, 58906, 562211, 355244, 442014, 462255 |
| 1150 | 43603, 19936 | How are the best way to make and save money at 22 years old | 10476, 529444, 433986, 353369, 319760, 595287, 216365, 494815, 38269, 45353 |
| 7705 | 195191 | Why would I pick a specific ETF over an equivalent Mutual Fund? | 500486, 161019, 539263, 312015, 358997, 153112, 377429, 364735, 270992, 580802 |
| 9808 | 40702, 431946 | Selling To Close | 416307, 557582, 345368, 151587, 362473, 374204, 218423, 107045, 43087, 152719 |
| 3888 | 319213, 239632 | Why I can't view my debit card pre-authorized amounts? | 208169, 316652, 432077, 276733, 185434, 294077, 440527, 181757, 281129, 418580 |
| 10109 | 506374, 499849 | Why does Charles Schwab have a Mandatory Settlement Period after selling stocks? | 28314, 93231, 332243, 266725, 370635, 563826, 98302, 570248, 121465, 537212 |
| 2790 | 279329, 469125 | Should I pay more than 20% down on a home? | 472484, 357200, 400896, 75961, 385343, 207564, 64400, 215103, 27073, 487593 |
| 9882 | 65702 | Money-market or cash-type ETFs for foreigners with U.S brokerage account | 391876, 131059, 389581, 313775, 363378, 188524, 535340, 179527, 451729, 349417 |
| 4999 | 314898 | Looking for a good source for Financial Statements | 9938, 171964, 431459, 597241, 11263, 465971, 295738, 146076, 520165, 343964 |
| 3189 | 225395 | Diversify my retirement investments with a Roth IRA | 122222, 287225, 240975, 404800, 187124, 347651, 423658, 88311, 18436, 246109 |
| 5134 | 158523 | Why does Yahoo Finance's data for a Vanguard fund's dividend per share not match the info from Vanguard? | 46774, 532616, 206727, 405474, 465536, 584128, 215486, 239137, 559884, 54225 |
| 1321 | 216456, 292065 | Are social media accounts (e.g. YouTube, Twitter, Instagram, etc.) considered assets? | 392379, 381322, 203715, 235082, 210019, 481036, 559738, 426184, 229640, 521823 |
| 4539 | 370879 | How should I save money if the real interest rate (after inflation) is negative? | 42475, 472837, 203926, 275925, 32744, 449745, 61586, 220720, 503394, 322816 |
| 715 | 579763, 546538, 187404 | what would you do with $100K saving? | 548758, 427032, 337561, 133120, 113885, 203201, 387647, 333784, 508343, 162633 |
| 504 | 344203, 498751 | Have plenty of cash flow but bad credit | 22807, 93573, 569240, 70806, 41875, 546097, 495431, 68431, 368247, 252473 |
| 2296 | 83330, 366594, 253563 | How does a bank make money on an interest free secured loan? | 400009, 396853, 119298, 259919, 94230, 279897, 175824, 546874, 580147, 106424 |
| 10975 | 61022 | How to contribute to Roth IRA when income is at the maximum limit & you have employer-sponsored 401k plans? | 163865, 81148, 441632, 140330, 360533, 101490, 447482, 446226, 110114, 53028 |
| 1994 | 156640 | Does the IRS reprieve those who have to commute for work? | 231990, 434846, 192843, 263259, 243356, 63919, 51491, 380635, 544381, 22819 |
| 9164 | 365298, 263390 | Bonds vs equities: crash theory | 115648, 309326, 149900, 287656, 321941, 506743, 16924, 599420, 296516, 305274 |
| 10462 | 8266, 11378, 35680, 437879, 204035, 581204 | Is it okay to be married, 30 years old and have no retirement? | 268023, 66376, 151774, 160105, 152478, 122333, 361049, 15322, 391583, 519856 |
| 8855 | 208165, 312821 | How do i get into investing stocks [duplicate] | 155677, 312445, 403092, 560395, 555521, 367415, 142320, 152286, 67327, 218285 |
| 7071 | 124230 | ESPP strategy - Sell right away or hold? | 133644, 294573, 511678, 127702, 71713, 361345, 35575, 434812, 575213, 387035 |
| 8974 | 523331, 356595, 170625 | As a 22-year-old, how risky should I be with my 401(k) investments? | 216365, 10476, 140738, 134931, 102501, 452126, 336144, 124762, 199237, 246253 |
| 5178 | 240261 | Formula that predicts whether one is better off investing or paying down debt | 39819, 557506, 396889, 262772, 111815, 393833, 290434, 373554, 257016, 473865 |
| 5061 | 23747 | What fiscal scrutiny can be expected from IRS in early retirement? | 502150, 151263, 25481, 398078, 96720, 513392, 149742, 517903, 31699, 598378 |
| 2075 | 170042, 359580, 14967 | Are stories of turning a few thousands into millions by trading stocks real? | 65667, 519619, 44417, 33357, 506149, 285147, 519501, 555521, 188129, 523393 |
| 4335 | 357013 | What is the US Fair Tax? | 318583, 419466, 322246, 363178, 585212, 589161, 484148, 412877, 3181, 479451 |
| 7533 | 93853 | Investing tax (savings) | 8012, 142658, 563986, 516756, 162668, 480827, 526664, 250603, 550468, 284411 |
| 1393 | 352838, 539133 | Which is better when working as a contractor, 1099 or incorporating? | 220022, 586026, 234436, 578196, 32072, 68524, 352640, 532932, 277812, 93638 |
| 9733 | 110163, 38655 | Due Diligence - Dilution? | 301880, 316321, 121262, 23414, 526073, 135798, 108965, 267266, 450132, 154841 |
| 7311 | 323768 | Finance, Social Capital IPOA.U | 507841, 199746, 479752, 290325, 583646, 419735, 264740, 261231, 327911, 579110 |
| 744 | 566480, 78176 | What options are available for a home loan with poor credit but a good rental history? | 573276, 310790, 67066, 415425, 490443, 596272, 80607, 289450, 92397, 313623 |
| 7141 | 132288 | Do investors go long option contracts when they cannot cover the exercise of the options? | 538054, 570046, 243714, 507828, 288289, 383328, 305676, 255927, 388571, 44530 |
| 4071 | 129875 | If our economy crashes, and cash is worthless, should i buy gold or silver | 524142, 473965, 291862, 505136, 566669, 487817, 502634, 308332, 53538, 506780 |
| 7512 | 191060 | understanding the process/payment of short sale dividends | 222320, 487329, 115553, 480949, 568166, 202985, 298284, 13631, 409432, 259450 |
| 1391 | 562176 | How is taxation for youtube/twitch etc monetization handled in the UK? | 267067, 254151, 510599, 527951, 223170, 266229, 77171, 454208, 440745, 21136 |
| 7534 | 358125 | Can you explain why it's better to invest now rather than waiting for the market to dip? | 175821, 426157, 145539, 310218, 89714, 114806, 221869, 419747, 94302, 474006 |
| 5356 | 312405 | Historical stock prices: Where to find free / low cost data for offline analysis? | 535343, 560108, 279785, 240086, 529877, 546379, 596106, 537111, 47798, 226749 |
| 2579 | 432020 | What to do when a job offer is made but with a salary less than what was asked for? | 423070, 200946, 181213, 524471, 432808, 256802, 157919, 364159, 559900, 190077 |
| 5790 | 134794 | FX losses on non-UK mortgage for UK property - tax deductable? | 11122, 141738, 403948, 484375, 356884, 369419, 256395, 241136, 342903, 51497 |
| 7823 | 583549 | Retirement Funds: Betterment vs Vanguard Life strategy vs Target Retirement | 451196, 105666, 175927, 331492, 347825, 268731, 172336, 57070, 11094, 374225 |
| 689 | 411044 | Receive credit card payment sending my customer details to a credit card processing company? | 446932, 104079, 204288, 63366, 96547, 171761, 195852, 421803, 438032, 567201 |
| 9174 | 535317, 160218 | Which U.S. online discount broker is the best value for money? | 192910, 236931, 200052, 413856, 451729, 513281, 31936, 515144, 563334, 358770 |
| 6867 | 443804, 540799, 445258 | Will there always be somebody selling/buying in every stock? | 230343, 301985, 61006, 466143, 18532, 226197, 482739, 543589, 429196, 349147 |
| 2383 | 232199 | Should I Purchase Health Insurance Through My S-Corp | 17215, 527620, 224406, 546634, 457034, 476085, 327232, 534277, 423074, 521489 |
| 5083 | 138845 | Co-signer deceased | 369075, 18257, 270952, 273759, 447983, 453263, 495482, 305509, 518681, 334606 |
| 10526 | 39185 | What extra information might be obtained from the next highest bids in an order book? | 546493, 485973, 283008, 427747, 298551, 251100, 138830, 467852, 322798, 146125 |
| 2181 | 376631, 397329 | What are the risks & rewards of being a self-employed independent contractor / consultant vs. being a permanent emplo... | 37725, 488755, 139501, 584218, 383088, 406656, 197870, 525360, 260603, 524788 |
| 5903 | 231863 | Fees aside, what factors could account for performance differences between U.S. large-cap index ETFs? | 408524, 395842, 159471, 372233, 246996, 20504, 501153, 230997, 517391, 402091 |
| 5620 | 448784, 329552, 548740 | What's the fuss about identity theft? | 260580, 158285, 90632, 598801, 551747, 91986, 423809, 98993, 581889, 5860 |
| 5254 | 392851 | How do I calculate the quarterly returns of a stock index? | 402466, 574974, 563169, 96697, 559168, 238817, 99708, 422904, 531066, 183100 |
| 2472 | 370334 | How do I deal with a mistaken attempt to collect a debt from me that is owed by someone else? | 180601, 201758, 200263, 49321, 62109, 500671, 595441, 435006, 584582, 161422 |
| 2306 | 315875 | To whom should I report fraud on both of my credit cards? | 581889, 531137, 596284, 289706, 90632, 164729, 412542, 270449, 298729, 202457 |
| 7633 | 197839 | Can a trade happen “in between” the bid and ask price? | 494727, 353396, 281844, 137175, 402482, 554207, 458933, 179258, 505244, 164008 |
| 2400 | 564271 | Will I be paid dividends if I own shares? | 1198, 456470, 95889, 1034, 97942, 311214, 501931, 481169, 365627, 587689 |
| 5549 | 286227, 309361 | Pros / cons of being more involved with IRA investments [duplicate] | 429106, 90294, 105468, 336394, 546150, 561636, 471204, 87260, 324012, 181624 |
| 3801 | 307776 | Can a bunch of wealthy people force Facebook to go public? | 390529, 69017, 264498, 209242, 168565, 171236, 371293, 92014, 394734, 134902 |
| 4605 | 453941 | If the U.S. defaults on its debt, what will happen to my bank money? | 41312, 400826, 313306, 169691, 526384, 229310, 479527, 598030, 581054, 373717 |
| 2885 | 367360, 359579, 85229, 454810 | Merits of buying apartment houses and renting them | 430672, 159403, 451849, 502291, 581251, 507029, 80838, 502514, 129149, 343917 |
| 6110 | 94117, 259706 | Why does short selling require borrowing? | 188531, 320450, 226496, 67107, 314478, 79764, 35500, 384252, 107045, 329662 |
| 3694 | 282442 | Has anyone created a documentary about folks who fail to save enough for retirement? | 204747, 242237, 222260, 383051, 311931, 385932, 91911, 40821, 489480, 582307 |
| 8 | 566392 | How to deposit a cheque issued to an associate in my business into my business account? | 65404, 508754, 261856, 590102, 564553, 308938, 301833, 188893, 489199, 29372 |
| 1309 | 156162, 489401 | Why does FlagStar Bank harass you about payments within grace period? | 471630, 489368, 271040, 336792, 15824, 438869, 329817, 526989, 173919, 366594 |
| 7109 | 447781 | How do I analyse moving averages? | 489933, 140804, 42620, 227669, 193012, 257185, 565501, 221627, 518932, 35006 |
| 5080 | 256055 | Is there a standard or best practice way to handle money from an expiring UTMA account? | 445521, 279291, 414429, 451189, 490991, 69841, 236186, 470928, 267206, 324564 |
| 4981 | 247894 | Where can I find open source portfolio management software? | 102684, 45218, 259463, 587792, 557861, 232736, 196432, 200683, 529790, 419171 |
| 7445 | 153178, 104343 | IS it the wrong time to get into the equity market immediately after large gains? | 573612, 483025, 89714, 590902, 350068, 284075, 79111, 33155, 516880, 488207 |
| 2895 | 328691 | Where should a young student put their money? | 426461, 354551, 55841, 148453, 517313, 332749, 496170, 5188, 502170, 442896 |
| 6787 | 587120 | Would it make sense to sell a stock, then repurchase it for tax purposes? | 23217, 219762, 400730, 263751, 474981, 390864, 221715, 328073, 17184, 374867 |
| 6041 | 241308 | Most effective Fundamental Analysis indicators for market entry | 425020, 81655, 224695, 96910, 528034, 108579, 331530, 194240, 263464, 542765 |
| 7700 | 273761, 2653, 179328 | Should I re-allocate my portfolio now or let it balance out over time? | 224392, 269169, 253268, 395208, 131127, 22221, 422051, 434014, 552298, 126836 |
| 547 | 6349 | What percentage of my company should I have if I only put money? | 396694, 68088, 95243, 523158, 80913, 156747, 213399, 498681, 447353, 445353 |
| 3394 | 342258, 129319 | What is the easiest way to back-test index funds and ETFs? | 172374, 448745, 71230, 528034, 224765, 408524, 445971, 571913, 159471, 99568 |
| 4102 | 448699 | How can I determine if my rate of return is “good” for the market I am in? | 597437, 535737, 88801, 554734, 369439, 71219, 162488, 461082, 554237, 135176 |

### Lowest MRR@10

| Query ID | MRR@10 | Recall@100 | First Relevant Rank | Query |
| --- | ---: | ---: | ---: | --- |
| 10246 | 0.0000 | 0.0000 | missing | Understanding the T + 3 settlement days rule |
| 10462 | 0.0000 | 0.0000 | missing | Is it okay to be married, 30 years old and have no retirement? |
| 10482 | 0.0000 | 0.0000 | missing | Rollover into bond fund to do dollar cost averaging [duplicate] |
| 1085 | 0.0000 | 0.0000 | missing | How do disputed debts work on credit reports? |
| 10975 | 0.0000 | 0.0000 | missing | How to contribute to Roth IRA when income is at the maximum limit & you have employer-sponsored 401k plans? |
| 1159 | 0.0000 | 0.0000 | missing | what is the best way to do a freelancing job over the summer for a student |
| 1321 | 0.0000 | 0.0000 | missing | Are social media accounts (e.g. YouTube, Twitter, Instagram, etc.) considered assets? |
| 1783 | 0.0000 | 0.0000 | missing | Freelancing Tax implication |
| 1920 | 0.0000 | 0.0000 | missing | Clarification on student expenses - To file the tax for the next year |
| 2181 | 0.0000 | 0.0000 | missing | What are the risks & rewards of being a self-employed independent contractor / consultant vs. being a permanent emplo... |
