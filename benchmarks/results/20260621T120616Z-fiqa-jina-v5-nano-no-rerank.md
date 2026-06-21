# Mnemosyne BEIR fiqa Results

- Generated: 2026-06-21T12:06:16Z
- Dataset: `fiqa`
- Queries: 648
- Corpus documents: 57638
- Search limit: 100
- Rerank candidates: 100
- Rerank enabled: False
- Run label: `jina-v5-nano`
- Mnemosyne config: `/Users/gandazgul/Documents/web/mnemosyne/cli/configs/jina-v5-text-nano-retrieval.yaml`

| Metric | Score |
| --- | ---: |
| `ndcg@10` | 0.4614 |
| `mrr@10` | 0.5418 |
| `recall@10` | 0.5324 |
| `recall@100` | 0.7901 |
| `map@100` | 0.3996 |

## Breakdown

| Detail | Count |
| --- | ---: |
| `queries` | 648 |
| `queries_with_rank_1_hit` | 293 |
| `queries_with_top_10_hit` | 471 |
| `queries_with_top_100_hit` | 589 |
| `queries_missing_at_100` | 59 |

### First Relevant Rank

| Bucket | Queries |
| --- | ---: |
| `rank_1` | 293 |
| `rank_2_3` | 101 |
| `rank_4_10` | 77 |
| `rank_11_100` | 118 |
| `missing@100` | 59 |

### Missing At 100

