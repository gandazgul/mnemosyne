# Mnemosyne BEIR fiqa Results

- Generated: 2026-06-21T22:57:39Z
- Dataset: `fiqa`
- Queries: 648
- Corpus documents: 57638
- Search limit: 100
- Source mode: `hybrid`
- Fusion: `vector-bm25`
- BM25 weight: `0.1`
- Rerank candidates: 150
- Rerank enabled: False
- Run label: `jina-v5-nano`
- Mnemosyne config: `/Users/gandazgul/Documents/web/mnemosyne/cli/configs/jina-v5-text-nano-retrieval.yaml`

| Metric | Score |
| --- | ---: |
| `ndcg@10` | 0.4859 |
| `mrr@10` | 0.5701 |
| `recall@10` | 0.5493 |
| `recall@100` | 0.7872 |
| `map@100` | 0.4251 |

## Breakdown

| Detail | Count |
| --- | ---: |
| `queries` | 648 |
| `queries_with_rank_1_hit` | 318 |
| `queries_with_top_10_hit` | 481 |
| `queries_with_top_100_hit` | 587 |
| `queries_missing_at_100` | 61 |

### First Relevant Rank

| Bucket | Queries |
| --- | ---: |
| `rank_1` | 318 |
| `rank_2_3` | 92 |
| `rank_4_10` | 71 |
| `rank_11_100` | 106 |
| `missing@100` | 61 |

### Missing At 100

