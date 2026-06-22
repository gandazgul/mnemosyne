# Mnemosyne BEIR fiqa Results

- Generated: 2026-06-21T22:18:18Z
- Dataset: `fiqa`
- Queries: 648
- Corpus documents: 57638
- Search limit: 100
- Source mode: `hybrid`
- Fusion: `vector-bm25`
- BM25 weight: `0.075`
- Rerank candidates: 300
- Rerank enabled: False
- Run label: `jina-v5-nano`
- Mnemosyne config: `/Users/gandazgul/Documents/web/mnemosyne/cli/configs/jina-v5-text-nano-retrieval.yaml`

| Metric | Score |
| --- | ---: |
| `ndcg@10` | 0.4879 |
| `mrr@10` | 0.5752 |
| `recall@10` | 0.5535 |
| `recall@100` | 0.7914 |
| `map@100` | 0.4258 |

## Breakdown

| Detail | Count |
| --- | ---: |
| `queries` | 648 |
| `queries_with_rank_1_hit` | 323 |
| `queries_with_top_10_hit` | 485 |
| `queries_with_top_100_hit` | 589 |
| `queries_missing_at_100` | 59 |

### First Relevant Rank

| Bucket | Queries |
| --- | ---: |
| `rank_1` | 323 |
| `rank_2_3` | 89 |
| `rank_4_10` | 73 |
| `rank_11_100` | 104 |
| `missing@100` | 59 |

### Missing At 100

