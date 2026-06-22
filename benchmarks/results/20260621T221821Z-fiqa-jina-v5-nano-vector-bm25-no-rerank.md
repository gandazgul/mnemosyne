# Mnemosyne BEIR fiqa Results

- Generated: 2026-06-21T22:18:21Z
- Dataset: `fiqa`
- Queries: 648
- Corpus documents: 57638
- Search limit: 100
- Source mode: `hybrid`
- Fusion: `vector-bm25`
- BM25 weight: `0.15`
- Rerank candidates: 300
- Rerank enabled: False
- Run label: `jina-v5-nano`
- Mnemosyne config: `/Users/gandazgul/Documents/web/mnemosyne/cli/configs/jina-v5-text-nano-retrieval.yaml`

| Metric | Score |
| --- | ---: |
| `ndcg@10` | 0.4733 |
| `mrr@10` | 0.5512 |
| `recall@10` | 0.5391 |
| `recall@100` | 0.7806 |
| `map@100` | 0.4131 |

## Breakdown

| Detail | Count |
| --- | ---: |
| `queries` | 648 |
| `queries_with_rank_1_hit` | 303 |
| `queries_with_top_10_hit` | 471 |
| `queries_with_top_100_hit` | 586 |
| `queries_missing_at_100` | 62 |

### First Relevant Rank

| Bucket | Queries |
| --- | ---: |
| `rank_1` | 303 |
| `rank_2_3` | 99 |
| `rank_4_10` | 69 |
| `rank_11_100` | 115 |
| `missing@100` | 62 |

### Missing At 100