| Query ID | Missing Doc IDs | Query | Top 10 Doc IDs |
| --- | --- | --- | --- |
| 4641 | 319954, 397358 | Where should I park my rainy-day / emergency fund? | 580025, 497993, 527939, 285812, 406219, 44594, 400503, 583695, 538023, 254572 |
| 3451 | 26292, 192307, 588448 | Should you keep your stocks if you are too late to sell? | 251536, 310683, 420974, 471911, 545284, 41852, 165970, 306460, 528518, 27015 |
| 753 | 243503 | Taxes due for hobbyist Group Buy | 466718, 132780, 451020, 203791, 283505, 466442, 547941, 172745, 299211, 293310 |
| 4465 | 376575 | How to donate to charity that will make a difference? | 106786, 132881, 174033, 90591, 46381, 427322, 379353, 266342, 326167, 174543 |
| 5951 | 497260 | Why can't house prices be out of tune with salaries | 599860, 418034, 31663, 259777, 374480, 62702, 596834, 373188, 529715, 589470 |
| 10827 | 160786, 7748, 107554, 95282 | How much should I be contributing to my 401k given my employer's contribution? | 290105, 296405, 140330, 436071, 555377, 452592, 576391, 312369, 242556, 38532 |
| 9771 | 263955, 28740 | Is there any emprical research done on 'adding to a loser' | 490479, 18671, 38704, 375290, 393272, 137898, 455631, 468717, 473652, 356175 |
| 3724 | 508921, 279570, 497216, 552887, 199970 | Should you always max out contributions to your 401k? | 302512, 3104, 122910, 430931, 273497, 576391, 135790, 488673, 497561, 43573 |
| 5853 | 476663, 160105, 495699, 424598 | Paying Off Principal of Home vs. Investing In Mutual Fund | 284318, 182612, 494148, 387722, 473647, 31525, 406325, 349355, 473865, 301224 |
| 570 | 363591 | Employer options when setting up 401k for employees | 532839, 117845, 79375, 289064, 387876, 150883, 301616, 15841, 555377, 508457 |
| 10122 | 273718 | Why diversify stocks/investments? | 259084, 297100, 144261, 459970, 89084, 508540, 502495, 564069, 153989, 331008 |
| 1783 | 332314 | Freelancing Tax implication | 421924, 14609, 156063, 159709, 179359, 445298, 383172, 549870, 223624, 15270 |
| 6004 | 149555 | Put-Call parity - what is the difference between the two representations? | 122432, 278373, 374797, 13260, 118762, 271109, 345410, 21768, 5257, 195152 |
| 4837 | 531841, 20958 | When applying for a mortgage, can it also cover outstanding debts? | 97925, 245005, 32749, 366448, 516578, 30311, 204035, 274870, 254454, 433171 |
| 4415 | 67676 | How much is inflation? | 117578, 206580, 468089, 267176, 501743, 381675, 513249, 290585, 59225, 306166 |
| 8378 | 125298 | Should I wait a few days to sell ESPP Stock? | 511678, 133644, 434812, 495568, 294573, 546125, 387035, 178684, 361345, 12232 |
| 8102 | 552707, 378173, 90294 | When do I sell a stock that I hold as a long-term position? | 306460, 165970, 35752, 537212, 557877, 310683, 203638, 171819, 219762, 250873 |
| 7928 | 118633 | If I believe a stock is going to fall, what options do I have to invest on this? | 427808, 410404, 501504, 260384, 173374, 47053, 24563, 480967, 30305, 137073 |
| 4777 | 590710 | How to finance necessary repairs to our home in order to sell it? | 365342, 52351, 570318, 416382, 478413, 482638, 487179, 26837, 293464, 432961 |
| 620 | 331332 | Is it wise to have plenty of current accounts in different banks? | 543921, 467059, 572848, 489959, 180673, 164801, 146317, 38628, 535817, 244097 |
| 8271 | 415511 | Income in zero-interest environment | 376709, 249558, 52047, 380368, 119298, 83330, 106424, 546874, 378437, 503040 |
| 864 | 211364, 152072 | Why use accounting software like Quickbooks instead of Excel spreadsheets? | 30142, 472924, 566337, 24890, 2436, 157751, 78117, 329774, 222380, 402174 |
| 5888 | 540806 | Interest charges on balance transfer when purchases are involved | 263647, 125497, 543776, 336792, 579601, 545327, 358445, 429746, 490529, 231105 |
| 2568 | 388798, 127353 | How to pay with cash when car shopping? | 346042, 9146, 15696, 355310, 108739, 166314, 514238, 269898, 258247, 214749 |
| 7124 | 74615 | How come we can find stocks with a Price-to-Book ratio less than 1? | 558617, 526110, 154725, 583708, 278582, 533818, 504243, 157597, 192686, 123263 |
| 2857 | 295864 | I have around 60K $. Thinking about investing in Oil, how to proceed? | 233732, 127566, 501384, 542051, 76996, 49235, 474575, 117451, 379140, 58065 |
| 8539 | 196304, 396038 | Can the risk of investing in an asset be different for different investors? | 283074, 483123, 391861, 502495, 471817, 385881, 387515, 499166, 500863, 292609 |
| 1085 | 467737, 393710 | How do disputed debts work on credit reports? | 372039, 450031, 268777, 319276, 161422, 398258, 78328, 384608, 1218, 339365 |
| 5422 | 151973 | What are some good books for learning stocks, bonds, derivatives e.t.c for beginner with a math background? | 165294, 172587, 276786, 273906, 191688, 552371, 221319, 193555, 79517, 241423 |
| 8247 | 321114, 465313 | Tax on Stocks or ETF's | 586010, 161019, 153112, 518735, 190687, 474745, 350317, 437907, 195767, 580802 |
| 10482 | 549072 | Rollover into bond fund to do dollar cost averaging [duplicate] | 330023, 447567, 18436, 434466, 439757, 211765, 564787, 224782, 175576, 265817 |
| 8789 | 70853 | What does “profits to the shareholders jumped to 15 cents a share” mean? | 87349, 41912, 341424, 573079, 20076, 144079, 317363, 14870, 325669, 234040 |
| 2423 | 538023 | At what age should I start or stop saving money? | 529444, 417787, 396127, 234846, 272328, 337561, 104457, 328157, 553288, 235855 |
| 8475 | 293320 | Why I cannot find a “Pure Cash” option in 401k investments? | 494655, 481728, 104793, 240562, 437427, 444107, 525426, 363652, 332152, 554739 |
| 4523 | 594257 | What should I do with my $25k to invest as a 20 years old? | 129255, 465819, 171712, 332022, 332203, 272070, 216365, 72578, 10476, 286746 |
| 8513 | 270573 | Buy on dip when earnings fail? | 572622, 175821, 203873, 351396, 335626, 53431, 534734, 391043, 73872, 494877 |
| 5054 | 28119 | How to stress test an investment plan? | 588481, 101124, 377186, 127263, 458183, 220665, 263390, 582899, 403916, 534010 |
| 1159 | 496064 | what is the best way to do a freelancing job over the summer for a student | 55064, 347181, 163881, 460648, 237282, 594182, 469972, 449155, 112669, 132287 |
| 68 | 19183 | Intentions of Deductible Amount for Small Business | 519473, 97719, 86134, 70315, 54333, 146657, 462831, 192516, 97233, 18850 |
| 9701 | 387141, 357739 | How to bet against the London housing market? | 473883, 408865, 258048, 225682, 516214, 108399, 588608, 70799, 412502, 367103 |
| 6199 | 239214 | How can all these countries owe so much money?  Why & where did they borrow it from? | 414693, 584273, 47163, 10399, 169921, 351853, 49602, 490042, 298794, 370186 |
| 9126 | 514831 | Short an option - random assignment? | 334473, 477588, 308859, 82194, 228810, 469382, 267113, 349974, 307518, 277760 |
| 34 | 599545 | 401k Transfer After Business Closure | 458917, 551545, 490867, 483268, 144109, 226547, 168890, 424766, 104134, 387876 |
| 6395 | 166227 | Option settlement for calendar spreads | 584223, 22916, 467463, 401447, 276314, 505223, 111301, 516790, 8177, 189144 |
| 9115 | 207325 | Why does the calculation for percentage profit vary based on whether a position is short vs. long? | 245082, 158520, 419897, 422467, 428786, 314478, 385220, 330041, 99131, 232880 |
| 4767 | 280805, 568670, 224057, 22804 | New car: buy with cash or 0% financing | 420018, 429153, 306834, 584106, 56867, 451092, 9146, 166314, 408932, 395590 |
| 8351 | 472516 | What happens when a calendar spread is assigned in a non-margin account? | 102316, 141213, 45674, 516790, 273142, 23609, 527654, 111301, 100628, 272754 |
| 6131 | 381720, 170204, 416679, 2460 | Is it ever a good idea to close credit cards? | 326094, 368806, 391384, 339030, 334111, 218088, 533288, 99449, 35625, 258465 |
| 858 | 122485, 278450 | Is it bad practice to invest in stocks that fluctuate by single points throughout the day? | 293027, 30774, 146632, 433730, 573612, 241860, 208932, 519501, 514780, 5572 |
| 4019 | 6881, 125477 | How and Should I Invest (As a college 18 year old with minimal living expenses)? | 85977, 332938, 332749, 379948, 269671, 426461, 571044, 66864, 20304, 10476 |
| 6080 | 164513 | Is ScholarShare a legitimate entity for a 529 plan in California? | 22856, 236732, 201500, 277581, 468527, 115175, 233401, 83080, 2809, 535357 |
| 6959 | 205010 | What is the term for the quantity (high price minus low price) for a stock? | 468025, 428117, 303325, 169954, 577573, 599523, 229573, 317365, 472537, 115134 |
| 5402 | 491350 | Is it impossible to get a home loan with a poor credit history after a divorce? | 445163, 51728, 90579, 227485, 595029, 44105, 52250, 129177, 67066, 122807 |
| 6562 | 501157 | Cheapest way to “wire” money in an Australian bank account to a person in England, while I'm in Laos? | 473605, 582414, 60446, 282744, 183880, 182443, 324817, 135675, 549684, 177714 |
| 701 | 389446 | What are the ins/outs of writing-off part of one's rent for working at home? | 231990, 436505, 349672, 344955, 337706, 456234, 339488, 177074, 124507, 243306 |
| 7674 | 519390 | Choosing the limit when making a limit order? | 249279, 447886, 15917, 155151, 514841, 278630, 94653, 437436, 31933, 184756 |
| 5940 | 486243 | How does investment into a private company work? | 512609, 46842, 250354, 454465, 535314, 473154, 530, 182226, 252853, 135411 |
| 6612 | 205522, 322900 | If I have a lot of debt and the housing market is rising, should I rent and slowly pay off my debt or buy and roll th... | 502594, 301192, 180192, 431481, 254454, 198442, 290900, 433171, 115408, 473427 |
| 8456 | 486333 | What typically happens to unvested stock during an acquisition? | 257853, 534755, 93215, 174321, 469036, 391156, 555276, 186869, 475019, 235779 |
| 10213 | 270221, 545712 | Looking for good investment vehicle for seasonal work and savings | 446186, 38269, 100517, 112499, 195373, 223551, 518664, 390556, 488100, 94680 |
| 5196 | 172128, 114829 | I might use a credit card convenience check. What should I consider? | 565745, 393866, 85252, 402543, 289483, 456098, 302823, 2875, 85517, 481052 |
| 3006 | 269851, 568473, 328300 | Strategies for putting away money for a child's future (college, etc.)? | 512096, 127838, 258704, 471019, 211713, 597627, 490382, 303432, 372900, 290441 |
| 3909 | 312248, 404356, 245616, 353028 | How to rescue my money from negative interest? | 514003, 83330, 61586, 472837, 328499, 362730, 202527, 404352, 574011, 42475 |
| 6907 | 251604 | Nominal value of shares | 303112, 69506, 170652, 487738, 4854, 111827, 91870, 480515, 21786, 481761 |
| 5464 | 86691 | Resources on Buying Rental Properties | 222095, 372274, 423438, 325722, 383921, 315972, 26339, 536126, 294549, 155964 |
| 10034 | 480749 | Tax implications of holding EWU (or other such UK ETFs) as a US citizen? | 528880, 181942, 565296, 447197, 141585, 430868, 197478, 562007, 44955, 104128 |
| 5090 | 436493 | Should I take a student loan to pursue my undergraduate studies in France? | 12988, 246286, 92430, 217831, 560681, 21913, 586289, 455666, 213328, 287507 |
| 9391 | 503637 | Should I replace bonds in a passive investment strategy | 535518, 136515, 248158, 107424, 283202, 155242, 342485, 577832, 112369, 171669 |
| 3148 | 178127, 438000 | Can a car company refuse to give me a copy of my contract or balance details? | 172855, 584305, 430100, 29721, 357280, 65046, 92888, 164702, 560325, 395995 |
| 4678 | 305153 | Finance, Cash or Lease? | 427884, 185405, 311748, 376016, 215225, 504918, 311446, 85373, 260095, 487678 |
| 2398 | 224654, 590489 | Frustrated Landlord | 556453, 44058, 96538, 487094, 168089, 98372, 393883, 310992, 422579, 112535 |
| 5511 | 169893, 560325, 383193, 278699, 12746 | Pay off car loan entirely or leave $1 until the end of the loan period? | 334559, 529123, 38786, 155843, 500946, 324269, 107898, 139788, 51873, 329137 |
| 8834 | 12232 | Pros/Cons of Buying Discounted Company Stock | 203139, 599156, 569303, 569224, 371821, 528827, 133644, 469100, 349567, 67625 |
| 988 | 226053, 107688 | Where should I invest my savings? | 501384, 168402, 571218, 60093, 223551, 450558, 223872, 82482, 320675, 194080 |
| 3369 | 231012, 145716, 411910, 395840 | Why should one only contribute up to the employer's match in a 401(k)? | 296405, 555377, 341493, 15841, 436071, 242556, 92370, 240373, 24231, 240259 |
| 9245 | 194561 | Stock Options for a company bought out in cash and stock | 207253, 186869, 39345, 178497, 131488, 248393, 265111, 472516, 259560, 261487 |
| 5763 | 462019 | What is the best way to get a “rough” home appraisal prior to starting the refinance process? | 570318, 326214, 38712, 218144, 89964, 515361, 67379, 251466, 11572, 331255 |
| 4962 | 599925 | Net Cash Flows from Selling the Bond and Investing | 158363, 416839, 52149, 537603, 408661, 308276, 535518, 395208, 196173, 34949 |
| 4846 | 151104 | Is there anything comparable to/resembling CNN's Fear and Greed Index? | 335892, 538974, 98096, 489352, 320059, 415161, 183597, 533311, 3533, 309245 |
| 9403 | 6666, 328086, 345199 | Abundance of Cash - What should I do? | 410450, 570632, 14349, 159235, 103447, 551986, 186332, 215296, 357887, 372223 |
| 5993 | 367375, 272866, 55084, 352638, 426120, 63501 | Why would anyone want to pay off their debts in a way other than “highest interest” first? | 94373, 160193, 287571, 416796, 128574, 156195, 494306, 353911, 886, 431212 |
| 5710 | 232311 | Bucketing investments to track individual growths | 227364, 516267, 88417, 534323, 227733, 412830, 508610, 472328, 11979, 355108 |
| 7529 | 66607 | Does the expense ratio of a fund-of-funds include the expense ratios of its holdings? | 514529, 464337, 89297, 293626, 287537, 59249, 218261, 518402, 387980, 451855 |
| 5021 | 589285 | Is there a more flexible stock chart service, e.g. permitting choice of colours when comparing multiple stocks? | 528576, 584801, 189341, 465971, 60284, 211444, 555506, 511861, 49893, 517935 |
| 3612 | 259625 | How can I buy and sell the same stock on the same day? | 567383, 522658, 165548, 310636, 429418, 390864, 402726, 483676, 584291, 460937 |
| 4409 | 499128, 100306, 147439 | My friend wants to put my name down for a house he's buying. What risks would I be taking? | 243732, 223841, 102088, 268078, 360682, 102326, 514790, 84732, 244278, 426676 |
| 2070 | 363678 | Advantage of credit union or local community bank over larger nationwide banks such as BOA, Chase, etc.? | 550303, 578357, 469515, 587737, 597571, 249839, 590209, 30253, 38038, 261389 |
| 11039 | 53544, 249063 | Pay off credit card debt or earn employer 401(k) match? | 91183, 287876, 552383, 508534, 5203, 79363, 163287, 345895, 105557, 281049 |
| 5460 | 184337, 21174, 108514, 463885 | Paying off a loan with a loan to get a better interest rate | 77052, 327115, 470716, 529418, 106495, 344812, 343208, 58432, 482798, 522341 |
| 7925 | 318185, 402482 | Can I sell a stock immediately? | 438974, 584291, 591436, 407602, 219762, 133644, 310636, 238215, 44461, 581579 |
| 4286 | 566069 | Given advice “buy term insurance and invest the rest”, how should one “invest the rest”? | 70460, 229239, 511386, 391243, 206830, 10531, 155640, 79142, 564675, 151817 |
| 2685 | 154113, 370300, 303293, 468923 | What ways are there for us to earn a little extra side money? | 382005, 468086, 594182, 269380, 186889, 4992, 558832, 237950, 576047, 543275 |
| 1090 | 518896 | Need a formula to determine monthly payments received at time t if I'm reinvesting my returns | 446454, 393987, 296146, 573928, 179365, 520217, 19999, 281329, 16051, 521590 |
| 6122 | 44344 | Better to rent condo to daughter or put her on title? | 496166, 403515, 316794, 558251, 182039, 118246, 53840, 577658, 423398, 566184 |
| 4514 | 69485, 337764, 209804 | What intrinsic, non-monetary value does gold have as a commodity? | 471825, 426270, 156211, 317429, 99089, 146573, 408336, 240894, 80141, 532381 |
| 8507 | 370995 | When to sell a stock? | 251536, 99132, 303724, 545284, 272091, 236415, 420974, 217837, 88813, 368348 |
| 6221 | 257248, 76414, 455614, 115717 | To pay off a student loan, should I save up a lump sum payoff payment or pay extra each month? | 448791, 110081, 352363, 254245, 414534, 394474, 541313, 274108, 69150, 414288 |
| 3008 | 180192, 323406 | What are my chances at getting a mortgage with Terrible credit but High income | 231688, 47441, 102266, 251846, 407401, 44105, 227485, 285694, 455952, 78176 |
| 4007 | 521657 | What is a reasonable salary for the owner and sole member of a small S-Corp? | 556220, 260385, 205341, 370542, 170933, 388704, 515233, 315552, 543085, 521933 |
| 6644 | 175035 | How to know precisely when a SWIFT is issued by a bank? | 110198, 475527, 218761, 271596, 41383, 118396, 355870, 327623, 39783, 187497 |
| 10267 | 460398 | How should I prepare for the next financial crisis? | 178693, 143393, 569632, 326398, 87520, 305600, 436091, 182442, 96017, 279456 |
| 7622 | 253369, 378594 | Best way to pay off debt? | 220241, 353911, 457945, 271525, 480773, 157923, 388095, 416796, 115499, 373554 |
| 3767 | 153922, 392060 | What should I be doing to protect myself from identity theft? | 90632, 423809, 260580, 97686, 125204, 368679, 171510, 581889, 158008, 587778 |
| 6410 | 471723 | Will an ETF immediately reflect a reconstitution of underlying index | 454610, 71230, 87261, 227324, 214281, 147282, 200360, 295993, 408524, 428187 |
| 5030 | 215540 | Why pay for end-of-day historical prices? | 13511, 227192, 149420, 295344, 560108, 327974, 532178, 378994, 347992, 471131 |
| 6252 | 394551, 160932, 293624, 233294, 243268, 379487, 62868 | Is this mortgage advice good, or is it hooey? | 213713, 120061, 473647, 47565, 139366, 104988, 205906, 316149, 272634, 443852 |
| 885 | 337165, 409184 | How long do credit cards keep working after you disappear? | 516678, 254968, 89888, 472336, 251643, 99449, 301792, 251701, 588719, 59228 |
| 4031 | 115741 | 28 years old and just inherited large amount of money and real estate - unsure what to do with it | 318864, 140002, 568629, 65180, 375708, 193171, 163197, 588316, 266481, 80844 |
| 766 | 550172 | Will the ex-homeowner still owe money after a foreclosure? | 2996, 552768, 299591, 163711, 27987, 333583, 104955, 268865, 56126, 427110 |
| 8202 | 513258, 93971 | What accounted for DXJR's huge drop in stock price? | 457689, 122542, 337001, 317363, 162047, 5220, 537862, 41271, 67237, 363451 |
| 7345 | 237645 | What do these numbers mean? (futures) | 9274, 508821, 527080, 206895, 460331, 276381, 366526, 108, 164001, 529996 |
| 776 | 583640, 127263, 597880, 496899 | Can saving/investing 15% of your income starting age 25, likely make you a millionaire? | 124027, 10440, 417787, 467044, 143591, 41960, 191148, 554833, 374266, 434972 |
| 89 | 248624 | How can I deposit a check made out to my business into my personal account? | 508754, 309023, 526817, 188893, 308938, 135196, 29372, 590102, 65404, 521540 |
| 8013 | 496159, 224231 | Frequency of investments to maximise returns (and minimise fees) | 384983, 81652, 537626, 48678, 57033, 446948, 224816, 388389, 278607, 24846 |
| 3759 | 527966, 67167, 522358 | Simplifying money management | 455457, 490065, 373772, 214248, 248663, 372743, 122378, 145812, 526159, 231412 |
| 7295 | 244749 | Selling non-dividend for dividend stocks | 542667, 292559, 501931, 71511, 352484, 437907, 478291, 472470, 217837, 517577 |
| 10639 | 431799, 495774, 278453, 187039 | Short term parking of a large inheritance? | 171196, 163353, 318864, 235628, 375708, 111048, 163197, 140002, 590276, 405115 |
| 6635 | 156358 | Why don't share prices of a company rise every other Friday when the company buys shares for its own employees? | 587137, 3656, 533712, 343452, 235531, 579037, 12560, 125298, 95806, 569224 |
| 4312 | 399149 | Is it true that 90% of investors lose their money? | 282435, 222639, 167950, 532485, 116647, 497786, 285945, 170628, 431735, 300770 |
| 6525 | 181985 | Does it make sense to trade my GOOGL shares for GOOG and pocket the difference? | 106541, 550661, 98150, 105542, 221795, 378906, 362473, 192696, 498014, 156467 |
| 2590 | 589625 | Are non-residents or foreigners permitted to buy or own shares of UK companies? | 296528, 48269, 209493, 55999, 158923, 310103, 458730, 307776, 456999, 188776 |
| 5374 | 152688 | What were the main causes of the spike and drop of DRYS's stock price? | 283106, 133204, 457689, 122542, 253339, 261522, 317363, 589533, 5220, 380894 |
| 2994 | 419319 | Work on the side for my wife's company | 510317, 569145, 321743, 138428, 460721, 491844, 5840, 562780, 523372, 506991 |
| 3683 | 454501 | Can I trust the Motley Fool? | 276975, 105973, 408995, 428848, 500338, 538086, 565016, 522713, 6607, 192912 |
| 7206 | 441155, 532211, 553066 | Who Bought A Large Number Of Shares? | 65667, 351570, 558703, 34882, 573846, 350214, 444752, 358164, 54424, 498676 |
| 10246 | 77573 | Understanding the T + 3 settlement days rule | 370635, 156029, 176717, 327080, 340263, 28314, 293389, 332243, 11927, 315205 |
| 5241 | 322157, 27489 | Mortgage vs. Cash for U.S. home buy now | 344740, 281675, 111184, 213713, 438073, 390976, 309420, 273735, 107350, 234286 |
| 98 | 575929 | How can I make $250,000.00 from trading/investing/business within 5 years? | 527522, 555630, 519619, 66034, 438279, 373119, 336661, 449745, 506149, 102113 |
| 4615 | 262934 | Are solar cell panels and wind mills worth the money? | 261900, 69523, 496427, 455798, 120384, 425595, 19179, 510872, 271015, 385028 |
| 3264 | 598807, 134764 | Pros and Cons of Interest Only Loans | 486525, 260383, 316230, 400009, 7243, 453847, 526169, 106424, 119298, 25466 |
| 6467 | 453256, 23217, 346641, 367313 | Advice on strategy for when to sell | 88813, 240089, 217837, 368348, 203873, 99857, 109455, 250873, 303724, 251536 |
| 4289 | 24881 | Does the currency exchange rate contain any additional information at all? | 288330, 114886, 226102, 17469, 416975, 324546, 517345, 356465, 256996, 104480 |
| 4394 | 441582 | Transfer $50k to another person's account (in California, USA) | 322838, 93386, 431462, 305907, 462585, 415655, 412258, 293653, 524752, 521753 |
| 4047 | 77652 | Does doing your “research”/“homework” on stocks make any sense? | 417840, 324779, 452175, 110733, 170966, 284165, 210241, 358846, 474296, 312445 |
| 7344 | 108403 | How is the Dow divisor calculated? | 159166, 14368, 150430, 378974, 253926, 65618, 313421, 591089, 501032, 96697 |
| 6875 | 224392 | Where to find free Thailand stock recommendations and research? | 567500, 110733, 352557, 224366, 79337, 512895, 77502, 232460, 556770, 249679 |
| 10447 | 152096, 300721 | Is there an advantage to a traditional but non-deductable IRA over a taxable account? [duplicate] | 144751, 500175, 382236, 532657, 259150, 447482, 382894, 114912, 540389, 588134 |
| 5782 | 595455 | Pay off credit cards in one lump sum, or spread over a few months? | 487621, 262026, 172084, 117007, 273631, 440620, 114592, 529312, 559523, 113822 |
| 9871 | 448890, 40051, 170594 | What should I do with the 50k I have sitting in a European bank? | 367207, 73741, 292714, 433003, 387723, 231521, 74668, 76562, 219477, 212464 |
| 3625 | 384469, 414295 | What should I do with my paper financial documents? | 509617, 500751, 569812, 380263, 513248, 113830, 344236, 163168, 63422, 305065 |
| 6005 | 478457, 345895, 73310, 384626, 390689 | Why might it be advisable to keep student debt vs. paying it off quickly? | 149500, 571198, 507544, 96268, 431884, 25190, 572272, 564206, 422704, 414288 |
| 3822 | 385090, 418900, 308837 | How to change a large quantity of U.S. dollars into Euros? | 292714, 417917, 541608, 239876, 194730, 174406, 478065, 208499, 144181, 305907 |
| 7879 | 372551, 421285 | Any Tips on How to Get the Highest Returns Within 4 Months by Investing in Stocks? | 58186, 102029, 272174, 43088, 174313, 540919, 105391, 599217, 367415, 160170 |
| 3115 | 234950, 389028, 316794 | How can I live outside of the rat race of American life with 300k? | 233562, 183869, 267892, 129364, 174272, 366961, 369742, 559852, 220023, 136035 |
| 3995 | 278734 | I have more than $250,000 in a US Bank account… mistake? | 485507, 171720, 404954, 479918, 264934, 14349, 505461, 303367, 21420, 403288 |
| 10136 | 526115 | How to minimise the risk of a reduction in purchase power in case of Brexit for money held in a bank account? | 466950, 290930, 417740, 304007, 583903, 265453, 168137, 205685, 393823, 150543 |
| 8635 | 67107, 240215 | Is there any flaw in this investment scheme? | 46818, 575918, 493841, 447619, 303619, 151238, 365816, 202638, 203729, 480802 |
| 5206 | 563030, 28230, 201982, 117276, 300660 | Is it a good idea to get an unsecured loan to pay off a credit card that won't lower a high rate? | 298908, 595455, 225522, 69938, 340520, 153088, 2519, 375780, 498728, 504293 |
| 2713 | 388147 | Physical Checks - Mailing | 284528, 29372, 41944, 584170, 216200, 20791, 199069, 564301, 229546, 78139 |
| 9060 | 40447 | Buying puts without owning underlying | 511093, 528052, 228217, 345851, 7743, 359778, 181924, 338782, 3062, 118762 |
| 4105 | 25096 | As an investor what are side effects of Quantitative Easing in US and in EU? | 416483, 345910, 176262, 293104, 305029, 108519, 369038, 246687, 185300, 30946 |
| 2465 | 570680, 81046 | Can capital expenses for volunteer purposes be deducted from income? | 37382, 432545, 398536, 202645, 510716, 146657, 18889, 310612, 107213, 541809 |
| 4640 | 101369 | What can my relatives do to minimize their out of pocket expenses on their fathers estate | 295246, 331534, 372808, 356035, 360816, 17110, 422994, 144965, 375708, 367404 |
| 9275 | 338754 | Do I have to pay a capital gains tax if I rebuy the same stock within 30 days? | 400730, 343219, 390864, 376800, 23217, 537916, 102443, 448659, 407602, 263751 |
| 3500 | 174019 | Why invest in becoming a landlord? | 273187, 528206, 557478, 71424, 422331, 141935, 578597, 159156, 423438, 112535 |
| 1306 | 484437, 204075 | I made an investment with a company that contacted me, was it safe? | 309851, 594206, 160611, 538086, 205665, 537698, 450779, 567973, 407663, 112259 |
| 6262 | 26799 | Help required on estimating SSA benefit amounts | 34538, 390877, 83338, 2648, 15322, 118707, 498444, 192585, 430407, 320362 |
| 4500 | 533623 | What to ask Warren Buffet at the Berkshire Hathaway shareholder meeting? | 153649, 113800, 461355, 375671, 100283, 41809, 404341, 245168, 39149, 3533 |
| 8632 | 213976 | Is it best to exercise options shares when they vest, or wait | 43497, 237718, 340730, 420722, 382381, 259560, 104188, 293959, 237783, 488207 |
| 6133 | 415705 | What happens to all of the options when they expire? | 7733, 575408, 11456, 73256, 132288, 581672, 242298, 358492, 428399, 481070 |
| 3771 | 488948, 198349, 217683, 49601 | Best way to buy Japanese yen for travel? | 521712, 495826, 490384, 96211, 306130, 128471, 434201, 350245, 152695, 217715 |
| 1736 | 25543, 443419 | How can people have such high credit card debts? | 399406, 437610, 475668, 562896, 372993, 517050, 298908, 569056, 174941, 246882 |
| 6814 | 340214, 223206 | Selling Stock - All or Nothing? | 590188, 513734, 400614, 878, 154976, 3095, 349147, 394454, 345368, 66834 |
| 1322 | 64138 | Is this follow-up after a car crash a potential scam? | 114231, 397852, 567973, 226090, 283917, 91463, 318941, 33914, 98356, 44635 |
| 5185 | 210236, 317354 | Invest in low cost small cap index funds when saving towards retirement? | 196992, 376485, 262180, 241202, 580313, 503725, 524525, 434279, 59670, 513474 |
| 6909 | 127012 | Why do stocks priced above $2.00 on the ASX sometimes move in $0.005 increments? | 72633, 118232, 490584, 47217, 298551, 64943, 450489, 112946, 168080, 43432 |
| 2348 | 211867, 566573, 211622, 474234, 352271, 265874 | Why can't you just have someone invest for you and split the profits (and losses) with him? | 447619, 247486, 389004, 306430, 420544, 381757, 151412, 309851, 177194, 64410 |
| 4499 | 323363, 76996 | Is investing exlusively in a small-cap index fund a wise investment? | 196992, 501153, 517391, 52274, 235917, 238963, 14748, 372233, 335136, 513818 |
| 3530 | 239998 | How to exclude stock from mutual fund | 184299, 24029, 378075, 209879, 370754, 449124, 479420, 574383, 110343, 332152 |
| 659 | 584685 | Buying from an aggressive salesperson | 168796, 449079, 13139, 120279, 186206, 235925, 429119, 466122, 109852, 365240 |
| 9108 | 272021, 472585 | Starting an investment portfolio with Rs 5,000/- | 290757, 171189, 312821, 323067, 356552, 414116, 51848, 46967, 122679, 461193 |
| 6835 | 102243 | Are bond ETF capital gains taxed similar to stock or stock funds if held for more than 1 year? | 149305, 5710, 84238, 287950, 153112, 29502, 570546, 543842, 586010, 225536 |
| 3067 | 406156 | Should I make extra payments to my under water mortgage or increase my savings? | 477907, 560915, 476068, 468831, 423403, 90009, 131365, 341837, 3092, 440719 |
| 4125 | 344648, 72046 | Alternative means of salary for my employees | 174787, 73999, 365558, 58906, 36608, 302310, 245451, 414694, 287200, 361954 |
| 1150 | 43603, 19936 | How are the best way to make and save money at 22 years old | 10476, 529444, 353369, 319760, 595287, 328157, 433986, 66864, 204479, 417787 |
| 7431 | 372921 | Pay off mortgage or invest in high value saving account | 2393, 435576, 468831, 589256, 157414, 477907, 557734, 25877, 537721, 440719 |
| 9808 | 40702, 431946 | Selling To Close | 416307, 557582, 218423, 414636, 125079, 151587, 345368, 43087, 152719, 449280 |
| 3888 | 319213, 239632 | Why I can't view my debit card pre-authorized amounts? | 208169, 185434, 294077, 432077, 316652, 276733, 440527, 281129, 521010, 448086 |
| 10109 | 506374, 499849 | Why does Charles Schwab have a Mandatory Settlement Period after selling stocks? | 93231, 28314, 332243, 370635, 266725, 98302, 121465, 124188, 293389, 563826 |
| 2790 | 279329, 469125 | Should I pay more than 20% down on a home? | 472484, 400896, 207564, 357200, 75961, 385343, 64400, 100483, 487593, 330229 |
| 9882 | 65702 | Money-market or cash-type ETFs for foreigners with U.S brokerage account | 389581, 391876, 188524, 313775, 363378, 535340, 386173, 131059, 466845, 586151 |
| 4999 | 314898 | Looking for a good source for Financial Statements | 9938, 171964, 431459, 11263, 597241, 146076, 295738, 465971, 520165, 516379 |
| 3189 | 225395 | Diversify my retirement investments with a Roth IRA | 287225, 240975, 122222, 423658, 404800, 88311, 458168, 187124, 272840, 347651 |
| 5134 | 158523 | Why does Yahoo Finance's data for a Vanguard fund's dividend per share not match the info from Vanguard? | 46774, 532616, 374330, 584128, 206727, 215486, 405474, 221477, 239137, 559884 |
| 1321 | 216456, 292065 | Are social media accounts (e.g. YouTube, Twitter, Instagram, etc.) considered assets? | 381322, 392379, 203715, 235082, 481036, 210019, 559738, 100849, 351672, 426184 |
| 4539 | 370879 | How should I save money if the real interest rate (after inflation) is negative? | 42475, 275925, 472837, 194080, 203926, 32744, 479659, 61586, 503394, 449745 |
| 715 | 579763, 546538, 187404 | what would you do with $100K saving? | 548758, 133120, 113885, 427032, 203201, 200273, 260677, 556545, 337561, 328770 |
| 504 | 500755, 344203, 498751 | Have plenty of cash flow but bad credit | 22807, 546097, 93573, 368247, 486334, 99463, 503419, 495431, 41875, 432040 |
| 2296 | 83330, 366594 | How does a bank make money on an interest free secured loan? | 400009, 396853, 119298, 94230, 259919, 580147, 106424, 172303, 279897, 249831 |
| 10975 | 61022 | How to contribute to Roth IRA when income is at the maximum limit & you have employer-sponsored 401k plans? | 441632, 163865, 81148, 353009, 360533, 140330, 446226, 110114, 101490, 270818 |
| 1994 | 156640 | Does the IRS reprieve those who have to commute for work? | 434846, 231990, 192843, 51491, 585706, 263259, 63919, 147624, 380635, 319965 |
| 9164 | 365298, 263390 | Bonds vs equities: crash theory | 115648, 309326, 149900, 16924, 506743, 321941, 599420, 305274, 287656, 296516 |
| 10462 | 8266, 11378, 35680, 437879, 204035, 581204 | Is it okay to be married, 30 years old and have no retirement? | 66376, 268023, 391583, 122333, 152478, 593377, 139595, 51914, 221295, 361049 |
| 8855 | 208165, 312821 | How do i get into investing stocks [duplicate] | 155677, 312445, 403092, 560395, 313570, 484327, 555521, 328477, 142320, 367415 |
| 7071 | 124230 | ESPP strategy - Sell right away or hold? | 511678, 133644, 294573, 434812, 361345, 71713, 127702, 495568, 387035, 575213 |
| 3534 | 340329 | Why do dishonour fees exist? | 386745, 149845, 312349, 234260, 166431, 517497, 286064, 333149, 34108, 171374 |
| 8974 | 356595, 170625 | As a 22-year-old, how risky should I be with my 401(k) investments? | 134931, 216365, 10476, 102501, 452126, 140738, 336144, 96110, 124762, 327925 |
| 5178 | 240261 | Formula that predicts whether one is better off investing or paying down debt | 39819, 557506, 373554, 393833, 111815, 396889, 290434, 257016, 262772, 154449 |
| 5061 | 23747 | What fiscal scrutiny can be expected from IRS in early retirement? | 31699, 502150, 25481, 598378, 149742, 513392, 484683, 537049, 253735, 170916 |
| 2075 | 170042, 359580, 14967 | Are stories of turning a few thousands into millions by trading stocks real? | 519619, 44417, 523393, 65667, 506149, 555521, 375876, 285147, 41926, 519501 |
| 7533 | 162668, 93853 | Investing tax (savings) | 8012, 142658, 563986, 315105, 250603, 365193, 550468, 526664, 66717, 187571 |
| 1393 | 352838, 539133 | Which is better when working as a contractor, 1099 or incorporating? | 220022, 352640, 234436, 32072, 68524, 586026, 578196, 277812, 532932, 506108 |
| 9733 | 110163, 38655 | Due Diligence - Dilution? | 301880, 316321, 23414, 526073, 135798, 108965, 267266, 419505, 154841, 423260 |
| 7311 | 323768 | Finance, Social Capital IPOA.U | 479752, 507841, 290325, 579110, 419735, 583646, 584135, 199746, 327911, 264740 |
| 744 | 566480 | What options are available for a home loan with poor credit but a good rental history? | 490443, 67066, 310790, 415425, 80607, 92397, 313623, 289231, 47441, 573276 |
| 7141 | 132288 | Do investors go long option contracts when they cannot cover the exercise of the options? | 538054, 507828, 255927, 243714, 41967, 557356, 293767, 357324, 570046, 288289 |
| 4071 | 129875 | If our economy crashes, and cash is worthless, should i buy gold or silver | 524142, 505136, 291862, 487817, 473965, 506780, 502634, 53538, 308332, 302010 |
| 7512 | 191060 | understanding the process/payment of short sale dividends | 487329, 222320, 568166, 298284, 13631, 115553, 202985, 480949, 409432, 527636 |
| 1391 | 562176 | How is taxation for youtube/twitch etc monetization handled in the UK? | 254151, 510599, 527951, 266229, 223170, 21136, 440745, 77171, 315516, 454208 |
| 7534 | 358125 | Can you explain why it's better to invest now rather than waiting for the market to dip? | 175821, 145539, 114806, 310218, 89714, 33155, 103622, 426157, 474006, 350068 |
| 5356 | 312405 | Historical stock prices: Where to find free / low cost data for offline analysis? | 279785, 240086, 529877, 596106, 535343, 560108, 189341, 546379, 47798, 167586 |
| 2579 | 432020 | What to do when a job offer is made but with a salary less than what was asked for? | 423070, 200946, 524471, 190077, 432808, 256802, 283825, 364159, 181213, 157919 |
| 7823 | 583549 | Retirement Funds: Betterment vs Vanguard Life strategy vs Target Retirement | 451196, 105666, 175927, 331492, 268731, 172336, 347825, 293679, 57070, 329425 |
| 689 | 411044 | Receive credit card payment sending my customer details to a credit card processing company? | 446932, 63366, 553418, 421803, 104079, 567201, 195852, 438032, 171761, 96547 |
| 9174 | 535317, 160218 | Which U.S. online discount broker is the best value for money? | 236931, 192910, 200052, 563334, 31936, 544576, 515144, 513281, 405217, 358770 |
| 6867 | 443804, 445258 | Will there always be somebody selling/buying in every stock? | 466143, 301985, 482739, 429196, 321639, 349147, 560273, 230343, 400614, 224672 |
| 2383 | 232199 | Should I Purchase Health Insurance Through My S-Corp | 17215, 527620, 224406, 327232, 546634, 308255, 423074, 476085, 457034, 222665 |
| 5083 | 138845 | Co-signer deceased | 369075, 447983, 270952, 18257, 305509, 273759, 518681, 453263, 495482, 288701 |
| 10526 | 39185 | What extra information might be obtained from the next highest bids in an order book? | 546493, 283008, 485973, 467852, 427747, 251100, 146125, 557770, 577573, 298551 |
| 5903 | 231863 | Fees aside, what factors could account for performance differences between U.S. large-cap index ETFs? | 408524, 159471, 246996, 395842, 20504, 372233, 14185, 230997, 501153, 402091 |
| 5620 | 448784, 329552, 548740 | What's the fuss about identity theft? | 260580, 158285, 90632, 98993, 551747, 598801, 423809, 5860, 91986, 440524 |
| 2472 | 370334 | How do I deal with a mistaken attempt to collect a debt from me that is owed by someone else? | 180601, 201758, 49321, 161422, 200263, 595441, 500671, 435006, 330507, 131926 |
| 2306 | 315875 | To whom should I report fraud on both of my credit cards? | 581889, 531137, 289706, 596284, 90632, 270449, 164729, 226590, 412542, 360586 |
| 7633 | 197839 | Can a trade happen “in between” the bid and ask price? | 494727, 281844, 402482, 137175, 353396, 554207, 179258, 394244, 164008, 175831 |
| 2400 | 564271 | Will I be paid dividends if I own shares? | 1198, 97942, 456470, 95889, 587689, 311214, 501931, 1034, 377007, 400497 |
| 5549 | 286227, 309361 | Pros / cons of being more involved with IRA investments [duplicate] | 429106, 105468, 336394, 561636, 181624, 32009, 396852, 382894, 222082, 505362 |
| 3801 | 307776 | Can a bunch of wealthy people force Facebook to go public? | 390529, 69017, 264498, 209242, 168565, 371293, 394734, 570634, 171236, 92014 |
| 4605 | 453941 | If the U.S. defaults on its debt, what will happen to my bank money? | 41312, 313306, 229310, 169691, 400826, 526384, 479527, 581054, 354896, 598030 |
| 2885 | 367360, 359579, 85229, 454810 | Merits of buying apartment houses and renting them | 430672, 451849, 159403, 80838, 358687, 581251, 343917, 315972, 4739, 502291 |
| 6110 | 94117, 259706 | Why does short selling require borrowing? | 188531, 320450, 226496, 384252, 67107, 35500, 79764, 501984, 84761, 247313 |
| 8 | 566392 | How to deposit a cheque issued to an associate in my business into my business account? | 65404, 508754, 261856, 590102, 564553, 188893, 308938, 309023, 536686, 489199 |
| 1309 | 156162, 489401 | Why does FlagStar Bank harass you about payments within grace period? | 471630, 489368, 271040, 336792, 15824, 438869, 329817, 75108, 173919, 526989 |
| 7109 | 447781 | How do I analyse moving averages? | 489933, 140804, 221627, 42620, 257185, 227669, 193012, 565501, 518932, 180428 |
| 5080 | 256055 | Is there a standard or best practice way to handle money from an expiring UTMA account? | 445521, 279291, 414429, 451189, 69841, 236186, 267206, 132288, 324564, 150756 |
| 4981 | 247894 | Where can I find open source portfolio management software? | 45218, 102684, 259463, 557861, 587792, 232736, 81865, 226628, 296401, 55845 |
| 7445 | 153178, 104343 | IS it the wrong time to get into the equity market immediately after large gains? | 89714, 573612, 33155, 590902, 483025, 350068, 284075, 356623, 14543, 127160 |
| 2895 | 328691 | Where should a young student put their money? | 426461, 354551, 332749, 256055, 148453, 55841, 502170, 36190, 5188, 123256 |
| 6787 | 587120 | Would it make sense to sell a stock, then repurchase it for tax purposes? | 219762, 23217, 400730, 263751, 106104, 221715, 17184, 474981, 311782, 374867 |
| 5862 | 130209 | Can I get a discount on merchandise by paying with cash instead of credit? | 21194, 301643, 495751, 503171, 170141, 299840, 418801, 535015, 562511, 420622 |
| 6041 | 241308 | Most effective Fundamental Analysis indicators for market entry | 425020, 81655, 96910, 528034, 108579, 224695, 194240, 115087, 263464, 542765 |
| 7700 | 273761, 2653, 179328 | Should I re-allocate my portfolio now or let it balance out over time? | 269169, 224392, 253268, 22221, 434014, 395208, 131127, 28425, 422051, 126836 |
| 547 | 6349 | What percentage of my company should I have if I only put money? | 68088, 396694, 156747, 523158, 95243, 80913, 498681, 445353, 559522, 213399 |
| 3394 | 342258, 129319 | What is the easiest way to back-test index funds and ETFs? | 172374, 71230, 408524, 528034, 159471, 188855, 445971, 507777, 99568, 503725 |
| 4102 | 448699 | How can I determine if my rate of return is “good” for the market I am in? | 597437, 554734, 554237, 135176, 369439, 484688, 461082, 46394, 535737, 88801 |

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
| 34 | 0.0000 | 0.0000 | missing | 401k Transfer After Business Closure |
| 3534 | 0.0000 | 0.0000 | missing | Why do dishonour fees exist? |
| 3759 | 0.0000 | 0.0000 | missing | Simplifying money management |