| Query ID | Missing Doc IDs | Query | Top 10 Doc IDs |
| --- | --- | --- | --- |
| 4641 | 319954, 397358 | Where should I park my rainy-day / emergency fund? | 580025, 406219, 376148, 497993, 400503, 44594, 527939, 254572, 95598, 323686 |
| 7017 | 28271 | Basic Algorithmic Trading Strategy | 446629, 329182, 49345, 179068, 513055, 514562, 476887, 222363, 407385, 177347 |
| 3451 | 26292, 192307, 588448 | Should you keep your stocks if you are too late to sell? | 310683, 165970, 250873, 420974, 471911, 272091, 251536, 545284, 306460, 41852 |
| 753 | 243503 | Taxes due for hobbyist Group Buy | 132780, 451020, 283505, 466718, 466442, 293310, 172745, 547941, 540634, 313361 |
| 4465 | 376575 | How to donate to charity that will make a difference? | 371717, 106786, 375708, 132881, 134037, 90591, 379353, 46381, 174033, 326167 |
| 5951 | 497260 | Why can't house prices be out of tune with salaries | 599860, 31663, 418034, 259777, 62702, 374480, 373188, 557220, 596834, 529715 |
| 10827 | 160786, 7748, 107554, 95282 | How much should I be contributing to my 401k given my employer's contribution? | 290105, 296405, 576391, 240259, 312369, 436071, 452592, 38532, 42301, 41330 |
| 9771 | 263955, 28740 | Is there any emprical research done on 'adding to a loser' | 80481, 220190, 147581, 93714, 222639, 436765, 490479, 46825, 147848, 141475 |
| 3724 | 508921, 434190, 279570, 497216, 552887, 341146, 199970 | Should you always max out contributions to your 401k? | 302512, 216243, 328754, 430931, 122910, 3104, 140330, 481793, 38532, 309200 |
| 5853 | 476663, 160105, 495699, 424598 | Paying Off Principal of Home vs. Investing In Mutual Fund | 284318, 31525, 494148, 182612, 387722, 473647, 491923, 349355, 406325, 63883 |
| 10122 | 273718 | Why diversify stocks/investments? | 144261, 259084, 558088, 153989, 551719, 459970, 178875, 297100, 570787, 293679 |
| 1783 | 332314 | Freelancing Tax implication | 421924, 179359, 14609, 159709, 549870, 15270, 156063, 223624, 419319, 445298 |
| 4837 | 531841, 20958 | When applying for a mortgage, can it also cover outstanding debts? | 366448, 97925, 274870, 187739, 30311, 28230, 254454, 245005, 433171, 300660 |
| 4415 | 238234, 67676 | How much is inflation? | 117578, 553634, 184776, 519596, 32023, 511977, 189006, 593820, 53637, 468089 |
| 8378 | 125298 | Should I wait a few days to sell ESPP Stock? | 511678, 133644, 434812, 495568, 387035, 294573, 178684, 361345, 546125, 12232 |
| 8102 | 552707, 378173, 90294 | When do I sell a stock that I hold as a long-term position? | 171819, 306460, 35752, 422183, 537212, 457584, 165970, 310683, 250873, 368348 |
| 7928 | 118633 | If I believe a stock is going to fall, what options do I have to invest on this? | 501504, 427808, 260384, 410404, 24563, 173374, 30305, 102757, 292338, 47053 |
| 4777 | 590710 | How to finance necessary repairs to our home in order to sell it? | 365342, 570318, 482638, 118491, 26837, 487179, 478413, 432961, 293464, 52351 |
| 620 | 331332 | Is it wise to have plenty of current accounts in different banks? | 543921, 467059, 572848, 180673, 146317, 489959, 244097, 449630, 422908, 21420 |
| 8271 | 415511 | Income in zero-interest environment | 376709, 538898, 249558, 305539, 52047, 380368, 83330, 119298, 106424, 546874 |
| 864 | 211364, 152072 | Why use accounting software like Quickbooks instead of Excel spreadsheets? | 30142, 472924, 2436, 566337, 24890, 78117, 157751, 329774, 505361, 222380 |
| 5888 | 540806 | Interest charges on balance transfer when purchases are involved | 263647, 125497, 579601, 545327, 358445, 543776, 429746, 231105, 506132, 65537 |
| 2568 | 127353 | How to pay with cash when car shopping? | 258247, 346042, 264586, 9146, 355310, 465801, 269898, 15696, 514238, 108739 |
| 7124 | 74615 | How come we can find stocks with a Price-to-Book ratio less than 1? | 558617, 154725, 526110, 583708, 33357, 278582, 416732, 94591, 210268, 192686 |
| 2857 | 295864 | I have around 60K $. Thinking about investing in Oil, how to proceed? | 233732, 501384, 127566, 76996, 49235, 329930, 542051, 241800, 51848, 379140 |
| 8539 | 218728, 196304, 396038 | Can the risk of investing in an asset be different for different investors? | 391861, 483123, 469599, 283074, 292609, 113786, 471817, 128077, 240178, 387515 |
| 1085 | 467737, 393710 | How do disputed debts work on credit reports? | 268777, 398258, 319276, 384608, 1218, 450031, 161422, 372039, 189662, 223128 |
| 5422 | 151973 | What are some good books for learning stocks, bonds, derivatives e.t.c for beginner with a math background? | 172587, 165294, 273906, 276786, 221319, 191688, 552371, 79517, 13513, 325212 |
| 8247 | 321114, 465313 | Tax on Stocks or ETF's | 459953, 270992, 159076, 586010, 244813, 350317, 161019, 299690, 29502, 402046 |
| 10482 | 549072 | Rollover into bond fund to do dollar cost averaging [duplicate] | 18436, 330023, 447567, 434466, 175576, 564787, 439757, 224782, 265817, 127401 |
| 8789 | 70853 | What does “profits to the shareholders jumped to 15 cents a share” mean? | 87349, 41912, 341424, 573079, 144079, 20076, 325669, 317363, 411617, 14870 |
| 2423 | 538023 | At what age should I start or stop saving money? | 529444, 417787, 396127, 234846, 328157, 104457, 272328, 553288, 204479, 524414 |
| 8475 | 293320 | Why I cannot find a “Pure Cash” option in 401k investments? | 494655, 481728, 240562, 104793, 444107, 437427, 79375, 565432, 554739, 225395 |
| 4523 | 594257 | What should I do with my $25k to invest as a 20 years old? | 129255, 171712, 332203, 465819, 332022, 58065, 478242, 153231, 579901, 537711 |
| 8513 | 270573 | Buy on dip when earnings fail? | 572622, 175821, 351396, 384893, 53431, 335626, 221869, 432665, 106740, 28740 |
| 5054 | 28119 | How to stress test an investment plan? | 418281, 101124, 263390, 534010, 458183, 127263, 233717, 377186, 463831, 403916 |
| 1159 | 496064 | what is the best way to do a freelancing job over the summer for a student | 163881, 594182, 55064, 347181, 469972, 237282, 460648, 186889, 271812, 449155 |
| 68 | 19183 | Intentions of Deductible Amount for Small Business | 97719, 70315, 107213, 54333, 14255, 146657, 524134, 347935, 519473, 541809 |
| 9701 | 387141, 357739 | How to bet against the London housing market? | 473883, 408865, 258048, 70799, 516214, 588608, 108399, 225682, 367103, 150632 |
| 6199 | 239214 | How can all these countries owe so much money?  Why & where did they borrow it from? | 414693, 47163, 584273, 351853, 490042, 169921, 10399, 596959, 298794, 49602 |
| 9126 | 514831 | Short an option - random assignment? | 308859, 334473, 267113, 277760, 477588, 469382, 510006, 349974, 281533, 82194 |
| 34 | 599545 | 401k Transfer After Business Closure | 551545, 168890, 483268, 529845, 490867, 92941, 591168, 255277, 458917, 6471 |
| 6395 | 166227 | Option settlement for calendar spreads | 584223, 401447, 22916, 505223, 151546, 189144, 111301, 4845, 276314, 230355 |
| 9115 | 207325 | Why does the calculation for percentage profit vary based on whether a position is short vs. long? | 245082, 158520, 419897, 422467, 428786, 99131, 385220, 314478, 209359, 331850 |
| 4767 | 280805, 224057, 22804 | New car: buy with cash or 0% financing | 296612, 285033, 420018, 306834, 56867, 429153, 584106, 9146, 396817, 395590 |
| 8351 | 472516 | What happens when a calendar spread is assigned in a non-margin account? | 141213, 102316, 111301, 516790, 23609, 238075, 527654, 4845, 100628, 272754 |
| 6131 | 381720, 170204, 416679, 2460 | Is it ever a good idea to close credit cards? | 391384, 102823, 326094, 368806, 533288, 339030, 218088, 35625, 334111, 99449 |
| 858 | 122485, 278450 | Is it bad practice to invest in stocks that fluctuate by single points throughout the day? | 293027, 433730, 519501, 146632, 514780, 289, 216964, 310683, 241860, 538352 |
| 4019 | 287991, 6881, 125477 | How and Should I Invest (As a college 18 year old with minimal living expenses)? | 85977, 269671, 332938, 426461, 379948, 571044, 269384, 347849, 332749, 211713 |
| 6080 | 164513 | Is ScholarShare a legitimate entity for a 529 plan in California? | 22856, 236732, 201500, 468527, 277581, 115175, 233401, 83080, 380557, 535357 |
| 6959 | 205010 | What is the term for the quantity (high price minus low price) for a stock? | 468025, 116846, 33357, 472537, 599523, 208331, 280172, 28604, 577573, 340607 |
| 5402 | 491350 | Is it impossible to get a home loan with a poor credit history after a divorce? | 445163, 227485, 51728, 595029, 44105, 90579, 129177, 52250, 67066, 289231 |
| 6562 | 501157 | Cheapest way to “wire” money in an Australian bank account to a person in England, while I'm in Laos? | 473605, 324817, 582414, 428696, 72965, 254257, 135675, 183880, 182443, 60446 |
| 701 | 389446 | What are the ins/outs of writing-off part of one's rent for working at home? | 231990, 436505, 339488, 456234, 337706, 344955, 288537, 177074, 443859, 509862 |
| 7674 | 519390 | Choosing the limit when making a limit order? | 249279, 546150, 155151, 514841, 278630, 94653, 437436, 184756, 447886, 236133 |
| 5940 | 486243 | How does investment into a private company work? | 46842, 507494, 469043, 535314, 127434, 454465, 250354, 512609, 473154, 530 |
| 6612 | 205522, 322900 | If I have a lot of debt and the housing market is rising, should I rent and slowly pay off my debt or buy and roll th... | 198442, 393002, 553277, 254454, 180192, 101363, 290900, 386348, 433171, 28230 |
| 8456 | 486333 | What typically happens to unvested stock during an acquisition? | 257853, 534755, 174321, 391156, 261487, 93215, 235779, 475019, 498676, 518340 |
| 10213 | 270221, 545712 | Looking for good investment vehicle for seasonal work and savings | 446186, 223551, 100517, 589394, 38269, 388252, 94680, 390556, 58065, 112499 |
| 5196 | 172128, 114829 | I might use a credit card convenience check. What should I consider? | 565745, 316652, 393866, 498775, 62862, 418580, 307807, 302823, 454842, 149032 |
| 3006 | 269851, 568473, 328300 | Strategies for putting away money for a child's future (college, etc.)? | 512096, 258704, 127838, 471019, 211713, 360285, 490382, 45451, 597627, 495282 |
| 3909 | 312248, 404356, 245616, 353028 | How to rescue my money from negative interest? | 514003, 288656, 61586, 83330, 526169, 364361, 472837, 12133, 346064, 362730 |
| 6907 | 251604 | Nominal value of shares | 427859, 546548, 303112, 400738, 69506, 207253, 487738, 312811, 170652, 285606 |
| 5464 | 86691 | Resources on Buying Rental Properties | 294549, 372274, 44855, 222095, 325722, 155964, 536126, 383921, 315972, 16013 |
| 5090 | 436493 | Should I take a student loan to pursue my undergraduate studies in France? | 12988, 586289, 217831, 213328, 246286, 551964, 92430, 317662, 117085, 455666 |
| 2088 | 599524 | How would I go about selling the stock of a privately held company? | 140835, 72846, 293687, 53993, 432727, 291886, 581127, 473798, 178521, 455168 |
| 9391 | 503637 | Should I replace bonds in a passive investment strategy | 136515, 535518, 283202, 342485, 466403, 112369, 248158, 269055, 107424, 20354 |
| 3148 | 178127, 438000 | Can a car company refuse to give me a copy of my contract or balance details? | 172855, 29721, 584305, 430100, 357280, 65046, 205984, 92888, 560325, 164702 |
| 4678 | 305153 | Finance, Cash or Lease? | 376016, 504918, 307158, 519950, 427884, 536773, 386151, 539805, 185405, 311748 |
| 2398 | 224654, 590489 | Frustrated Landlord | 556453, 168089, 44058, 96538, 310992, 504703, 146603, 393883, 422579, 487094 |
| 5511 | 169893, 560325, 383193, 114303, 278699, 12746 | Pay off car loan entirely or leave $1 until the end of the loan period? | 529123, 155843, 500946, 324269, 334559, 38786, 27693, 107898, 139788, 82952 |
| 8834 | 521095 | Pros/Cons of Buying Discounted Company Stock | 203139, 569303, 569224, 371821, 469100, 599156, 67625, 349567, 528827, 212470 |
| 988 | 226053, 107688 | Where should I invest my savings? | 531965, 82482, 579848, 179702, 320675, 574011, 571218, 589088, 194080, 501384 |
| 3369 | 231012, 145716, 411910, 395840 | Why should one only contribute up to the employer's match in a 401(k)? | 296405, 555377, 92370, 443002, 15841, 240259, 242556, 436071, 341493, 576391 |
| 5763 | 462019 | What is the best way to get a “rough” home appraisal prior to starting the refinance process? | 515361, 326214, 218144, 570318, 38712, 11572, 324874, 366591, 129439, 89964 |
| 4962 | 599925 | Net Cash Flows from Selling the Bond and Investing | 158363, 52149, 537603, 510268, 416839, 154707, 395208, 408661, 196173, 543842 |
| 9403 | 6666, 328086, 345199 | Abundance of Cash - What should I do? | 410450, 14349, 103447, 186332, 570632, 159235, 215296, 357887, 372223, 551986 |
| 5993 | 272866, 55084, 352638, 426120, 63501 | Why would anyone want to pay off their debts in a way other than “highest interest” first? | 94373, 287571, 160193, 128574, 156195, 431212, 544858, 416796, 353911, 886 |
| 5710 | 232311 | Bucketing investments to track individual growths | 227364, 516267, 11979, 399738, 7748, 353657, 88417, 492506, 583549, 54190 |
| 7529 | 66607 | Does the expense ratio of a fund-of-funds include the expense ratios of its holdings? | 514529, 464337, 518402, 293626, 89297, 287537, 387980, 451855, 218261, 59249 |
| 5021 | 589285 | Is there a more flexible stock chart service, e.g. permitting choice of colours when comparing multiple stocks? | 584801, 528576, 465971, 105717, 423177, 49893, 60284, 471643, 517935, 45218 |
| 3612 | 259625 | How can I buy and sell the same stock on the same day? | 547553, 360059, 165548, 336018, 567383, 483676, 353396, 525231, 460937, 538743 |
| 4409 | 499128, 100306, 147439 | My friend wants to put my name down for a house he's buying. What risks would I be taking? | 243732, 102088, 223841, 360682, 244278, 531299, 268078, 546288, 129855, 102326 |
| 5369 | 171339 | Paying for things on credit and immediately paying them off: any help for credit rating? | 540959, 572426, 409927, 319773, 153088, 55162, 529229, 448614, 114303, 577542 |
| 2070 | 363678 | Advantage of credit union or local community bank over larger nationwide banks such as BOA, Chase, etc.? | 578357, 469515, 597571, 249839, 587737, 550303, 261389, 197087, 258581, 366475 |
| 11039 | 53544, 249063 | Pay off credit card debt or earn employer 401(k) match? | 91183, 287876, 508534, 79363, 345895, 281049, 105557, 552383, 5203, 163287 |
| 5460 | 93248, 184337, 21174, 108514, 463885 | Paying off a loan with a loan to get a better interest rate | 243065, 343208, 139788, 399259, 482798, 196237, 5188, 58432, 94373, 197313 |
| 7925 | 318185, 251100, 503912, 402482 | Can I sell a stock immediately? | 44461, 438974, 457294, 584291, 66210, 133644, 95415, 219762, 17661, 407602 |
| 4286 | 566069 | Given advice “buy term insurance and invest the rest”, how should one “invest the rest”? | 70460, 511386, 229239, 79142, 206830, 221479, 391243, 56732, 207391, 107780 |
| 2685 | 154113, 370300, 303293, 468923 | What ways are there for us to earn a little extra side money? | 382005, 468086, 186889, 269380, 543275, 4992, 113690, 558832, 594182, 316925 |
| 1090 | 518896 | Need a formula to determine monthly payments received at time t if I'm reinvesting my returns | 446454, 296146, 573928, 521590, 203091, 435096, 198606, 179365, 281329, 395128 |
| 6122 | 44344 | Better to rent condo to daughter or put her on title? | 496166, 403515, 316794, 558251, 118246, 182039, 423398, 101816, 577658, 566184 |
| 4514 | 69485, 337764, 209804 | What intrinsic, non-monetary value does gold have as a commodity? | 471825, 426270, 317429, 156211, 99089, 408336, 352485, 532381, 180404, 479398 |
| 8507 | 509819, 370995 | When to sell a stock? | 88813, 251536, 303724, 236415, 273565, 351615, 102237, 554568, 234361, 89460 |
| 6221 | 257248, 76414, 455614, 115717 | To pay off a student loan, should I save up a lump sum payoff payment or pay extra each month? | 414534, 110081, 448791, 541313, 69150, 394474, 571198, 414288, 352363, 27625 |
| 3008 | 180192, 323406 | What are my chances at getting a mortgage with Terrible credit but High income | 231688, 251846, 227485, 47441, 44105, 394460, 364802, 597265, 407401, 455952 |
| 4007 | 521657 | What is a reasonable salary for the owner and sole member of a small S-Corp? | 556220, 260385, 205341, 515233, 370542, 521933, 388704, 315552, 249322, 170933 |
| 6644 | 175035 | How to know precisely when a SWIFT is issued by a bank? | 110198, 475527, 41383, 218761, 271596, 355870, 327623, 429172, 187497, 39783 |
| 10267 | 460398 | How should I prepare for the next financial crisis? | 178693, 569632, 143393, 87520, 326398, 436091, 577806, 33391, 279456, 540553 |
| 7622 | 253369, 378594 | Best way to pay off debt? | 480773, 345895, 373554, 340484, 460225, 136047, 271525, 130812, 353911, 457945 |
| 3767 | 153922, 392060 | What should I be doing to protect myself from identity theft? | 90632, 125204, 260580, 368679, 97686, 423809, 171510, 520395, 581889, 59228 |
| 6410 | 471723 | Will an ETF immediately reflect a reconstitution of underlying index | 227324, 454610, 87261, 71230, 147282, 408524, 428187, 120059, 200360, 46671 |
| 5030 | 215540 | Why pay for end-of-day historical prices? | 13511, 327974, 149420, 227192, 295344, 378994, 560108, 127015, 440959, 347992 |
| 6252 | 394551, 160932, 293624, 233294, 243268, 379487, 62868 | Is this mortgage advice good, or is it hooey? | 213713, 316149, 104988, 272634, 373966, 118532, 66676, 455952, 300660, 426638 |
| 885 | 337165, 409184 | How long do credit cards keep working after you disappear? | 516678, 89888, 254968, 59228, 301792, 588719, 251643, 288657, 296165, 472336 |
| 4031 | 115741 | 28 years old and just inherited large amount of money and real estate - unsure what to do with it | 318864, 140002, 163197, 375708, 171196, 568629, 257840, 193171, 65180, 531665 |
| 766 | 550172 | Will the ex-homeowner still owe money after a foreclosure? | 2996, 552768, 299591, 249788, 6882, 27987, 56126, 554171, 104955, 333583 |
| 8202 | 513258, 93971 | What accounted for DXJR's huge drop in stock price? | 457689, 162047, 41271, 5220, 122542, 83391, 253339, 129481, 18335, 245654 |
| 7345 | 237645 | What do these numbers mean? (futures) | 9274, 508821, 366526, 527080, 108, 164001, 529996, 475916, 460331, 206895 |
| 776 | 583640, 127263, 597880, 496899 | Can saving/investing 15% of your income starting age 25, likely make you a millionaire? | 124027, 10440, 467044, 143591, 191148, 417787, 41960, 554833, 319760, 449745 |
| 89 | 248624, 64556 | How can I deposit a check made out to my business into my personal account? | 508754, 188893, 526817, 308938, 309023, 65404, 29372, 181187, 590102, 45078 |
| 6629 | 444405 | Tax treatment of a boxed trade? | 300486, 295153, 385310, 30403, 103758, 318741, 589377, 530066, 127004, 123287 |
| 1889 | 388713 | Reporting financial gains from my online store | 584074, 243503, 355959, 352136, 427202, 599876, 128861, 172652, 33287, 333954 |
| 8013 | 496159, 224231 | Frequency of investments to maximise returns (and minimise fees) | 48678, 384983, 81652, 446948, 278607, 564787, 24846, 254910, 153178, 554237 |
| 3759 | 67167, 522358 | Simplifying money management | 455457, 490065, 551986, 591704, 248663, 526568, 214248, 250397, 197506, 218644 |
| 7295 | 244749 | Selling non-dividend for dividend stocks | 352484, 170318, 509879, 71511, 292559, 217837, 193303, 501931, 527636, 235391 |
| 10639 | 431799, 495774, 278453, 187039 | Short term parking of a large inheritance? | 163353, 318864, 171196, 375708, 140002, 405115, 522510, 422994, 331534, 360816 |
| 6635 | 156358 | Why don't share prices of a company rise every other Friday when the company buys shares for its own employees? | 587137, 3656, 533712, 343452, 579037, 235531, 125298, 467594, 12560, 569224 |
| 6525 | 181985 | Does it make sense to trade my GOOGL shares for GOOG and pocket the difference? | 106541, 105542, 221795, 550661, 192696, 98150, 378906, 504579, 76330, 362473 |
| 2590 | 589625 | Are non-residents or foreigners permitted to buy or own shares of UK companies? | 296528, 48269, 55999, 209493, 158923, 458730, 188776, 35191, 456999, 562007 |
| 5374 | 152688 | What were the main causes of the spike and drop of DRYS's stock price? | 283106, 457689, 122542, 253339, 162047, 133204, 5220, 589533, 380894, 490899 |
| 2994 | 419319 | Work on the side for my wife's company | 510317, 148346, 477357, 138428, 321743, 569145, 562780, 460721, 523372, 318491 |
| 3683 | 454501 | Can I trust the Motley Fool? | 131381, 276975, 240628, 105973, 408995, 114981, 538086, 565016, 428848, 500338 |
| 7206 | 441155, 532211, 553066 | Who Bought A Large Number Of Shares? | 65667, 260085, 367960, 558703, 351570, 554996, 100128, 364674, 158923, 54424 |
| 10246 | 77573 | Understanding the T + 3 settlement days rule | 156029, 370635, 327080, 176717, 293389, 315205, 340263, 28314, 332243, 11927 |
| 5241 | 376123, 322157, 27489 | Mortgage vs. Cash for U.S. home buy now | 281675, 111184, 344740, 273735, 107350, 234286, 210972, 213713, 309420, 504479 |
| 98 | 575929 | How can I make $250,000.00 from trading/investing/business within 5 years? | 527522, 555630, 438279, 519619, 227568, 449745, 373119, 596087, 465819, 66034 |
| 4615 | 262934 | Are solar cell panels and wind mills worth the money? | 261900, 69523, 496427, 455798, 120384, 425595, 176265, 271015, 385028, 158216 |
| 3264 | 598807, 134764 | Pros and Cons of Interest Only Loans | 486525, 260383, 453847, 316230, 7243, 400009, 106424, 526169, 322194, 25466 |
| 6467 | 453256, 23217, 346641, 367313 | Advice on strategy for when to sell | 109455, 498075, 250873, 368348, 76996, 303724, 320587, 251536, 295912, 52351 |
| 4394 | 441582 | Transfer $50k to another person's account (in California, USA) | 524752, 132693, 80657, 281000, 313141, 196365, 286866, 426216, 412258, 488566 |
| 4047 | 77652 | Does doing your “research”/“homework” on stocks make any sense? | 417840, 324779, 452175, 110733, 170966, 284165, 210241, 358846, 359718, 474296 |
| 2407 | 173929 | How long to wait after getting a mortgage to increase my credit limit? | 448368, 264341, 26397, 294327, 257644, 289231, 263999, 585654, 348313, 2064 |
| 7344 | 108403 | How is the Dow divisor calculated? | 313421, 159166, 14368, 378974, 150430, 65618, 96697, 591089, 169754, 253926 |
| 6875 | 224392 | Where to find free Thailand stock recommendations and research? | 567500, 352557, 512895, 224366, 110733, 79337, 105717, 408995, 249679, 500338 |
| 10447 | 152096, 300721 | Is there an advantage to a traditional but non-deductable IRA over a taxable account? [duplicate] | 144751, 500175, 382236, 259150, 532657, 114912, 382894, 26652, 588134, 540389 |
| 5782 | 595455 | Pay off credit cards in one lump sum, or spread over a few months? | 262026, 117007, 487621, 440620, 559523, 273631, 114592, 113822, 529312, 508453 |
| 9871 | 448890, 40051, 170594 | What should I do with the 50k I have sitting in a European bank? | 367207, 292714, 433003, 387723, 231521, 250800, 74668, 73741, 174406, 219477 |
| 3625 | 384469, 414295 | What should I do with my paper financial documents? | 500751, 509617, 569812, 344236, 380263, 513248, 305065, 113830, 206466, 63422 |
| 6005 | 478457, 345895, 73310, 384626, 390689 | Why might it be advisable to keep student debt vs. paying it off quickly? | 507544, 149500, 96268, 571198, 572272, 270856, 519075, 117085, 431884, 176498 |
| 3822 | 385090, 418900, 308837 | How to change a large quantity of U.S. dollars into Euros? | 541608, 292714, 239876, 174406, 144181, 417917, 478065, 200487, 208499, 369266 |
| 7879 | 372551, 421285 | Any Tips on How to Get the Highest Returns Within 4 Months by Investing in Stocks? | 58186, 43088, 477295, 377186, 59394, 272174, 102029, 388252, 105391, 367415 |
| 3115 | 234950, 389028, 316794 | How can I live outside of the rat race of American life with 300k? | 233562, 183869, 267892, 129364, 366961, 559852, 174272, 271818, 369742, 188750 |
| 3995 | 230208 | I have more than $250,000 in a US Bank account… mistake? | 485507, 479918, 171720, 403288, 297900, 264934, 404954, 156716, 14349, 463608 |
| 10136 | 526115 | How to minimise the risk of a reduction in purchase power in case of Brexit for money held in a bank account? | 290930, 466950, 168137, 314850, 393823, 583903, 304007, 25582, 265453, 594157 |
| 8635 | 67107, 240215 | Is there any flaw in this investment scheme? | 144200, 575918, 46818, 3750, 264822, 447619, 303619, 121798, 343518, 365816 |
| 5206 | 563030, 28230, 201982, 117276 | Is it a good idea to get an unsecured loan to pay off a credit card that won't lower a high rate? | 540959, 298908, 595455, 287157, 498728, 60996, 585288, 190553, 340520, 292038 |
| 2713 | 388147 | Physical Checks - Mailing | 584170, 29372, 564301, 284528, 229546, 324717, 41944, 487576, 165691, 333292 |
| 9060 | 40447 | Buying puts without owning underlying | 511093, 528052, 228217, 345851, 359778, 576364, 222498, 100628, 334473, 415705 |
| 4105 | 25096 | As an investor what are side effects of Quantitative Easing in US and in EU? | 185300, 416483, 459400, 345910, 489478, 598177, 134213, 293104, 467327, 400032 |
| 2465 | 570680, 81046 | Can capital expenses for volunteer purposes be deducted from income? | 37382, 398536, 310612, 432545, 18889, 107213, 510716, 480512, 202645, 121187 |
| 4640 | 101369 | What can my relatives do to minimize their out of pocket expenses on their fathers estate | 360816, 356035, 295246, 422994, 375708, 331534, 144965, 17110, 565428, 421652 |
| 9275 | 338754 | Do I have to pay a capital gains tax if I rebuy the same stock within 30 days? | 400730, 390864, 343219, 376800, 102443, 448659, 522319, 263751, 423929, 23217 |
| 1306 | 484437, 204075 | I made an investment with a company that contacted me, was it safe? | 538086, 205665, 309851, 567973, 450779, 124687, 428552, 407663, 112259, 537698 |
| 4500 | 533623 | What to ask Warren Buffet at the Berkshire Hathaway shareholder meeting? | 404341, 113800, 461355, 153649, 39149, 520677, 3533, 375671, 41809, 100283 |
| 8632 | 213976 | Is it best to exercise options shares when they vest, or wait | 43497, 340730, 420722, 178497, 488207, 388362, 382381, 576503, 334473, 595787 |
| 9979 | 337243 | What is the best way to invest in gold as a hedge against inflation without having to hold physical gold? | 35369, 13885, 96351, 10578, 483734, 556936, 30584, 326858, 399751, 572670 |
| 6133 | 415705 | What happens to all of the options when they expire? | 73256, 7733, 11456, 481070, 132288, 581672, 575408, 463254, 358492, 72694 |
| 3771 | 488948, 198349, 217683, 49601 | Best way to buy Japanese yen for travel? | 521712, 495826, 128471, 96211, 490384, 152695, 402006, 434201, 350245, 306130 |
| 1736 | 25543, 443419 | How can people have such high credit card debts? | 327441, 399406, 437610, 475668, 463065, 235646, 562896, 291024, 382591, 372993 |
| 6814 | 340214, 223206 | Selling Stock - All or Nothing? | 279782, 66834, 178497, 590188, 198583, 369166, 513734, 878, 250873, 137073 |
| 1322 | 64138 | Is this follow-up after a car crash a potential scam? | 114231, 318941, 283917, 397852, 91463, 33914, 567973, 264986, 98356, 219648 |
| 5185 | 210236, 317354 | Invest in low cost small cap index funds when saving towards retirement? | 503725, 376485, 262180, 59670, 196992, 434279, 545760, 241202, 524525, 312406 |
| 2348 | 211867, 566573, 211622, 474234, 352271, 265874, 134864 | Why can't you just have someone invest for you and split the profits (and losses) with him? | 389004, 247486, 447619, 420544, 309851, 197389, 381757, 306430, 387344, 420379 |
| 4499 | 323363, 76996 | Is investing exlusively in a small-cap index fund a wise investment? | 501153, 196992, 517391, 238963, 235917, 9512, 445322, 14748, 526859, 406768 |
| 3530 | 239998 | How to exclude stock from mutual fund | 24029, 378075, 184299, 209879, 449124, 574383, 332152, 287537, 475640, 479420 |
| 659 | 584685 | Buying from an aggressive salesperson | 168796, 449079, 120279, 429119, 13139, 365240, 519961, 264297, 466122, 291090 |
| 9108 | 272021, 472585 | Starting an investment portfolio with Rs 5,000/- | 290757, 312821, 356552, 51848, 171189, 323067, 122679, 527522, 548467, 465819 |
| 6835 | 102243 | Are bond ETF capital gains taxed similar to stock or stock funds if held for more than 1 year? | 149305, 5710, 84238, 543842, 29502, 287950, 225536, 153112, 586010, 137299 |
| 3067 | 406156 | Should I make extra payments to my under water mortgage or increase my savings? | 560915, 477907, 468831, 341837, 423403, 90009, 476068, 131365, 58588, 440719 |
| 4125 | 344648, 72046 | Alternative means of salary for my employees | 365558, 396255, 414694, 302310, 65795, 287200, 36608, 174787, 245451, 58906 |
| 1150 | 43603, 19936 | How are the best way to make and save money at 22 years old | 529444, 319760, 353369, 328157, 595287, 204479, 66864, 147745, 10476, 365167 |
| 7431 | 372921 | Pay off mortgage or invest in high value saving account | 2393, 589256, 557734, 468831, 440719, 364099, 220733, 537721, 157414, 403314 |
| 7747 | 97729 | What happens to bonds values when interest rates rise? [duplicate] | 296420, 559157, 34925, 74287, 31581, 29073, 212732, 127082, 180958, 274818 |
| 9808 | 40702, 431946 | Selling To Close | 416307, 557582, 151587, 345368, 362473, 374204, 107045, 149420, 399394, 229573 |
| 3888 | 319213, 239632 | Why I can't view my debit card pre-authorized amounts? | 208169, 185434, 294077, 448086, 521010, 440527, 281129, 168283, 146909, 209718 |
| 10109 | 506374, 156029, 499849 | Why does Charles Schwab have a Mandatory Settlement Period after selling stocks? | 93231, 28314, 332243, 370635, 121465, 98302, 119161, 293389, 124188, 584291 |
| 2790 | 279329, 27268, 469125 | Should I pay more than 20% down on a home? | 357200, 64400, 296906, 234890, 145862, 286656, 472484, 400896, 207564, 385343 |
| 9882 | 65702 | Money-market or cash-type ETFs for foreigners with U.S brokerage account | 389581, 188524, 386173, 98461, 535340, 161966, 529638, 528880, 551809, 508895 |
| 4999 | 314898 | Looking for a good source for Financial Statements | 356515, 9938, 171964, 11263, 431459, 146076, 338803, 520165, 295738, 516379 |
| 3189 | 225395 | Diversify my retirement investments with a Roth IRA | 287225, 423658, 240975, 404800, 458168, 553031, 46986, 403556, 272840, 88311 |
| 5134 | 158523 | Why does Yahoo Finance's data for a Vanguard fund's dividend per share not match the info from Vanguard? | 374330, 46774, 532616, 60098, 584128, 215486, 206727, 221477, 263088, 239137 |
| 1321 | 216456 | Are social media accounts (e.g. YouTube, Twitter, Instagram, etc.) considered assets? | 351672, 381322, 278538, 235082, 481036, 64727, 369884, 203715, 561226, 100849 |
| 4539 | 370879 | How should I save money if the real interest rate (after inflation) is negative? | 275925, 585494, 194080, 479659, 42475, 589286, 514003, 472837, 328499, 61586 |
| 715 | 579763, 546538, 187404 | what would you do with $100K saving? | 273925, 200273, 133120, 121621, 548758, 372223, 260677, 556545, 113885, 190929 |
| 504 | 500755, 344203, 498751 | Have plenty of cash flow but bad credit | 77248, 22807, 546097, 486334, 99463, 503419, 432040, 52250, 368247, 74448 |
| 2296 | 83330, 366594 | How does a bank make money on an interest free secured loan? | 462036, 400009, 94230, 119298, 580147, 106424, 172303, 249831, 396853, 344165 |
| 10975 | 61022 | How to contribute to Roth IRA when income is at the maximum limit & you have employer-sponsored 401k plans? | 353009, 441632, 110114, 446226, 387338, 75766, 270818, 452592, 261369, 292230 |
| 1994 | 156640 | Does the IRS reprieve those who have to commute for work? | 434846, 585706, 231990, 51491, 147624, 192843, 63919, 319965, 585562, 263259 |
| 9164 | 365298, 263390 | Bonds vs equities: crash theory | 115648, 149900, 16924, 506743, 309326, 599420, 305274, 321941, 409859, 303037 |
| 10462 | 8266, 11378, 35680, 437879, 204035, 581204 | Is it okay to be married, 30 years old and have no retirement? | 66376, 593377, 391583, 122333, 152478, 386404, 564860, 139595, 51914, 221295 |
| 8855 | 208165, 312821 | How do i get into investing stocks [duplicate] | 363719, 312445, 313570, 155677, 403092, 484327, 274400, 560395, 328477, 525318 |
| 7071 | 124230 | ESPP strategy - Sell right away or hold? | 511678, 434812, 133644, 361345, 294573, 495568, 387035, 71713, 575213, 12232 |
| 3534 | 340329 | Why do dishonour fees exist? | 149845, 517497, 312349, 34108, 198599, 269584, 234260, 286064, 166431, 333149 |
| 8974 | 356595, 170625 | As a 22-year-old, how risky should I be with my 401(k) investments? | 134931, 216365, 102501, 452126, 10476, 96110, 199970, 327925, 324066, 53100 |
| 5061 | 23747 | What fiscal scrutiny can be expected from IRS in early retirement? | 31699, 484683, 598378, 253735, 537049, 25481, 493461, 200914, 170916, 205232 |
| 2075 | 170042, 14967 | Are stories of turning a few thousands into millions by trading stocks real? | 523393, 519619, 375876, 44417, 555521, 506149, 249055, 438279, 39927, 41926 |
| 7533 | 93853 | Investing tax (savings) | 8012, 526664, 171596, 371176, 316230, 315105, 516756, 142658, 563986, 262772 |
| 1393 | 352838, 539133 | Which is better when working as a contractor, 1099 or incorporating? | 220022, 352640, 234436, 68524, 32072, 194899, 277812, 510913, 194955, 209974 |
| 9733 | 110163, 38655 | Due Diligence - Dilution? | 450132, 23414, 526073, 135798, 108965, 267266, 301880, 419505, 316321, 154841 |
| 7311 | 323768 | Finance, Social Capital IPOA.U | 479752, 579110, 507841, 419735, 538443, 116804, 299021, 482228, 401004, 290325 |
| 744 | 566480 | What options are available for a home loan with poor credit but a good rental history? | 490443, 67066, 415425, 80607, 313623, 47441, 92397, 256981, 289231, 231688 |
| 7141 | 132288 | Do investors go long option contracts when they cannot cover the exercise of the options? | 538054, 557356, 293767, 255927, 41967, 334473, 31587, 357324, 576364, 507828 |
| 4071 | 129875 | If our economy crashes, and cash is worthless, should i buy gold or silver | 505136, 487817, 291862, 506780, 53538, 502634, 302010, 524142, 135113, 308332 |
| 7512 | 191060 | understanding the process/payment of short sale dividends | 487329, 222320, 568166, 298284, 13631, 115553, 202985, 480949, 501931, 409432 |
| 1391 | 562176 | How is taxation for youtube/twitch etc monetization handled in the UK? | 440745, 77171, 454208, 223170, 378060, 254151, 111131, 245753, 527951, 510599 |
| 7534 | 358125 | Can you explain why it's better to invest now rather than waiting for the market to dip? | 175821, 145539, 114806, 33155, 103622, 350068, 474006, 310218, 14543, 222444 |
| 5356 | 312405 | Historical stock prices: Where to find free / low cost data for offline analysis? | 596106, 279785, 240086, 529877, 189341, 71553, 546379, 47798, 167586, 466255 |
| 2579 | 432020 | What to do when a job offer is made but with a salary less than what was asked for? | 423070, 190077, 283825, 524471, 200946, 256802, 364159, 559900, 68081, 157919 |
| 7823 | 583549 | Retirement Funds: Betterment vs Vanguard Life strategy vs Target Retirement | 451196, 105666, 175927, 331492, 268731, 172336, 293679, 329425, 571217, 347825 |
| 9174 | 535317, 160218 | Which U.S. online discount broker is the best value for money? | 236931, 192910, 544576, 563334, 274987, 31936, 158091, 405217, 200052, 110608 |
| 6867 | 443804, 445258, 538750 | Will there always be somebody selling/buying in every stock? | 573077, 466143, 301985, 321639, 400614, 482739, 429196, 560273, 349147, 224672 |
| 5083 | 138845 | Co-signer deceased | 305509, 369075, 447983, 518681, 270952, 288701, 253697, 18257, 495482, 273759 |
| 10526 | 39185 | What extra information might be obtained from the next highest bids in an order book? | 283008, 467852, 146125, 546493, 577573, 427747, 137175, 281844, 251100, 557770 |
| 5903 | 231863 | Fees aside, what factors could account for performance differences between U.S. large-cap index ETFs? | 159471, 408524, 246996, 14185, 20504, 230997, 395842, 148721, 301580, 358997 |
| 5620 | 448784, 329552, 548740 | What's the fuss about identity theft? | 260580, 98993, 90632, 158285, 440524, 547189, 423809, 5860, 208402, 551747 |
| 2472 | 370334 | How do I deal with a mistaken attempt to collect a debt from me that is owed by someone else? | 180601, 49321, 161422, 546028, 201758, 435006, 144922, 330507, 233535, 494116 |
| 2306 | 315875 | To whom should I report fraud on both of my credit cards? | 581889, 531137, 289706, 360586, 596284, 139518, 125204, 90632, 226590, 270449 |
| 7633 | 197839 | Can a trade happen “in between” the bid and ask price? | 494727, 353396, 458933, 505244, 560558, 1203, 281844, 402482, 137175, 179258 |
| 2400 | 564271 | Will I be paid dividends if I own shares? | 91870, 1198, 188531, 97942, 453582, 377007, 587689, 350110, 184077, 152014 |
| 5549 | 286227, 309361 | Pros / cons of being more involved with IRA investments [duplicate] | 105468, 181624, 561636, 32009, 336394, 222082, 505362, 364131, 299690, 396852 |
| 3801 | 307776 | Can a bunch of wealthy people force Facebook to go public? | 390529, 69017, 570634, 156903, 394734, 371293, 462517, 298666, 92014, 171385 |
| 4605 | 453941 | If the U.S. defaults on its debt, what will happen to my bank money? | 313306, 41312, 229310, 581054, 479527, 169691, 354896, 526384, 598030, 373717 |
| 2885 | 367360, 414692, 359579, 85229, 454810 | Merits of buying apartment houses and renting them | 430672, 451849, 315972, 4739, 358687, 538062, 150893, 343917, 483285, 159403 |
| 6110 | 94117, 259706 | Why does short selling require borrowing? | 188531, 384252, 226496, 320450, 35500, 501984, 84761, 67107, 79764, 247313 |
| 1309 | 156162, 489401 | Why does FlagStar Bank harass you about payments within grace period? | 471630, 489368, 271040, 75108, 336792, 173919, 234632, 151506, 181855, 568636 |
| 7109 | 447781 | How do I analyse moving averages? | 221627, 489933, 257185, 140804, 193012, 42620, 227669, 180428, 518932, 565501 |
| 5080 | 256055 | Is there a standard or best practice way to handle money from an expiring UTMA account? | 445521, 279291, 182645, 326451, 274834, 69841, 267206, 414429, 523952, 223652 |
| 4981 | 247894 | Where can I find open source portfolio management software? | 45218, 102684, 259463, 81865, 226628, 232736, 296401, 55845, 557861, 587792 |
| 7445 | 153178, 104343 | IS it the wrong time to get into the equity market immediately after large gains? | 89714, 33155, 356623, 350068, 114806, 590902, 573612, 284075, 103622, 483025 |
| 2895 | 521996, 328691 | Where should a young student put their money? | 444568, 256055, 332749, 354551, 400500, 36190, 571044, 426461, 141032, 123256 |
| 6787 | 587120 | Would it make sense to sell a stock, then repurchase it for tax purposes? | 219762, 400730, 23217, 311782, 106104, 38287, 263751, 448659, 468047, 156092 |
| 1748 | 576295 | How high should I set my KickStarter funding goal in order to have $35,000 left over? | 18001, 401505, 47949, 451492, 61485, 425435, 429555, 527373, 234506, 367957 |
| 5862 | 130209 | Can I get a discount on merchandise by paying with cash instead of credit? | 503171, 170141, 301643, 21194, 299840, 420622, 495751, 394658, 535015, 557862 |
| 6041 | 241308 | Most effective Fundamental Analysis indicators for market entry | 425020, 81655, 5054, 263464, 323749, 115087, 542765, 194240, 528034, 108579 |
| 7700 | 273761, 2653, 179328 | Should I re-allocate my portfolio now or let it balance out over time? | 269169, 224392, 28425, 22221, 253268, 434014, 507468, 470758, 441176, 588607 |
| 547 | 6349 | What percentage of my company should I have if I only put money? | 68088, 96110, 156747, 160170, 368587, 396694, 559522, 498681, 523158, 119077 |
| 3394 | 129319 | What is the easiest way to back-test index funds and ETFs? | 172374, 188855, 408524, 507777, 159471, 503725, 71230, 99568, 386299, 276593 |
| 4102 | 448699 | How can I determine if my rate of return is “good” for the market I am in? | 597437, 554734, 135176, 554237, 484688, 369439, 46394, 249360, 461082, 559168 |

### Lowest MRR@10

| Query ID | MRR@10 | Recall@100 | First Relevant Rank | Query |
| --- | ---: | ---: | ---: | --- |
| 10462 | 0.0000 | 0.0000 | missing | Is it okay to be married, 30 years old and have no retirement? |
| 10482 | 0.0000 | 0.0000 | missing | Rollover into bond fund to do dollar cost averaging [duplicate] |
| 1085 | 0.0000 | 0.0000 | missing | How do disputed debts work on credit reports? |
| 10975 | 0.0000 | 0.0000 | missing | How to contribute to Roth IRA when income is at the maximum limit & you have employer-sponsored 401k plans? |
| 1159 | 0.0000 | 0.0000 | missing | what is the best way to do a freelancing job over the summer for a student |
| 1783 | 0.0000 | 0.0000 | missing | Freelancing Tax implication |
| 1889 | 0.0000 | 0.0000 | missing | Reporting financial gains from my online store |
| 2885 | 0.0000 | 0.0000 | missing | Merits of buying apartment houses and renting them |
| 2895 | 0.0000 | 0.0000 | missing | Where should a young student put their money? |
| 34 | 0.0000 | 0.0000 | missing | 401k Transfer After Business Closure |