| Query ID | Missing Doc IDs | Query | Top 10 Doc IDs |
| --- | --- | --- | --- |
| 4641 | 319954, 397358 | Where should I park my rainy-day / emergency fund? | 497993, 580025, 527939, 406219, 44594, 285812, 400503, 583695, 254572, 538023 |
| 3451 | 26292, 192307, 588448, 490170 | Should you keep your stocks if you are too late to sell? | 251536, 420974, 471911, 310683, 528518, 545284, 301757, 41852, 306460, 27015 |
| 753 | 243503 | Taxes due for hobbyist Group Buy | 466718, 203791, 122185, 132780, 114418, 451020, 299211, 127004, 170632, 465447 |
| 4465 | 376575 | How to donate to charity that will make a difference? | 174033, 132881, 106786, 90591, 427322, 46381, 379353, 326167, 266342, 174543 |
| 7326 | 584295 | Do brokers execute every trade on the exchange? | 404339, 593445, 163333, 35340, 272008, 37040, 147573, 294718, 257656, 395357 |
| 5951 | 497260 | Why can't house prices be out of tune with salaries | 599860, 418034, 31663, 259777, 374480, 62702, 196405, 596834, 589470, 285525 |
| 10827 | 160786, 7748, 107554, 95282 | How much should I be contributing to my 401k given my employer's contribution? | 290105, 296405, 140330, 555377, 436071, 452592, 15841, 242556, 497561, 576391 |
| 9771 | 263955, 28740 | Is there any emprical research done on 'adding to a loser' | 137534, 18671, 490479, 130349, 38704, 375290, 137898, 393272, 356175, 468717 |
| 3724 | 508921, 279570, 497216, 552887, 199970 | Should you always max out contributions to your 401k? | 302512, 3104, 273497, 122910, 430931, 576391, 43573, 497561, 488673, 135790 |
| 5853 | 476663, 160105, 495699, 424598 | Paying Off Principal of Home vs. Investing In Mutual Fund | 284318, 182612, 387722, 473647, 494148, 64456, 473865, 406325, 186071, 31525 |
| 570 | 363591 | Employer options when setting up 401k for employees | 532839, 117845, 79375, 289064, 301616, 387876, 150883, 15841, 555377, 242529 |
| 594 | 534059 | Should a retail trader bother about reading SEC filings | 377322, 64881, 11148, 97837, 213214, 548596, 314898, 86281, 161411, 38863 |
| 10122 | 273718 | Why diversify stocks/investments? | 144261, 297100, 259084, 459970, 89084, 508540, 502495, 180855, 331008, 139368 |
| 1783 | 332314 | Freelancing Tax implication | 421924, 14609, 156063, 159709, 179359, 445298, 383172, 549870, 223624, 15270 |
| 6004 | 149555 | Put-Call parity - what is the difference between the two representations? | 374797, 122432, 13260, 278373, 345410, 216065, 118762, 271109, 247738, 21768 |
| 4837 | 531841, 20958 | When applying for a mortgage, can it also cover outstanding debts? | 97925, 245005, 204035, 32749, 516578, 294167, 30311, 274870, 366448, 254454 |
| 4415 | 67676 | How much is inflation? | 117578, 206580, 381675, 468089, 501743, 290585, 513249, 267176, 519596, 59225 |
| 8378 | 125298 | Should I wait a few days to sell ESPP Stock? | 511678, 133644, 434812, 495568, 546125, 294573, 178684, 387035, 361345, 71713 |
| 8102 | 552707, 378173, 90294 | When do I sell a stock that I hold as a long-term position? | 306460, 165970, 35752, 537212, 171819, 557877, 203638, 310683, 219762, 496209 |
| 7928 | 118633 | If I believe a stock is going to fall, what options do I have to invest on this? | 427808, 410404, 501504, 260384, 67415, 47053, 480967, 173374, 137073, 221869 |
| 4777 | 590710 | How to finance necessary repairs to our home in order to sell it? | 365342, 52351, 570318, 416382, 478413, 171631, 482638, 426678, 360872, 378384 |
| 620 | 331332 | Is it wise to have plenty of current accounts in different banks? | 543921, 467059, 535817, 572848, 164801, 489959, 38628, 180673, 61734, 380786 |
| 8271 | 415511 | Income in zero-interest environment | 376709, 249558, 380368, 52047, 119298, 83330, 136262, 65179, 106424, 137225 |
| 864 | 211364, 152072 | Why use accounting software like Quickbooks instead of Excel spreadsheets? | 30142, 472924, 24890, 566337, 2436, 329774, 157751, 78117, 222380, 402174 |
| 5888 | 540806 | Interest charges on balance transfer when purchases are involved | 263647, 125497, 543776, 336792, 579601, 545327, 358445, 490529, 429746, 213242 |
| 2568 | 388798, 127353 | How to pay with cash when car shopping? | 346042, 9146, 15696, 166314, 258247, 355310, 108739, 514238, 269898, 453301 |
| 7124 | 74615 | How come we can find stocks with a Price-to-Book ratio less than 1? | 558617, 526110, 583708, 533818, 278582, 154725, 504243, 157597, 123263, 226070 |
| 2857 | 295864 | I have around 60K $. Thinking about investing in Oil, how to proceed? | 233732, 127566, 501384, 474575, 542051, 117451, 49235, 379140, 76996, 316444 |
| 8539 | 218728, 196304, 396038 | Can the risk of investing in an asset be different for different investors? | 283074, 483123, 391861, 502495, 471817, 385881, 499166, 500863, 387515, 378403 |
| 6896 | 251704 | Selling high, pay capital gains, re-purchase later | 448659, 468047, 343219, 522319, 376800, 66834, 216827, 169240, 561999, 196427 |
| 1085 | 467737, 393710 | How do disputed debts work on credit reports? | 372039, 450031, 161422, 78328, 268777, 319276, 398258, 242013, 384608, 574122 |
| 5422 | 151973 | What are some good books for learning stocks, bonds, derivatives e.t.c for beginner with a math background? | 165294, 172587, 273906, 276786, 193555, 552371, 191688, 221319, 79517, 241423 |
| 8247 | 42521, 321114, 465313 | Tax on Stocks or ETF's | 586010, 161019, 153112, 518735, 437907, 580802, 474745, 195767, 190687, 528880 |
| 10482 | 549072 | Rollover into bond fund to do dollar cost averaging [duplicate] | 330023, 447567, 224782, 439757, 265817, 564787, 134005, 175576, 211765, 18436 |
| 8789 | 70853 | What does “profits to the shareholders jumped to 15 cents a share” mean? | 87349, 41912, 341424, 20076, 573079, 14870, 234040, 219927, 317363, 144079 |
| 2423 | 538023 | At what age should I start or stop saving money? | 529444, 417787, 396127, 234846, 272328, 337561, 235855, 553288, 104457, 328157 |
| 4523 | 594257 | What should I do with my $25k to invest as a 20 years old? | 129255, 272070, 465819, 171712, 332022, 442776, 216365, 10476, 286746, 72578 |
| 8513 | 270573 | Buy on dip when earnings fail? | 572622, 175821, 203873, 351396, 462135, 573767, 335626, 391043, 73872, 53431 |
| 5054 | 28119 | How to stress test an investment plan? | 588481, 377186, 564007, 101124, 127263, 458183, 582899, 220665, 448745, 403916 |
| 1159 | 496064 | what is the best way to do a freelancing job over the summer for a student | 55064, 347181, 460648, 163881, 156063, 112669, 132287, 449155, 237282, 594182 |
| 68 | 19183 | Intentions of Deductible Amount for Small Business | 519473, 86134, 97233, 192516, 97719, 354716, 462831, 381151, 54333, 146657 |
| 9701 | 387141, 357739 | How to bet against the London housing market? | 473883, 408865, 225682, 258048, 516214, 108399, 588608, 73283, 412502, 70799 |
| 6199 | 239214 | How can all these countries owe so much money?  Why & where did they borrow it from? | 414693, 584273, 47163, 10399, 49602, 169921, 490042, 351853, 298794, 380714 |
| 9126 | 514831 | Short an option - random assignment? | 334473, 477588, 82194, 308859, 228810, 307518, 166227, 102316, 590453, 469382 |
| 34 | 599545 | 401k Transfer After Business Closure | 458917, 551545, 144109, 490867, 483268, 226547, 424766, 15728, 168890, 104134 |
| 6395 | 166227 | Option settlement for calendar spreads | 584223, 467463, 276314, 22916, 401447, 505223, 111301, 8177, 516790, 273142 |
| 9115 | 207325 | Why does the calculation for percentage profit vary based on whether a position is short vs. long? | 419897, 158520, 422467, 245082, 428786, 314478, 331606, 154665, 385220, 232880 |
| 4767 | 280805, 568670, 224057, 22804 | New car: buy with cash or 0% financing | 420018, 584106, 429153, 306834, 56867, 9146, 451092, 166314, 408932, 475440 |
| 8351 | 472516 | What happens when a calendar spread is assigned in a non-margin account? | 102316, 141213, 45674, 273142, 527654, 516790, 23609, 272754, 100628, 300139 |
| 6131 | 381720, 170204, 416679, 2460 | Is it ever a good idea to close credit cards? | 326094, 368806, 391384, 334111, 339030, 218088, 99449, 258465, 35625, 533288 |
| 858 | 45185, 122485, 278450 | Is it bad practice to invest in stocks that fluctuate by single points throughout the day? | 293027, 146632, 567608, 30774, 433730, 573612, 241860, 208932, 519501, 214281 |
| 4019 | 6881 | How and Should I Invest (As a college 18 year old with minimal living expenses)? | 85977, 332938, 332749, 379948, 426461, 269671, 20304, 66864, 571044, 10476 |
| 6080 | 164513 | Is ScholarShare a legitimate entity for a 529 plan in California? | 22856, 201500, 236732, 277581, 468527, 83080, 233401, 115175, 2809, 535357 |
| 6959 | 205010 | What is the term for the quantity (high price minus low price) for a stock? | 428117, 468025, 303325, 169954, 577573, 229573, 599523, 412223, 373034, 304399 |
| 5402 | 491350 | Is it impossible to get a home loan with a poor credit history after a divorce? | 445163, 90579, 51728, 227485, 595029, 44105, 52250, 440063, 180214, 122807 |
| 6562 | 501157 | Cheapest way to “wire” money in an Australian bank account to a person in England, while I'm in Laos? | 473605, 60446, 282744, 582414, 183880, 182443, 549684, 308837, 385182, 135675 |
| 701 | 389446 | What are the ins/outs of writing-off part of one's rent for working at home? | 231990, 349672, 436505, 344955, 337706, 456234, 124507, 339488, 243306, 177074 |
| 7674 | 519390 | Choosing the limit when making a limit order? | 249279, 447886, 15917, 155151, 514841, 278630, 94653, 526235, 184756, 437436 |
| 5940 | 486243, 93936 | How does investment into a private company work? | 512609, 250354, 46842, 454465, 182226, 473154, 535314, 252853, 530, 135411 |
| 6612 | 205522, 322900 | If I have a lot of debt and the housing market is rising, should I rent and slowly pay off my debt or buy and roll th... | 502594, 301192, 431481, 180192, 254454, 198442, 433171, 290900, 317945, 229572 |
| 4714 | 450819 | Personal finance app where I can mark transactions as “reviewed”? | 505057, 584450, 218793, 353915, 344473, 529790, 29812, 479390, 65957, 386349 |
| 8456 | 486333 | What typically happens to unvested stock during an acquisition? | 257853, 93215, 534755, 555276, 469036, 186869, 174321, 492428, 475019, 391156 |
| 10213 | 270221, 545712 | Looking for good investment vehicle for seasonal work and savings | 195373, 446186, 38269, 100517, 112499, 272198, 96949, 386305, 263390, 488100 |
| 5196 | 172128, 114829 | I might use a credit card convenience check. What should I consider? | 565745, 402543, 85517, 85252, 289483, 456098, 393866, 2875, 481052, 302823 |
| 3006 | 269851, 568473, 328300 | Strategies for putting away money for a child's future (college, etc.)? | 512096, 127838, 258704, 372900, 332749, 211713, 303432, 290441, 471019, 8266 |
| 3909 | 312248, 404356, 245616, 353028 | How to rescue my money from negative interest? | 514003, 83330, 61586, 472837, 328499, 362730, 574011, 42475, 404352, 438403 |
| 6907 | 251604 | Nominal value of shares | 4854, 303112, 170652, 111827, 69506, 91870, 481761, 487738, 480515, 275392 |
| 5464 | 350399, 86691 | Resources on Buying Rental Properties | 222095, 423438, 26339, 372274, 325722, 383921, 315972, 545341, 536126, 426705 |
| 10034 | 480749 | Tax implications of holding EWU (or other such UK ETFs) as a US citizen? | 528880, 181942, 565296, 44955, 197478, 447197, 141585, 430868, 180146, 85926 |
| 5090 | 436493 | Should I take a student loan to pursue my undergraduate studies in France? | 12988, 246286, 92430, 21913, 217831, 560681, 586289, 58005, 287507, 455666 |
| 2088 | 399875 | How would I go about selling the stock of a privately held company? | 53993, 455168, 291886, 293687, 140835, 238215, 72846, 188776, 530, 413672 |
| 9391 | 503637 | Should I replace bonds in a passive investment strategy | 535518, 136515, 248158, 107424, 577832, 283202, 155242, 171669, 545760, 494653 |
| 3148 | 178127, 438000 | Can a car company refuse to give me a copy of my contract or balance details? | 172855, 584305, 430100, 29721, 357280, 65046, 92888, 395995, 164702, 560325 |
| 4678 | 305153 | Finance, Cash or Lease? | 185405, 311748, 215225, 376016, 427884, 504918, 311446, 260095, 487678, 522532 |
| 2398 | 363810, 224654, 590489 | Frustrated Landlord | 556453, 44058, 487094, 96538, 98372, 393883, 168089, 310992, 201705, 395770 |
| 6746 | 210887 | What happens if stock purchased on margin plummets below what I have in the brokerage? | 279185, 333674, 247680, 231221, 283982, 115918, 527654, 176822, 469830, 399903 |
| 5511 | 169893, 560325, 478426, 383193, 278699, 12746 | Pay off car loan entirely or leave $1 until the end of the loan period? | 334559, 38786, 376016, 529123, 329137, 51873, 155843, 500946, 107898, 139788 |
| 8834 | 12232 | Pros/Cons of Buying Discounted Company Stock | 203139, 599156, 528827, 133644, 569303, 569224, 371821, 57387, 163396, 469100 |
| 988 | 226053, 107688 | Where should I invest my savings? | 168402, 501384, 60093, 571218, 450558, 347651, 223872, 285812, 223551, 388252 |
| 3369 | 163834, 231012, 145716, 411910, 395840 | Why should one only contribute up to the employer's match in a 401(k)? | 341493, 555377, 296405, 436071, 242556, 15841, 92370, 240373, 24231, 463892 |
| 9296 | 435746 | Why would Two ETFs tracking Identical Indexes Produce different Returns? | 148721, 206744, 159471, 368124, 408524, 428187, 492212, 410123, 209996, 285135 |
| 9245 | 194561 | Stock Options for a company bought out in cash and stock | 207253, 186869, 39345, 131488, 178497, 265111, 248393, 259560, 261487, 409818 |
| 3490 | 420529 | Tax Witholding for Stock Sale | 447651, 591157, 361482, 311782, 537371, 400730, 152960, 42521, 407602, 367742 |
| 5763 | 462019 | What is the best way to get a “rough” home appraisal prior to starting the refinance process? | 570318, 38712, 326214, 67379, 218144, 251466, 89964, 563380, 331255, 215647 |
| 4962 | 599925 | Net Cash Flows from Selling the Bond and Investing | 416839, 158363, 52149, 408661, 537603, 34949, 308276, 431386, 187110, 535518 |
| 4846 | 151104 | Is there anything comparable to/resembling CNN's Fear and Greed Index? | 98096, 335892, 538974, 3533, 415161, 320059, 183597, 489352, 270305, 533311 |
| 9403 | 6666, 328086, 345199 | Abundance of Cash - What should I do? | 410450, 570632, 159235, 14349, 551986, 103447, 215296, 186332, 357887, 372223 |
| 5993 | 367375, 272866, 55084, 352638, 426120, 63501 | Why would anyone want to pay off their debts in a way other than “highest interest” first? | 94373, 416796, 160193, 287571, 128574, 353911, 494306, 156195, 886, 431212 |
| 5710 | 232311 | Bucketing investments to track individual growths | 227364, 516267, 88417, 411856, 534323, 412830, 227733, 508610, 135765, 472328 |
| 7529 | 66607 | Does the expense ratio of a fund-of-funds include the expense ratios of its holdings? | 514529, 464337, 89297, 59249, 293626, 287537, 102904, 218261, 387980, 464668 |
| 5021 | 589285 | Is there a more flexible stock chart service, e.g. permitting choice of colours when comparing multiple stocks? | 528576, 584801, 189341, 555506, 465971, 211444, 60284, 511861, 494939, 49893 |
| 3612 | 259625 | How can I buy and sell the same stock on the same day? | 522658, 567383, 310636, 390864, 402726, 165548, 429418, 584291, 367873, 460937 |
| 4409 | 499128, 100306, 147439 | My friend wants to put my name down for a house he's buying. What risks would I be taking? | 243732, 223841, 268078, 102088, 514790, 102326, 84732, 60135, 360682, 341947 |
| 2070 | 363678 | Advantage of credit union or local community bank over larger nationwide banks such as BOA, Chase, etc.? | 550303, 578357, 469515, 587737, 597571, 590209, 249839, 30253, 38038, 408166 |
| 11039 | 53544, 249063 | Pay off credit card debt or earn employer 401(k) match? | 91183, 287876, 552383, 5203, 163287, 79363, 508534, 345895, 105557, 437706 |
| 5460 | 184337, 21174, 108514, 463885 | Paying off a loan with a loan to get a better interest rate | 77052, 106495, 470716, 327115, 344812, 343208, 529418, 555280, 243065, 194557 |
| 7925 | 318185, 402482 | Can I sell a stock immediately? | 591436, 438974, 584291, 227399, 310636, 44461, 332467, 81721, 315760, 581866 |
| 4286 | 566069 | Given advice “buy term insurance and invest the rest”, how should one “invest the rest”? | 70460, 229239, 10531, 391243, 155640, 511386, 206830, 151817, 564675, 79142 |
| 2685 | 154113, 370300, 468923 | What ways are there for us to earn a little extra side money? | 382005, 468086, 594182, 576047, 269380, 237950, 186889, 4992, 280099, 558832 |
| 1090 | 518896 | Need a formula to determine monthly payments received at time t if I'm reinvesting my returns | 446454, 393987, 16051, 179365, 296146, 573928, 520217, 19999, 281329, 209238 |
| 6122 | 44344 | Better to rent condo to daughter or put her on title? | 496166, 316794, 403515, 558251, 182039, 118246, 53840, 577658, 566184, 80269 |
| 4514 | 69485, 337764, 209804 | What intrinsic, non-monetary value does gold have as a commodity? | 471825, 426270, 156211, 317429, 146573, 99089, 408336, 240894, 80141, 532381 |
| 8507 | 370995 | When to sell a stock? | 99132, 251536, 303724, 545284, 272091, 236415, 217837, 88813, 420974, 102237 |
| 6221 | 257248, 519675, 76414, 455614, 115717 | To pay off a student loan, should I save up a lump sum payoff payment or pay extra each month? | 448791, 352363, 254245, 110081, 124705, 529551, 274108, 394474, 541313, 414534 |
| 3008 | 180192, 323406 | What are my chances at getting a mortgage with Terrible credit but High income | 102266, 231688, 407401, 285694, 47441, 251846, 44105, 455952, 78176, 574438 |
| 4007 | 521657 | What is a reasonable salary for the owner and sole member of a small S-Corp? | 556220, 260385, 370542, 205341, 170933, 388704, 315552, 543085, 458431, 515233 |
| 6644 | 175035 | How to know precisely when a SWIFT is issued by a bank? | 110198, 475527, 218761, 118396, 271596, 355870, 41383, 327623, 39783, 298587 |
| 10267 | 460398 | How should I prepare for the next financial crisis? | 178693, 143393, 569632, 305600, 326398, 87520, 182442, 36961, 96017, 369470 |
| 7622 | 253369, 378594 | Best way to pay off debt? | 220241, 457945, 353911, 480773, 388095, 157923, 373554, 345895, 271525, 416796 |
| 3767 | 153922, 392060 | What should I be doing to protect myself from identity theft? | 90632, 423809, 260580, 97686, 581889, 171510, 125204, 158008, 158285, 587778 |
| 6410 | 471723 | Will an ETF immediately reflect a reconstitution of underlying index | 454610, 71230, 214281, 87261, 330729, 295993, 200360, 87238, 227324, 357127 |
| 5030 | 215540 | Why pay for end-of-day historical prices? | 227192, 13511, 532178, 149420, 295344, 560108, 471131, 327974, 378994, 370569 |
| 6252 | 394551, 160932, 293624, 233294, 243268, 379487, 62868 | Is this mortgage advice good, or is it hooey? | 213713, 473647, 120061, 47565, 139366, 205906, 495089, 104988, 27268, 443852 |
| 885 | 337165, 409184 | How long do credit cards keep working after you disappear? | 254968, 516678, 472336, 99449, 89888, 251701, 251643, 181757, 301792, 588719 |
| 4031 | 115741 | 28 years old and just inherited large amount of money and real estate - unsure what to do with it | 318864, 140002, 588316, 568629, 65180, 80844, 375708, 266481, 193171, 387717 |
| 766 | 550172 | Will the ex-homeowner still owe money after a foreclosure? | 2996, 299591, 552768, 163711, 333583, 427110, 104955, 27987, 212827, 268865 |
| 8202 | 513258, 93971 | What accounted for DXJR's huge drop in stock price? | 457689, 317363, 337001, 122542, 71924, 537862, 67237, 355167, 412584, 421248 |
| 7345 | 237645 | What do these numbers mean? (futures) | 9274, 527080, 276381, 206895, 508821, 460331, 108, 529996, 164001, 354429 |
| 776 | 583640, 127263, 496899 | Can saving/investing 15% of your income starting age 25, likely make you a millionaire? | 124027, 10440, 417787, 143591, 41960, 467044, 374266, 191148, 434972, 554833 |
| 89 | 248624 | How can I deposit a check made out to my business into my personal account? | 508754, 309023, 526817, 135196, 308938, 188893, 29372, 80538, 590102, 521540 |
| 1920 | 269943 | Clarification on student expenses - To file the tax for the next year | 263485, 481114, 585356, 551187, 585023, 128980, 295562, 416117, 446889, 316482 |
| 8013 | 496159, 224231 | Frequency of investments to maximise returns (and minimise fees) | 384983, 537626, 81652, 57033, 388389, 48678, 224816, 40652, 446948, 8759 |
| 3759 | 527966, 67167, 522358 | Simplifying money management | 455457, 373772, 214248, 490065, 248663, 122378, 372743, 145812, 526159, 231412 |
| 10639 | 431799, 495774, 278453, 187039 | Short term parking of a large inheritance? | 171196, 163353, 235628, 318864, 111048, 590276, 163197, 375708, 131391, 546538 |
| 6635 | 156358 | Why don't share prices of a company rise every other Friday when the company buys shares for its own employees? | 587137, 3656, 533712, 235531, 343452, 12560, 579037, 95806, 491064, 125298 |
| 4312 | 399149 | Is it true that 90% of investors lose their money? | 282435, 222639, 167950, 285945, 116647, 497786, 532485, 170628, 431735, 300770 |
| 6525 | 181985 | Does it make sense to trade my GOOGL shares for GOOG and pocket the difference? | 106541, 550661, 98150, 105542, 362473, 156467, 498014, 378906, 221795, 192696 |
| 2590 | 589625 | Are non-residents or foreigners permitted to buy or own shares of UK companies? | 296528, 48269, 209493, 55999, 158923, 310103, 458730, 307776, 456999, 188776 |
| 5374 | 152688 | What were the main causes of the spike and drop of DRYS's stock price? | 283106, 133204, 457689, 122542, 261522, 317363, 362462, 467594, 253339, 380894 |
| 2994 | 419319 | Work on the side for my wife's company | 510317, 569145, 321743, 138428, 491844, 460721, 5840, 382005, 269380, 506991 |
| 3683 | 185909, 454501 | Can I trust the Motley Fool? | 276975, 105973, 408995, 428848, 500338, 538086, 565016, 6607, 522713, 192912 |
| 7206 | 441155, 532211, 553066 | Who Bought A Large Number Of Shares? | 65667, 34882, 351570, 573846, 558703, 358164, 350214, 444752, 498676, 552375 |
| 10246 | 77573 | Understanding the T + 3 settlement days rule | 370635, 156029, 176717, 327080, 340263, 28314, 332243, 11927, 226984, 293389 |
| 5241 | 322157, 27489 | Mortgage vs. Cash for U.S. home buy now | 344740, 281675, 213713, 111184, 438073, 390976, 309420, 273735, 42604, 234286 |
| 98 | 575929 | How can I make $250,000.00 from trading/investing/business within 5 years? | 527522, 555630, 66034, 102113, 373119, 519619, 336661, 438279, 506149, 121161 |
| 4615 | 262934 | Are solar cell panels and wind mills worth the money? | 261900, 69523, 496427, 455798, 425595, 120384, 271015, 510872, 376430, 385028 |
| 6467 | 453256, 23217, 346641, 367313 | Advice on strategy for when to sell | 88813, 240089, 217837, 203873, 109455, 368348, 99857, 130941, 250873, 83807 |
| 4289 | 24881 | Does the currency exchange rate contain any additional information at all? | 288330, 17469, 114886, 517345, 226102, 324546, 356465, 416975, 119316, 135220 |
| 4394 | 336045, 441582 | Transfer $50k to another person's account (in California, USA) | 322838, 93386, 462585, 305907, 293653, 431462, 415655, 412258, 521753, 307404 |
| 7344 | 108403 | How is the Dow divisor calculated? | 14368, 159166, 150430, 378974, 253926, 313421, 65618, 591089, 501032, 69655 |
| 6875 | 224392 | Where to find free Thailand stock recommendations and research? | 567500, 110733, 352557, 224366, 284539, 79337, 556770, 9354, 232460, 77502 |
| 10447 | 152096, 300721 | Is there an advantage to a traditional but non-deductable IRA over a taxable account? [duplicate] | 144751, 500175, 382236, 532657, 447482, 259150, 382894, 540389, 299690, 114912 |
| 9871 | 448890, 40051, 170594 | What should I do with the 50k I have sitting in a European bank? | 367207, 73741, 292714, 433003, 293179, 76562, 74668, 231521, 387723, 175139 |
| 3625 | 414295 | What should I do with my paper financial documents? | 509617, 500751, 569812, 380263, 513248, 163168, 113830, 37582, 123366, 44204 |
| 6005 | 135415, 478457, 345895, 73310, 384626, 390689 | Why might it be advisable to keep student debt vs. paying it off quickly? | 149500, 571198, 507544, 96268, 431884, 25190, 572272, 564206, 422704, 414288 |
| 3822 | 385090, 418900, 308837 | How to change a large quantity of U.S. dollars into Euros? | 292714, 194730, 417917, 239876, 340777, 174406, 19618, 541608, 478065, 208499 |
| 7879 | 372551, 421285 | Any Tips on How to Get the Highest Returns Within 4 Months by Investing in Stocks? | 58186, 102029, 272174, 43088, 540919, 174313, 105391, 228488, 7625, 367415 |
| 3115 | 234950, 389028, 316794 | How can I live outside of the rat race of American life with 300k? | 233562, 183869, 267892, 129364, 174272, 136035, 369742, 366961, 220023, 475736 |
| 3995 | 278734, 230208 | I have more than $250,000 in a US Bank account… mistake? | 485507, 404954, 171720, 479918, 264934, 14349, 352883, 303367, 505461, 506909 |
| 10136 | 526115 | How to minimise the risk of a reduction in purchase power in case of Brexit for money held in a bank account? | 466950, 290930, 417740, 304007, 583903, 205685, 35511, 265453, 150543, 152316 |
| 8635 | 67107, 240215 | Is there any flaw in this investment scheme? | 46818, 493841, 151238, 447619, 575918, 303619, 203729, 365816, 202638, 510565 |
| 5206 | 563030, 28230, 117276, 300660 | Is it a good idea to get an unsecured loan to pay off a credit card that won't lower a high rate? | 298908, 595455, 69938, 225522, 340520, 153088, 375780, 2519, 504293, 516397 |
| 2713 | 388147 | Physical Checks - Mailing | 284528, 29372, 41944, 216200, 20791, 199069, 584170, 78139, 229546, 350237 |
| 9060 | 40447 | Buying puts without owning underlying | 511093, 528052, 228217, 181924, 338782, 345851, 7743, 359778, 3062, 316037 |
| 4105 | 25096 | As an investor what are side effects of Quantitative Easing in US and in EU? | 416483, 345910, 176262, 305029, 108519, 393791, 369038, 293104, 30946, 339640 |
| 2465 | 570680, 81046 | Can capital expenses for volunteer purposes be deducted from income? | 37382, 202645, 432545, 510716, 146657, 398536, 598646, 18889, 541809, 275543 |
| 4640 | 101369 | What can my relatives do to minimize their out of pocket expenses on their fathers estate | 295246, 331534, 372808, 356035, 17110, 360816, 367404, 144965, 422994, 375708 |
| 9275 | 338754 | Do I have to pay a capital gains tax if I rebuy the same stock within 30 days? | 400730, 23217, 343219, 537916, 376800, 390864, 161155, 407602, 102443, 263751 |
| 3500 | 174019 | Why invest in becoming a landlord? | 273187, 528206, 71424, 557478, 422331, 578597, 141935, 572061, 159156, 11601 |
| 1306 | 484437, 204075 | I made an investment with a company that contacted me, was it safe? | 594206, 160611, 309851, 538086, 205665, 537698, 309590, 519038, 167684, 316321 |
| 6262 | 26799 | Help required on estimating SSA benefit amounts | 34538, 390877, 118707, 2648, 83338, 498444, 15322, 430407, 320362, 529927 |
| 8632 | 213976 | Is it best to exercise options shares when they vest, or wait | 43497, 237718, 104188, 382381, 259560, 237783, 420722, 340730, 293959, 255927 |
| 6133 | 415705 | What happens to all of the options when they expire? | 7733, 575408, 11456, 581672, 242298, 132288, 73256, 358492, 428399, 116436 |
| 3771 | 488948, 198349, 217683, 49601 | Best way to buy Japanese yen for travel? | 490384, 495826, 521712, 96211, 306130, 434201, 217715, 350245, 128471, 120604 |
| 1736 | 25543, 443419 | How can people have such high credit card debts? | 399406, 562896, 437610, 569056, 372993, 475668, 517050, 174941, 298908, 99463 |
| 6814 | 340214, 223206 | Selling Stock - All or Nothing? | 590188, 513734, 400614, 154976, 66834, 67107, 279782, 276883, 3095, 878 |
| 1322 | 64138 | Is this follow-up after a car crash a potential scam? | 226090, 114231, 397852, 567973, 44635, 283917, 91463, 332916, 98356, 524723 |
| 5185 | 210236, 317354 | Invest in low cost small cap index funds when saving towards retirement? | 196992, 376485, 241202, 262180, 580313, 524525, 503725, 106620, 434279, 268731 |
| 6909 | 127012 | Why do stocks priced above $2.00 on the ASX sometimes move in $0.005 increments? | 72633, 118232, 490584, 47217, 112946, 375019, 64943, 168080, 298551, 450489 |
| 2348 | 211867, 566573, 211622, 474234, 352271, 265874 | Why can't you just have someone invest for you and split the profits (and losses) with him? | 447619, 306430, 247486, 389004, 420544, 381757, 151412, 177194, 309851, 64410 |
| 4499 | 76996 | Is investing exlusively in a small-cap index fund a wise investment? | 196992, 501153, 52274, 517391, 235917, 228111, 372233, 335136, 14748, 376485 |
| 42 | 272709 | What are the ins/outs of writing equipment purchases off as business expenses in a home based business? | 305222, 510863, 581265, 88967, 571711, 28764, 443859, 35379, 456234, 47260 |
| 3530 | 239998 | How to exclude stock from mutual fund | 184299, 24029, 378075, 370754, 209879, 449124, 479420, 110343, 574383, 332152 |
| 9108 | 272021, 472585 | Starting an investment portfolio with Rs 5,000/- | 290757, 46967, 171189, 414116, 312821, 240351, 336218, 323067, 51848, 356552 |
| 6835 | 102243 | Are bond ETF capital gains taxed similar to stock or stock funds if held for more than 1 year? | 149305, 5710, 84238, 570546, 287950, 153112, 29502, 586010, 110343, 543842 |
| 3067 | 406156 | Should I make extra payments to my under water mortgage or increase my savings? | 477907, 476068, 560915, 3092, 131365, 423403, 468831, 90009, 341837, 336276 |
| 4125 | 344648, 72046 | Alternative means of salary for my employees | 174787, 73999, 365558, 58906, 361954, 127347, 562211, 36608, 245451, 355244 |
| 1150 | 43603, 19936 | How are the best way to make and save money at 22 years old | 10476, 529444, 353369, 319760, 595287, 433986, 328157, 66864, 38269, 494815 |
| 9808 | 40702, 431946 | Selling To Close | 416307, 557582, 151587, 345368, 218423, 43087, 362473, 152719, 374204, 107045 |
| 3888 | 319213, 239632 | Why I can't view my debit card pre-authorized amounts? | 208169, 316652, 432077, 185434, 276733, 294077, 440527, 281129, 181757, 209718 |
| 10109 | 506374, 499849 | Why does Charles Schwab have a Mandatory Settlement Period after selling stocks? | 28314, 93231, 332243, 266725, 370635, 98302, 121465, 563826, 226984, 124188 |
| 2790 | 279329, 469125 | Should I pay more than 20% down on a home? | 472484, 400896, 357200, 75961, 207564, 385343, 64400, 215103, 487593, 330229 |
| 9882 | 65702 | Money-market or cash-type ETFs for foreigners with U.S brokerage account | 391876, 389581, 131059, 313775, 363378, 188524, 535340, 466845, 179527, 94477 |
| 4999 | 314898 | Looking for a good source for Financial Statements | 9938, 171964, 431459, 11263, 597241, 295738, 146076, 465971, 520165, 516379 |
| 3189 | 225395 | Diversify my retirement investments with a Roth IRA | 287225, 240975, 122222, 404800, 423658, 187124, 88311, 347651, 458168, 272840 |
| 5134 | 158523 | Why does Yahoo Finance's data for a Vanguard fund's dividend per share not match the info from Vanguard? | 46774, 532616, 206727, 584128, 374330, 405474, 215486, 465536, 221477, 239137 |
| 1321 | 216456, 292065 | Are social media accounts (e.g. YouTube, Twitter, Instagram, etc.) considered assets? | 392379, 381322, 203715, 235082, 481036, 210019, 559738, 426184, 100849, 517726 |
| 4539 | 370879 | How should I save money if the real interest rate (after inflation) is negative? | 42475, 472837, 275925, 203926, 32744, 449745, 61586, 194080, 220720, 503394 |
| 715 | 579763, 546538, 187404 | what would you do with $100K saving? | 548758, 133120, 427032, 113885, 203201, 337561, 333784, 328770, 387647, 121108 |
| 504 | 344203, 498751 | Have plenty of cash flow but bad credit | 22807, 93573, 546097, 569240, 70806, 41875, 368247, 495431, 68431, 252473 |
| 2296 | 83330, 366594, 253563 | How does a bank make money on an interest free secured loan? | 400009, 396853, 119298, 259919, 94230, 580147, 106424, 279897, 172303, 546874 |
| 10975 | 61022 | How to contribute to Roth IRA when income is at the maximum limit & you have employer-sponsored 401k plans? | 163865, 441632, 81148, 360533, 140330, 101490, 446226, 353009, 110114, 447482 |
| 1994 | 156640 | Does the IRS reprieve those who have to commute for work? | 434846, 231990, 192843, 263259, 51491, 63919, 243356, 380635, 585706, 147624 |
| 9164 | 365298, 263390 | Bonds vs equities: crash theory | 115648, 309326, 149900, 506743, 321941, 16924, 599420, 287656, 305274, 296516 |
| 10462 | 8266, 11378, 35680, 437879, 204035, 581204 | Is it okay to be married, 30 years old and have no retirement? | 268023, 66376, 122333, 152478, 391583, 151774, 361049, 139595, 51914, 593377 |
| 8855 | 208165, 312821 | How do i get into investing stocks [duplicate] | 155677, 312445, 403092, 560395, 555521, 313570, 484327, 142320, 367415, 152286 |
| 7071 | 124230 | ESPP strategy - Sell right away or hold? | 511678, 133644, 294573, 127702, 434812, 361345, 71713, 575213, 387035, 495568 |
| 8974 | 523331, 356595, 170625 | As a 22-year-old, how risky should I be with my 401(k) investments? | 216365, 134931, 10476, 102501, 452126, 140738, 336144, 124762, 199237, 96110 |
| 5178 | 240261 | Formula that predicts whether one is better off investing or paying down debt | 39819, 557506, 373554, 393833, 111815, 396889, 290434, 262772, 257016, 154449 |
| 5061 | 23747 | What fiscal scrutiny can be expected from IRS in early retirement? | 502150, 25481, 31699, 149742, 513392, 151263, 598378, 398078, 96720, 517903 |
| 2075 | 170042, 359580, 14967 | Are stories of turning a few thousands into millions by trading stocks real? | 519619, 44417, 65667, 506149, 523393, 555521, 285147, 519501, 375876, 33357 |
| 4335 | 357013 | What is the US Fair Tax? | 318583, 419466, 322246, 589161, 484148, 363178, 585212, 412877, 3181, 479451 |
| 7533 | 93853 | Investing tax (savings) | 8012, 142658, 563986, 315105, 250603, 526664, 550468, 365193, 516756, 187571 |
| 1393 | 352838, 539133 | Which is better when working as a contractor, 1099 or incorporating? | 220022, 586026, 234436, 352640, 32072, 68524, 578196, 277812, 532932, 506108 |
| 9733 | 110163, 38655 | Due Diligence - Dilution? | 301880, 316321, 23414, 526073, 121262, 135798, 108965, 267266, 154841, 419505 |
| 7311 | 323768 | Finance, Social Capital IPOA.U | 507841, 479752, 290325, 583646, 199746, 579110, 419735, 264740, 261231, 327911 |
| 744 | 566480 | What options are available for a home loan with poor credit but a good rental history? | 310790, 67066, 490443, 573276, 415425, 80607, 92397, 313623, 289231, 289450 |
| 7141 | 132288 | Do investors go long option contracts when they cannot cover the exercise of the options? | 538054, 570046, 243714, 507828, 255927, 288289, 41967, 383328, 44530, 388571 |
| 4071 | 129875 | If our economy crashes, and cash is worthless, should i buy gold or silver | 524142, 291862, 505136, 473965, 487817, 502634, 566669, 53538, 506780, 308332 |
| 7512 | 191060 | understanding the process/payment of short sale dividends | 222320, 487329, 568166, 298284, 115553, 202985, 13631, 480949, 409432, 527636 |
| 1391 | 562176 | How is taxation for youtube/twitch etc monetization handled in the UK? | 254151, 267067, 510599, 223170, 77171, 527951, 440745, 454208, 266229, 111131 |
| 7534 | 358125 | Can you explain why it's better to invest now rather than waiting for the market to dip? | 175821, 145539, 426157, 114806, 310218, 89714, 33155, 474006, 419747, 103622 |
| 5356 | 312405 | Historical stock prices: Where to find free / low cost data for offline analysis? | 279785, 535343, 240086, 560108, 529877, 596106, 546379, 47798, 189341, 537111 |
| 2579 | 432020 | What to do when a job offer is made but with a salary less than what was asked for? | 423070, 200946, 524471, 181213, 432808, 256802, 157919, 364159, 190077, 559900 |
| 7823 | 583549 | Retirement Funds: Betterment vs Vanguard Life strategy vs Target Retirement | 451196, 105666, 175927, 331492, 268731, 347825, 172336, 293679, 57070, 11094 |
| 689 | 411044 | Receive credit card payment sending my customer details to a credit card processing company? | 446932, 63366, 104079, 421803, 171761, 438032, 195852, 553418, 567201, 96547 |
| 9174 | 535317, 160218 | Which U.S. online discount broker is the best value for money? | 192910, 236931, 200052, 563334, 31936, 513281, 515144, 413856, 451729, 47579 |
| 6867 | 443804, 540799, 445258 | Will there always be somebody selling/buying in every stock? | 466143, 301985, 230343, 482739, 18532, 226197, 429196, 61006, 349147, 543589 |
| 2383 | 232199 | Should I Purchase Health Insurance Through My S-Corp | 17215, 224406, 527620, 546634, 327232, 476085, 457034, 423074, 308255, 521489 |
| 5083 | 138845 | Co-signer deceased | 369075, 18257, 270952, 447983, 273759, 305509, 453263, 518681, 495482, 334606 |
| 10526 | 39185 | What extra information might be obtained from the next highest bids in an order book? | 546493, 485973, 283008, 427747, 467852, 251100, 298551, 138830, 146125, 322798 |
| 2181 | 376631, 397329 | What are the risks & rewards of being a self-employed independent contractor / consultant vs. being a permanent emplo... | 37725, 488755, 584218, 139501, 383088, 176777, 525360, 392124, 406656, 524788 |
| 5903 | 231863 | Fees aside, what factors could account for performance differences between U.S. large-cap index ETFs? | 408524, 159471, 395842, 246996, 372233, 20504, 230997, 501153, 14185, 402091 |
| 5620 | 448784, 329552, 548740 | What's the fuss about identity theft? | 260580, 158285, 90632, 598801, 551747, 98993, 91986, 423809, 5860, 581889 |
| 2472 | 370334 | How do I deal with a mistaken attempt to collect a debt from me that is owed by someone else? | 180601, 201758, 49321, 161422, 200263, 435006, 330507, 500671, 595441, 62109 |
| 2306 | 315875 | To whom should I report fraud on both of my credit cards? | 581889, 531137, 596284, 289706, 90632, 164729, 412542, 270449, 298729, 125204 |
| 7633 | 197839 | Can a trade happen “in between” the bid and ask price? | 494727, 281844, 353396, 402482, 137175, 554207, 179258, 164008, 394244, 458933 |
| 2400 | 564271 | Will I be paid dividends if I own shares? | 1198, 456470, 95889, 97942, 1034, 311214, 501931, 587689, 365627, 481169 |
| 5549 | 286227, 309361 | Pros / cons of being more involved with IRA investments [duplicate] | 429106, 105468, 336394, 561636, 181624, 546150, 87260, 90294, 324012, 396852 |
| 3801 | 307776 | Can a bunch of wealthy people force Facebook to go public? | 390529, 69017, 209242, 264498, 168565, 171236, 371293, 394734, 92014, 570634 |
| 4605 | 453941 | If the U.S. defaults on its debt, what will happen to my bank money? | 41312, 313306, 400826, 229310, 169691, 526384, 479527, 581054, 598030, 354896 |
| 2885 | 367360, 359579, 85229, 454810 | Merits of buying apartment houses and renting them | 430672, 451849, 159403, 581251, 502291, 80838, 343917, 358687, 502514, 507029 |
| 6110 | 94117, 259706 | Why does short selling require borrowing? | 188531, 320450, 226496, 67107, 384252, 35500, 79764, 107045, 84761, 501984 |
| 8 | 566392 | How to deposit a cheque issued to an associate in my business into my business account? | 65404, 508754, 261856, 590102, 564553, 308938, 188893, 309023, 301833, 489199 |
| 1309 | 156162, 489401 | Why does FlagStar Bank harass you about payments within grace period? | 471630, 489368, 271040, 336792, 15824, 438869, 329817, 173919, 526989, 75108 |
| 7109 | 447781 | How do I analyse moving averages? | 489933, 140804, 42620, 227669, 257185, 193012, 221627, 565501, 518932, 35006 |
| 5080 | 256055 | Is there a standard or best practice way to handle money from an expiring UTMA account? | 445521, 279291, 414429, 451189, 69841, 236186, 490991, 267206, 324564, 470928 |
| 4981 | 247894 | Where can I find open source portfolio management software? | 102684, 45218, 259463, 587792, 557861, 232736, 200683, 81865, 196432, 226628 |
| 7445 | 153178, 104343 | IS it the wrong time to get into the equity market immediately after large gains? | 89714, 573612, 590902, 483025, 350068, 33155, 284075, 127160, 356623, 79111 |
| 2895 | 328691 | Where should a young student put their money? | 426461, 354551, 332749, 148453, 55841, 502170, 5188, 256055, 517313, 496170 |
| 6787 | 587120 | Would it make sense to sell a stock, then repurchase it for tax purposes? | 219762, 23217, 400730, 263751, 474981, 221715, 390864, 17184, 106104, 374867 |
| 6041 | 241308 | Most effective Fundamental Analysis indicators for market entry | 425020, 81655, 96910, 224695, 528034, 108579, 194240, 263464, 331530, 542765 |
| 7700 | 273761, 2653, 179328 | Should I re-allocate my portfolio now or let it balance out over time? | 224392, 269169, 253268, 22221, 395208, 131127, 434014, 422051, 441176, 126836 |
| 547 | 6349 | What percentage of my company should I have if I only put money? | 396694, 68088, 523158, 95243, 156747, 80913, 498681, 213399, 445353, 559522 |
| 3394 | 342258, 129319 | What is the easiest way to back-test index funds and ETFs? | 172374, 71230, 528034, 448745, 408524, 159471, 445971, 99568, 224765, 571913 |
| 4102 | 448699 | How can I determine if my rate of return is “good” for the market I am in? | 597437, 554734, 135176, 369439, 554237, 535737, 88801, 484688, 461082, 162488 |

