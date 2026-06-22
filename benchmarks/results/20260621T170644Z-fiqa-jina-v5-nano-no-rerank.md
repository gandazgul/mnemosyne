# Mnemosyne BEIR fiqa Results

- Generated: 2026-06-21T17:06:44Z
- Dataset: `fiqa`
- Queries: 648
- Corpus documents: 57638
- Search limit: 100
- Source mode: `hybrid`
- Rerank candidates: 100
- Rerank enabled: False
- Run label: `jina-v5-nano`
- Mnemosyne config: `/Users/gandazgul/Documents/web/mnemosyne/cli/configs/jina-v5-text-nano-retrieval.yaml`

| Metric | Score |
| --- | ---: |
| `ndcg@10` | 0.3720 |
| `mrr@10` | 0.4483 |
| `recall@10` | 0.4485 |
| `recall@100` | 0.7574 |
| `map@100` | 0.3145 |

## Breakdown

| Detail | Count |
| --- | ---: |
| `queries` | 648 |
| `queries_with_rank_1_hit` | 230 |
| `queries_with_top_10_hit` | 424 |
| `queries_with_top_100_hit` | 580 |
| `queries_missing_at_100` | 68 |

### First Relevant Rank

| Bucket | Queries |
| --- | ---: |
| `rank_1` | 230 |
| `rank_2_3` | 103 |
| `rank_4_10` | 91 |
| `rank_11_100` | 156 |
| `missing@100` | 68 |

### Missing At 100