| Query ID | Missing Doc IDs | Query | Top 10 Doc IDs |
| --- | --- | --- | --- |
| 4641 | 319954, 397358 | Where should I park my rainy-day / emergency fund? | 497993, 580025, 406219, 527939, 44594, 400503, 254572, 285812, 95598, 323686 |
| 3451 | 26292, 192307, 588448 | Should you keep your stocks if you are too late to sell? | 251536, 310683, 420974, 471911, 545284, 165970, 306460, 41852, 250873, 272091 |
| 753 | 243503 | Taxes due for hobbyist Group Buy | 466718, 132780, 451020, 283505, 466442, 172745, 293310, 547941, 299211, 203791 |
| 4465 | 376575 | How to donate to charity that will make a difference? | 106786, 132881, 90591, 174033, 46381, 379353, 427322, 326167, 266342, 174543 |
| 5951 | 497260 | Why can't house prices be out of tune with salaries | 418034, 599860, 31663, 259777, 374480, 62702, 596834, 373188, 529715, 557220 |
| 10827 | 160786, 7748, 107554, 95282 | How much should I be contributing to my 401k given my employer's contribution? | 290105, 296405, 576391, 436071, 140330, 555377, 312369, 452592, 38532, 242556 |
| 9771 | 263955, 28740 | Is there any emprical research done on 'adding to a loser' | 490479, 18671, 436765, 222639, 375290, 393272, 137898, 241605, 80481, 220190 |
| 3724 | 508921, 279570, 497216, 552887, 199970 | Should you always max out contributions to your 401k? | 302512, 430931, 122910, 3104, 273497, 576391, 135790, 488673, 497561, 140330 |
| 5853 | 476663, 160105, 495699, 424598 | Paying Off Principal of Home vs. Investing In Mutual Fund | 284318, 182612, 494148, 387722, 473647, 31525, 406325, 349355, 491923, 473865 |
| 570 | 363591 | Employer options when setting up 401k for employees | 532839, 117845, 79375, 387876, 289064, 150883, 555377, 15841, 508457, 301616 |
| 10122 | 273718 | Why diversify stocks/investments? | 259084, 297100, 459970, 144261, 89084, 508540, 153989, 502495, 564069, 570787 |
| 1783 | 332314 | Freelancing Tax implication | 421924, 14609, 156063, 159709, 179359, 445298, 549870, 383172, 223624, 15270 |
| 6004 | 149555 | Put-Call parity - what is the difference between the two representations? | 122432, 278373, 118762, 271109, 13260, 374797, 345410, 21768, 5257, 195152 |
| 4837 | 531841, 20958 | When applying for a mortgage, can it also cover outstanding debts? | 97925, 245005, 366448, 274870, 30311, 254454, 32749, 516578, 433171, 187739 |
| 4415 | 67676 | How much is inflation? | 117578, 468089, 267176, 206580, 59225, 306166, 501743, 513249, 381675, 184776 |
| 8378 | 125298 | Should I wait a few days to sell ESPP Stock? | 511678, 133644, 434812, 495568, 294573, 387035, 178684, 546125, 361345, 71713 |
| 8102 | 552707, 378173, 90294 | When do I sell a stock that I hold as a long-term position? | 306460, 165970, 35752, 537212, 557877, 310683, 203638, 219762, 250873, 171819 |
| 7928 | 118633 | If I believe a stock is going to fall, what options do I have to invest on this? | 427808, 501504, 260384, 410404, 24563, 173374, 30305, 47053, 480967, 102757 |
| 4777 | 590710 | How to finance necessary repairs to our home in order to sell it? | 365342, 52351, 570318, 478413, 482638, 416382, 26837, 487179, 118491, 432961 |
| 620 | 331332 | Is it wise to have plenty of current accounts in different banks? | 543921, 467059, 572848, 489959, 180673, 146317, 244097, 38628, 449630, 21420 |
| 8271 | 415511 | Income in zero-interest environment | 376709, 249558, 52047, 380368, 119298, 83330, 106424, 546874, 503040, 136262 |
| 864 | 211364, 152072 | Why use accounting software like Quickbooks instead of Excel spreadsheets? | 30142, 472924, 566337, 24890, 2436, 157751, 78117, 329774, 222380, 505361 |
| 5888 | 540806 | Interest charges on balance transfer when purchases are involved | 263647, 125497, 579601, 543776, 545327, 358445, 336792, 429746, 231105, 506132 |
| 2568 | 388798, 127353 | How to pay with cash when car shopping? | 346042, 9146, 15696, 355310, 269898, 514238, 108739, 166314, 214749, 78865 |
| 7124 | 74615 | How come we can find stocks with a Price-to-Book ratio less than 1? | 558617, 526110, 154725, 583708, 278582, 533818, 123263, 157597, 192686, 33357 |
| 2857 | 295864 | I have around 60K $. Thinking about investing in Oil, how to proceed? | 233732, 127566, 501384, 76996, 542051, 49235, 117451, 379140, 329930, 58065 |
| 8539 | 218728, 196304, 396038 | Can the risk of investing in an asset be different for different investors? | 283074, 483123, 471817, 391861, 502495, 387515, 292609, 240178, 346735, 499166 |
| 1085 | 467737, 393710 | How do disputed debts work on credit reports? | 450031, 372039, 268777, 319276, 161422, 398258, 384608, 1218, 78328, 339365 |
| 5422 | 151973 | What are some good books for learning stocks, bonds, derivatives e.t.c for beginner with a math background? | 165294, 172587, 273906, 276786, 191688, 221319, 552371, 79517, 193555, 13513 |
| 8247 | 321114, 465313 | Tax on Stocks or ETF's | 586010, 161019, 153112, 350317, 518735, 190687, 474745, 492053, 343708, 29502 |
| 10482 | 549072 | Rollover into bond fund to do dollar cost averaging [duplicate] | 330023, 447567, 18436, 439757, 564787, 224782, 175576, 434466, 265817, 127401 |
| 8789 | 70853 | What does “profits to the shareholders jumped to 15 cents a share” mean? | 87349, 41912, 341424, 573079, 20076, 144079, 325669, 317363, 14870, 411617 |
| 2423 | 538023 | At what age should I start or stop saving money? | 529444, 417787, 396127, 234846, 272328, 104457, 328157, 553288, 337561, 468010 |
| 8475 | 293320 | Why I cannot find a “Pure Cash” option in 401k investments? | 494655, 481728, 104793, 240562, 437427, 444107, 525426, 363652, 554739, 332152 |
| 4523 | 594257 | What should I do with my $25k to invest as a 20 years old? | 129255, 465819, 171712, 332022, 332203, 272070, 286746, 72578, 10476, 216365 |
| 8513 | 270573 | Buy on dip when earnings fail? | 572622, 175821, 203873, 351396, 335626, 53431, 534734, 432665, 73872, 494877 |
| 5054 | 28119 | How to stress test an investment plan? | 101124, 588481, 377186, 458183, 127263, 263390, 220665, 534010, 403916, 463831 |
| 1159 | 496064 | what is the best way to do a freelancing job over the summer for a student | 55064, 347181, 163881, 460648, 594182, 237282, 469972, 112669, 449155, 132287 |
| 68 | 19183 | Intentions of Deductible Amount for Small Business | 97719, 519473, 54333, 70315, 146657, 86134, 462831, 107213, 192516, 14255 |
| 9701 | 387141, 357739 | How to bet against the London housing market? | 473883, 408865, 258048, 225682, 516214, 588608, 108399, 70799, 367103, 150632 |
| 6199 | 239214 | How can all these countries owe so much money?  Why & where did they borrow it from? | 414693, 47163, 584273, 351853, 10399, 490042, 169921, 49602, 298794, 596959 |
| 9126 | 514831 | Short an option - random assignment? | 334473, 477588, 308859, 82194, 267113, 469382, 277760, 228810, 349974, 510006 |
| 34 | 599545 | 401k Transfer After Business Closure | 551545, 483268, 458917, 490867, 168890, 226547, 144109, 591168, 104134, 255277 |
| 6395 | 166227 | Option settlement for calendar spreads | 584223, 22916, 276314, 401447, 505223, 467463, 111301, 516790, 189144, 151546 |
| 9115 | 207325 | Why does the calculation for percentage profit vary based on whether a position is short vs. long? | 245082, 158520, 419897, 422467, 428786, 314478, 385220, 99131, 232880, 330041 |
| 4767 | 280805, 568670, 224057, 22804 | New car: buy with cash or 0% financing | 420018, 306834, 429153, 56867, 584106, 9146, 451092, 408932, 166314, 475440 |
| 8351 | 472516 | What happens when a calendar spread is assigned in a non-margin account? | 141213, 102316, 516790, 45674, 23609, 111301, 527654, 100628, 272754, 300139 |
| 6131 | 381720, 170204, 416679, 2460 | Is it ever a good idea to close credit cards? | 326094, 368806, 391384, 339030, 218088, 334111, 533288, 99449, 35625, 258465 |
| 858 | 122485, 278450 | Is it bad practice to invest in stocks that fluctuate by single points throughout the day? | 293027, 433730, 146632, 519501, 514780, 241860, 30774, 289, 216964, 573612 |
| 4019 | 6881, 125477 | How and Should I Invest (As a college 18 year old with minimal living expenses)? | 85977, 332938, 269671, 379948, 426461, 332749, 571044, 211713, 347849, 204479 |
| 6080 | 164513 | Is ScholarShare a legitimate entity for a 529 plan in California? | 22856, 236732, 201500, 277581, 468527, 115175, 233401, 83080, 2809, 535357 |
| 6959 | 205010 | What is the term for the quantity (high price minus low price) for a stock? | 468025, 428117, 577573, 599523, 303325, 169954, 472537, 229573, 116846, 317365 |
| 5402 | 491350 | Is it impossible to get a home loan with a poor credit history after a divorce? | 445163, 51728, 227485, 90579, 595029, 44105, 52250, 129177, 67066, 122807 |
| 6562 | 501157 | Cheapest way to “wire” money in an Australian bank account to a person in England, while I'm in Laos? | 473605, 582414, 60446, 183880, 324817, 182443, 135675, 282744, 428696, 549684 |
| 701 | 389446 | What are the ins/outs of writing-off part of one's rent for working at home? | 231990, 436505, 344955, 349672, 337706, 456234, 339488, 177074, 288537, 124507 |
| 7674 | 519390 | Choosing the limit when making a limit order? | 249279, 447886, 155151, 514841, 278630, 94653, 15917, 437436, 184756, 236133 |
| 5940 | 486243 | How does investment into a private company work? | 46842, 512609, 250354, 535314, 454465, 473154, 530, 182226, 252853, 135411 |
| 6612 | 205522, 322900 | If I have a lot of debt and the housing market is rising, should I rent and slowly pay off my debt or buy and roll th... | 502594, 301192, 180192, 198442, 254454, 431481, 290900, 433171, 393002, 386348 |
| 8456 | 486333 | What typically happens to unvested stock during an acquisition? | 257853, 534755, 93215, 174321, 391156, 469036, 475019, 555276, 235779, 261487 |
| 10213 | 270221, 545712 | Looking for good investment vehicle for seasonal work and savings | 446186, 100517, 38269, 223551, 112499, 488100, 589394, 390556, 94680, 58065 |
| 5196 | 172128, 114829 | I might use a credit card convenience check. What should I consider? | 565745, 393866, 85252, 289483, 302823, 456098, 498775, 402543, 316652, 2875 |
| 3006 | 269851, 568473, 328300 | Strategies for putting away money for a child's future (college, etc.)? | 512096, 258704, 127838, 471019, 211713, 490382, 597627, 303432, 372900, 360285 |
| 3909 | 312248, 404356, 245616, 353028 | How to rescue my money from negative interest? | 514003, 83330, 61586, 472837, 362730, 328499, 288656, 404352, 526169, 438403 |
| 6907 | 251604 | Nominal value of shares | 303112, 69506, 170652, 487738, 111827, 21786, 91870, 480515, 481761, 275392 |
| 5464 | 86691 | Resources on Buying Rental Properties | 222095, 372274, 325722, 423438, 383921, 294549, 536126, 315972, 44855, 155964 |
| 10034 | 480749 | Tax implications of holding EWU (or other such UK ETFs) as a US citizen? | 528880, 181942, 565296, 447197, 141585, 562007, 430868, 104128, 197478, 586010 |
| 5090 | 436493 | Should I take a student loan to pursue my undergraduate studies in France? | 12988, 246286, 217831, 92430, 586289, 560681, 21913, 213328, 455666, 551964 |
| 9391 | 503637 | Should I replace bonds in a passive investment strategy | 535518, 136515, 248158, 107424, 283202, 342485, 155242, 112369, 577832, 466403 |
| 3148 | 178127, 438000 | Can a car company refuse to give me a copy of my contract or balance details? | 172855, 584305, 430100, 29721, 357280, 65046, 92888, 164702, 205984, 560325 |
| 4678 | 305153 | Finance, Cash or Lease? | 427884, 185405, 311748, 376016, 215225, 504918, 311446, 85373, 487678, 536773 |
| 2398 | 224654, 590489 | Frustrated Landlord | 556453, 44058, 168089, 96538, 487094, 393883, 310992, 98372, 422579, 482991 |
| 5511 | 169893, 560325, 383193, 278699, 12746 | Pay off car loan entirely or leave $1 until the end of the loan period? | 529123, 334559, 155843, 500946, 38786, 324269, 107898, 139788, 334964, 27693 |
| 988 | 226053, 107688 | Where should I invest my savings? | 501384, 571218, 223872, 82482, 320675, 194080, 223551, 60093, 168402, 450558 |
| 3369 | 231012, 145716, 411910, 395840 | Why should one only contribute up to the employer's match in a 401(k)? | 296405, 555377, 341493, 15841, 242556, 92370, 436071, 240259, 443002, 240373 |
| 9245 | 194561 | Stock Options for a company bought out in cash and stock | 207253, 186869, 178497, 131488, 39345, 248393, 472516, 451613, 259560, 409818 |
| 5763 | 462019 | What is the best way to get a “rough” home appraisal prior to starting the refinance process? | 570318, 326214, 38712, 218144, 515361, 67379, 89964, 11572, 251466, 331255 |
| 4962 | 599925 | Net Cash Flows from Selling the Bond and Investing | 158363, 416839, 52149, 537603, 408661, 154707, 510268, 308276, 395208, 196173 |
| 4846 | 151104 | Is there anything comparable to/resembling CNN's Fear and Greed Index? | 538974, 335892, 489352, 98096, 320059, 533311, 415161, 183597, 309245, 269701 |
| 9403 | 6666, 328086, 345199 | Abundance of Cash - What should I do? | 410450, 14349, 570632, 159235, 103447, 186332, 215296, 551986, 357887, 372223 |
| 5993 | 367375, 272866, 55084, 352638, 426120, 63501 | Why would anyone want to pay off their debts in a way other than “highest interest” first? | 94373, 160193, 287571, 416796, 128574, 156195, 353911, 431212, 494306, 544858 |
| 5710 | 232311 | Bucketing investments to track individual growths | 227364, 516267, 88417, 11979, 534323, 7748, 227733, 399738, 412830, 583549 |
| 7529 | 66607 | Does the expense ratio of a fund-of-funds include the expense ratios of its holdings? | 514529, 464337, 89297, 293626, 287537, 59249, 518402, 218261, 387980, 451855 |
| 5021 | 589285 | Is there a more flexible stock chart service, e.g. permitting choice of colours when comparing multiple stocks? | 528576, 584801, 465971, 60284, 49893, 189341, 511861, 105717, 423177, 211444 |
| 3612 | 259625 | How can I buy and sell the same stock on the same day? | 567383, 165548, 522658, 310636, 429418, 483676, 460937, 402726, 390864, 584291 |
| 4409 | 499128, 100306, 147439 | My friend wants to put my name down for a house he's buying. What risks would I be taking? | 243732, 223841, 102088, 268078, 360682, 244278, 102326, 84732, 514790, 426676 |
| 2070 | 363678 | Advantage of credit union or local community bank over larger nationwide banks such as BOA, Chase, etc.? | 578357, 469515, 550303, 587737, 597571, 249839, 590209, 261389, 38038, 197087 |
| 11039 | 53544, 249063 | Pay off credit card debt or earn employer 401(k) match? | 91183, 287876, 508534, 79363, 552383, 5203, 345895, 163287, 105557, 281049 |
| 5460 | 184337, 21174, 108514, 463885 | Paying off a loan with a loan to get a better interest rate | 77052, 327115, 482798, 58432, 470716, 529418, 522341, 106495, 490648, 503723 |
| 7925 | 318185, 402482 | Can I sell a stock immediately? | 438974, 584291, 591436, 133644, 219762, 407602, 581579, 238215, 568200, 310636 |
| 4286 | 566069 | Given advice “buy term insurance and invest the rest”, how should one “invest the rest”? | 70460, 511386, 229239, 391243, 206830, 79142, 10531, 155640, 221479, 56732 |
| 2685 | 154113, 370300, 303293, 468923 | What ways are there for us to earn a little extra side money? | 382005, 468086, 594182, 269380, 186889, 4992, 558832, 543275, 237950, 399882 |
| 1090 | 518896 | Need a formula to determine monthly payments received at time t if I'm reinvesting my returns | 446454, 296146, 573928, 393987, 179365, 520217, 521590, 281329, 203091, 198606 |
| 6122 | 44344 | Better to rent condo to daughter or put her on title? | 496166, 403515, 316794, 558251, 182039, 118246, 423398, 577658, 566184, 101816 |
| 4514 | 69485, 337764, 209804 | What intrinsic, non-monetary value does gold have as a commodity? | 471825, 426270, 156211, 317429, 99089, 408336, 352485, 80141, 532381, 146573 |
| 8507 | 370995 | When to sell a stock? | 251536, 99132, 545284, 303724, 236415, 27015, 272091, 368348, 420974, 217837 |
| 6221 | 257248, 76414, 455614, 115717 | To pay off a student loan, should I save up a lump sum payoff payment or pay extra each month? | 448791, 110081, 352363, 414534, 394474, 541313, 69150, 254245, 414288, 571198 |
| 3008 | 180192, 323406 | What are my chances at getting a mortgage with Terrible credit but High income | 231688, 251846, 47441, 227485, 44105, 102266, 407401, 364802, 455952, 285694 |
| 4007 | 521657 | What is a reasonable salary for the owner and sole member of a small S-Corp? | 556220, 260385, 205341, 370542, 388704, 170933, 515233, 315552, 521933, 458431 |
| 6644 | 175035 | How to know precisely when a SWIFT is issued by a bank? | 110198, 475527, 218761, 271596, 41383, 355870, 327623, 39783, 187497, 118396 |
| 10267 | 460398 | How should I prepare for the next financial crisis? | 178693, 569632, 143393, 326398, 87520, 436091, 305600, 182442, 279456, 577806 |
| 7622 | 253369, 378594 | Best way to pay off debt? | 220241, 353911, 457945, 271525, 157923, 388095, 480773, 544858, 115499, 416796 |
| 3767 | 153922, 392060 | What should I be doing to protect myself from identity theft? | 90632, 423809, 260580, 125204, 97686, 368679, 171510, 581889, 158008, 587778 |
| 6410 | 471723 | Will an ETF immediately reflect a reconstitution of underlying index | 454610, 71230, 87261, 227324, 147282, 200360, 408524, 295993, 428187, 120059 |
| 5030 | 215540 | Why pay for end-of-day historical prices? | 13511, 227192, 149420, 295344, 560108, 327974, 378994, 532178, 370569, 471131 |
| 6252 | 394551, 160932, 293624, 233294, 243268, 379487, 62868 | Is this mortgage advice good, or is it hooey? | 213713, 120061, 47565, 473647, 316149, 104988, 272634, 139366, 426638, 205906 |
| 885 | 337165, 409184 | How long do credit cards keep working after you disappear? | 516678, 254968, 89888, 472336, 251643, 301792, 588719, 99449, 59228, 251701 |
| 4031 | 115741 | 28 years old and just inherited large amount of money and real estate - unsure what to do with it | 318864, 140002, 568629, 65180, 375708, 163197, 193171, 387717, 171196, 266481 |
| 766 | 550172 | Will the ex-homeowner still owe money after a foreclosure? | 2996, 552768, 299591, 27987, 333583, 104955, 56126, 163711, 554171, 6882 |
| 8202 | 513258, 93971 | What accounted for DXJR's huge drop in stock price? | 457689, 162047, 122542, 41271, 337001, 5220, 317363, 363451, 537862, 253339 |
| 7345 | 237645 | What do these numbers mean? (futures) | 9274, 508821, 527080, 206895, 460331, 108, 366526, 164001, 529996, 475916 |
| 776 | 583640, 127263, 597880, 496899 | Can saving/investing 15% of your income starting age 25, likely make you a millionaire? | 124027, 10440, 467044, 417787, 143591, 41960, 191148, 554833, 374266, 434972 |
| 89 | 248624 | How can I deposit a check made out to my business into my personal account? | 508754, 309023, 188893, 526817, 308938, 135196, 29372, 65404, 590102, 181187 |
| 1889 | 388713 | Reporting financial gains from my online store | 584074, 355959, 352136, 243503, 553540, 41509, 427202, 599876, 299211, 560776 |
| 8013 | 496159, 224231 | Frequency of investments to maximise returns (and minimise fees) | 384983, 48678, 81652, 537626, 446948, 278607, 24846, 225815, 8759, 57033 |
| 3759 | 527966, 67167, 522358 | Simplifying money management | 455457, 490065, 214248, 373772, 248663, 372743, 551986, 591704, 526159, 526568 |
| 10639 | 431799, 495774, 278453, 187039 | Short term parking of a large inheritance? | 171196, 318864, 163353, 375708, 235628, 140002, 163197, 405115, 422994, 360816 |
| 6635 | 156358 | Why don't share prices of a company rise every other Friday when the company buys shares for its own employees? | 587137, 3656, 533712, 343452, 235531, 579037, 125298, 12560, 467594, 569224 |
| 4312 | 399149 | Is it true that 90% of investors lose their money? | 282435, 222639, 167950, 532485, 116647, 497786, 285945, 170628, 431735, 544857 |
| 6525 | 181985 | Does it make sense to trade my GOOGL shares for GOOG and pocket the difference? | 106541, 550661, 105542, 98150, 221795, 192696, 378906, 362473, 498014, 76330 |
| 2590 | 589625 | Are non-residents or foreigners permitted to buy or own shares of UK companies? | 296528, 48269, 55999, 209493, 158923, 458730, 188776, 310103, 456999, 307776 |
| 5374 | 152688 | What were the main causes of the spike and drop of DRYS's stock price? | 283106, 457689, 133204, 122542, 253339, 589533, 317363, 380894, 5220, 261522 |
| 2994 | 419319 | Work on the side for my wife's company | 510317, 569145, 138428, 321743, 460721, 491844, 562780, 523372, 5840, 506991 |
| 3683 | 454501 | Can I trust the Motley Fool? | 276975, 105973, 408995, 428848, 538086, 500338, 565016, 522713, 47812, 261679 |
| 7206 | 441155, 532211, 553066 | Who Bought A Large Number Of Shares? | 65667, 351570, 558703, 34882, 350214, 498676, 54424, 444752, 573846, 207853 |
| 10246 | 77573 | Understanding the T + 3 settlement days rule | 370635, 156029, 176717, 327080, 340263, 28314, 293389, 332243, 315205, 11927 |
| 5241 | 322157, 27489 | Mortgage vs. Cash for U.S. home buy now | 344740, 281675, 111184, 213713, 273735, 438073, 390976, 309420, 107350, 234286 |
| 98 | 575929 | How can I make $250,000.00 from trading/investing/business within 5 years? | 527522, 555630, 438279, 519619, 373119, 66034, 449745, 336661, 506149, 465819 |
| 4615 | 262934 | Are solar cell panels and wind mills worth the money? | 261900, 69523, 496427, 455798, 120384, 425595, 271015, 385028, 158216, 19179 |
| 6467 | 453256, 23217, 346641, 367313 | Advice on strategy for when to sell | 88813, 368348, 240089, 217837, 250873, 99857, 303724, 203873, 251536, 109455 |
| 4289 | 24881 | Does the currency exchange rate contain any additional information at all? | 288330, 114886, 226102, 416975, 17469, 324546, 256996, 517345, 104480, 356465 |
| 4394 | 441582 | Transfer $50k to another person's account (in California, USA) | 322838, 524752, 93386, 412258, 305907, 415655, 431462, 293653, 462585, 80657 |
| 4047 | 77652 | Does doing your “research”/“homework” on stocks make any sense? | 417840, 324779, 452175, 110733, 170966, 284165, 210241, 358846, 474296, 312445 |
| 7344 | 108403 | How is the Dow divisor calculated? | 159166, 14368, 378974, 150430, 65618, 253926, 313421, 591089, 96697, 501032 |
| 6875 | 224392 | Where to find free Thailand stock recommendations and research? | 567500, 110733, 352557, 224366, 79337, 512895, 249679, 77502, 472062, 232460 |
| 10447 | 152096, 300721 | Is there an advantage to a traditional but non-deductable IRA over a taxable account? [duplicate] | 144751, 500175, 382236, 259150, 532657, 382894, 114912, 447482, 26652, 588134 |
| 5782 | 595455 | Pay off credit cards in one lump sum, or spread over a few months? | 262026, 487621, 172084, 117007, 440620, 273631, 114592, 559523, 529312, 113822 |
| 9871 | 448890, 40051, 170594 | What should I do with the 50k I have sitting in a European bank? | 367207, 292714, 433003, 73741, 387723, 231521, 74668, 76562, 219477, 231246 |
| 3625 | 414295 | What should I do with my paper financial documents? | 500751, 509617, 569812, 380263, 513248, 113830, 344236, 305065, 63422, 123366 |
| 6005 | 478457, 345895, 73310, 384626, 390689 | Why might it be advisable to keep student debt vs. paying it off quickly? | 149500, 507544, 571198, 96268, 572272, 431884, 25190, 422704, 117085, 564206 |
| 3822 | 385090, 418900, 308837 | How to change a large quantity of U.S. dollars into Euros? | 292714, 541608, 239876, 417917, 174406, 194730, 478065, 208499, 144181, 305907 |
| 7879 | 372551, 421285 | Any Tips on How to Get the Highest Returns Within 4 Months by Investing in Stocks? | 58186, 102029, 43088, 272174, 174313, 105391, 477295, 377186, 367415, 540919 |
| 3115 | 234950, 389028, 316794 | How can I live outside of the rat race of American life with 300k? | 233562, 183869, 267892, 129364, 366961, 174272, 559852, 369742, 220023, 136035 |
| 3995 | 278734, 230208 | I have more than $250,000 in a US Bank account… mistake? | 485507, 171720, 479918, 404954, 264934, 14349, 403288, 505461, 297900, 156716 |
| 10136 | 526115 | How to minimise the risk of a reduction in purchase power in case of Brexit for money held in a bank account? | 466950, 290930, 417740, 304007, 583903, 168137, 393823, 265453, 205685, 152316 |
| 8635 | 67107, 240215 | Is there any flaw in this investment scheme? | 46818, 575918, 447619, 303619, 493841, 151238, 365816, 202638, 203729, 264822 |
| 5206 | 563030, 28230, 117276, 300660 | Is it a good idea to get an unsecured loan to pay off a credit card that won't lower a high rate? | 298908, 595455, 540959, 498728, 60996, 287157, 340520, 585288, 69938, 153088 |
| 2713 | 388147 | Physical Checks - Mailing | 29372, 284528, 584170, 41944, 564301, 216200, 229546, 20791, 199069, 78139 |
| 9060 | 40447 | Buying puts without owning underlying | 511093, 528052, 228217, 345851, 359778, 7743, 3062, 118762, 415705, 316037 |
| 4105 | 25096 | As an investor what are side effects of Quantitative Easing in US and in EU? | 416483, 345910, 176262, 305029, 293104, 185300, 108519, 459400, 369038, 30946 |
| 2465 | 570680, 81046 | Can capital expenses for volunteer purposes be deducted from income? | 37382, 398536, 432545, 202645, 510716, 18889, 310612, 107213, 146657, 364938 |
| 4640 | 101369 | What can my relatives do to minimize their out of pocket expenses on their fathers estate | 295246, 331534, 356035, 360816, 422994, 17110, 375708, 372808, 144965, 421652 |
| 9275 | 338754 | Do I have to pay a capital gains tax if I rebuy the same stock within 30 days? | 400730, 343219, 390864, 376800, 23217, 102443, 537916, 448659, 263751, 522319 |
| 3500 | 174019 | Why invest in becoming a landlord? | 273187, 528206, 557478, 422331, 71424, 141935, 159156, 578597, 536126, 431110 |
| 1306 | 484437, 204075 | I made an investment with a company that contacted me, was it safe? | 309851, 594206, 538086, 205665, 160611, 537698, 567973, 450779, 407663, 112259 |
| 6262 | 26799 | Help required on estimating SSA benefit amounts | 34538, 390877, 83338, 15322, 2648, 118707, 498444, 186602, 107963, 192585 |
| 8632 | 213976 | Is it best to exercise options shares when they vest, or wait | 43497, 340730, 420722, 382381, 237718, 488207, 388362, 178497, 259560, 104188 |
| 6133 | 415705 | What happens to all of the options when they expire? | 7733, 11456, 575408, 73256, 581672, 132288, 242298, 481070, 358492, 428399 |
| 3771 | 488948, 198349, 217683, 49601 | Best way to buy Japanese yen for travel? | 521712, 495826, 490384, 96211, 128471, 306130, 434201, 350245, 152695, 402006 |
| 1736 | 25543, 443419 | How can people have such high credit card debts? | 399406, 437610, 475668, 562896, 372993, 517050, 463065, 298908, 382591, 235646 |
| 6814 | 340214, 223206 | Selling Stock - All or Nothing? | 590188, 513734, 400614, 878, 394454, 349147, 3095, 345368, 124188, 250873 |
| 1322 | 64138 | Is this follow-up after a car crash a potential scam? | 114231, 397852, 283917, 567973, 91463, 318941, 33914, 226090, 98356, 264986 |
| 5185 | 210236, 317354 | Invest in low cost small cap index funds when saving towards retirement? | 196992, 376485, 262180, 241202, 503725, 580313, 524525, 434279, 59670, 545760 |
| 6909 | 127012 | Why do stocks priced above $2.00 on the ASX sometimes move in $0.005 increments? | 72633, 118232, 450489, 298551, 47217, 490584, 64943, 555226, 596821, 43432 |
| 2348 | 211867, 566573, 211622, 474234, 352271, 265874 | Why can't you just have someone invest for you and split the profits (and losses) with him? | 447619, 247486, 389004, 306430, 420544, 381757, 309851, 151412, 177194, 197389 |
| 4499 | 323363, 76996 | Is investing exlusively in a small-cap index fund a wise investment? | 196992, 501153, 517391, 235917, 52274, 238963, 9512, 14748, 445322, 513818 |
| 3530 | 239998 | How to exclude stock from mutual fund | 24029, 184299, 378075, 209879, 449124, 370754, 574383, 332152, 479420, 110343 |
| 9108 | 272021, 472585 | Starting an investment portfolio with Rs 5,000/- | 290757, 171189, 312821, 356552, 51848, 323067, 122679, 414116, 527522, 461193 |
| 6835 | 102243 | Are bond ETF capital gains taxed similar to stock or stock funds if held for more than 1 year? | 149305, 5710, 84238, 287950, 29502, 543842, 153112, 586010, 570546, 225536 |
| 3067 | 406156 | Should I make extra payments to my under water mortgage or increase my savings? | 477907, 560915, 476068, 468831, 423403, 90009, 131365, 341837, 440719, 323475 |
| 4125 | 344648, 72046 | Alternative means of salary for my employees | 174787, 365558, 58906, 73999, 302310, 36608, 414694, 245451, 287200, 65795 |
| 1150 | 43603, 19936 | How are the best way to make and save money at 22 years old | 529444, 10476, 353369, 319760, 595287, 328157, 433986, 66864, 204479, 147745 |
| 7431 | 372921 | Pay off mortgage or invest in high value saving account | 2393, 589256, 468831, 435576, 557734, 157414, 537721, 477907, 440719, 364099 |
| 9808 | 40702, 431946 | Selling To Close | 416307, 125079, 557582, 218423, 414636, 43087, 151587, 152719, 345368, 449280 |
| 3888 | 319213, 239632 | Why I can't view my debit card pre-authorized amounts? | 208169, 185434, 294077, 440527, 276733, 432077, 281129, 316652, 521010, 448086 |
| 10109 | 506374, 499849 | Why does Charles Schwab have a Mandatory Settlement Period after selling stocks? | 93231, 28314, 332243, 370635, 121465, 98302, 266725, 124188, 293389, 119161 |
| 2790 | 279329, 469125 | Should I pay more than 20% down on a home? | 472484, 400896, 207564, 385343, 75961, 357200, 64400, 100483, 443413, 487593 |
| 9882 | 65702 | Money-market or cash-type ETFs for foreigners with U.S brokerage account | 389581, 391876, 188524, 535340, 313775, 386173, 363378, 466845, 98461, 161966 |
| 4999 | 314898 | Looking for a good source for Financial Statements | 9938, 171964, 11263, 431459, 146076, 295738, 597241, 520165, 465971, 338803 |
| 3189 | 225395 | Diversify my retirement investments with a Roth IRA | 287225, 240975, 423658, 404800, 122222, 458168, 88311, 272840, 271949, 403556 |
| 5134 | 158523 | Why does Yahoo Finance's data for a Vanguard fund's dividend per share not match the info from Vanguard? | 46774, 532616, 374330, 584128, 206727, 215486, 221477, 60098, 239137, 317666 |
| 1321 | 216456, 292065 | Are social media accounts (e.g. YouTube, Twitter, Instagram, etc.) considered assets? | 381322, 392379, 235082, 203715, 481036, 351672, 210019, 559738, 278538, 100849 |
| 4539 | 370879 | How should I save money if the real interest rate (after inflation) is negative? | 42475, 275925, 194080, 472837, 479659, 203926, 32744, 61586, 503394, 328499 |
| 715 | 579763, 546538, 187404 | what would you do with $100K saving? | 548758, 133120, 200273, 113885, 260677, 203201, 556545, 427032, 372223, 328770 |
| 504 | 344203, 498751 | Have plenty of cash flow but bad credit | 22807, 546097, 486334, 99463, 368247, 503419, 93573, 432040, 495431, 52250 |
| 2296 | 83330, 366594 | How does a bank make money on an interest free secured loan? | 400009, 94230, 119298, 396853, 580147, 106424, 259919, 172303, 249831, 279897 |
| 10975 | 61022 | How to contribute to Roth IRA when income is at the maximum limit & you have employer-sponsored 401k plans? | 441632, 353009, 163865, 81148, 446226, 110114, 360533, 270818, 140330, 75766 |
| 1994 | 156640 | Does the IRS reprieve those who have to commute for work? | 434846, 231990, 192843, 51491, 585706, 147624, 63919, 263259, 380635, 319965 |
| 9164 | 365298, 263390 | Bonds vs equities: crash theory | 115648, 149900, 309326, 16924, 506743, 599420, 321941, 305274, 296516, 409859 |
| 10462 | 8266, 11378, 35680, 437879, 204035, 581204 | Is it okay to be married, 30 years old and have no retirement? | 66376, 268023, 391583, 122333, 593377, 152478, 139595, 51914, 221295, 386404 |
| 8855 | 208165, 312821 | How do i get into investing stocks [duplicate] | 312445, 155677, 313570, 403092, 560395, 484327, 555521, 328477, 210241, 178017 |
| 7071 | 124230 | ESPP strategy - Sell right away or hold? | 511678, 133644, 294573, 434812, 361345, 71713, 495568, 387035, 575213, 127702 |
| 3534 | 340329 | Why do dishonour fees exist? | 149845, 386745, 312349, 234260, 166431, 517497, 34108, 286064, 333149, 149944 |
| 8974 | 356595, 170625 | As a 22-year-old, how risky should I be with my 401(k) investments? | 134931, 216365, 102501, 452126, 10476, 96110, 140738, 336144, 327925, 124762 |
| 5061 | 23747 | What fiscal scrutiny can be expected from IRS in early retirement? | 31699, 502150, 25481, 598378, 149742, 513392, 484683, 537049, 253735, 170916 |
| 2075 | 170042, 14967 | Are stories of turning a few thousands into millions by trading stocks real? | 519619, 523393, 44417, 375876, 506149, 555521, 65667, 41926, 249055, 39927 |
| 7533 | 93853 | Investing tax (savings) | 315105, 8012, 142658, 563986, 365193, 250603, 550468, 187571, 66717, 106501 |
| 1393 | 352838, 539133 | Which is better when working as a contractor, 1099 or incorporating? | 220022, 352640, 234436, 68524, 32072, 586026, 277812, 578196, 510913, 194899 |
| 9733 | 110163, 38655 | Due Diligence - Dilution? | 23414, 526073, 301880, 316321, 135798, 108965, 267266, 419505, 154841, 423260 |
| 7311 | 323768 | Finance, Social Capital IPOA.U | 479752, 507841, 579110, 290325, 419735, 583646, 584135, 571001, 538727, 261231 |
| 744 | 566480 | What options are available for a home loan with poor credit but a good rental history? | 490443, 67066, 415425, 310790, 80607, 313623, 92397, 47441, 289231, 256981 |
| 7141 | 132288 | Do investors go long option contracts when they cannot cover the exercise of the options? | 538054, 255927, 41967, 507828, 557356, 293767, 357324, 243714, 334473, 31587 |
| 4071 | 129875 | If our economy crashes, and cash is worthless, should i buy gold or silver | 524142, 505136, 291862, 487817, 473965, 502634, 53538, 506780, 302010, 308332 |
| 7512 | 191060 | understanding the process/payment of short sale dividends | 487329, 222320, 568166, 298284, 13631, 115553, 202985, 480949, 409432, 527636 |
| 1391 | 562176 | How is taxation for youtube/twitch etc monetization handled in the UK? | 440745, 77171, 254151, 454208, 223170, 510599, 527951, 378060, 111131, 245753 |
| 7534 | 358125 | Can you explain why it's better to invest now rather than waiting for the market to dip? | 175821, 145539, 114806, 33155, 310218, 103622, 89714, 474006, 350068, 222444 |
| 5356 | 312405 | Historical stock prices: Where to find free / low cost data for offline analysis? | 279785, 240086, 529877, 596106, 189341, 535343, 560108, 546379, 47798, 167586 |
| 2579 | 432020 | What to do when a job offer is made but with a salary less than what was asked for? | 423070, 200946, 524471, 190077, 283825, 256802, 364159, 432808, 559900, 157919 |
| 7823 | 583549 | Retirement Funds: Betterment vs Vanguard Life strategy vs Target Retirement | 451196, 105666, 175927, 331492, 268731, 172336, 293679, 347825, 329425, 57070 |
| 689 | 411044 | Receive credit card payment sending my customer details to a credit card processing company? | 446932, 63366, 553418, 421803, 567201, 438032, 104079, 184924, 195852, 171761 |
| 9174 | 535317, 160218 | Which U.S. online discount broker is the best value for money? | 236931, 192910, 563334, 200052, 544576, 31936, 515144, 405217, 513281, 47579 |
| 6867 | 443804, 445258 | Will there always be somebody selling/buying in every stock? | 466143, 301985, 482739, 429196, 321639, 349147, 560273, 400614, 224672, 547553 |
| 2383 | 232199 | Should I Purchase Health Insurance Through My S-Corp | 17215, 224406, 527620, 327232, 209159, 546634, 222665, 308255, 281268, 41793 |
| 5083 | 138845 | Co-signer deceased | 369075, 447983, 270952, 305509, 18257, 518681, 273759, 288701, 495482, 453263 |
| 10526 | 39185 | What extra information might be obtained from the next highest bids in an order book? | 283008, 546493, 467852, 485973, 427747, 146125, 251100, 577573, 137175, 557770 |
| 5903 | 231863 | Fees aside, what factors could account for performance differences between U.S. large-cap index ETFs? | 408524, 159471, 246996, 20504, 395842, 14185, 230997, 372233, 501153, 580820 |
| 5620 | 448784, 329552, 548740 | What's the fuss about identity theft? | 260580, 90632, 158285, 98993, 551747, 423809, 598801, 5860, 440524, 547189 |
| 2472 | 370334 | How do I deal with a mistaken attempt to collect a debt from me that is owed by someone else? | 180601, 49321, 161422, 201758, 435006, 546028, 330507, 200263, 500671, 233535 |
| 2306 | 315875 | To whom should I report fraud on both of my credit cards? | 581889, 531137, 289706, 596284, 90632, 270449, 360586, 125204, 226590, 274019 |
| 7633 | 197839 | Can a trade happen “in between” the bid and ask price? | 494727, 281844, 402482, 137175, 554207, 179258, 353396, 394244, 382067, 175831 |
| 2400 | 564271 | Will I be paid dividends if I own shares? | 1198, 97942, 456470, 95889, 587689, 311214, 501931, 1034, 377007, 152014 |
| 5549 | 286227, 309361 | Pros / cons of being more involved with IRA investments [duplicate] | 429106, 105468, 336394, 561636, 181624, 32009, 222082, 505362, 396852, 364131 |
| 3801 | 307776 | Can a bunch of wealthy people force Facebook to go public? | 390529, 69017, 209242, 570634, 394734, 371293, 168565, 156903, 92014, 171236 |
| 4605 | 453941 | If the U.S. defaults on its debt, what will happen to my bank money? | 41312, 313306, 229310, 169691, 526384, 479527, 581054, 400826, 354896, 598030 |
| 2885 | 367360, 359579, 85229, 454810 | Merits of buying apartment houses and renting them | 430672, 451849, 159403, 4739, 315972, 358687, 343917, 80838, 164059, 150893 |
| 6110 | 94117, 259706 | Why does short selling require borrowing? | 188531, 320450, 226496, 384252, 35500, 67107, 501984, 79764, 84761, 247313 |
| 8 | 566392 | How to deposit a cheque issued to an associate in my business into my business account? | 65404, 508754, 261856, 590102, 564553, 188893, 308938, 309023, 536686, 489199 |
| 1309 | 156162, 489401 | Why does FlagStar Bank harass you about payments within grace period? | 471630, 489368, 271040, 336792, 75108, 173919, 151506, 329817, 526989, 15824 |
| 7109 | 447781 | How do I analyse moving averages? | 489933, 221627, 140804, 257185, 42620, 193012, 227669, 565501, 518932, 180428 |
| 5080 | 256055 | Is there a standard or best practice way to handle money from an expiring UTMA account? | 445521, 279291, 414429, 69841, 451189, 267206, 236186, 182645, 274834, 326451 |
| 4981 | 247894 | Where can I find open source portfolio management software? | 45218, 102684, 259463, 557861, 587792, 232736, 81865, 226628, 296401, 55845 |
| 7445 | 153178, 104343 | IS it the wrong time to get into the equity market immediately after large gains? | 89714, 33155, 573612, 590902, 350068, 483025, 284075, 356623, 114806, 127160 |
| 2895 | 328691 | Where should a young student put their money? | 426461, 354551, 332749, 256055, 36190, 123256, 148453, 55841, 502170, 579901 |
| 6787 | 587120 | Would it make sense to sell a stock, then repurchase it for tax purposes? | 219762, 400730, 23217, 263751, 106104, 311782, 221715, 17184, 38287, 474981 |
| 5862 | 130209 | Can I get a discount on merchandise by paying with cash instead of credit? | 21194, 301643, 503171, 495751, 170141, 299840, 535015, 418801, 420622, 557862 |
| 6041 | 241308 | Most effective Fundamental Analysis indicators for market entry | 425020, 81655, 96910, 528034, 108579, 194240, 263464, 542765, 115087, 5054 |
| 7700 | 273761, 2653, 179328 | Should I re-allocate my portfolio now or let it balance out over time? | 269169, 224392, 253268, 22221, 434014, 28425, 131127, 395208, 441176, 126836 |
| 547 | 6349 | What percentage of my company should I have if I only put money? | 68088, 396694, 156747, 523158, 498681, 95243, 80913, 559522, 445353, 569111 |
| 3394 | 129319 | What is the easiest way to back-test index funds and ETFs? | 172374, 408524, 71230, 188855, 159471, 528034, 507777, 99568, 503725, 445971 |
| 4102 | 448699 | How can I determine if my rate of return is “good” for the market I am in? | 597437, 554734, 135176, 554237, 484688, 369439, 46394, 461082, 535737, 88801 |

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
| 1889 | 0.0000 | 0.0000 | missing | Reporting financial gains from my online store |
| 34 | 0.0000 | 0.0000 | missing | 401k Transfer After Business Closure |
| 3534 | 0.0000 | 0.0000 | missing | Why do dishonour fees exist? |
