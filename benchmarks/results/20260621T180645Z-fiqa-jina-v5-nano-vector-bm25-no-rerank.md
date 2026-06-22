# Mnemosyne BEIR fiqa Results

- Generated: 2026-06-21T18:06:45Z
- Dataset: `fiqa`
- Queries: 648
- Corpus documents: 57638
- Search limit: 100
- Source mode: `hybrid`
- Fusion: `vector-bm25`
- BM25 weight: `0.3`
- Rerank candidates: 300
- Rerank enabled: False
- Run label: `jina-v5-nano`
- Mnemosyne config: `/Users/gandazgul/Documents/web/mnemosyne/cli/configs/jina-v5-text-nano-retrieval.yaml`

| Metric | Score |
| --- | ---: |
| `ndcg@10` | 0.4075 |
| `mrr@10` | 0.4833 |
| `recall@10` | 0.4800 |
| `recall@100` | 0.7704 |
| `map@100` | 0.3496 |

## Breakdown

| Detail | Count |
| --- | ---: |
| `queries` | 648 |
| `queries_with_rank_1_hit` | 252 |
| `queries_with_top_10_hit` | 447 |
| `queries_with_top_100_hit` | 583 |
| `queries_missing_at_100` | 65 |

### First Relevant Rank

| Bucket | Queries |
| --- | ---: |
| `rank_1` | 252 |
| `rank_2_3` | 100 |
| `rank_4_10` | 95 |
| `rank_11_100` | 136 |
| `missing@100` | 65 |

### Missing At 100