| Query ID | Missing Doc IDs | Query | Top 10 Doc IDs |
| --- | --- | --- | --- |
| 4641 | 319954, 397358 | Where should I park my rainy-day / emergency fund? | 497993, 580025, 527939, 285812, 108978, 242011, 538023, 583695, 179702, 282623 |
| 7017 | 28271 | Basic Algorithmic Trading Strategy | 514562, 446629, 36723, 177347, 350460, 179068, 301849, 212390, 216588, 352144 |
| 3451 | 192307, 588448 | Should you keep your stocks if you are too late to sell? | 251536, 41852, 27015, 310683, 303724, 528518, 417506, 165970, 250873, 495568 |
| 753 | 243503 | Taxes due for hobbyist Group Buy | 203791, 170632, 132780, 370857, 79511, 451020, 283505, 122185, 539244, 466718 |
| 4465 | 376575 | How to donate to charity that will make a difference? | 174033, 90591, 132881, 326167, 106786, 427322, 266342, 379353, 174543, 46381 |
| 5951 | 497260 | Why can't house prices be out of tune with salaries | 599860, 418034, 259777, 31663, 589470, 285525, 196405, 62702, 159204, 81592 |
| 10827 | 160786, 7748, 107554, 95282 | How much should I be contributing to my 401k given my employer's contribution? | 555377, 452592, 15841, 97805, 290105, 427997, 140330, 361510, 247473, 17166 |
| 9771 | 263955, 28740 | Is there any emprical research done on 'adding to a loser' | 137898, 356175, 38287, 80481, 220190, 589507, 131258, 147581, 370632, 93714 |
| 3724 | 508921, 434190, 279570, 497216, 552887, 341146, 199970 | Should you always max out contributions to your 401k? | 497561, 488673, 198825, 135790, 273497, 340624, 480036, 122910, 302512, 163834 |
| 5853 | 476663, 160105, 495699, 424598 | Paying Off Principal of Home vs. Investing In Mutual Fund | 284318, 476517, 301224, 336218, 182612, 387722, 383682, 186071, 173431, 439459 |
| 570 | 363591 | Employer options when setting up 401k for employees | 79375, 117845, 242529, 15841, 150883, 387876, 289064, 168890, 477175, 97805 |
| 594 | 534059 | Should a retail trader bother about reading SEC filings | 377322, 64881, 548596, 314898, 97837, 86281, 161411, 38863, 11148, 213214 |
| 10122 | 273718 | Why diversify stocks/investments? | 144261, 297100, 508540, 404949, 286227, 187124, 336722, 89084, 456761, 324914 |
| 4037 | 10275 | How separate individual expenses from family expenses in Gnucash? | 355415, 244555, 347137, 400291, 273071, 202224, 304971, 123384, 287293, 557186 |
| 1783 | 332314 | Freelancing Tax implication | 14609, 156063, 421924, 445298, 159709, 383172, 179359, 423625, 417769, 136315 |
| 6004 | 149555 | Put-Call parity - what is the difference between the two representations? | 13260, 345410, 374797, 122432, 21768, 118762, 37517, 271109, 387801, 232880 |
| 4837 | 531841, 20958 | When applying for a mortgage, can it also cover outstanding debts? | 97925, 32749, 516578, 204035, 245005, 445639, 565698, 294167, 257644, 249054 |
| 4415 | 67676 | How much is inflation? | 117578, 206580, 513249, 501743, 553634, 290585, 59225, 519596, 468089, 424220 |
| 8378 | 125298 | Should I wait a few days to sell ESPP Stock? | 495568, 294573, 133644, 361345, 127702, 178684, 567244, 546125, 387035, 511678 |
| 8102 | 552707, 378173, 90294 | When do I sell a stock that I hold as a long-term position? | 306460, 165970, 35752, 537212, 171819, 203638, 310683, 511678, 488207, 501214 |
| 7928 | 118633 | If I believe a stock is going to fall, what options do I have to invest on this? | 427808, 410404, 260384, 166597, 42438, 294688, 443926, 171784, 480967, 187214 |
| 4777 | 590710 | How to finance necessary repairs to our home in order to sell it? | 416382, 378384, 478413, 171631, 426678, 247083, 365342, 308802, 570318, 45773 |
| 620 | 331332 | Is it wise to have plenty of current accounts in different banks? | 543921, 38628, 404954, 189303, 535817, 278734, 456636, 164801, 104492, 48346 |
| 8271 | 415511 | Income in zero-interest environment | 376709, 137225, 249558, 380368, 531758, 136262, 65179, 16626, 533589, 283862 |
| 4865 | 100485 | Why are historical prices of stocks different on different websites? Which one should I believe? | 370569, 47738, 189341, 246733, 394151, 236321, 218842, 67365, 173986, 248564 |
| 864 | 211364, 152072 | Why use accounting software like Quickbooks instead of Excel spreadsheets? | 30142, 472924, 329774, 24890, 78117, 157751, 566337, 402174, 479625, 2436 |
| 5888 | 540806 | Interest charges on balance transfer when purchases are involved | 125497, 263647, 543776, 336792, 490529, 213242, 300990, 499098, 343961, 358445 |
| 2568 | 388798, 127353 | How to pay with cash when car shopping? | 346042, 166314, 258247, 15696, 453301, 9146, 27693, 584106, 351312, 514238 |
| 7124 | 74615 | How come we can find stocks with a Price-to-Book ratio less than 1? | 558617, 278582, 583708, 123263, 526110, 226070, 522813, 433905, 533818, 414940 |
| 2857 | 295864 | I have around 60K $. Thinking about investing in Oil, how to proceed? | 542051, 127566, 379140, 117451, 474575, 316444, 60175, 516818, 233732, 501384 |
| 8539 | 218728, 196304, 396038 | Can the risk of investing in an asset be different for different investors? | 283074, 391861, 502495, 483123, 499166, 385881, 500863, 497266, 378403, 242992 |
| 1085 | 467737, 393710 | How do disputed debts work on credit reports? | 372039, 450031, 574122, 422484, 242013, 78328, 259764, 161422, 391085, 339365 |
| 5422 | 151973 | What are some good books for learning stocks, bonds, derivatives e.t.c for beginner with a math background? | 273906, 193555, 412327, 552371, 241423, 276786, 221319, 378163, 463713, 246545 |
| 8247 | 42521, 321114, 465313 | Tax on Stocks or ETF's | 586010, 437907, 518735, 528880, 474745, 580802, 161019, 270992, 459953, 183910 |
| 10482 | 549072 | Rollover into bond fund to do dollar cost averaging [duplicate] | 330023, 447567, 224782, 525089, 265817, 134005, 237052, 473042, 544070, 439757 |
| 8789 | 70853 | What does “profits to the shareholders jumped to 15 cents a share” mean? | 87349, 41912, 14870, 20076, 339854, 113587, 317363, 217731, 444747, 411617 |
| 2423 | 538023 | At what age should I start or stop saving money? | 529444, 234846, 417787, 396127, 337561, 468010, 235855, 571044, 519174, 553288 |
| 8475 | 293320 | Why I cannot find a “Pure Cash” option in 401k investments? | 494655, 104793, 525426, 240562, 437427, 332152, 444044, 444107, 505617, 226547 |
| 4523 | 594257, 393009 | What should I do with my $25k to invest as a 20 years old? | 272070, 10476, 442776, 129255, 252942, 171712, 589103, 521847, 332203, 465819 |
| 8513 | 270573 | Buy on dip when earnings fail? | 175821, 572622, 203873, 391043, 73872, 347900, 351396, 104343, 273416, 261522 |
| 5054 | 28119 | How to stress test an investment plan? | 588481, 377186, 582899, 448745, 481859, 216915, 101124, 263390, 564007, 349475 |
| 1159 | 496064 | what is the best way to do a freelancing job over the summer for a student | 347181, 55064, 460648, 449155, 101329, 144948, 156063, 271812, 132287, 439899 |
| 68 | 19183 | Intentions of Deductible Amount for Small Business | 562777, 381151, 519473, 192516, 462831, 545296, 354716, 201954, 146657, 86134 |
| 7269 | 486696 | How do I track investment performance in Quicken across rollovers? | 152045, 563499, 473644, 357626, 343586, 277544, 151554, 294644, 272202, 308991 |
| 9701 | 387141, 357739 | How to bet against the London housing market? | 473883, 408865, 225682, 242321, 45583, 509464, 180453, 419058, 592612, 258048 |
| 6199 | 239214 | How can all these countries owe so much money?  Why & where did they borrow it from? | 36659, 124940, 169921, 47163, 380714, 490042, 367473, 392806, 414693, 558742 |
| 9126 | 514831 | Short an option - random assignment? | 334473, 477588, 228810, 166227, 307518, 102316, 238474, 292045, 82194, 349974 |
| 34 | 599545 | 401k Transfer After Business Closure | 122114, 483268, 64459, 398520, 168890, 217661, 120394, 262322, 338519, 226547 |
| 6395 | 166227 | Option settlement for calendar spreads | 22916, 467463, 584223, 276314, 401447, 516790, 182272, 8177, 154989, 543312 |
| 9115 | 207325 | Why does the calculation for percentage profit vary based on whether a position is short vs. long? | 419897, 422467, 154665, 232880, 314478, 158520, 233379, 331606, 428786, 226496 |
| 4767 | 280805, 568670, 224057, 22804 | New car: buy with cash or 0% financing | 420018, 408932, 584106, 306834, 9146, 256803, 475440, 59372, 166314, 56867 |
| 8351 | 472516 | What happens when a calendar spread is assigned in a non-margin account? | 141213, 102316, 45674, 23609, 527654, 272754, 273142, 273612, 547460, 511093 |
| 6131 | 381720, 170204, 416679, 2460 | Is it ever a good idea to close credit cards? | 326094, 334111, 339030, 368806, 258465, 391384, 99449, 508510, 38938, 143596 |
| 858 | 122485, 278450 | Is it bad practice to invest in stocks that fluctuate by single points throughout the day? | 42347, 208932, 139699, 293027, 508764, 433730, 567608, 127689, 519501, 146632 |
| 4019 | 6881, 125477 | How and Should I Invest (As a college 18 year old with minimal living expenses)? | 332938, 332749, 66864, 20304, 58065, 85977, 429432, 140135, 269671, 81016 |
| 6080 | 164513 | Is ScholarShare a legitimate entity for a 529 plan in California? | 22856, 201500, 277581, 236732, 83080, 2809, 380424, 445868, 233401, 380557 |
| 2580 | 503934 | Stock market vs. baseball card trading analogy | 344118, 11988, 408918, 371720, 163987, 206765, 363158, 257457, 516561, 577768 |
| 6959 | 205010 | What is the term for the quantity (high price minus low price) for a stock? | 229573, 577573, 468025, 304399, 69790, 477162, 517961, 116846, 33357, 301570 |
| 5402 | 491350 | Is it impossible to get a home loan with a poor credit history after a divorce? | 445163, 595029, 51728, 90579, 310790, 594595, 52250, 85697, 488609, 180214 |
| 701 | 389446 | What are the ins/outs of writing-off part of one's rent for working at home? | 436505, 349672, 231990, 344955, 124507, 376446, 210117, 378384, 145313, 339488 |
| 585 | 140226 | Following an investment guru a good idea? | 311117, 535737, 14748, 235772, 426550, 230733, 2235, 483185, 209843, 113221 |
| 7674 | 519390 | Choosing the limit when making a limit order? | 447886, 249279, 514841, 94653, 289045, 251596, 15917, 266785, 448713, 184756 |
| 5940 | 486243, 93936 | How does investment into a private company work? | 250354, 182226, 252853, 512609, 179664, 473154, 349684, 454465, 53993, 473798 |
| 6612 | 322900 | If I have a lot of debt and the housing market is rising, should I rent and slowly pay off my debt or buy and roll th... | 502594, 433171, 254454, 301192, 290900, 14083, 28230, 253880, 198442, 431481 |
| 4714 | 450819 | Personal finance app where I can mark transactions as “reviewed”? | 584450, 218793, 344473, 65957, 505057, 479390, 353915, 171409, 529790, 324833 |
| 8456 | 486333 | What typically happens to unvested stock during an acquisition? | 257853, 534755, 93215, 555276, 469036, 104188, 186869, 486696, 345388, 598051 |
| 10213 | 270221, 545712 | Looking for good investment vehicle for seasonal work and savings | 446186, 343693, 191658, 223551, 100517, 502281, 246991, 589394, 84267, 38269 |
| 5196 | 172128, 114829 | I might use a credit card convenience check. What should I consider? | 565745, 289483, 456098, 85252, 402543, 393866, 581976, 219033, 481052, 85517 |
| 3006 | 269851, 568473, 328300 | Strategies for putting away money for a child's future (college, etc.)? | 258704, 290441, 372900, 490382, 127838, 303432, 332749, 411686, 512096, 361821 |
| 3909 | 312248, 404356, 245616, 353028 | How to rescue my money from negative interest? | 472837, 328499, 83330, 362730, 42475, 438403, 61586, 404352, 514003, 46587 |
| 5464 | 86691 | Resources on Buying Rental Properties | 423438, 26339, 545341, 325722, 372274, 85229, 383921, 55634, 408029, 426705 |
| 2385 | 407654 | As director, can I invoice my self-owned company? | 373059, 210889, 247760, 496064, 217472, 224438, 55440, 82119, 496959, 120649 |
| 3177 | 268289 | Vanguard ETF vs mutual fund | 539263, 143238, 500486, 367960, 172703, 454224, 480315, 364735, 161019, 138383 |
| 10034 | 480749 | Tax implications of holding EWU (or other such UK ETFs) as a US citizen? | 528880, 181942, 565296, 197478, 447197, 495417, 180146, 51777, 141585, 381884 |
| 5090 | 436493 | Should I take a student loan to pursue my undergraduate studies in France? | 246286, 92430, 12988, 455666, 287507, 217831, 564206, 231283, 58005, 572272 |
| 2088 | 399875 | How would I go about selling the stock of a privately held company? | 53993, 455168, 62897, 575554, 413672, 345368, 260023, 140835, 377563, 72846 |
| 9391 | 503637 | Should I replace bonds in a passive investment strategy | 535518, 107424, 248158, 136515, 577832, 545760, 171669, 155242, 494653, 142631 |
| 3148 | 178127, 438000 | Can a car company refuse to give me a copy of my contract or balance details? | 172855, 584305, 65046, 430100, 164702, 327806, 395995, 103589, 76898, 29721 |
| 4678 | 305153 | Finance, Cash or Lease? | 185405, 376016, 504918, 311748, 215225, 427884, 487678, 488258, 522532, 311446 |
| 2398 | 224654, 590489 | Frustrated Landlord | 556453, 96538, 487094, 98372, 44058, 201705, 436875, 393883, 309231, 249195 |
| 5511 | 169893, 560325, 478426, 383193, 114303, 278699, 12746 | Pay off car loan entirely or leave $1 until the end of the loan period? | 352027, 329137, 543193, 376016, 179891, 38786, 206449, 479050, 43142, 502686 |
| 8834 | 12232 | Pros/Cons of Buying Discounted Company Stock | 203139, 599156, 528827, 163396, 553331, 268802, 569303, 471686, 569224, 417457 |
| 988 | 226053, 107688 | Where should I invest my savings? | 501384, 571218, 58065, 223872, 569528, 480827, 60093, 451782, 168402, 579901 |
| 3369 | 231012, 145716, 411910, 395840 | Why should one only contribute up to the employer's match in a 401(k)? | 296405, 341493, 240373, 175470, 463892, 519750, 301616, 565684, 38532, 75766 |
| 9296 | 435746 | Why would Two ETFs tracking Identical Indexes Produce different Returns? | 148721, 159471, 206744, 408524, 99568, 428187, 55751, 492212, 20504, 285135 |
| 9245 | 194561 | Stock Options for a company bought out in cash and stock | 207253, 259560, 117177, 186869, 220147, 340730, 510875, 287092, 344372, 555276 |
| 5763 | 462019 | What is the best way to get a “rough” home appraisal prior to starting the refinance process? | 67379, 331255, 38712, 215647, 440063, 218144, 326214, 60088, 497927, 251466 |
| 4962 | 599925 | Net Cash Flows from Selling the Bond and Investing | 416839, 308276, 431386, 408661, 152265, 187110, 537603, 34949, 158363, 535518 |
| 4846 | 151104 | Is there anything comparable to/resembling CNN's Fear and Greed Index? | 98096, 335892, 538974, 3533, 317666, 415161, 543239, 270305, 489352, 498754 |
| 9403 | 6666, 328086, 345199 | Abundance of Cash - What should I do? | 228403, 217865, 105089, 410450, 598526, 14349, 486692, 103447, 186332, 158614 |
| 5993 | 367375, 224918, 272866, 55084, 352638, 426120, 63501 | Why would anyone want to pay off their debts in a way other than “highest interest” first? | 94373, 416796, 494306, 128574, 160193, 353911, 31189, 445639, 886, 437538 |
| 9633 | 585447 | Video recommendation for stock market education | 534418, 162884, 259081, 353048, 235197, 157509, 524501, 481166, 501189, 171456 |
| 7529 | 66607 | Does the expense ratio of a fund-of-funds include the expense ratios of its holdings? | 89297, 464337, 514529, 293626, 59249, 102904, 65587, 135405, 361013, 464668 |
| 5021 | 589285 | Is there a more flexible stock chart service, e.g. permitting choice of colours when comparing multiple stocks? | 189341, 422453, 584801, 310218, 500527, 528576, 465971, 79357, 151678, 105717 |
| 3612 | 259625 | How can I buy and sell the same stock on the same day? | 392403, 522658, 142599, 212687, 568200, 584291, 331521, 352415, 216964, 591436 |
| 4409 | 499128, 100306, 147439 | My friend wants to put my name down for a house he's buying. What risks would I be taking? | 102326, 243732, 360682, 539432, 115066, 102088, 514790, 483480, 306926, 518242 |
| 5369 | 171339 | Paying for things on credit and immediately paying them off: any help for credit rating? | 540959, 190553, 574065, 250722, 577542, 375780, 272890, 393817, 403934, 188676 |
| 2070 | 363678 | Advantage of credit union or local community bank over larger nationwide banks such as BOA, Chase, etc.? | 550303, 587737, 469515, 597571, 590209, 30253, 578357, 484313, 18749, 408166 |
| 11039 | 53544, 249063 | Pay off credit card debt or earn employer 401(k) match? | 91183, 163287, 519750, 437706, 353625, 508534, 124042, 105557, 130691, 79363 |
| 5460 | 184337, 21174, 108514, 463885 | Paying off a loan with a loan to get a better interest rate | 470716, 31189, 490648, 202527, 93248, 98920, 503723, 33350, 243065, 589582 |
| 7925 | 318185, 251100, 402482 | Can I sell a stock immediately? | 591436, 310636, 44461, 81721, 227399, 339419, 221869, 315760, 488207, 377719 |
| 4286 | 566069 | Given advice “buy term insurance and invest the rest”, how should one “invest the rest”? | 70460, 229239, 151817, 564675, 155640, 10531, 206830, 71926, 377477, 56732 |
| 3789 | 571131 | How to work around the Owner Occupancy Affidavit to buy another home in less than a year? | 274573, 492856, 129862, 459724, 482963, 327428, 578906, 125613, 287458, 286632 |
| 5228 | 232451 | How does the bank/IRS know whether a bank transfer over $14k is a gift or loan repayment? | 117661, 214934, 232322, 32455, 344398, 553205, 553328, 28160, 131774, 93386 |
| 2685 | 154113, 370300, 37900, 303293, 468923 | What ways are there for us to earn a little extra side money? | 576047, 382005, 109880, 280099, 237950, 468086, 89624, 186889, 194540, 269380 |
| 1090 | 518896 | Need a formula to determine monthly payments received at time t if I'm reinvesting my returns | 446454, 393987, 179365, 16051, 19999, 584231, 209238, 340254, 132950, 166394 |
| 6122 | 169824, 44344 | Better to rent condo to daughter or put her on title? | 496166, 316794, 566184, 182039, 163289, 558251, 574432, 84732, 80269, 101816 |
| 4514 | 69485, 337764, 209804 | What intrinsic, non-monetary value does gold have as a commodity? | 471825, 156211, 408336, 240894, 426270, 99089, 146573, 80141, 317429, 479398 |
| 8507 | 509819, 370995 | When to sell a stock? | 99132, 273565, 217837, 251536, 303724, 545284, 53047, 554568, 102237, 66834 |
| 6221 | 257248, 519675, 76414, 169688, 455614, 115717 | To pay off a student loan, should I save up a lump sum payoff payment or pay extra each month? | 448791, 352363, 27625, 254245, 479050, 221364, 110081, 399863, 538014, 274108 |
| 3008 | 180192, 323406 | What are my chances at getting a mortgage with Terrible credit but High income | 44105, 47441, 310120, 285694, 407401, 231688, 257644, 312211, 2064, 102266 |
| 4007 | 521657 | What is a reasonable salary for the owner and sole member of a small S-Corp? | 556220, 260385, 205341, 370542, 170933, 388704, 458431, 508078, 521933, 543085 |
| 6644 | 175035 | How to know precisely when a SWIFT is issued by a bank? | 218761, 475527, 110198, 589616, 298587, 350396, 39783, 327623, 446807, 41383 |
| 7463 | 305287 | Pros/cons of borrowing money using a mortgage loan and investing it in a low-fee index fund? | 577951, 269770, 384747, 288504, 544236, 51640, 473427, 321637, 192910, 105634 |
| 10267 | 460398 | How should I prepare for the next financial crisis? | 178693, 143393, 96017, 305600, 549409, 182442, 326398, 36961, 125847, 588948 |
| 7622 | 253369, 378594 | Best way to pay off debt? | 457945, 353911, 480773, 157923, 345895, 373554, 115499, 249643, 416796, 529312 |
| 10979 | 148728 | Closing a futures position | 362762, 503505, 298833, 121158, 557582, 357324, 533408, 92695, 206683, 450067 |
| 3767 | 153922, 392060 | What should I be doing to protect myself from identity theft? | 90632, 423809, 260580, 97686, 581889, 158285, 587778, 171510, 538217, 551747 |
| 6410 | 471723 | Will an ETF immediately reflect a reconstitution of underlying index | 454610, 71230, 200360, 295993, 313897, 144033, 94076, 304023, 276593, 418150 |
| 5030 | 215540 | Why pay for end-of-day historical prices? | 532178, 560108, 227192, 113150, 370569, 471131, 394151, 378994, 330276, 587587 |
| 6252 | 394551, 160932, 293624, 233294, 243268, 379487, 62868 | Is this mortgage advice good, or is it hooey? | 213713, 139366, 47565, 120061, 473647, 205906, 78518, 27268, 71709, 466587 |
| 885 | 337165, 409184 | How long do credit cards keep working after you disappear? | 472336, 254968, 99449, 251643, 181757, 296165, 312618, 584237, 516678, 489501 |
| 4031 | 115741 | 28 years old and just inherited large amount of money and real estate - unsure what to do with it | 318864, 159652, 148335, 509565, 588316, 307426, 344186, 140002, 578597, 163197 |
| 766 | 550172 | Will the ex-homeowner still owe money after a foreclosure? | 2996, 163711, 299591, 104955, 578906, 350588, 333583, 274435, 212827, 62908 |
| 8202 | 513258, 93971 | What accounted for DXJR's huge drop in stock price? | 457689, 537862, 67237, 317363, 122542, 431814, 188839, 337001, 7593, 261522 |
| 7345 | 237645 | What do these numbers mean? (futures) | 527080, 9274, 354429, 273789, 261331, 108, 529996, 587111, 110966, 419864 |
| 776 | 583640, 127263, 220127, 597880 | Can saving/investing 15% of your income starting age 25, likely make you a millionaire? | 124027, 417787, 143591, 41960, 374266, 434972, 418281, 10440, 563284, 405985 |
| 89 | 248624 | How can I deposit a check made out to my business into my personal account? | 526817, 309023, 508754, 308938, 188893, 400230, 98636, 135196, 521540, 188167 |
| 6629 | 444405 | Tax treatment of a boxed trade? | 524649, 5762, 9353, 300486, 412404, 295153, 502874, 358090, 385310, 30403 |
| 1889 | 388713 | Reporting financial gains from my online store | 255281, 557885, 92593, 584074, 243503, 226997, 355959, 379661, 196520, 352136 |
| 8013 | 496159, 224231 | Frequency of investments to maximise returns (and minimise fees) | 81652, 537626, 384983, 388389, 224816, 28291, 270256, 57033, 450662, 8759 |
| 3759 | 527966, 67167, 522358 | Simplifying money management | 455457, 373772, 214248, 122378, 145812, 384631, 490065, 551986, 478914, 591704 |
| 7295 | 244749 | Selling non-dividend for dividend stocks | 352484, 472470, 501931, 542667, 71511, 509879, 115553, 222921, 352415, 170318 |
| 10639 | 431799, 495774, 278453, 187039 | Short term parking of a large inheritance? | 171196, 163353, 163197, 235628, 222030, 438874, 590276, 111048, 318864, 289326 |
| 6635 | 156358 | Why don't share prices of a company rise every other Friday when the company buys shares for its own employees? | 587137, 533712, 3656, 401818, 95806, 235531, 245654, 117082, 545036, 177093 |
| 4312 | 399149 | Is it true that 90% of investors lose their money? | 282435, 222639, 167950, 497786, 285945, 170628, 300770, 116647, 507284, 532485 |
| 859 | 18749 | Any reason to keep around my account with my old, 'big' bank? | 546779, 584419, 412036, 142809, 469731, 537593, 489959, 152151, 329819, 597571 |
| 6525 | 181985 | Does it make sense to trade my GOOGL shares for GOOG and pocket the difference? | 550661, 106541, 98150, 362473, 498014, 156467, 88652, 147002, 53263, 492503 |
| 2590 | 589625 | Are non-residents or foreigners permitted to buy or own shares of UK companies? | 209493, 188776, 296528, 307776, 528516, 310103, 262485, 536483, 263312, 415954 |
| 5374 | 152688 | What were the main causes of the spike and drop of DRYS's stock price? | 283106, 133204, 457689, 261522, 537862, 317363, 67237, 362462, 73857, 122542 |
| 8005 | 48800 | Difference between Vanguard sp500 UCITS and Vanguard sp500 | 27930, 172703, 56894, 371251, 241101, 271825, 220486, 408103, 237305, 454224 |
| 2994 | 419319 | Work on the side for my wife's company | 569145, 506991, 200211, 5840, 382005, 269380, 341909, 399882, 423625, 504399 |
| 3683 | 185909, 454501 | Can I trust the Motley Fool? | 276975, 408995, 428848, 500338, 105973, 192912, 6607, 522713, 538086, 218285 |
| 7206 | 441155, 532211, 553066 | Who Bought A Large Number Of Shares? | 351570, 34882, 369998, 358164, 327525, 482739, 152097, 436530, 65667, 529001 |
| 10246 | 77573 | Understanding the T + 3 settlement days rule | 370635, 28314, 11927, 89506, 176717, 332243, 226984, 327080, 340263, 36193 |
| 5241 | 376123, 322157, 27489 | Mortgage vs. Cash for U.S. home buy now | 344740, 390976, 111184, 281675, 78176, 213713, 438073, 133735, 444148, 15487 |
| 98 | 575929 | How can I make $250,000.00 from trading/investing/business within 5 years? | 336661, 209067, 102113, 527522, 352363, 221795, 555630, 438279, 385484, 519619 |
| 4615 | 262934 | Are solar cell panels and wind mills worth the money? | 69523, 425595, 455798, 158216, 271015, 385028, 496427, 249191, 596196, 120384 |
| 3264 | 134764 | Pros and Cons of Interest Only Loans | 482507, 486525, 268802, 471686, 260383, 315847, 453847, 316230, 301609, 10873 |
| 6467 | 453256, 23217, 346641, 367313 | Advice on strategy for when to sell | 217837, 240089, 88813, 109455, 203873, 83807, 130941, 99857, 596664, 88540 |
| 4289 | 24881 | Does the currency exchange rate contain any additional information at all? | 17469, 416975, 44843, 135220, 343489, 226102, 493576, 439779, 324546, 517345 |
| 4394 | 336045, 441582 | Transfer $50k to another person's account (in California, USA) | 322838, 412258, 305907, 293653, 431462, 521753, 462585, 415655, 307404, 93386 |
| 7344 | 108403 | How is the Dow divisor calculated? | 14368, 159166, 150430, 378974, 253926, 313421, 501032, 591089, 69655, 65618 |
| 10447 | 152096, 300721 | Is there an advantage to a traditional but non-deductable IRA over a taxable account? [duplicate] | 144751, 500175, 382236, 532657, 447482, 259150, 406239, 540389, 109305, 94496 |
| 5782 | 379891, 319773, 595455 | Pay off credit cards in one lump sum, or spread over a few months? | 487621, 172084, 114592, 117007, 261697, 262026, 511240, 440620, 254245, 559523 |
| 9871 | 448890, 40051, 170594 | What should I do with the 50k I have sitting in a European bank? | 73741, 367207, 76562, 175139, 212464, 293179, 103795, 292714, 354553, 433003 |
| 3404 | 498834, 277583 | In US, is it a good idea to hire a tax consultant for doing taxes? | 37725, 197870, 488574, 218986, 450933, 525360, 34338, 442110, 420295, 44152 |
| 3625 | 384469, 414295 | What should I do with my paper financial documents? | 509617, 500751, 163168, 123366, 489827, 37582, 44204, 245967, 278656, 569812 |
| 6005 | 135415, 478457, 345895, 73310, 384626, 390689 | Why might it be advisable to keep student debt vs. paying it off quickly? | 149500, 507544, 96268, 431884, 25190, 572272, 571198, 196237, 564206, 52136 |
| 3822 | 385090, 418900, 308837 | How to change a large quantity of U.S. dollars into Euros? | 292714, 194730, 417917, 549787, 19618, 340777, 531953, 79777, 245727, 390524 |
| 7879 | 372551, 421285 | Any Tips on How to Get the Highest Returns Within 4 Months by Investing in Stocks? | 102029, 58186, 540919, 43088, 174313, 120297, 201391, 367348, 404732, 593879 |
| 3115 | 234950, 389028, 316794 | How can I live outside of the rat race of American life with 300k? | 233562, 183869, 174272, 136035, 267892, 475736, 129364, 252852, 150066, 155374 |
| 3995 | 278734, 230208 | I have more than $250,000 in a US Bank account… mistake? | 404954, 146557, 352883, 583803, 14349, 303367, 506909, 190539, 307404, 89662 |
| 10136 | 526115 | How to minimise the risk of a reduction in purchase power in case of Brexit for money held in a bank account? | 466950, 417740, 304500, 152316, 284305, 290930, 51337, 168137, 448769, 314850 |
| 8635 | 67107, 240215 | Is there any flaw in this investment scheme? | 46818, 365816, 303619, 100694, 202638, 493841, 105973, 1907, 510565, 51615 |
| 5206 | 563030, 28230, 201982, 117276, 300660 | Is it a good idea to get an unsecured loan to pay off a credit card that won't lower a high rate? | 298908, 340520, 595455, 153088, 225522, 481822, 559523, 508510, 69938, 504293 |
| 2713 | 388147 | Physical Checks - Mailing | 29372, 41944, 284528, 20791, 199069, 48866, 190606, 216200, 513980, 50737 |
| 9060 | 40447 | Buying puts without owning underlying | 511093, 316037, 181924, 228217, 338782, 7743, 528052, 3062, 521644, 123320 |
| 4105 | 25096 | As an investor what are side effects of Quantitative Easing in US and in EU? | 345910, 416483, 305029, 176262, 108519, 369038, 393791, 30946, 339640, 37954 |
| 2465 | 570680, 81046 | Can capital expenses for volunteer purposes be deducted from income? | 202645, 146657, 541809, 432545, 82199, 18850, 421924, 216783, 81599, 275543 |
| 4640 | 101369, 540539 | What can my relatives do to minimize their out of pocket expenses on their fathers estate | 331534, 367404, 372808, 438874, 521803, 478966, 59749, 464269, 360816, 418328 |
| 9275 | 338754 | Do I have to pay a capital gains tax if I rebuy the same stock within 30 days? | 343219, 23217, 400730, 537916, 526661, 263751, 376800, 407602, 561999, 318321 |
| 3500 | 174019 | Why invest in becoming a landlord? | 273187, 71424, 11601, 528206, 557478, 487094, 112535, 249195, 41356, 560195 |
| 1306 | 484437, 204075 | I made an investment with a company that contacted me, was it safe? | 594206, 160611, 316321, 355792, 538086, 87632, 70389, 205665, 309851, 471138 |
| 6262 | 26799 | Help required on estimating SSA benefit amounts | 34538, 118707, 430407, 320362, 450765, 529927, 2648, 599485, 489376, 183349 |
| 8632 | 213976 | Is it best to exercise options shares when they vest, or wait | 43497, 237718, 163396, 237783, 382381, 104188, 340264, 61919, 220147, 313372 |
| 9979 | 337243 | What is the best way to invest in gold as a hedge against inflation without having to hold physical gold? | 35369, 483734, 96351, 30584, 556936, 327271, 326858, 13885, 257881, 399751 |
| 6133 | 415705 | What happens to all of the options when they expire? | 7733, 581672, 575408, 358492, 116436, 11456, 177559, 480879, 293605, 176786 |
| 3771 | 488948, 198349, 217683, 49601 | Best way to buy Japanese yen for travel? | 490384, 96211, 217715, 306130, 575495, 350245, 120604, 495826, 274683, 59994 |
| 1736 | 25543, 443419 | How can people have such high credit card debts? | 399406, 372993, 562896, 569056, 421379, 76248, 315275, 2018, 89622, 298908 |
| 6814 | 340214, 223206 | Selling Stock - All or Nothing? | 154976, 345368, 394454, 118633, 66834, 449280, 220608, 279782, 489254, 292338 |
| 1322 | 64138 | Is this follow-up after a car crash a potential scam? | 283917, 98356, 332916, 50368, 219119, 204735, 153640, 226090, 491537, 44635 |
| 5185 | 210236, 317354 | Invest in low cost small cap index funds when saving towards retirement? | 196992, 376485, 262180, 241202, 524525, 106620, 523331, 268731, 7208, 311527 |
| 2348 | 211867, 566573, 410166, 211622, 474234, 352271, 543714, 265874, 134864 | Why can't you just have someone invest for you and split the profits (and losses) with him? | 447619, 306430, 151412, 381757, 14870, 268261, 247486, 128435, 177194, 153491 |
| 687 | 146021 | Online tool to connect to my bank account and tell me what I spend in different categories? | 447597, 478807, 396222, 584175, 258423, 291278, 357113, 140657, 273308, 196299 |
| 4499 | 76996 | Is investing exlusively in a small-cap index fund a wise investment? | 517391, 501153, 196992, 52274, 272790, 513818, 241202, 335136, 98978, 14748 |
| 42 | 272709 | What are the ins/outs of writing equipment purchases off as business expenses in a home based business? | 28764, 305222, 88967, 510863, 47260, 35379, 494000, 220063, 581265, 363495 |
| 3530 | 239998 | How to exclude stock from mutual fund | 184299, 479420, 110343, 226967, 346474, 467575, 153660, 407185, 580802, 180196 |
| 9108 | 272021, 472585 | Starting an investment portfolio with Rs 5,000/- | 290757, 46967, 7748, 414116, 183074, 563986, 312821, 240351, 356552, 69012 |
| 6835 | 102243 | Are bond ETF capital gains taxed similar to stock or stock funds if held for more than 1 year? | 149305, 153112, 110343, 5710, 570546, 23217, 195767, 423929, 586010, 169240 |
| 3067 | 406156 | Should I make extra payments to my under water mortgage or increase my savings? | 131365, 476068, 423403, 477907, 560915, 90009, 202987, 372921, 336276, 11791 |
| 4125 | 344648, 72046 | Alternative means of salary for my employees | 365558, 174787, 245451, 479718, 361954, 36608, 75642, 70357, 355244, 50944 |
| 1150 | 531698, 43603, 19936 | How are the best way to make and save money at 22 years old | 10476, 38269, 433986, 494815, 519346, 529444, 319760, 353369, 305946, 328157 |
| 7705 | 195191 | Why would I pick a specific ETF over an equivalent Mutual Fund? | 500486, 539263, 161019, 153112, 106863, 270992, 370244, 253971, 377429, 367960 |
| 9808 | 40702, 431946 | Selling To Close | 416307, 557582, 151587, 345368, 362473, 374204, 107045, 149420, 125079, 229573 |
| 3888 | 319213, 239632 | Why I can't view my debit card pre-authorized amounts? | 208169, 185434, 432077, 440527, 316652, 418580, 181757, 50347, 276733, 21223 |
| 10109 | 506374, 156029, 499849 | Why does Charles Schwab have a Mandatory Settlement Period after selling stocks? | 28314, 93231, 266725, 332243, 563826, 370635, 121465, 537212, 296475, 226984 |
| 2790 | 279329, 469125 | Should I pay more than 20% down on a home? | 472484, 357200, 64400, 475632, 215103, 234890, 23533, 420707, 352339, 296906 |
| 9882 | 65702 | Money-market or cash-type ETFs for foreigners with U.S brokerage account | 389581, 391876, 94477, 131059, 97836, 188524, 241661, 386173, 216441, 98461 |
| 4999 | 314898 | Looking for a good source for Financial Statements | 9938, 171964, 431459, 597241, 295738, 465971, 11263, 520165, 146076, 46211 |
| 3189 | 225395 | Diversify my retirement investments with a Roth IRA | 287225, 404800, 561636, 122222, 272840, 271949, 568322, 262322, 455398, 1219 |
| 5134 | 158523 | Why does Yahoo Finance's data for a Vanguard fund's dividend per share not match the info from Vanguard? | 532616, 206727, 263088, 584128, 491472, 437465, 559884, 54225, 221477, 368848 |
| 1321 | 216456, 292065 | Are social media accounts (e.g. YouTube, Twitter, Instagram, etc.) considered assets? | 381322, 203715, 392379, 235082, 426184, 559738, 210019, 59972, 481036, 106614 |
| 3049 | 127974 | How to calculate my estimated taxes. 1099 MISC + Self Employment | 174025, 477476, 434351, 582864, 446117, 279538, 406789, 569645, 406042, 254151 |
| 715 | 579763, 546538, 187404 | what would you do with $100K saving? | 337561, 489179, 273925, 200273, 93936, 133120, 54574, 283635, 174336, 548758 |
| 5534 | 423272 | How does “taking over payments” work? | 297535, 421136, 397917, 67716, 70209, 293501, 316852, 490991, 416192, 202768 |
| 504 | 500755, 344203, 498751 | Have plenty of cash flow but bad credit | 22807, 368247, 93573, 41875, 68431, 569240, 23949, 77248, 546097, 78754 |
| 2296 | 83330, 366594, 253563 | How does a bank make money on an interest free secured loan? | 396853, 400009, 259919, 175824, 395769, 272150, 513079, 119298, 249831, 568047 |
| 10975 | 61022 | How to contribute to Roth IRA when income is at the maximum limit & you have employer-sponsored 401k plans? | 441632, 110114, 81148, 140330, 446226, 360533, 163865, 53028, 270818, 189989 |
| 1994 | 156640 | Does the IRS reprieve those who have to commute for work? | 231990, 434846, 380635, 243356, 63919, 319965, 544381, 192843, 22819, 453257 |
| 9164 | 365298, 263390 | Bonds vs equities: crash theory | 115648, 309326, 149900, 321941, 296516, 418528, 599420, 506743, 305274, 287656 |
| 1812 | 530570 | splitting a joint mortgage - one owner in home | 159590, 319159, 75235, 219274, 60163, 287458, 280177, 595029, 570964, 313896 |
| 10462 | 8266, 11378, 35680, 437879, 204035, 581204 | Is it okay to be married, 30 years old and have no retirement? | 268023, 66376, 151774, 519856, 259227, 280967, 591705, 198394, 217365, 309840 |
| 8855 | 208165, 312821 | How do i get into investing stocks [duplicate] | 210241, 155677, 367415, 67327, 235772, 555521, 142320, 386803, 426021, 341148 |
| 7071 | 124230 | ESPP strategy - Sell right away or hold? | 511678, 294573, 133644, 361345, 127702, 71713, 575213, 387035, 495568, 35575 |
| 3534 | 340329 | Why do dishonour fees exist? | 386745, 234260, 470032, 28720, 149845, 3279, 517497, 312349, 314499, 34108 |
| 8974 | 523331, 356595, 170625 | As a 22-year-old, how risky should I be with my 401(k) investments? | 102501, 216365, 452126, 140738, 165981, 357555, 134931, 279570, 437706, 92941 |
| 5178 | 240261 | Formula that predicts whether one is better off investing or paying down debt | 557506, 257016, 262772, 396889, 106215, 473865, 77153, 154449, 14028, 290434 |
| 5061 | 23747 | What fiscal scrutiny can be expected from IRS in early retirement? | 513392, 25481, 502150, 149742, 537049, 79610, 545172, 523831, 513981, 192738 |
| 2075 | 170042, 359580, 14967 | Are stories of turning a few thousands into millions by trading stocks real? | 519619, 188129, 285147, 39927, 65667, 519501, 562784, 523393, 464810, 78224 |
| 4335 | 357013 | What is the US Fair Tax? | 318583, 322246, 363178, 585212, 419466, 589161, 3181, 33764, 412877, 294157 |
| 7533 | 93853 | Investing tax (savings) | 142658, 8012, 250603, 284411, 526664, 399203, 315105, 316230, 516756, 278638 |
| 1393 | 352838, 539133 | Which is better when working as a contractor, 1099 or incorporating? | 234436, 220022, 277812, 586026, 32072, 578196, 352640, 68524, 232544, 77245 |
| 5264 | 576564 | Does a company's stock price give any indication to or affect their revenue? | 52579, 431814, 233988, 8643, 580920, 505694, 218326, 111076, 115134, 28662 |
| 9733 | 110163, 38655 | Due Diligence - Dilution? | 316321, 301880, 450132, 121262, 23414, 135798, 108965, 526073, 267266, 78236 |
| 7311 | 323768 | Finance, Social Capital IPOA.U | 507841, 583646, 463599, 479752, 261231, 377247, 116958, 207928, 538727, 571001 |
| 744 | 566480 | What options are available for a home loan with poor credit but a good rental history? | 310790, 573276, 495431, 82472, 517633, 409573, 305049, 360872, 596272, 526477 |
| 7141 | 132288 | Do investors go long option contracts when they cannot cover the exercise of the options? | 538054, 41967, 288289, 44530, 243714, 103147, 507828, 477588, 388571, 255927 |
| 4071 | 129875 | If our economy crashes, and cash is worthless, should i buy gold or silver | 291862, 487817, 505136, 502634, 524142, 53538, 470357, 356259, 474187, 97964 |
| 4103 | 440270 | What causes US Treasury I bond fixed interest to increase? | 180958, 29073, 182249, 372657, 104160, 567749, 256663, 388391, 379492, 74287 |
| 3254 | 484891 | Why do people buy US dollars on the black market? | 443511, 313706, 443380, 117010, 148454, 528234, 79777, 475756, 21896, 402006 |
| 7512 | 191060 | understanding the process/payment of short sale dividends | 480949, 202985, 115553, 107045, 527636, 84761, 228810, 241425, 188531, 549528 |
| 1391 | 562176 | How is taxation for youtube/twitch etc monetization handled in the UK? | 510599, 198692, 243939, 382623, 440745, 267067, 77171, 367015, 454208, 560548 |
| 7534 | 358125 | Can you explain why it's better to invest now rather than waiting for the market to dip? | 175821, 89714, 393009, 426157, 419747, 133380, 94302, 458843, 71219, 145539 |
| 5356 | 312405 | Historical stock prices: Where to find free / low cost data for offline analysis? | 279785, 560108, 535343, 529877, 226749, 391171, 240086, 546379, 537111, 47798 |
| 2579 | 432020 | What to do when a job offer is made but with a salary less than what was asked for? | 524471, 157919, 366282, 181213, 200946, 552290, 364159, 489554, 280838, 331029 |
| 7823 | 583549 | Retirement Funds: Betterment vs Vanguard Life strategy vs Target Retirement | 451196, 175927, 268731, 105666, 331492, 347825, 11094, 571217, 293679, 436120 |
| 689 | 411044 | Receive credit card payment sending my customer details to a credit card processing company? | 446932, 438032, 131177, 195852, 375652, 235808, 171761, 567201, 9814, 521688 |
| 9174 | 535317, 160218 | Which U.S. online discount broker is the best value for money? | 192910, 200052, 236931, 413856, 451729, 515144, 477683, 47579, 129466, 522798 |
| 6867 | 443804, 540799, 445258, 538750 | Will there always be somebody selling/buying in every stock? | 230343, 229573, 61006, 369166, 208070, 466143, 301985, 573077, 321639, 400614 |
| 2383 | 232199 | Should I Purchase Health Insurance Through My S-Corp | 17215, 224406, 476085, 527620, 154931, 546634, 490223, 507408, 588509, 41793 |
| 5083 | 138845 | Co-signer deceased | 369075, 18257, 270952, 273759, 447983, 453263, 495482, 334606, 142876, 138419 |
| 10526 | 39185 | What extra information might be obtained from the next highest bids in an order book? | 546493, 427747, 283008, 251100, 298551, 485973, 138830, 322798, 136822, 260153 |
| 2181 | 397329 | What are the risks & rewards of being a self-employed independent contractor / consultant vs. being a permanent emplo... | 37725, 139501, 524788, 488755, 197870, 260603, 383088, 406656, 234436, 77245 |
| 5903 | 231863 | Fees aside, what factors could account for performance differences between U.S. large-cap index ETFs? | 408524, 395842, 372233, 246996, 20504, 159471, 230997, 112187, 402091, 501153 |
| 5620 | 448784, 329552, 548740 | What's the fuss about identity theft? | 90632, 260580, 158285, 598801, 91986, 551747, 423809, 5860, 97686, 98993 |
| 2472 | 370334 | How do I deal with a mistaken attempt to collect a debt from me that is owed by someone else? | 62109, 435006, 543607, 584582, 200263, 101149, 525967, 205865, 304179, 475396 |
| 2306 | 315875 | To whom should I report fraud on both of my credit cards? | 596284, 581889, 270449, 90632, 423809, 226590, 531137, 298729, 249960, 556219 |
| 7633 | 197839 | Can a trade happen “in between” the bid and ask price? | 494727, 137175, 402482, 353396, 284235, 175831, 164008, 179258, 394244, 5018 |
| 2400 | 564271 | Will I be paid dividends if I own shares? | 311214, 97942, 1198, 400497, 501931, 152014, 1034, 29306, 456470, 126479 |
| 5549 | 286227, 309361 | Pros / cons of being more involved with IRA investments [duplicate] | 429106, 382894, 199544, 219208, 105468, 202140, 181624, 110465, 471686, 561636 |
| 3801 | 307776 | Can a bunch of wealthy people force Facebook to go public? | 390529, 69017, 92014, 209242, 168565, 394734, 371293, 362905, 556264, 186354 |
| 9088 | 569461, 561377 | Brokerage account for charity | 329849, 236186, 174543, 390089, 326167, 266342, 242027, 31703, 434224, 340947 |
| 4605 | 453941 | If the U.S. defaults on its debt, what will happen to my bank money? | 169691, 526384, 373717, 354896, 400826, 41312, 538582, 462668, 127268, 3040 |
| 2885 | 367360, 414692, 359579, 85229, 454810 | Merits of buying apartment houses and renting them | 159403, 451849, 430672, 502291, 164059, 80838, 387715, 343917, 4739, 358687 |
| 6110 | 331850, 94117, 259706 | Why does short selling require borrowing? | 188531, 320450, 67107, 226496, 79764, 329662, 49794, 107045, 314478, 84761 |
| 8 | 566392 | How to deposit a cheque issued to an associate in my business into my business account? | 65404, 261856, 508754, 590102, 564553, 301833, 25397, 489199, 309023, 318108 |
| 1309 | 156162, 489401 | Why does FlagStar Bank harass you about payments within grace period? | 489368, 471630, 271040, 336792, 75108, 526989, 173919, 438869, 151506, 15824 |
| 7109 | 447781 | How do I analyse moving averages? | 489933, 42620, 140804, 565501, 227669, 257185, 193012, 518932, 35006, 488285 |
| 5080 | 256055 | Is there a standard or best practice way to handle money from an expiring UTMA account? | 445521, 279291, 69841, 414429, 451189, 267206, 324564, 236186, 533122, 62079 |
| 4981 | 247894 | Where can I find open source portfolio management software? | 102684, 45218, 259463, 419171, 529790, 587792, 557861, 78436, 160700, 196432 |
| 7445 | 153178, 104343, 296231 | IS it the wrong time to get into the equity market immediately after large gains? | 483025, 350068, 590902, 221869, 598478, 434252, 89714, 33155, 79111, 400782 |
| 2895 | 521996, 328691 | Where should a young student put their money? | 426461, 517313, 55841, 148453, 496170, 502170, 241952, 561832, 204479, 307162 |
| 6787 | 587120 | Would it make sense to sell a stock, then repurchase it for tax purposes? | 23217, 219762, 263751, 400730, 390864, 474981, 311782, 580534, 328073, 343219 |
| 1748 | 576295 | How high should I set my KickStarter funding goal in order to have $35,000 left over? | 18001, 451492, 528564, 114443, 349926, 55714, 401505, 124705, 47949, 240351 |
| 5862 | 130209 | Can I get a discount on merchandise by paying with cash instead of credit? | 495751, 21194, 562511, 122908, 503171, 301643, 564180, 418801, 535015, 299840 |
| 6041 | 241308 | Most effective Fundamental Analysis indicators for market entry | 425020, 81655, 108579, 96910, 194240, 263464, 224695, 528034, 5054, 111091 |
| 7700 | 273761, 2653, 179328 | Should I re-allocate my portfolio now or let it balance out over time? | 224392, 131127, 422051, 22221, 355738, 395208, 441176, 175226, 569849, 551719 |
| 547 | 6349 | What percentage of my company should I have if I only put money? | 523158, 213399, 80913, 445353, 95243, 417838, 285041, 418003, 135411, 346537 |
| 3394 | 342258, 129319, 570664 | What is the easiest way to back-test index funds and ETFs? | 391215, 408524, 159471, 99568, 172374, 364735, 437875, 41176, 571913, 528034 |
| 5410 | 507813, 368802 | Dealership made me the secondary owner to my own car | 247371, 197138, 13975, 295355, 147530, 70456, 494323, 357280, 65046, 120080 |
| 4102 | 448699 | How can I determine if my rate of return is “good” for the market I am in? | 597437, 162488, 369439, 88801, 135176, 554734, 154450, 461082, 249360, 556235 |

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
| 2885 | 0.0000 | 0.0000 | missing | Merits of buying apartment houses and renting them |
| 2895 | 0.0000 | 0.0000 | missing | Where should a young student put their money? |