### Lowest MRR@10

| Query ID | MRR@10 | Recall@100 | First Relevant Rank | Query |
| --- | ---: | ---: | ---: | --- |
| 10462 | 0.0000 | 0.0000 | missing | Is it okay to be married, 30 years old and have no retirement? |
| 10482 | 0.0000 | 0.0000 | missing | Rollover into bond fund to do dollar cost averaging [duplicate] |
| 1085 | 0.0000 | 0.0000 | missing | How do disputed debts work on credit reports? |
| 10975 | 0.0000 | 0.0000 | missing | How to contribute to Roth IRA when income is at the maximum limit & you have employer-sponsored 401k plans? |
| 1159 | 0.0000 | 0.0000 | missing | what is the best way to do a freelancing job over the summer for a student |
| 1321 | 0.0000 | 0.0000 | missing | Are social media accounts (e.g. YouTube, Twitter, Instagram, etc.) considered assets? |
| 1783 | 0.0000 | 0.0000 | missing | Freelancing Tax implication |
| 1920 | 0.0000 | 0.0000 | missing | Clarification on student expenses - To file the tax for the next year |
| 2181 | 0.0000 | 0.0000 | missing | What are the risks & rewards of being a self-employed independent contractor / consultant vs. being a permanent emplo... |
| 34 | 0.0000 | 0.0000 | missing | 401k Transfer After Business Closure |