| Query ID | Missing Doc IDs | Query | Top 10 Doc IDs |
| --- | --- | --- | --- |
| 4641 | 319954, 397358 | Where should I park my rainy-day / emergency fund? | 580025, 497993, 376148, 32833, 583695, 538023, 285812, 527939, 282623, 108978 |
| 3451 | 26292, 192307, 588448, 490170 | Should you keep your stocks if you are too late to sell? | 251536, 528518, 301757, 238215, 420974, 471911, 545284, 302808, 509819, 310683 |
| 753 | 243503 | Taxes due for hobbyist Group Buy | 466718, 122185, 203791, 114418, 198090, 158738, 170632, 94088, 444486, 127004 |
| 7326 | 584295 | Do brokers execute every trade on the exchange? | 404339, 35340, 163333, 37040, 257656, 384850, 272008, 593445, 227542, 395357 |
| 5951 | 497260 | Why can't house prices be out of tune with salaries | 599860, 418034, 31663, 259777, 196405, 285525, 589470, 66017, 374480, 596834 |
| 10827 | 160786, 7748, 107554, 95282 | How much should I be contributing to my 401k given my employer's contribution? | 290105, 436930, 555377, 140330, 497561, 15841, 242556, 481793, 296405, 436071 |
| 9771 | 263955, 28740 | Is there any emprical research done on 'adding to a loser' | 137534, 130349, 18671, 38704, 541485, 490479, 356175, 428133, 137898, 375290 |
| 3724 | 508921, 279570, 497216, 552887, 341146, 199970 | Should you always max out contributions to your 401k? | 302512, 273497, 43573, 3104, 576391, 122910, 497561, 488673, 480036, 135790 |
| 5853 | 476663, 160105, 495699, 424598 | Paying Off Principal of Home vs. Investing In Mutual Fund | 284318, 64456, 473865, 186071, 336218, 279229, 472159, 301224, 473647, 387722 |
| 570 | 363591 | Employer options when setting up 401k for employees | 532839, 289064, 79375, 301616, 117845, 150883, 396968, 387876, 15841, 101490 |
| 594 | 534059 | Should a retail trader bother about reading SEC filings | 377322, 64881, 97837, 11148, 314898, 548596, 213214, 161411, 579110, 86281 |
| 10122 | 273718 | Why diversify stocks/investments? | 144261, 297100, 89084, 286227, 259084, 180855, 187124, 456761, 459970, 331008 |
| 1783 | 332314 | Freelancing Tax implication | 156063, 14609, 421924, 159709, 445298, 383172, 270426, 179359, 423625, 72519 |
| 6004 | 149555 | Put-Call parity - what is the difference between the two representations? | 374797, 216065, 122432, 247738, 13260, 345410, 232261, 278373, 100021, 37517 |
| 4837 | 531841, 20958 | When applying for a mortgage, can it also cover outstanding debts? | 97925, 204035, 245005, 32749, 58353, 225874, 372993, 565972, 257644, 249054 |
| 4415 | 67676 | How much is inflation? | 206580, 117578, 290585, 519596, 381675, 501461, 352013, 501743, 513249, 238234 |
| 8378 | 125298 | Should I wait a few days to sell ESPP Stock? | 584291, 546125, 511678, 133644, 495568, 294573, 178684, 332243, 434812, 361345 |
| 8102 | 552707, 378173, 90294 | When do I sell a stock that I hold as a long-term position? | 306460, 165970, 171819, 35752, 203638, 537212, 557877, 310683, 292045, 496209 |
| 7928 | 118633 | If I believe a stock is going to fall, what options do I have to invest on this? | 410404, 427808, 67415, 260384, 501504, 294688, 171784, 480967, 222444, 47053 |
| 4777 | 590710 | How to finance necessary repairs to our home in order to sell it? | 52351, 365342, 171631, 416382, 323449, 478413, 570318, 378384, 247083, 426678 |
| 5741 | 25943 | Learning investing and the stock market | 64168, 379546, 47973, 85558, 367845, 241423, 167088, 124219, 34752, 341148 |
| 620 | 331332 | Is it wise to have plenty of current accounts in different banks? | 535817, 164801, 543921, 104492, 48346, 467059, 456636, 380786, 61734, 214248 |
| 8271 | 415511 | Income in zero-interest environment | 376709, 249558, 380368, 137225, 127074, 52047, 65179, 136262, 119298, 283862 |
| 864 | 211364, 152072 | Why use accounting software like Quickbooks instead of Excel spreadsheets? | 472924, 30142, 24890, 329774, 566337, 157751, 203446, 78117, 2436, 222380 |
| 5888 | 540806 | Interest charges on balance transfer when purchases are involved | 125497, 336792, 263647, 213242, 543776, 373271, 490529, 579601, 358445, 545327 |
| 2568 | 388798, 127353 | How to pay with cash when car shopping? | 346042, 258247, 166314, 60261, 196870, 15696, 260159, 9146, 289483, 38893 |
| 7124 | 74615 | How come we can find stocks with a Price-to-Book ratio less than 1? | 558617, 533818, 526110, 583708, 251100, 278582, 567034, 504243, 226070, 414940 |
| 2857 | 295864 | I have around 60K $. Thinking about investing in Oil, how to proceed? | 127566, 474575, 233732, 117451, 591705, 501384, 542051, 316444, 60175, 58065 |
| 8539 | 218728, 196304, 396038 | Can the risk of investing in an asset be different for different investors? | 283074, 483123, 391861, 385881, 502495, 150475, 500863, 378403, 142320, 469599 |
| 6896 | 251704 | Selling high, pay capital gains, re-purchase later | 468047, 448659, 343219, 169240, 66834, 366560, 522319, 257625, 581780, 97081 |
| 1085 | 467737, 393710 | How do disputed debts work on credit reports? | 372039, 242013, 450031, 78328, 161422, 574122, 319276, 422484, 339365, 233535 |
| 5422 | 151973 | What are some good books for learning stocks, bonds, derivatives e.t.c for beginner with a math background? | 193555, 165294, 552371, 241423, 273906, 276786, 191688, 463713, 172587, 221319 |
| 8247 | 42521, 475170, 321114, 465313 | Tax on Stocks or ETF's | 586010, 437907, 580802, 270992, 518735, 54947, 153112, 195767, 161019, 244813 |
| 10482 | 549072 | Rollover into bond fund to do dollar cost averaging [duplicate] | 330023, 447567, 525089, 211765, 224782, 134005, 237052, 473042, 265817, 439757 |
| 8789 | 70853 | What does “profits to the shareholders jumped to 15 cents a share” mean? | 87349, 41912, 234040, 341424, 339854, 212470, 219927, 14870, 20076, 317363 |
| 2423 | 538023 | At what age should I start or stop saving money? | 529444, 417787, 234846, 396127, 337561, 235855, 272328, 85977, 400419, 553288 |
| 4523 | 594257, 119165 | What should I do with my $25k to invest as a 20 years old? | 129255, 442776, 272070, 216365, 192669, 468485, 10476, 465819, 286746, 171712 |
| 8513 | 270573 | Buy on dip when earnings fail? | 572622, 462135, 175821, 203873, 573767, 321205, 391043, 278268, 351396, 73872 |
| 5054 | 28119 | How to stress test an investment plan? | 588481, 564007, 377186, 582899, 448745, 101124, 114092, 220665, 127263, 450558 |
| 1159 | 496064 | what is the best way to do a freelancing job over the summer for a student | 156063, 55064, 460648, 347181, 439899, 112669, 132287, 449155, 163881, 144948 |
| 68 | 19183 | Intentions of Deductible Amount for Small Business | 97233, 519473, 192516, 86134, 354716, 447231, 381151, 545296, 296098, 285502 |
| 9701 | 387141, 357739 | How to bet against the London housing market? | 473883, 408865, 73283, 410404, 253359, 225682, 412502, 70931, 242321, 45583 |
| 6199 | 239214 | How can all these countries owe so much money?  Why & where did they borrow it from? | 414693, 584273, 47163, 10399, 49602, 380714, 69518, 169921, 151971, 400046 |
| 9126 | 514831 | Short an option - random assignment? | 334473, 477588, 82194, 228810, 166227, 307518, 102316, 362473, 292045, 449147 |
| 34 | 599545 | 401k Transfer After Business Closure | 15728, 477603, 458917, 144109, 424766, 154839, 387876, 288074, 333616, 336917 |
| 6395 | 166227 | Option settlement for calendar spreads | 584223, 467463, 8177, 276314, 273142, 22916, 516790, 255927, 505223, 401447 |
| 9115 | 207325 | Why does the calculation for percentage profit vary based on whether a position is short vs. long? | 419897, 158520, 422467, 331606, 154665, 410822, 544857, 314478, 245082, 232880 |
| 4767 | 280805, 568670, 224057, 22804 | New car: buy with cash or 0% financing | 420018, 584106, 59372, 256803, 166314, 429153, 306834, 311446, 427884, 451092 |
| 8351 | 472516 | What happens when a calendar spread is assigned in a non-margin account? | 273142, 102316, 45674, 141213, 228810, 596567, 477588, 35102, 399903, 527654 |
| 6131 | 381720, 170204, 416679, 2460 | Is it ever a good idea to close credit cards? | 326094, 368806, 334111, 201982, 391384, 38938, 299819, 504293, 339030, 258465 |
| 858 | 45185, 122485, 278450 | Is it bad practice to invest in stocks that fluctuate by single points throughout the day? | 567608, 214281, 55751, 293027, 417062, 139699, 197856, 208932, 30774, 573612 |
| 4019 | 6881 | How and Should I Invest (As a college 18 year old with minimal living expenses)? | 332938, 493034, 332749, 85977, 287991, 55841, 20304, 553288, 168983, 66864 |
| 6080 | 164513 | Is ScholarShare a legitimate entity for a 529 plan in California? | 22856, 201500, 236732, 277581, 2809, 83080, 233401, 468527, 313392, 115175 |
| 6959 | 205010 | What is the term for the quantity (high price minus low price) for a stock? | 373034, 428117, 169954, 303325, 412223, 304399, 229573, 577573, 468025, 506460 |
| 5402 | 491350 | Is it impossible to get a home loan with a poor credit history after a divorce? | 445163, 90579, 51728, 440063, 595029, 227485, 52250, 180214, 310790, 460548 |
| 701 | 389446 | What are the ins/outs of writing-off part of one's rent for working at home? | 349672, 231990, 436505, 344955, 124507, 243306, 337706, 279339, 456234, 87113 |
| 5231 | 146188 | Where to find CSV or JSON data for publicly traded companies listed with their IPO date? | 510163, 146076, 558286, 9938, 592484, 548596, 53993, 542721, 480748, 200894 |
| 7674 | 519390 | Choosing the limit when making a limit order? | 526235, 447886, 15917, 200666, 249279, 260153, 514841, 31933, 164008, 94653 |
| 5940 | 486243, 433827, 93936 | How does investment into a private company work? | 512609, 250354, 567492, 182226, 252853, 135411, 179664, 473154, 454465, 119470 |
| 6612 | 205522, 322900 | If I have a lot of debt and the housing market is rising, should I rent and slowly pay off my debt or buy and roll th... | 502594, 301192, 431481, 180192, 343917, 373554, 443536, 317945, 254454, 14083 |
| 4714 | 450819 | Personal finance app where I can mark transactions as “reviewed”? | 505057, 584450, 218793, 529790, 29812, 353915, 344473, 479390, 374518, 65957 |
| 8456 | 486333 | What typically happens to unvested stock during an acquisition? | 257853, 93215, 555276, 492428, 243886, 534755, 469036, 186869, 87331, 104188 |
| 10213 | 270221, 545712 | Looking for good investment vehicle for seasonal work and savings | 195373, 263390, 38269, 272198, 80844, 386305, 405344, 391605, 285812, 96949 |
| 5196 | 172128, 114829 | I might use a credit card convenience check. What should I consider? | 565745, 85517, 402543, 2875, 85252, 456098, 289483, 467581, 572396, 481052 |
| 3006 | 269851, 568473, 328300 | Strategies for putting away money for a child's future (college, etc.)? | 332749, 8266, 538282, 127838, 512096, 290441, 372900, 258704, 303432, 353186 |
| 3909 | 312248, 404356, 193459, 245616, 353028 | How to rescue my money from negative interest? | 514003, 574011, 184838, 328499, 472837, 374400, 83330, 590968, 42475, 61586 |
| 5464 | 350399, 86691 | Resources on Buying Rental Properties | 222095, 26339, 423438, 87324, 465256, 383921, 325722, 545341, 207815, 372274 |
| 2385 | 407654 | As director, can I invoice my self-owned company? | 373059, 55440, 210889, 496064, 572690, 217472, 247760, 82119, 224438, 519321 |
| 10034 | 480749 | Tax implications of holding EWU (or other such UK ETFs) as a US citizen? | 44955, 181942, 528880, 197478, 565296, 180146, 447197, 437584, 141585, 145999 |
| 5090 | 436493 | Should I take a student loan to pursue my undergraduate studies in France? | 12988, 246286, 21913, 92430, 560681, 58005, 217831, 547246, 287507, 464356 |
| 2088 | 399875 | How would I go about selling the stock of a privately held company? | 53993, 455168, 599524, 488207, 291886, 168001, 413672, 530, 238215, 188776 |
| 9391 | 503637 | Should I replace bonds in a passive investment strategy | 535518, 248158, 107424, 545760, 577832, 136515, 171669, 494653, 155242, 543619 |
| 3148 | 178127, 438000 | Can a car company refuse to give me a copy of my contract or balance details? | 172855, 584305, 430100, 358631, 357280, 29721, 340287, 395995, 574398, 562282 |
| 4678 | 305153 | Finance, Cash or Lease? | 215225, 185405, 260095, 311748, 504918, 311446, 376016, 487678, 316035, 522532 |
| 2398 | 509391, 363810, 224654, 590489 | Frustrated Landlord | 556453, 487094, 96538, 44058, 98372, 201705, 436875, 395770, 309231, 249195 |
| 6746 | 210887 | What happens if stock purchased on margin plummets below what I have in the brokerage? | 279185, 247680, 231221, 251704, 333674, 115918, 392481, 283982, 399903, 434212 |
| 5511 | 169893, 560325, 478426, 383193, 278699, 12746 | Pay off car loan entirely or leave $1 until the end of the loan period? | 376016, 179891, 329137, 334559, 51873, 38786, 107898, 139788, 352027, 334964 |
| 8834 | 12232 | Pros/Cons of Buying Discounted Company Stock | 203139, 599156, 553331, 521095, 528827, 133644, 57387, 163396, 569303, 569224 |
| 988 | 226053, 107688 | Where should I invest my savings? | 168402, 60093, 347651, 167438, 279288, 501384, 211767, 285812, 388252, 450558 |
| 3369 | 163834, 231012, 145716, 411910, 395840 | Why should one only contribute up to the employer's match in a 401(k)? | 341493, 555377, 436071, 296405, 24231, 242556, 463892, 15841, 200898, 240373 |
| 9296 | 435746 | Why would Two ETFs tracking Identical Indexes Produce different Returns? | 148721, 206744, 159471, 428187, 285135, 368124, 408524, 410123, 492212, 261902 |
| 9245 | 194561 | Stock Options for a company bought out in cash and stock | 207253, 39345, 186869, 265111, 261487, 131488, 117177, 555276, 259560, 100628 |
| 3490 | 420529 | Tax Witholding for Stock Sale | 591157, 400730, 361482, 447651, 367742, 537371, 311782, 571124, 152960, 279606 |
| 5763 | 462019 | What is the best way to get a “rough” home appraisal prior to starting the refinance process? | 563380, 251466, 570318, 67379, 565868, 38712, 503444, 89964, 331255, 215647 |
| 4962 | 599925 | Net Cash Flows from Selling the Bond and Investing | 416839, 34949, 393838, 408661, 490077, 187110, 308276, 431386, 152265, 158363 |
| 4846 | 151104 | Is there anything comparable to/resembling CNN's Fear and Greed Index? | 98096, 3533, 335892, 538974, 415161, 183597, 317666, 320059, 270305, 92650 |
| 9403 | 6666, 328086, 345199 | Abundance of Cash - What should I do? | 410450, 159235, 570632, 14349, 228403, 273978, 217865, 70806, 551986, 499548 |
| 5993 | 367375, 224918, 272866, 230215, 55084, 5827, 352638, 426120, 63501 | Why would anyone want to pay off their debts in a way other than “highest interest” first? | 94373, 416796, 160193, 287571, 494306, 353911, 886, 128574, 40522, 590145 |
| 5710 | 232311 | Bucketing investments to track individual growths | 411856, 227364, 135765, 516267, 88417, 412830, 177036, 508610, 406192, 227733 |
| 7529 | 66607 | Does the expense ratio of a fund-of-funds include the expense ratios of its holdings? | 89297, 464337, 102904, 514529, 59249, 293626, 65587, 534019, 218261, 88823 |
| 5021 | 589285 | Is there a more flexible stock chart service, e.g. permitting choice of colours when comparing multiple stocks? | 528576, 189341, 494939, 555506, 584801, 211444, 535928, 252084, 511861, 465971 |
| 3612 | 259625 | How can I buy and sell the same stock on the same day? | 522658, 390864, 567383, 292159, 402726, 310636, 352415, 378821, 200879, 367873 |
| 4409 | 499128, 100306, 147439 | My friend wants to put my name down for a house he's buying. What risks would I be taking? | 268078, 223841, 243732, 60135, 514790, 341947, 102326, 84732, 102088, 306926 |
| 2070 | 363678 | Advantage of credit union or local community bank over larger nationwide banks such as BOA, Chase, etc.? | 550303, 587737, 30253, 469515, 578357, 590209, 408166, 597571, 38038, 249839 |
| 11039 | 53544, 249063 | Pay off credit card debt or earn employer 401(k) match? | 91183, 552383, 5203, 163287, 287876, 420727, 457945, 79363, 508534, 105557 |
| 5460 | 184337, 21174, 108514, 463885 | Paying off a loan with a loan to get a better interest rate | 344812, 343208, 77052, 576609, 106495, 470716, 243065, 194557, 327115, 399259 |
| 7925 | 318185, 402482 | Can I sell a stock immediately? | 591436, 332467, 227399, 581866, 221869, 315760, 310636, 81721, 44461, 339419 |
| 4286 | 566069 | Given advice “buy term insurance and invest the rest”, how should one “invest the rest”? | 70460, 151817, 10531, 155640, 564675, 391243, 229239, 206830, 71926, 272590 |
| 2685 | 154113, 370300, 468923 | What ways are there for us to earn a little extra side money? | 382005, 576047, 109880, 594182, 280099, 237950, 431751, 269380, 77088, 468086 |
| 1090 | 518896 | Need a formula to determine monthly payments received at time t if I'm reinvesting my returns | 393987, 16051, 446454, 19999, 166394, 179365, 520217, 478339, 584231, 209238 |
| 6122 | 44344 | Better to rent condo to daughter or put her on title? | 496166, 316794, 403515, 53840, 80269, 182039, 267592, 577658, 424125, 566184 |
| 4514 | 69485, 337764, 209804 | What intrinsic, non-monetary value does gold have as a commodity? | 156211, 146573, 471825, 426270, 240894, 277694, 317429, 80141, 408336, 213517 |
| 8507 | 370995 | When to sell a stock? | 99132, 251536, 303724, 273565, 272091, 545284, 102237, 217837, 236415, 165970 |
| 6221 | 257248, 519675, 76414, 169688, 455614, 115717 | To pay off a student loan, should I save up a lump sum payoff payment or pay extra each month? | 254245, 529551, 352363, 448791, 124705, 188384, 110081, 274108, 413313, 192602 |
| 3008 | 180192, 323406 | What are my chances at getting a mortgage with Terrible credit but High income | 102266, 285694, 2064, 407401, 574438, 16563, 112435, 78176, 317461, 47441 |
| 4007 | 521657 | What is a reasonable salary for the owner and sole member of a small S-Corp? | 556220, 260385, 334603, 370542, 170933, 205341, 325677, 352838, 388704, 543085 |
| 6644 | 175035 | How to know precisely when a SWIFT is issued by a bank? | 475527, 110198, 118396, 554518, 218761, 298587, 355870, 589616, 350396, 143677 |
| 10267 | 328556, 460398 | How should I prepare for the next financial crisis? | 178693, 143393, 569632, 305600, 36961, 326398, 369470, 60166, 96017, 182442 |
| 7622 | 253369, 378594 | Best way to pay off debt? | 220241, 373554, 345895, 480773, 457945, 508952, 388095, 353911, 105557, 329249 |
| 3767 | 153922, 392060 | What should I be doing to protect myself from identity theft? | 423809, 90632, 91986, 260580, 97686, 581889, 158285, 158008, 587778, 320246 |
| 6410 | 471723 | Will an ETF immediately reflect a reconstitution of underlying index | 330729, 214281, 454610, 87238, 71230, 47276, 195089, 295993, 133196, 313897 |
| 5030 | 215540 | Why pay for end-of-day historical prices? | 227192, 532178, 352415, 471131, 13511, 560108, 149420, 295344, 113150, 370569 |
| 6252 | 394551, 160932, 293624, 233294, 243268, 379487, 62868 | Is this mortgage advice good, or is it hooey? | 473647, 495089, 213713, 120061, 47565, 205906, 139366, 423403, 27268, 71709 |
| 885 | 337165, 409184 | How long do credit cards keep working after you disappear? | 254968, 472336, 99449, 181757, 251701, 516678, 596284, 434082, 319213, 261697 |
| 766 | 550172 | Will the ex-homeowner still owe money after a foreclosure? | 2996, 212827, 163711, 299591, 427110, 333583, 408213, 552768, 268865, 578906 |
| 2183 | 571625 | Why are there many small banks and more banks in the U.S.? | 132678, 248863, 543921, 550924, 24344, 18749, 498146, 88196, 38038, 155960 |
| 8202 | 513258, 93971 | What accounted for DXJR's huge drop in stock price? | 317363, 457689, 71924, 337001, 421248, 412584, 355167, 537862, 122542, 238634 |
| 7345 | 237645 | What do these numbers mean? (futures) | 276381, 127845, 206895, 460331, 527080, 9274, 273789, 87548, 285648, 354429 |
| 776 | 583640, 127263, 496899 | Can saving/investing 15% of your income starting age 25, likely make you a millionaire? | 124027, 417787, 10440, 41960, 143591, 374266, 434972, 418281, 467044, 554833 |
| 89 | 248624 | How can I deposit a check made out to my business into my personal account? | 508754, 309023, 135196, 526817, 80538, 308938, 521540, 188893, 400230, 590102 |
| 9329 | 523913 | Interactive Brokers: IOPTS and list of structured products | 326991, 89351, 93727, 453524, 238173, 229582, 262591, 289853, 381512, 334383 |
| 1920 | 269943 | Clarification on student expenses - To file the tax for the next year | 263485, 481114, 585356, 316482, 295562, 150271, 551187, 446889, 202019, 304248 |
| 8013 | 496159, 224231 | Frequency of investments to maximise returns (and minimise fees) | 384983, 537626, 57033, 81652, 40652, 388389, 385955, 224816, 308764, 28291 |
| 3759 | 527966, 67167, 522358 | Simplifying money management | 373772, 455457, 214248, 122378, 145812, 490065, 372743, 248663, 309650, 129681 |
| 10639 | 431799, 495774, 278453, 187039 | Short term parking of a large inheritance? | 163353, 235628, 131391, 111048, 590276, 171196, 178386, 7625, 546538, 431443 |
| 4312 | 399149 | Is it true that 90% of investors lose their money? | 282435, 222639, 167950, 285945, 497786, 116647, 170628, 532485, 507284, 300770 |
| 6525 | 181985 | Does it make sense to trade my GOOGL shares for GOOG and pocket the difference? | 106541, 550661, 98150, 156467, 362473, 498014, 147002, 374867, 53263, 488920 |
| 2590 | 589625 | Are non-residents or foreigners permitted to buy or own shares of UK companies? | 296528, 209493, 528880, 262485, 307776, 310103, 158923, 536483, 73457, 456999 |
| 5374 | 152688 | What were the main causes of the spike and drop of DRYS's stock price? | 283106, 133204, 457689, 467594, 362462, 261522, 122542, 137079, 50141, 78648 |
| 2994 | 419319 | Work on the side for my wife's company | 569145, 510317, 547793, 382005, 491844, 5840, 113098, 399882, 423625, 269380 |
| 3683 | 185909, 454501 | Can I trust the Motley Fool? | 276975, 408995, 105973, 428848, 500338, 301739, 6607, 192912, 534080, 526015 |
| 7206 | 441155, 532211, 553066 | Who Bought A Large Number Of Shares? | 34882, 573846, 358164, 351570, 65667, 518526, 327525, 558703, 350214, 23108 |
| 10246 | 512984, 77573 | Understanding the T + 3 settlement days rule | 370635, 176717, 179520, 327080, 156029, 340263, 28314, 483704, 11927, 332243 |
| 5241 | 322157, 27489 | Mortgage vs. Cash for U.S. home buy now | 344740, 213713, 281675, 266649, 426211, 438073, 390976, 4739, 42604, 111184 |
| 98 | 575929 | How can I make $250,000.00 from trading/investing/business within 5 years? | 102113, 66034, 209067, 129540, 336661, 527522, 373119, 572563, 506149, 519619 |
| 4615 | 262934 | Are solar cell panels and wind mills worth the money? | 261900, 69523, 464715, 455798, 496427, 425595, 510872, 210811, 376430, 271015 |
| 6467 | 453256, 23217, 346641, 367313 | Advice on strategy for when to sell | 88813, 109455, 203873, 217837, 240089, 504235, 498075, 130941, 83807, 99857 |
| 4289 | 24881 | Does the currency exchange rate contain any additional information at all? | 288330, 119316, 517345, 17469, 356465, 135220, 300213, 439779, 324546, 535617 |
| 4394 | 336045, 441582 | Transfer $50k to another person's account (in California, USA) | 322838, 93386, 462585, 293653, 305907, 495827, 521753, 431462, 415655, 572842 |
| 7344 | 108403 | How is the Dow divisor calculated? | 14368, 150430, 159166, 378974, 253926, 313421, 501032, 65618, 591089, 418150 |
| 6875 | 224392 | Where to find free Thailand stock recommendations and research? | 567500, 110733, 352557, 9354, 284539, 224366, 556770, 232460, 452983, 559105 |
| 10447 | 152096, 300721 | Is there an advantage to a traditional but non-deductable IRA over a taxable account? [duplicate] | 447482, 144751, 500175, 382236, 532657, 259150, 299690, 406239, 382894, 403103 |
| 9871 | 448890, 40051, 170594 | What should I do with the 50k I have sitting in a European bank? | 73741, 293179, 367207, 76562, 455856, 175139, 212464, 362035, 524615, 164908 |
| 3404 | 556976 | In US, is it a good idea to hire a tax consultant for doing taxes? | 197870, 488574, 395483, 37725, 364463, 251564, 432619, 158058, 525360, 420295 |
| 3625 | 414295 | What should I do with my paper financial documents? | 509617, 500751, 163168, 569812, 380263, 245967, 305544, 37582, 513248, 113830 |
| 6005 | 135415, 478457, 345895, 73310, 384626, 390689 | Why might it be advisable to keep student debt vs. paying it off quickly? | 149500, 571198, 25190, 431884, 507544, 96268, 564206, 67091, 52136, 414288 |
| 3822 | 385090, 308837 | How to change a large quantity of U.S. dollars into Euros? | 340777, 19618, 194730, 292714, 390524, 417917, 79777, 549787, 531953, 239876 |
| 7879 | 372551, 421285 | Any Tips on How to Get the Highest Returns Within 4 Months by Investing in Stocks? | 58186, 102029, 228488, 7625, 540919, 272174, 43088, 593879, 174313, 4772 |
| 3115 | 234950, 316794 | How can I live outside of the rat race of American life with 300k? | 233562, 183869, 252852, 136035, 129364, 267892, 475736, 174272, 264976, 150066 |
| 3995 | 278734, 230208 | I have more than $250,000 in a US Bank account… mistake? | 485507, 404954, 352883, 200690, 171720, 234934, 14349, 583803, 303367, 17378 |
| 8002 | 34767 | What is the tax treatment of scrip dividends in the UK? | 118786, 32600, 217006, 162454, 110983, 267067, 115333, 97842, 263312, 216694 |
| 10136 | 526115 | How to minimise the risk of a reduction in purchase power in case of Brexit for money held in a bank account? | 466950, 417740, 182195, 304007, 284305, 290930, 205685, 432565, 583903, 35511 |
| 8635 | 67107, 240215 | Is there any flaw in this investment scheme? | 493841, 46818, 510565, 151238, 440305, 480802, 203729, 447619, 303619, 202638 |
| 5206 | 563030, 28230, 117276, 300660 | Is it a good idea to get an unsecured loan to pay off a credit card that won't lower a high rate? | 298908, 287571, 516397, 595455, 225522, 69938, 504293, 153088, 1472, 340520 |
| 2713 | 388147 | Physical Checks - Mailing | 284528, 41944, 29372, 20791, 268257, 402898, 216200, 199069, 78139, 350237 |
| 9060 | 40447 | Buying puts without owning underlying | 511093, 181924, 338782, 401447, 7743, 528052, 228217, 369031, 374797, 294688 |
| 4105 | 25096 | As an investor what are side effects of Quantitative Easing in US and in EU? | 416483, 176262, 345910, 305029, 393791, 339640, 369038, 30946, 108519, 239214 |
| 2465 | 570680, 81046, 546509 | Can capital expenses for volunteer purposes be deducted from income? | 37382, 598646, 490176, 202645, 275543, 146657, 510716, 432545, 541809, 82199 |
| 4640 | 101369 | What can my relatives do to minimize their out of pocket expenses on their fathers estate | 372808, 295246, 331534, 522619, 117960, 338539, 45819, 521803, 37823, 367404 |
| 9275 | 338754, 14364 | Do I have to pay a capital gains tax if I rebuy the same stock within 30 days? | 400730, 537916, 23217, 161155, 407602, 343219, 376800, 596518, 390864, 292762 |
| 3500 | 174019 | Why invest in becoming a landlord? | 528206, 273187, 71424, 572061, 557478, 11601, 41356, 578597, 487094, 76283 |
| 1306 | 484437, 204075 | I made an investment with a company that contacted me, was it safe? | 160611, 594206, 333050, 309851, 519038, 12391, 309590, 78632, 316321, 40870 |
| 6262 | 26799 | Help required on estimating SSA benefit amounts | 34538, 118707, 2648, 529927, 498444, 320362, 430407, 83338, 151153, 489376 |
| 8632 | 213976 | Is it best to exercise options shares when they vest, or wait | 43497, 237718, 104188, 237783, 259560, 255927, 293959, 382381, 163396, 200784 |
| 6133 | 415705 | What happens to all of the options when they expire? | 575408, 7733, 11456, 581672, 177559, 242298, 186869, 176786, 132288, 480879 |
| 3771 | 488948, 198349, 217683, 49601 | Best way to buy Japanese yen for travel? | 490384, 495826, 96211, 217715, 306130, 521712, 120604, 434201, 575495, 210587 |
| 1736 | 25543, 443419 | How can people have such high credit card debts? | 569056, 399406, 174941, 562896, 372993, 233892, 517050, 298908, 99463, 201447 |
| 6814 | 340214, 223206 | Selling Stock - All or Nothing? | 400614, 67107, 66834, 276883, 154976, 279782, 513734, 590188, 337736, 178497 |
| 1322 | 399418, 64138 | Is this follow-up after a car crash a potential scam? | 226090, 44635, 332916, 538217, 114231, 219119, 524723, 567973, 519692, 397852 |
| 5185 | 210236, 317354 | Invest in low cost small cap index funds when saving towards retirement? | 196992, 376485, 241202, 580313, 262180, 524525, 268731, 106620, 567256, 523331 |
| 6909 | 127012 | Why do stocks priced above $2.00 on the ASX sometimes move in $0.005 increments? | 72633, 118232, 375019, 376399, 112946, 490584, 168080, 102026, 47217, 64943 |
| 2348 | 211867, 566573, 211622, 474234, 352271, 265874 | Why can't you just have someone invest for you and split the profits (and losses) with him? | 447619, 306430, 151412, 420544, 247486, 381757, 177194, 389004, 64410, 134864 |
| 687 | 146021 | Online tool to connect to my bank account and tell me what I spend in different categories? | 478807, 584175, 447597, 258423, 396222, 291278, 140657, 45718, 479390, 476363 |
| 4499 | 76996 | Is investing exlusively in a small-cap index fund a wise investment? | 52274, 228111, 196992, 501153, 517391, 312591, 428197, 335136, 78837, 372233 |
| 42 | 272709, 327263 | What are the ins/outs of writing equipment purchases off as business expenses in a home based business? | 305222, 88967, 510863, 398536, 581265, 28764, 283079, 226519, 443859, 571711 |
| 3530 | 239998 | How to exclude stock from mutual fund | 184299, 370754, 24029, 530938, 322070, 110343, 479420, 378075, 226967, 209879 |
| 9108 | 472585 | Starting an investment portfolio with Rs 5,000/- | 290757, 46967, 336218, 240351, 414116, 171189, 282947, 7748, 262196, 95282 |
| 6835 | 102243 | Are bond ETF capital gains taxed similar to stock or stock funds if held for more than 1 year? | 149305, 570546, 5710, 110343, 581632, 537916, 287950, 153112, 84238, 423929 |
| 3067 | 406156 | Should I make extra payments to my under water mortgage or increase my savings? | 3092, 476068, 477907, 83543, 131365, 560915, 423403, 11791, 336276, 382629 |
| 4125 | 344648, 72046 | Alternative means of salary for my employees | 174787, 73999, 127347, 361954, 355244, 562211, 345795, 365558, 455524, 87231 |
| 1150 | 43603, 19936 | How are the best way to make and save money at 22 years old | 10476, 529444, 433986, 216365, 353369, 305946, 595287, 319760, 494815, 45353 |
| 7705 | 195191 | Why would I pick a specific ETF over an equivalent Mutual Fund? | 500486, 161019, 580802, 539263, 358997, 153112, 312015, 377429, 364735, 253971 |
| 9808 | 40702, 431946 | Selling To Close | 416307, 557582, 345368, 151587, 362473, 107045, 374204, 594614, 43087, 31465 |
| 3888 | 319213, 239632 | Why I can't view my debit card pre-authorized amounts? | 208169, 316652, 432077, 276733, 185434, 294077, 181757, 418580, 440527, 425487 |
| 10109 | 506374, 406974, 499849 | Why does Charles Schwab have a Mandatory Settlement Period after selling stocks? | 28314, 93231, 266725, 332243, 570248, 563826, 370635, 537212, 36193, 296475 |
| 2790 | 279329, 469125 | Should I pay more than 20% down on a home? | 472484, 357200, 400896, 75961, 64400, 234890, 385343, 207564, 296906, 215103 |
| 4999 | 314898 | Looking for a good source for Financial Statements | 9938, 171964, 431459, 597241, 465971, 295738, 425624, 11263, 546115, 343964 |
| 3189 | 225395 | Diversify my retirement investments with a Roth IRA | 122222, 240975, 287225, 187124, 347651, 327060, 404800, 18436, 246109, 88311 |
| 5134 | 158523 | Why does Yahoo Finance's data for a Vanguard fund's dividend per share not match the info from Vanguard? | 532616, 46774, 405474, 465536, 206727, 54225, 100485, 584128, 559884, 239137 |
| 1321 | 216456, 292065 | Are social media accounts (e.g. YouTube, Twitter, Instagram, etc.) considered assets? | 392379, 229640, 203715, 381322, 521823, 210019, 235082, 426184, 559738, 481036 |
| 4539 | 370879 | How should I save money if the real interest rate (after inflation) is negative? | 42475, 472837, 203926, 32744, 449745, 257547, 322816, 275925, 220720, 61586 |
| 715 | 579763, 546538, 187404 | what would you do with $100K saving? | 548758, 337561, 427032, 508343, 387647, 162633, 333784, 366961, 273925, 524615 |
| 504 | 344203, 498751 | Have plenty of cash flow but bad credit | 22807, 70806, 569240, 93573, 41875, 68431, 495431, 252473, 368247, 231412 |
| 2296 | 83330, 366594, 253563 | How does a bank make money on an interest free secured loan? | 400009, 396853, 259919, 119298, 198007, 175824, 395769, 279897, 10948, 272150 |
| 10975 | 61022 | How to contribute to Roth IRA when income is at the maximum limit & you have employer-sponsored 401k plans? | 163865, 81148, 140330, 360533, 441632, 447482, 101490, 17166, 189989, 446226 |
| 1994 | 156640 | Does the IRS reprieve those who have to commute for work? | 231990, 434846, 192843, 263259, 243356, 544381, 63919, 479649, 22819, 380635 |
| 9164 | 365298, 263390 | Bonds vs equities: crash theory | 309326, 115648, 287656, 321941, 149900, 296516, 506743, 599420, 464824, 418528 |
| 1812 | 530570 | splitting a joint mortgage - one owner in home | 159590, 219274, 319159, 75235, 60163, 287458, 570964, 544663, 395520, 150598 |
| 10462 | 8266, 11378, 35680, 437879, 204035, 581204 | Is it okay to be married, 30 years old and have no retirement? | 268023, 151774, 66376, 160105, 267204, 108845, 15322, 147422, 361049, 361126 |
| 8855 | 208165, 312821 | How do i get into investing stocks [duplicate] | 155677, 555521, 367415, 142320, 560395, 152286, 403092, 67327, 312445, 339553 |
| 7071 | 124230 | ESPP strategy - Sell right away or hold? | 133644, 294573, 127702, 511678, 35575, 71713, 361345, 575213, 434812, 387035 |
| 8974 | 523331, 356595, 170625 | As a 22-year-old, how risky should I be with my 401(k) investments? | 216365, 140738, 10476, 336144, 102501, 452126, 134931, 199237, 279570, 442776 |
| 5178 | 240261 | Formula that predicts whether one is better off investing or paying down debt | 557506, 139788, 39819, 262772, 279329, 396889, 473865, 111815, 290434, 106215 |
| 5061 | 23747 | What fiscal scrutiny can be expected from IRS in early retirement? | 151263, 502150, 398078, 96720, 517903, 25481, 499864, 513392, 149742, 521843 |
| 7880 | 85319 | Are there index tracking funds that avoid the “buy high - sell low” problem? | 299284, 184299, 272790, 81652, 36534, 148721, 70194, 83807, 416839, 241202 |
| 2075 | 170042, 359580, 14967 | Are stories of turning a few thousands into millions by trading stocks real? | 33357, 65667, 285147, 519501, 506149, 188129, 44417, 519619, 555521, 78224 |
| 4335 | 357013 | What is the US Fair Tax? | 318583, 419466, 322246, 363178, 585212, 589161, 3181, 412877, 470879, 516548 |
| 7533 | 93853 | Investing tax (savings) | 8012, 516756, 162668, 480827, 142658, 82482, 526664, 563986, 571218, 250603 |
| 1393 | 352838, 539133 | Which is better when working as a contractor, 1099 or incorporating? | 586026, 578196, 220022, 32072, 234436, 532932, 68524, 352640, 277812, 93638 |
| 5264 | 576564 | Does a company's stock price give any indication to or affect their revenue? | 52579, 431814, 218326, 226063, 505694, 8643, 270926, 371720, 264396, 111076 |
| 9733 | 110163, 38655 | Due Diligence - Dilution? | 121262, 301880, 316321, 23414, 450132, 135798, 19354, 526073, 78236, 108965 |
| 7311 | 323768 | Finance, Social Capital IPOA.U | 199746, 507841, 290325, 583646, 264740, 261231, 275494, 327911, 479752, 463599 |
| 744 | 566480, 78176 | What options are available for a home loan with poor credit but a good rental history? | 573276, 310790, 596272, 415425, 289450, 80607, 67066, 517633, 92397, 468104 |
| 7141 | 132288 | Do investors go long option contracts when they cannot cover the exercise of the options? | 570046, 538054, 243714, 305676, 383328, 288289, 507828, 394066, 401447, 273142 |
| 4071 | 129875 | If our economy crashes, and cash is worthless, should i buy gold or silver | 524142, 566669, 473965, 291862, 505136, 70575, 487817, 362102, 502634, 467825 |
| 7512 | 191060 | understanding the process/payment of short sale dividends | 480949, 222320, 115553, 202985, 487329, 259450, 568166, 409432, 241425, 298284 |
| 1391 | 562176 | How is taxation for youtube/twitch etc monetization handled in the UK? | 267067, 254151, 510599, 266229, 527951, 357419, 223170, 21136, 315516, 251649 |
| 7534 | 358125 | Can you explain why it's better to invest now rather than waiting for the market to dip? | 175821, 426157, 221869, 89714, 419747, 310218, 145539, 71219, 94302, 114806 |
| 5356 | 312405 | Historical stock prices: Where to find free / low cost data for offline analysis? | 535343, 560108, 279785, 391171, 240086, 529877, 546379, 537111, 226749, 47798 |
| 2579 | 432020 | What to do when a job offer is made but with a salary less than what was asked for? | 423070, 181213, 200946, 432808, 524471, 203626, 552290, 256802, 157919, 489554 |
| 5790 | 134794 | FX losses on non-UK mortgage for UK property - tax deductable? | 403948, 11122, 141738, 33117, 369419, 149341, 507107, 256395, 356884, 129695 |
| 7823 | 583549 | Retirement Funds: Betterment vs Vanguard Life strategy vs Target Retirement | 451196, 105666, 175927, 347825, 331492, 374225, 11094, 268731, 57070, 172336 |
| 689 | 411044 | Receive credit card payment sending my customer details to a credit card processing company? | 446932, 204288, 104079, 96547, 171761, 115868, 195852, 438032, 521688, 421803 |
| 9174 | 535317 | Which U.S. online discount broker is the best value for money? | 192910, 236931, 200052, 451729, 413856, 365465, 513281, 515144, 522798, 31936 |
| 6867 | 443804, 540799, 445258 | Will there always be somebody selling/buying in every stock? | 230343, 61006, 18532, 543589, 226197, 229573, 301985, 88157, 208070, 466143 |
| 2383 | 232199 | Should I Purchase Health Insurance Through My S-Corp | 17215, 527620, 224406, 534277, 546634, 457034, 476085, 349348, 170933, 521489 |
| 5083 | 138845 | Co-signer deceased | 369075, 18257, 273759, 270952, 453263, 447983, 267182, 495482, 142876, 334606 |
| 10526 | 39185 | What extra information might be obtained from the next highest bids in an order book? | 485973, 546493, 283008, 427747, 298551, 138830, 251100, 322798, 467852, 109345 |
| 2181 | 376631, 397329 | What are the risks & rewards of being a self-employed independent contractor / consultant vs. being a permanent emplo... | 37725, 488755, 139501, 197870, 406656, 383088, 260603, 145016, 584218, 524788 |
| 5903 | 231863 | Fees aside, what factors could account for performance differences between U.S. large-cap index ETFs? | 408524, 395842, 372233, 517391, 370244, 246996, 159471, 402091, 501153, 188497 |
| 5620 | 448784, 329552, 548740 | What's the fuss about identity theft? | 158285, 260580, 90632, 598801, 91986, 551747, 581889, 423809, 531137, 97686 |
| 5254 | 392851 | How do I calculate the quarterly returns of a stock index? | 402466, 574974, 563169, 559168, 96697, 508405, 99708, 531066, 422904, 457873 |
| 2472 | 370334, 307315 | How do I deal with a mistaken attempt to collect a debt from me that is owed by someone else? | 568230, 62109, 200263, 180601, 191660, 201758, 584582, 595441, 500671, 87977 |
| 2306 | 315875 | To whom should I report fraud on both of my credit cards? | 581889, 596284, 531137, 90632, 289706, 164729, 412542, 290069, 59023, 298729 |
| 7633 | 197839 | Can a trade happen “in between” the bid and ask price? | 494727, 353396, 505244, 281844, 458933, 137175, 402482, 554207, 164008, 442048 |
| 2400 | 564271 | Will I be paid dividends if I own shares? | 1198, 1034, 95889, 456470, 311214, 481169, 501931, 97942, 365627, 348922 |
| 5549 | 286227, 309361 | Pros / cons of being more involved with IRA investments [duplicate] | 429106, 90294, 471204, 546150, 87260, 555821, 336394, 358371, 324012, 182305 |
| 4605 | 453941 | If the U.S. defaults on its debt, what will happen to my bank money? | 400826, 41312, 169691, 526384, 313306, 229310, 598030, 559618, 373717, 479527 |
| 2885 | 367360, 359579, 85229, 454810 | Merits of buying apartment houses and renting them | 430672, 502291, 507029, 581251, 159403, 502514, 129149, 80838, 451849, 3217 |
| 6110 | 331850, 94117, 259706 | Why does short selling require borrowing? | 188531, 320450, 314478, 67107, 226496, 107045, 79764, 49794, 35500, 329662 |
| 4823 | 104726 | Close to retirement & we may move within 7 years. Should we re-finance our mortgage, or not? | 561929, 187590, 302093, 376084, 551099, 175200, 7094, 434519, 322804, 361126 |
| 3694 | 282442 | Has anyone created a documentary about folks who fail to save enough for retirement? | 204747, 242237, 385932, 222260, 91911, 383051, 489480, 568527, 582307, 311931 |
| 8 | 566392 | How to deposit a cheque issued to an associate in my business into my business account? | 65404, 508754, 261856, 590102, 564553, 301833, 29372, 25397, 489199, 308938 |
| 1309 | 156162, 489401 | Why does FlagStar Bank harass you about payments within grace period? | 471630, 271040, 489368, 15824, 438869, 329817, 336792, 125497, 490427, 366594 |
| 7109 | 447781 | How do I analyse moving averages? | 489933, 42620, 140804, 227669, 565501, 193012, 257185, 35006, 518932, 221627 |
| 5080 | 256055 | Is there a standard or best practice way to handle money from an expiring UTMA account? | 490991, 445521, 414429, 279291, 451189, 470928, 69841, 236186, 324564, 174406 |
| 4981 | 247894 | Where can I find open source portfolio management software? | 102684, 45218, 259463, 587792, 557861, 78436, 196432, 529790, 419171, 193805 |
| 7445 | 153178, 104343 | IS it the wrong time to get into the equity market immediately after large gains? | 79111, 573612, 483025, 488207, 590902, 516880, 350068, 89714, 8135, 284075 |
| 2895 | 328691 | Where should a young student put their money? | 426461, 354551, 517313, 55841, 442896, 496170, 422946, 148453, 31452, 332022 |
| 6787 | 587120 | Would it make sense to sell a stock, then repurchase it for tax purposes? | 23217, 219762, 400730, 474981, 263751, 328073, 390864, 221715, 374867, 17184 |
| 6041 | 241308 | Most effective Fundamental Analysis indicators for market entry | 425020, 81655, 224695, 96910, 528034, 108579, 331530, 194240, 204297, 55002 |
| 7700 | 273761, 2653, 179328 | Should I re-allocate my portfolio now or let it balance out over time? | 224392, 253268, 269169, 395208, 422051, 131127, 552298, 154552, 22221, 394941 |
| 547 | 6349 | What percentage of my company should I have if I only put money? | 447353, 396694, 95243, 523158, 556421, 213399, 80913, 68088, 146547, 409039 |
| 3394 | 342258, 129319 | What is the easiest way to back-test index funds and ETFs? | 448745, 528034, 172374, 224765, 71230, 571913, 570664, 472663, 445971, 428187 |
| 4102 | 448699 | How can I determine if my rate of return is “good” for the market I am in? | 535737, 597437, 88801, 71219, 160786, 162488, 461082, 369439, 554734, 46394 |

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
