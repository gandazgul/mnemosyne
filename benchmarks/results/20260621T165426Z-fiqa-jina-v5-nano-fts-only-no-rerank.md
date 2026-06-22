# Mnemosyne BEIR fiqa Results

- Generated: 2026-06-21T16:54:26Z
- Dataset: `fiqa`
- Queries: 648
- Corpus documents: 57638
- Search limit: 100
- Source mode: `fts-only`
- Rerank candidates: 100
- Rerank enabled: False
- Run label: `jina-v5-nano`
- Mnemosyne config: `/Users/gandazgul/Documents/web/mnemosyne/cli/configs/jina-v5-text-nano-retrieval.yaml`

| Metric | Score |
| --- | ---: |
| `ndcg@10` | 0.2297 |
| `mrr@10` | 0.2798 |
| `recall@10` | 0.2958 |
| `recall@100` | 0.5161 |
| `map@100` | 0.1833 |

## Breakdown

| Detail | Count |
| --- | ---: |
| `queries` | 648 |
| `queries_with_rank_1_hit` | 132 |
| `queries_with_top_10_hit` | 306 |
| `queries_with_top_100_hit` | 458 |
| `queries_missing_at_100` | 190 |

### First Relevant Rank

| Bucket | Queries |
| --- | ---: |
| `rank_1` | 132 |
| `rank_2_3` | 71 |
| `rank_4_10` | 103 |
| `rank_11_100` | 152 |
| `missing@100` | 190 |

### Missing At 100

| Query ID | Missing Doc IDs | Query | Top 10 Doc IDs |
| --- | --- | --- | --- |
| 4641 | 319954, 397358, 88327 | Where should I park my rainy-day / emergency fund? | 376148, 497993, 580025, 253614, 32833, 386305, 583695, 416189, 538023, 108978 |
| 5503 | 64279 | Tax considerations for selling a property below appraised value to family? | 387218, 181360, 519633, 319307, 116446, 299811, 497530, 375632, 309077, 565868 |
| 7803 | 565926 | Can the Delta be used to calculate the option premium given a certain target? | 202432, 482238, 371621, 197863, 16292, 416286, 185401, 62151, 270345, 273598 |
| 3451 | 192307, 276883, 588448, 165970, 490170, 99132 | Should you keep your stocks if you are too late to sell? | 251536, 417506, 495568, 466121, 290490, 371044, 475457, 190619, 437380, 536345 |
| 4804 | 583651, 104395 | How do financial services aimed at women differ from conventional services? | 418888, 483485, 227910, 447167, 76640, 78259, 246191, 401677, 137984, 583631 |
| 10809 | 103362 | Definitions of leverage and of leverage factor | 314410, 286141, 62955, 543811, 153179, 536043, 379314, 511587, 541718, 325787 |
| 753 | 466718, 243503 | Taxes due for hobbyist Group Buy | 370857, 79511, 122185, 539244, 51729, 158738, 566171, 88947, 530850, 502126 |
| 4465 | 376575 | How to donate to charity that will make a difference? | 81046, 174033, 475896, 569461, 375708, 326167, 8849, 90591, 221427, 561377 |
| 7326 | 584295 | Do brokers execute every trade on the exchange? | 35340, 257656, 486058, 37040, 404339, 499536, 87535, 503981, 343638, 150735 |
| 5951 | 374480, 473765, 298065, 596834, 497260 | Why can't house prices be out of tune with salaries | 599860, 418034, 259777, 285525, 27444, 179600, 196405, 589470, 82767, 537664 |
| 10827 | 160786, 42301, 7748, 107554, 95282 | How much should I be contributing to my 401k given my employer's contribution? | 509772, 555377, 403930, 63532, 15841, 81148, 529543, 139059, 181624, 538238 |
| 9771 | 263955, 28740 | Is there any emprical research done on 'adding to a loser' | 38287, 589507, 131258, 370632, 107545, 273718, 450355, 541485, 226646, 169028 |
| 3724 | 309200, 508921, 434190, 279570, 497216, 552887, 341146, 199970 | Should you always max out contributions to your 401k? | 480036, 159921, 497561, 53028, 7904, 411910, 340624, 198825, 488673, 336917 |
| 5853 | 476663, 160105, 495699, 473647, 431811, 424598 | Paying Off Principal of Home vs. Investing In Mutual Fund | 549270, 535581, 367960, 87722, 472159, 336218, 476517, 284318, 569948, 290757 |
| 570 | 363591 | Employer options when setting up 401k for employees | 79375, 346387, 117845, 242529, 396968, 477175, 15841, 168890, 150883, 110114 |
| 7279 | 322456, 464502 | If I invest in securities denominated in a foreign currency, should I hedge my currency risk? | 213373, 575046, 19618, 98461, 83860, 322168, 66119, 128048, 197918, 334495 |
| 594 | 534059 | Should a retail trader bother about reading SEC filings | 377322, 64881, 579110, 314898, 97837, 161411, 548596, 86281, 591558, 349181 |
| 10122 | 570787, 273718, 259084 | Why diversify stocks/investments? | 144261, 286227, 146181, 122222, 160170, 170147, 187124, 15169, 387723, 246253 |
| 3091 | 108302 | Am I considered in debt if I pay a mortgage? | 559370, 534099, 413976, 287571, 159936, 187739, 402659, 593373, 102155, 95778 |
| 4037 | 10275, 298099 | How separate individual expenses from family expenses in Gnucash? | 273071, 557186, 347137, 355415, 354716, 202224, 349674, 123384, 244555, 562777 |
| 8537 | 6574 | What is an “Options Account”? | 564353, 511093, 370494, 333784, 266840, 547460, 100628, 310207, 22263, 596567 |
| 1783 | 332314 | Freelancing Tax implication | 14609, 156063, 532259, 306800, 445298, 72519, 167660, 225666, 421924, 270426 |
| 6004 | 149555 | Put-Call parity - what is the difference between the two representations? | 374797, 314508, 13260, 354559, 238075, 336618, 345410, 365228, 247738, 343613 |
| 2923 | 82744, 390089, 264271 | Should I give to charity by check or credit card? | 46381, 407313, 123570, 161667, 521070, 250166, 481052, 290714, 106831, 302823 |
| 5271 | 58599, 221439, 67663 | Why are auto leases stubbornly strict about visa status and how to work around that? | 571295, 494439, 229354, 95189, 88229, 372771, 205088, 317781, 283494, 570095 |
| 2108 | 525200, 309171, 155389 | Can I pay taxes using bill pay from my on-line checking account? | 167202, 9970, 509075, 365851, 338701, 93523, 278678, 373180, 108734, 483265 |
| 4837 | 531841, 20958 | When applying for a mortgage, can it also cover outstanding debts? | 97925, 84267, 189881, 204035, 546190, 405695, 215232, 284403, 377477, 32749 |
| 4415 | 147646, 414188, 67676 | How much is inflation? | 206580, 108859, 290585, 117578, 519596, 75680, 41960, 29761, 513249, 456669 |
| 1915 | 433801 | Should I pay a company who failed to collect VAT from me over 6 months ago? | 141111, 544995, 380532, 182168, 393953, 306144, 156405, 378060, 452248, 463595 |
| 8378 | 125298 | Should I wait a few days to sell ESPP Stock? | 127702, 584291, 332243, 495568, 294573, 567244, 407602, 361345, 35575, 178684 |
| 2316 | 348955 | What exchange rate does El Al use when converting final payment amount to shekels? | 482343, 250281, 154417, 551451, 94315, 16051, 116930, 493576, 126565, 196653 |
| 8102 | 552707, 90294 | When do I sell a stock that I hold as a long-term position? | 171819, 306460, 165970, 181069, 226496, 84891, 501214, 240215, 218695, 35752 |
| 7928 | 118633, 501504, 499811 | If I believe a stock is going to fall, what options do I have to invest on this? | 67415, 410404, 271336, 427808, 595211, 493336, 340730, 206765, 42438, 125098 |
| 603 | 456440 | Will one’s education loan application be rejected if one doesn't have a payslip providing collateral? | 64103, 310056, 195498, 86952, 437149, 121731, 374331, 95869, 183237, 376403 |
| 4777 | 590710 | How to finance necessary repairs to our home in order to sell it? | 171631, 308802, 45773, 378384, 205217, 22360, 78252, 482112, 321361, 323449 |
| 5741 | 25943 | Learning investing and the stock market | 379546, 64168, 167088, 85558, 386803, 276960, 350497, 241433, 325212, 47973 |
| 620 | 331332, 417301, 180673, 487067 | Is it wise to have plenty of current accounts in different banks? | 535817, 511139, 174196, 201280, 588571, 456636, 189303, 483265, 509073, 142966 |
| 8271 | 415511 | Income in zero-interest environment | 376709, 260756, 508055, 538898, 137225, 175265, 211026, 90546, 503596, 72510 |
| 4865 | 100485 | Why are historical prices of stocks different on different websites? Which one should I believe? | 370569, 47738, 151610, 67415, 246733, 422476, 54118, 67365, 319606, 236321 |
| 864 | 211364, 152072 | Why use accounting software like Quickbooks instead of Excel spreadsheets? | 203446, 329774, 224438, 30142, 472924, 24890, 112728, 71569, 91671, 333616 |
| 5888 | 540806 | Interest charges on balance transfer when purchases are involved | 336792, 213242, 490529, 125497, 543776, 263647, 343961, 117602, 533589, 106448 |
| 2568 | 139047, 388798, 296769, 108739, 590082, 127353 | How to pay with cash when car shopping? | 258247, 166314, 278671, 260159, 196870, 60261, 453301, 288904, 68604, 2034 |
| 5150 | 229937 | What credit card information are offline US merchants allowed to collect for purposes other than the transaction? | 493638, 299840, 34108, 48404, 20261, 318132, 144580, 236194, 350655, 308889 |
| 7124 | 74615, 154725 | How come we can find stocks with a Price-to-Book ratio less than 1? | 487074, 558617, 533140, 433905, 251100, 278582, 317365, 414940, 226070, 583708 |
| 2857 | 233732, 501384 | I have around 60K $. Thinking about investing in Oil, how to proceed? | 516818, 474575, 251843, 542051, 6384, 340586, 379140, 137648, 18366, 127566 |
| 8539 | 218728, 292609, 196304, 396038 | Can the risk of investing in an asset be different for different investors? | 67415, 313493, 469599, 186575, 150475, 142320, 271691, 293679, 283074, 113786 |
| 6896 | 251704 | Selling high, pay capital gains, re-purchase later | 169240, 366560, 564129, 378594, 553253, 337993, 272028, 448659, 526661, 343219 |
| 1085 | 467737, 393710 | How do disputed debts work on credit reports? | 242013, 372039, 153220, 450031, 574122, 124452, 422484, 240038, 92977, 364361 |
| 5422 | 79517, 151973, 13513 | What are some good books for learning stocks, bonds, derivatives e.t.c for beginner with a math background? | 463713, 17186, 376692, 193555, 273906, 13573, 180380, 412327, 241423, 496429 |
| 8247 | 42521, 475170, 321114, 465313 | Tax on Stocks or ETF's | 437907, 260305, 244813, 270992, 120133, 243931, 473658, 54947, 528880, 299690 |
| 10482 | 549072 | Rollover into bond fund to do dollar cost averaging [duplicate] | 525089, 473042, 98704, 252677, 447567, 330023, 544070, 464264, 384893, 237052 |
| 8789 | 70853 | What does “profits to the shareholders jumped to 15 cents a share” mean? | 87349, 41912, 339854, 328341, 58290, 345400, 580512, 341293, 14870, 107857 |
| 8034 | 597351 | What is the average cost of a portfolio on a trading site? | 282376, 284075, 452736, 181013, 76330, 414335, 81865, 176335, 221627, 322725 |
| 2423 | 81106, 538023 | At what age should I start or stop saving money? | 529444, 234846, 396127, 417787, 519174, 326948, 235855, 357242, 337561, 471472 |
| 8475 | 293320, 481728 | Why I cannot find a “Pure Cash” option in 401k investments? | 494655, 332152, 104793, 262322, 525426, 538238, 21311, 104134, 410675, 168890 |
| 4523 | 594257, 119165, 393009, 129255, 66626 | What should I do with my $25k to invest as a 20 years old? | 252942, 589103, 521847, 390582, 420792, 296980, 492589, 272070, 363679, 482031 |
| 8513 | 270573 | Buy on dip when earnings fail? | 321205, 462135, 363178, 203873, 69790, 175821, 176384, 84334, 173026, 572622 |
| 9291 | 96926 | Are there any consequences for investing in Vanguard's Admiral Shares funds instead of ETF's in a Roth IRA? | 138383, 106611, 266457, 143238, 88823, 236006, 17208, 179737, 569342, 454610 |
| 8532 | 27401, 320101 | What do these options trading terms mean? | 560245, 224714, 473015, 120358, 2959, 42438, 245931, 344372, 399367, 278102 |
| 1824 | 244808 | Is there a way to open a U.S. bank account for my LLC remotely? | 195207, 452983, 43217, 364378, 43716, 131483, 154988, 88124, 20987, 75195 |
| 5054 | 28119 | How to stress test an investment plan? | 216915, 564007, 349475, 252017, 250844, 146629, 52438, 404881, 588481, 344596 |
| 1159 | 496064 | what is the best way to do a freelancing job over the summer for a student | 156063, 477205, 149415, 352320, 347181, 238015, 460648, 449155, 55064, 14609 |
| 68 | 19183 | Intentions of Deductible Amount for Small Business | 300510, 528564, 354716, 160612, 192516, 381151, 259227, 477940, 110577, 545296 |
| 7269 | 486696 | How do I track investment performance in Quicken across rollovers? | 152045, 401498, 37306, 563499, 473644, 343586, 239484, 111451, 534323, 558779 |
| 9701 | 387141, 357739 | How to bet against the London housing market? | 473883, 592612, 599389, 408865, 253359, 422609, 361884, 225682, 403212, 242321 |
| 6199 | 414693, 584273, 239214 | How can all these countries owe so much money?  Why & where did they borrow it from? | 392806, 558742, 36659, 124940, 167473, 141876, 96211, 355156, 336725, 243656 |
| 9126 | 514831 | Short an option - random assignment? | 334473, 581672, 477588, 166227, 228810, 362473, 292045, 575408, 494186, 244754 |
| 34 | 599545 | 401k Transfer After Business Closure | 231449, 494783, 232049, 262322, 536703, 383029, 367355, 122114, 161176, 134187 |
| 4011 | 136367, 470, 67699 | How can I deal with a spouse who compulsively spends? | 507630, 408327, 235520, 373946, 548626, 454661, 111531, 528077, 262400, 156835 |
| 6395 | 166227 | Option settlement for calendar spreads | 194730, 285605, 576976, 467463, 8177, 100021, 543312, 182272, 154989, 300139 |
| 9115 | 207325 | Why does the calculation for percentage profit vary based on whether a position is short vs. long? | 501214, 154665, 419897, 232880, 422467, 358586, 27236, 97402, 226496, 89602 |
| 4767 | 280805, 568670, 224057, 22804, 125986 | New car: buy with cash or 0% financing | 484979, 59372, 427884, 256803, 179960, 252859, 420018, 419851, 371720, 223502 |
| 2593 | 528132 | Am I “cheating the system” by opening up a tiny account with a credit union and then immediately applying for a huge... | 231614, 443134, 457135, 476923, 175522, 597291, 205196, 336518, 91471, 571295 |
| 11088 | 437100 | Am I required to have a lawyer create / oversee creation of my will? | 167828, 113674, 310250, 358640, 304023, 513642, 209846, 462942, 37183, 54007 |
| 8351 | 472516 | What happens when a calendar spread is assigned in a non-margin account? | 45674, 273142, 141213, 102316, 596567, 269043, 499060, 402726, 35102, 273612 |
| 6131 | 365263, 218088, 381720, 235452, 170204, 416679, 2460 | Is it ever a good idea to close credit cards? | 326094, 38938, 201982, 334111, 258465, 504293, 339030, 486376, 453147, 143596 |
| 3480 | 75568 | Why is OkPay not allowed in the United States? | 221015, 53468, 200457, 148078, 103552, 46791, 445084, 482332, 418352, 126615 |
| 858 | 146632, 45185, 122485, 278450 | Is it bad practice to invest in stocks that fluctuate by single points throughout the day? | 508764, 567608, 127689, 214281, 273978, 115719, 378994, 234893, 137255, 474296 |
| 4019 | 379948, 6881, 125477 | How and Should I Invest (As a college 18 year old with minimal living expenses)? | 429432, 140135, 81016, 287991, 357242, 61936, 401232, 556079, 553288, 188524 |
| 6080 | 164513 | Is ScholarShare a legitimate entity for a 529 plan in California? | 22856, 2809, 313392, 201500, 380424, 83080, 135128, 277581, 69841, 329812 |
| 2580 | 503934 | Stock market vs. baseball card trading analogy | 11988, 344118, 371720, 143125, 506078, 312801, 287656, 557137, 408918, 254910 |
| 7513 | 115372, 273861, 567362 | Where are Bogleheadian World ETFs or Index funds? | 408524, 391215, 41176, 437875, 99568, 364735, 159471, 507507, 575509, 450848 |
| 6959 | 208331, 205010 | What is the term for the quantity (high price minus low price) for a stock? | 229573, 517961, 301570, 120288, 574365, 583378, 445943, 506460, 536988, 333605 |
| 5402 | 491350, 227485 | Is it impossible to get a home loan with a poor credit history after a divorce? | 440063, 313896, 310790, 595029, 90579, 51728, 445163, 495431, 153088, 509739 |
| 4504 | 566458 | what is the best way of investment which gives returns forever? | 2981, 571579, 373554, 586756, 71926, 305633, 133486, 384893, 88417, 141365 |
| 6562 | 135675, 561344 | Cheapest way to “wire” money in an Australian bank account to a person in England, while I'm in Laos? | 185146, 282744, 183880, 308665, 60446, 184386, 462050, 263174, 323389, 473605 |
| 9565 | 292559 | What are the tax benefits of dividends vs selling stock | 478291, 227046, 352484, 444747, 199544, 63882, 291996, 206442, 314342, 167356 |
| 10734 | 589970, 263481, 470289 | How do you translate a per year salary into a part-time per hour job? | 375748, 426278, 517201, 220198, 175889, 225718, 394649, 355720, 120576, 62882 |
| 701 | 288537, 339488, 389446 | What are the ins/outs of writing-off part of one's rent for working at home? | 376446, 210117, 145313, 349672, 122569, 288564, 547302, 8464, 272987, 133076 |
| 585 | 140226, 552375 | Following an investment guru a good idea? | 426550, 230733, 483185, 113221, 222153, 506344, 52274, 531698, 79337, 277810 |
| 9598 | 379311 | How do index funds actually work? | 41176, 452122, 47985, 169627, 151741, 266194, 328691, 112701, 220486, 528034 |
| 4265 | 293111, 349710 | Does it make any sense to directly contribute to reducing the US national debt? | 157553, 463876, 315964, 52404, 411788, 392163, 103344, 110948, 149974, 496385 |
| 90 | 31793 | Filing personal with 1099s versus business s-corp? | 154931, 482165, 248629, 370542, 509111, 423074, 11401, 218257, 490489, 297965 |
| 8592 | 441718 | Tax implications of exercising ISOs and using proceeds to exercise more ISOs | 193717, 212025, 34269, 261258, 313372, 183166, 163566, 76556, 362874, 290468 |
| 5231 | 146188 | Where to find CSV or JSON data for publicly traded companies listed with their IPO date? | 510163, 526073, 71230, 365092, 592484, 199600, 200894, 182226, 420556, 481203 |
| 7674 | 193318, 519390, 494295 | Choosing the limit when making a limit order? | 447886, 251596, 200666, 167151, 164008, 486692, 15917, 448713, 591534, 289045 |
| 10137 | 57168 | F-1 student investing in foreign markets | 78632, 535340, 504612, 382295, 273861, 89351, 226568, 538940, 577628, 489113 |
| 5940 | 486243, 433827, 93936 | How does investment into a private company work? | 179664, 182226, 567492, 459051, 250354, 403615, 110200, 191756, 252853, 134995 |
| 6612 | 322900 | If I have a lot of debt and the housing market is rising, should I rent and slowly pay off my debt or buy and roll th... | 502594, 532153, 14083, 44360, 433171, 373554, 425963, 301192, 274870, 157728 |
| 945 | 352052 | Paid by an American company but working from France: where should I pay taxes? | 248167, 597154, 151763, 398472, 189238, 26182, 58406, 349424, 195571, 259881 |
| 2549 | 58451, 21103 | How to graph the market year over year? for example Dow Jones Index | 401939, 418150, 83391, 249273, 264850, 313421, 570247, 195089, 580313, 500695 |
| 4714 | 450819 | Personal finance app where I can mark transactions as “reviewed”? | 505057, 29812, 357023, 529790, 584450, 427229, 344473, 65957, 218793, 256693 |
| 588 | 570546 | Is there a reason to buy a 0% yield bond? | 203710, 367137, 431386, 395111, 237317, 45468, 394702, 192193, 582553, 480318 |
| 4233 | 31377 | Are personal finance / money management classes taught in high school, anywhere? | 187129, 590836, 223896, 356425, 275481, 233312, 16270, 145143, 273552, 545123 |
| 8456 | 510875, 486333 | What typically happens to unvested stock during an acquisition? | 555276, 483695, 257853, 575875, 243886, 34437, 416125, 469036, 83012, 104188 |
| 10213 | 380942, 270221, 545712 | Looking for good investment vehicle for seasonal work and savings | 343693, 191658, 502281, 246991, 84267, 92370, 424841, 112259, 138935, 130118 |
| 4306 | 221435, 61819 | How do currency markets work? What factors are behind why currencies go up or down? | 450694, 116921, 126431, 233718, 561535, 226102, 432283, 499875, 321941, 136073 |
| 5196 | 172128, 114829 | I might use a credit card convenience check. What should I consider? | 85517, 402543, 467581, 289483, 565745, 69938, 132636, 115712, 76149, 456098 |
| 3006 | 269851, 403137, 568473, 328300 | Strategies for putting away money for a child's future (college, etc.)? | 34929, 290441, 332749, 451189, 78648, 204227, 372900, 53003, 329269, 321281 |
| 3909 | 312248, 404356, 193459, 245616, 353028 | How to rescue my money from negative interest? | 227134, 42475, 328499, 472837, 572387, 30584, 438403, 252901, 587508, 391576 |
| 4116 | 234674 | Would the effects of an anticipated default by a nation be mostly symbolic? | 109149, 372474, 87436, 249851, 525793, 325374, 313306, 36659, 277230, 206491 |
| 4942 | 357685 | find stock composition of a publicly traded fund | 53993, 481203, 585139, 526073, 408524, 454596, 506618, 276983, 502091, 265447 |
| 5464 | 350399, 294549, 86691 | Resources on Buying Rental Properties | 26339, 87324, 465256, 423438, 3217, 92397, 207815, 433905, 38473, 296355 |
| 2385 | 407654 | As director, can I invoice my self-owned company? | 373059, 210889, 55440, 576384, 73421, 506766, 55108, 229110, 339658, 571265 |
| 3177 | 268289 | Vanguard ETF vs mutual fund | 138383, 539263, 367960, 172703, 143238, 161019, 367014, 135405, 536120, 364735 |
| 10034 | 480749 | Tax implications of holding EWU (or other such UK ETFs) as a US citizen? | 44955, 197478, 180146, 381884, 528880, 181942, 565296, 297051, 495417, 198532 |
| 5090 | 436493 | Should I take a student loan to pursue my undergraduate studies in France? | 246286, 217365, 287507, 589582, 92430, 354621, 541386, 6339, 547246, 67851 |
| 2088 | 399875 | How would I go about selling the stock of a privately held company? | 53993, 377563, 410404, 229838, 160430, 163387, 159703, 181582, 488207, 254591 |
| 9391 | 283202, 503637 | Should I replace bonds in a passive investment strategy | 535518, 107424, 545760, 225235, 273575, 144030, 108859, 248158, 494653, 472011 |
| 529 | 510701 | Sole proprietorship or LLC? | 66356, 151023, 504317, 334603, 213185, 297965, 34087, 243949, 246461, 552756 |
| 2118 | 411061 | What happened in Argentina in 2001 bank sector? did the banks closed? all or some? | 409562, 373726, 195315, 304007, 567891, 41271, 479351, 340436, 150543, 90958 |
| 3148 | 178127, 438000, 92888 | Can a car company refuse to give me a copy of my contract or balance details? | 584305, 172855, 521144, 374334, 5940, 159273, 443436, 38963, 343485, 248965 |
| 4955 | 581318 | How to calculate the value of a bond that is priced to yield X% | 152265, 290562, 198465, 327556, 453480, 316132, 19184, 416513, 211308, 142822 |
| 4678 | 305153 | Finance, Cash or Lease? | 215225, 522532, 307158, 504918, 185405, 376016, 487678, 311748, 311446, 179891 |
| 2398 | 118730, 509391, 363810, 224654, 590489 | Frustrated Landlord | 168922, 380618, 556041, 68320, 249195, 436875, 505900, 201705, 145016, 560195 |
| 3569 | 450135 | Funds in closed bank account have gone to the government | 19853, 225964, 67146, 151261, 87141, 397655, 68144, 80403, 11429, 567891 |
| 8947 | 455398, 461592 | Can a Roth IRA be used as a savings account? | 162592, 179408, 553031, 270818, 424841, 32671, 416240, 53996, 140330, 175968 |
| 6746 | 210887 | What happens if stock purchased on margin plummets below what I have in the brokerage? | 231221, 279185, 251704, 247680, 125659, 115918, 432111, 537153, 310326, 584291 |
| 5511 | 169893, 107898, 529123, 560325, 478426, 383193, 114303, 278699, 12746, 51873 | Pay off car loan entirely or leave $1 until the end of the loan period? | 179891, 376016, 479050, 352027, 10873, 329137, 488258, 16559, 543193, 516631 |
| 8834 | 133644, 12232, 197151, 569303 | Pros/Cons of Buying Discounted Company Stock | 268802, 471686, 417457, 56027, 315847, 192910, 301609, 324273, 171253, 488037 |
| 1157 | 272425 | Personal taxes for Shopify / Paypal shop? | 405777, 365192, 257311, 148288, 28974, 83346, 36783, 204870, 401677, 84596 |
| 988 | 226053, 107688 | Where should I invest my savings? | 315741, 501384, 137572, 377814, 480910, 79009, 480827, 569528, 100517, 197506 |
| 3369 | 231012, 145716, 411910, 395840 | Why should one only contribute up to the employer's match in a 401(k)? | 341493, 422979, 296405, 301616, 171196, 140917, 565684, 75766, 175470, 240373 |
| 9296 | 435746 | Why would Two ETFs tracking Identical Indexes Produce different Returns? | 148721, 524940, 429827, 99568, 285135, 145123, 55751, 261902, 144261, 327978 |
| 9245 | 194561 | Stock Options for a company bought out in cash and stock | 207253, 555276, 117177, 287092, 259560, 525390, 517873, 400644, 220147, 193398 |
| 3490 | 420529 | Tax Witholding for Stock Sale | 115333, 523360, 239233, 591157, 367742, 491028, 400730, 507276, 361482, 251704 |
| 5763 | 515361, 462019 | What is the best way to get a “rough” home appraisal prior to starting the refinance process? | 565868, 503444, 67379, 440063, 331255, 215647, 497927, 383272, 316878, 529186 |
| 26 | 285255, 350819 | Applying for and receiving business credit | 176284, 338406, 140116, 227910, 323389, 112793, 425352, 399013, 78230, 186504 |
| 4962 | 599925 | Net Cash Flows from Selling the Bond and Investing | 390810, 394702, 187110, 393838, 490077, 34949, 152265, 446214, 431386, 308276 |
| 4846 | 151104, 323749 | Is there anything comparable to/resembling CNN's Fear and Greed Index? | 98096, 3533, 498754, 114882, 559762, 317666, 558131, 228889, 78644, 24029 |
| 9403 | 6666, 328086, 345199 | Abundance of Cash - What should I do? | 105089, 598526, 486692, 158614, 155533, 158455, 243148, 181002, 257416, 404493 |
| 3503 | 345294 | Is there any instance where less leverage will get you a better return on a rental property? | 445887, 597679, 360621, 305065, 547196, 148299, 279488, 196295, 228657, 557324 |
| 929 | 367754 | Freelancer: Should I start a second bank account? | 30596, 576384, 179066, 148250, 545780, 467251, 396792, 144190, 368911, 234071 |
| 5993 | 367375, 287571, 224918, 431212, 272866, 230215, 55084, 5827, 352638, 63690, 426120, 63501 | Why would anyone want to pay off their debts in a way other than “highest interest” first? | 416796, 94373, 220241, 397538, 126965, 403450, 31189, 494306, 82741, 40522 |
| 8116 | 433032 | A-B-C Class Shares: What's the difference? | 483191, 147002, 550661, 9876, 454610, 98150, 591694, 251604, 554207, 74560 |
| 9633 | 585447 | Video recommendation for stock market education | 162884, 481166, 534418, 157509, 415121, 382631, 235197, 442830, 259081, 512895 |
| 9824 | 574777 | Where can end-of-day data be downloaded for corporate bonds? | 262291, 373170, 43060, 269550, 590233, 466255, 348886, 527001, 325370, 558924 |
| 7529 | 66607 | Does the expense ratio of a fund-of-funds include the expense ratios of its holdings? | 102904, 88823, 89297, 65587, 236507, 59249, 327925, 464337, 453542, 135405 |
| 5021 | 589285 | Is there a more flexible stock chart service, e.g. permitting choice of colours when comparing multiple stocks? | 310218, 500527, 189341, 151678, 226070, 66034, 174019, 418083, 237066, 513818 |
| 3612 | 402726, 259625 | How can I buy and sell the same stock on the same day? | 524238, 142599, 392403, 593644, 569912, 551108, 356490, 554568, 336018, 444668 |
| 4409 | 426676, 499128, 97925, 100306, 245903, 147439 | My friend wants to put my name down for a house he's buying. What risks would I be taking? | 102326, 306926, 518242, 539432, 283635, 11995, 148259, 56392, 115066, 53601 |
| 5369 | 171339, 44223 | Paying for things on credit and immediately paying them off: any help for credit rating? | 590370, 7928, 102326, 201447, 208909, 591995, 373497, 280131, 574065, 250722 |
| 2070 | 136438, 363678 | Advantage of credit union or local community bank over larger nationwide banks such as BOA, Chase, etc.? | 550303, 48866, 385343, 30253, 587737, 106981, 433933, 18749, 413078, 590209 |
| 11039 | 53544, 202768, 293531, 249063 | Pay off credit card debt or earn employer 401(k) match? | 175470, 595875, 124042, 301616, 341493, 422979, 371886, 171196, 519750, 463892 |
| 5460 | 21174, 108514, 463885 | Paying off a loan with a loan to get a better interest rate | 196237, 93248, 202527, 6339, 31189, 470716, 327700, 572272, 474184, 25190 |
| 7925 | 251100, 402482, 438974 | Can I sell a stock immediately? | 377719, 263751, 332467, 7561, 435715, 581866, 189874, 339419, 570112, 457294 |
| 3789 | 571131 | How to work around the Owner Occupancy Affidavit to buy another home in less than a year? | 274573, 492856, 129862, 286632, 327428, 119077, 296028, 482963, 578906, 57149 |
| 5228 | 232451 | How does the bank/IRS know whether a bank transfer over $14k is a gift or loan repayment? | 117661, 232322, 553205, 214934, 32455, 32880, 344398, 594595, 322424, 133548 |
| 2685 | 384532, 154113, 370300, 37900, 303293, 594182, 468923 | What ways are there for us to earn a little extra side money? | 576047, 109880, 89624, 194540, 446059, 229310, 252843, 597842, 541032, 549359 |
| 8702 | 345410 | Why is early exercise generally not recommended for an in-the-money option? | 135363, 6771, 157759, 194605, 73256, 529958, 193303, 44530, 193717, 36453 |
| 10645 | 588607 | Explain the details and benefits of rebalancing a retirement portfolio? | 28425, 355738, 516607, 36284, 22221, 569849, 217425, 340131, 421769, 85276 |
| 1090 | 518896, 203091 | Need a formula to determine monthly payments received at time t if I'm reinvesting my returns | 16051, 393987, 7540, 166394, 424079, 545902, 584231, 340254, 402875, 278638 |
| 6122 | 169824, 44344 | Better to rent condo to daughter or put her on title? | 496166, 566184, 203105, 316794, 258118, 294167, 182039, 590517, 515690, 35002 |
| 3682 | 356161, 329662 | Short selling - lender's motivation | 75686, 188531, 35500, 300391, 480967, 110046, 226496, 286535, 549040, 278821 |
| 4514 | 69485, 337764, 209804 | What intrinsic, non-monetary value does gold have as a commodity? | 240894, 156211, 1642, 146573, 471825, 512734, 156240, 408336, 520769, 474575 |
| 8507 | 509819, 370995 | When to sell a stock? | 554997, 343219, 99132, 365331, 53047, 554568, 503912, 424641, 273565, 320101 |
| 6221 | 470716, 257248, 519675, 76414, 169688, 455614, 115717 | To pay off a student loan, should I save up a lump sum payoff payment or pay extra each month? | 254245, 188713, 399863, 6339, 479050, 121505, 448791, 110503, 529551, 352363 |
| 1819 | 376499, 66058, 212713, 267362, 250285, 220691 | Found an old un-cashed paycheck. How long is it good for? What to do if it's expired? | 522332, 479781, 528661, 476233, 199069, 445739, 88223, 547773, 195532, 446070 |
| 3008 | 180192, 323406 | What are my chances at getting a mortgage with Terrible credit but High income | 287571, 2064, 312211, 233417, 407017, 298908, 389750, 212299, 317461, 285694 |
| 4007 | 521657 | What is a reasonable salary for the owner and sole member of a small S-Corp? | 334603, 556220, 170933, 325677, 352838, 370542, 260385, 205341, 388704, 454035 |
| 6644 | 175035 | How to know precisely when a SWIFT is issued by a bank? | 218761, 554518, 91534, 589616, 118396, 298587, 350396, 446807, 475527, 97988 |
| 6713 | 357571 | Will my father still be eligible for SNAP if I claim him as my dependent? | 147853, 481114, 692, 106735, 15319, 43647, 285493, 148204, 547087, 470066 |
| 7463 | 582005, 105634, 305287, 5152 | Pros/cons of borrowing money using a mortgage loan and investing it in a low-fee index fund? | 192910, 191741, 80272, 66453, 488037, 512827, 517391, 161445, 28060, 471686 |
| 10267 | 17652, 424511, 328556, 460398 | How should I prepare for the next financial crisis? | 178693, 36961, 518213, 415143, 60166, 551187, 125847, 143393, 588948, 204473 |
| 7622 | 253369, 378594 | Best way to pay off debt? | 130812, 373554, 345895, 313711, 480773, 329249, 213159, 457945, 349567, 402659 |
| 10979 | 148728, 164001 | Closing a futures position | 533408, 422922, 362762, 298833, 503505, 394886, 518214, 497928, 414636, 357324 |
| 3767 | 368679, 153922, 392060, 520395, 320246 | What should I be doing to protect myself from identity theft? | 423809, 91986, 90632, 252475, 470024, 598801, 260580, 97686, 158285, 551747 |
| 7145 | 17923, 116865 | Are there “buy and hold” passively managed funds? | 301580, 205280, 126151, 311192, 379311, 138383, 549364, 59736, 93882, 522007 |
| 6554 | 22469 | Mutual fund value went down, shares went up, no action taken by me | 88947, 184077, 183926, 287537, 583203, 1034, 212394, 179737, 454224, 162916 |
| 6410 | 471723 | Will an ETF immediately reflect a reconstitution of underlying index | 330729, 454610, 124900, 87238, 71230, 93836, 313897, 295993, 126146, 14440 |
| 5030 | 215540 | Why pay for end-of-day historical prices? | 330276, 113150, 532178, 352415, 343803, 524123, 391171, 105607, 371129, 110716 |
| 6252 | 394551, 160932, 293624, 233294, 243268, 379487, 62868 | Is this mortgage advice good, or is it hooey? | 148087, 423403, 495791, 473647, 205906, 77153, 103830, 139366, 138113, 213713 |
| 885 | 337165, 409184 | How long do credit cards keep working after you disappear? | 489501, 99449, 472336, 143596, 220241, 339030, 268895, 261697, 138954, 412542 |
| 4031 | 115741 | 28 years old and just inherited large amount of money and real estate - unsure what to do with it | 344186, 578597, 201797, 523949, 149178, 319307, 80797, 88398, 9243, 276381 |
| 766 | 550172 | Will the ex-homeowner still owe money after a foreclosure? | 163711, 350588, 104857, 62908, 92659, 212827, 409266, 2996, 276538, 559436 |
| 2183 | 124427, 571625, 24994, 24344 | Why are there many small banks and more banks in the U.S.? | 209493, 397679, 573518, 132678, 172567, 52441, 174523, 568090, 144304, 88196 |
| 8079 | 57138 | Growth rate plus dividend yieid total? | 350607, 95278, 125940, 187583, 248853, 11184, 553817, 136436, 112342, 354857 |
| 8202 | 513258, 93971 | What accounted for DXJR's huge drop in stock price? | 231295, 92267, 115553, 474670, 457689, 294095, 537862, 317363, 432511, 71924 |
| 7345 | 237645 | What do these numbers mean? (futures) | 87548, 553105, 222852, 285648, 189061, 127845, 273789, 354429, 527080, 444461 |
| 7441 | 514500, 117576 | Since many brokers disallow investors from shorting sub-$5 stocks, why don't all companies split their stock until it... | 20085, 80024, 359734, 125659, 37880, 404339, 343977, 97064, 319078, 537418 |
| 776 | 467044, 583640, 127263, 332373, 220127, 591516 | Can saving/investing 15% of your income starting age 25, likely make you a millionaire? | 124027, 425387, 563284, 417787, 3059, 374266, 41960, 231688, 434972, 143591 |
| 89 | 413229, 590102, 268026, 248624 | How can I deposit a check made out to my business into my personal account? | 526817, 309023, 358837, 188167, 521540, 98636, 400230, 308938, 346042, 135196 |
| 9329 | 523913 | Interactive Brokers: IOPTS and list of structured products | 326991, 168642, 15772, 234851, 262591, 334383, 339875, 89351, 478736, 562113 |
| 2516 | 199508, 566602 | Which banks have cash-deposit machines in Germany? | 505678, 491855, 362374, 229246, 535947, 82812, 296674, 320672, 444548, 557393 |
| 6629 | 444405 | Tax treatment of a boxed trade? | 412404, 502874, 358090, 219762, 474745, 294148, 115333, 25081, 389098, 398442 |
| 1889 | 388713 | Reporting financial gains from my online store | 92593, 226997, 379661, 196520, 202145, 228054, 9731, 516371, 255281, 125935 |
| 1920 | 269943 | Clarification on student expenses - To file the tax for the next year | 500357, 361974, 552810, 585356, 411063, 592192, 569628, 527776, 295250, 254514 |
| 8013 | 496159, 224231 | Frequency of investments to maximise returns (and minimise fees) | 562046, 385955, 388389, 583918, 338943, 478266, 381268, 220834, 56742, 81652 |
| 3759 | 527966, 67167, 522358 | Simplifying money management | 373772, 384631, 478914, 226691, 153033, 105949, 580445, 455457, 360356, 4783 |
| 7295 | 244749 | Selling non-dividend for dividend stocks | 270595, 352484, 509879, 234893, 438349, 328182, 93971, 39478, 124368, 407551 |
| 10639 | 431799, 495774, 278453, 187039 | Short term parking of a large inheritance? | 289326, 178386, 131391, 163353, 474007, 19551, 590276, 235628, 171196, 176061 |
| 6635 | 102449 | Why don't share prices of a company rise every other Friday when the company buys shares for its own employees? | 587137, 95806, 545036, 32040, 117082, 245654, 533712, 207253, 512914, 534478 |
| 4312 | 135845, 399149 | Is it true that 90% of investors lose their money? | 222639, 282435, 167950, 285945, 497786, 170628, 507284, 204690, 300770, 324779 |
| 1230 | 191649 | How does Walmart account their expired food | 572490, 510610, 121023, 344170, 170583, 156264, 213181, 307605, 51486, 19383 |
| 6121 | 394460 | What are my best options if I don't have a lot of credit lines for housing loans? | 143593, 452121, 336908, 408518, 577542, 68431, 67320, 438740, 336659, 251062 |
| 859 | 449630, 18749 | Any reason to keep around my account with my old, 'big' bank? | 58614, 301361, 513651, 64556, 185403, 84709, 11429, 412036, 219673, 395929 |
| 6525 | 181985 | Does it make sense to trade my GOOGL shares for GOOG and pocket the difference? | 550661, 98150, 156467, 106541, 488920, 147002, 197480, 144033, 362473, 9082 |
| 2590 | 589625 | Are non-residents or foreigners permitted to buy or own shares of UK companies? | 365907, 528880, 73457, 465292, 540394, 574055, 262485, 205556, 272944, 546190 |
| 5374 | 152688 | What were the main causes of the spike and drop of DRYS's stock price? | 283106, 133204, 73857, 92267, 501276, 115553, 543606, 474670, 457689, 520769 |
| 8005 | 48800 | Difference between Vanguard sp500 UCITS and Vanguard sp500 | 494236, 106314, 27930, 497174, 241101, 370193, 42620, 537945, 267119, 593378 |
| 2994 | 419319, 318491 | Work on the side for my wife's company | 403563, 148346, 113098, 569145, 129428, 253755, 200211, 564841, 399804, 68302 |
| 3125 | 89008 | Claiming mileage allowances, what are the rules/guidelines? | 227197, 577200, 409157, 390435, 34810, 12140, 594414, 379189, 451622, 338545 |
| 3683 | 185909, 454501, 565016 | Can I trust the Motley Fool? | 428848, 408995, 301739, 276975, 534080, 500338, 526015, 192912, 6607, 104924 |
| 7206 | 441155, 532211, 553066 | Who Bought A Large Number Of Shares? | 436530, 372153, 260085, 34882, 12885, 246586, 192900, 173836, 33357, 170700 |
| 10246 | 512984, 77573 | Understanding the T + 3 settlement days rule | 179520, 266767, 156143, 116420, 36193, 28314, 296475, 593445, 89506, 370635 |
| 5241 | 376123, 234286 | Mortgage vs. Cash for U.S. home buy now | 4739, 390976, 344740, 444369, 482963, 144304, 594051, 385139, 78176, 433633 |
| 8121 | 432424 | Can I calculate stock value with Williams%R if I know the last set? | 260085, 402889, 220327, 395128, 132171, 138830, 393496, 219888, 204761, 408465 |
| 98 | 575929, 527522 | How can I make $250,000.00 from trading/investing/business within 5 years? | 352363, 221795, 385484, 104924, 465380, 186313, 379387, 399367, 158343, 363043 |
| 4615 | 262934 | Are solar cell panels and wind mills worth the money? | 464715, 249191, 158216, 425595, 271015, 69523, 455798, 227013, 210811, 262196 |
| 503 | 367641 | Privacy preferences on creditworthiness data | 324686, 307595, 166921, 42956, 596046, 87879, 599376, 338406, 171761, 400103 |
| 3264 | 486525, 260383 | Pros and Cons of Interest Only Loans | 268802, 471686, 315847, 301609, 324273, 171253, 488037, 180311, 28942, 94710 |
| 6467 | 453256, 66834, 23217, 346641, 367313 | Advice on strategy for when to sell | 498075, 195222, 217837, 105973, 343977, 302315, 504235, 88813, 122404, 109455 |
| 4289 | 288330, 24881 | Does the currency exchange rate contain any additional information at all? | 135220, 17469, 442503, 439779, 147551, 100403, 128048, 342411, 22953, 484891 |
| 4394 | 336045, 441582 | Transfer $50k to another person's account (in California, USA) | 322838, 293653, 455005, 355796, 521753, 462585, 60020, 542828, 494783, 305907 |
| 8982 | 454610, 200360 | Are Exchange-Traded Funds (ETFs) less safe than regular mutual funds? | 129070, 351088, 370244, 427842, 454224, 529638, 388585, 4666, 522759, 276983 |
| 2407 | 294327, 319734 | How long to wait after getting a mortgage to increase my credit limit? | 448368, 142536, 319090, 34432, 2064, 119931, 286843, 1766, 396933, 264341 |
| 8937 | 469888 | “In-the-Money” vs “Out-of-the-Money” Call Options | 49766, 453582, 502164, 530446, 176883, 271109, 115973, 324564, 203040, 387801 |
| 7344 | 108403 | How is the Dow divisor calculated? | 14368, 150430, 378974, 159166, 195089, 313421, 253926, 51884, 418150, 144005 |
| 2856 | 231727, 213331, 110848, 342212 | How can I cash out a check internationally? | 186606, 510485, 338319, 40338, 94957, 297013, 18727, 103680, 84645, 66991 |
| 6891 | 463837 | What is the theory behind Rick Van Ness's risk calculation in the video about diversification? | 474351, 478509, 296567, 370777, 599436, 425452, 460373, 365189, 205070, 292219 |
| 10447 | 152096, 300721 | Is there an advantage to a traditional but non-deductable IRA over a taxable account? [duplicate] | 447482, 61022, 500175, 144751, 9845, 542166, 382236, 406239, 162592, 123027 |
| 5427 | 200603, 323284 | How do auto-loan payments factor into taxes for cars that are solely used by dependent(s)? | 179891, 59372, 396789, 346108, 287857, 549464, 374704, 18551, 453263, 79623 |
| 5782 | 379891, 319773, 448614, 595455 | Pay off credit cards in one lump sum, or spread over a few months? | 487621, 172084, 511240, 254245, 4734, 395590, 449543, 178989, 252677, 194030 |
| 9871 | 448890, 40051, 170594 | What should I do with the 50k I have sitting in a European bank? | 103795, 354553, 455856, 479861, 542783, 73741, 124479, 66431, 294297, 494233 |
| 6668 | 275902 | Approximate IT company valuation (to proximate stock options value) | 77124, 457294, 139387, 313372, 50798, 249320, 32198, 328794, 535605, 335543 |
| 3404 | 556976, 160301, 498834, 395483, 277583 | In US, is it a good idea to hire a tax consultant for doing taxes? | 197870, 124480, 277812, 37725, 95298, 244233, 190653, 406707, 219188, 322246 |
| 3625 | 414295 | What should I do with my paper financial documents? | 245967, 509617, 278656, 576564, 72953, 287999, 441038, 229157, 122185, 212367 |
| 6005 | 135415, 478457, 345895, 73310, 384626, 390689, 270856 | Why might it be advisable to keep student debt vs. paying it off quickly? | 149500, 40160, 555280, 103093, 25190, 507544, 37070, 431884, 67091, 586626 |
| 8544 | 267113 | Strategies to recover from a bad short-term call options purchase where the underlying dropped instead? | 281533, 507828, 214003, 502164, 238474, 37517, 171819, 528052, 242663, 578022 |
| 3822 | 305907, 385090, 418900, 308837 | How to change a large quantity of U.S. dollars into Euros? | 19618, 79777, 194730, 241661, 549787, 340777, 196596, 531953, 390524, 445943 |
| 7879 | 372551, 421285 | Any Tips on How to Get the Highest Returns Within 4 Months by Investing in Stocks? | 593879, 344783, 228488, 45029, 102029, 578028, 540919, 120297, 282483, 243837 |
| 3115 | 234950, 389028, 316794 | How can I live outside of the rat race of American life with 300k? | 233562, 183869, 80504, 252852, 599558, 327026, 91265, 136035, 551325, 212471 |
| 3995 | 427032, 278734, 297900, 230208 | I have more than $250,000 in a US Bank account… mistake? | 352883, 200690, 583803, 146557, 301745, 400845, 190539, 158, 249960, 158343 |
| 559 | 246459 | Challenged an apparently bogus credit card charge, what happens now? | 299840, 81416, 517440, 456098, 140500, 401254, 8570, 535015, 245481, 533933 |
| 3791 | 327432, 577201, 212222 | "When people say 'Interest rates are at all time low!"" … Which interest rate are they actually referring to?" |  |
| 8002 | 265159, 34767 | What is the tax treatment of scrip dividends in the UK? | 32600, 217006, 115333, 110983, 216694, 118786, 97842, 295082, 500744, 145999 |
| 853 | 260795 | What will my taxes be as self employed? | 56558, 252843, 585121, 483453, 477476, 3173, 510409, 446117, 111354, 15270 |
| 10136 | 526115, 290930 | How to minimise the risk of a reduction in purchase power in case of Brexit for money held in a bank account? | 51337, 417740, 448769, 466950, 229354, 191303, 252835, 295893, 137572, 537371 |
| 8635 | 67107, 240215 | Is there any flaw in this investment scheme? | 144200, 330694, 510565, 382754, 527522, 46818, 440305, 53544, 130881, 157509 |
| 5206 | 563030, 287157, 28230, 117276, 300660 | Is it a good idea to get an unsecured loan to pay off a credit card that won't lower a high rate? | 516397, 287571, 414534, 143593, 223166, 19479, 508510, 225522, 193641, 559523 |
| 4433 | 146181 | When should you use an actively managed mutual fund in a 401k? | 563728, 41176, 311527, 367960, 238360, 90858, 135596, 93882, 199493, 205280 |
| 4188 | 468108 | Why is the stock market rising after Trump's attack on the TPP? | 382034, 107891, 253574, 253030, 96228, 472269, 78644, 220533, 581464, 392979 |
| 2713 | 388147 | Physical Checks - Mailing | 41944, 20791, 502781, 48866, 29372, 467509, 268257, 125371, 513980, 402898 |
| 9060 | 40447 | Buying puts without owning underlying | 181924, 338782, 516790, 294688, 480879, 369031, 511093, 401447, 121334, 230666 |
| 6683 | 592187, 77631 | Who are the sellers for the new public stocks? | 301985, 103536, 335631, 222035, 373726, 212685, 201731, 543589, 19196, 117576 |
| 3829 | 523850, 291438 | Are all VISA cards connected with bank accounts? | 58511, 249960, 542783, 280699, 308383, 249505, 443487, 491793, 572796, 463829 |
| 5585 | 300117 | Is there any site you can find out about the 'bonus features' of credit cards? | 334111, 219303, 443487, 44802, 477226, 148423, 13656, 431435, 120954, 560004 |
| 4105 | 166412, 25096 | As an investor what are side effects of Quantitative Easing in US and in EU? | 117902, 239214, 197228, 393791, 345910, 305029, 339640, 282153, 558989, 416483 |
| 1074 | 443960 | How common is “pass-through” health insurance? | 544869, 145464, 187085, 573394, 46893, 500238, 17215, 207437, 298014, 157990 |
| 2465 | 570680, 81046, 546509 | Can capital expenses for volunteer purposes be deducted from income? | 275543, 490176, 598646, 183612, 360629, 421924, 202645, 541809, 124505, 146657 |
| 4640 | 322314, 101369, 540539 | What can my relatives do to minimize their out of pocket expenses on their fathers estate | 464269, 418328, 367404, 117960, 148228, 192781, 372808, 96725, 243822, 267859 |
| 9275 | 338754, 14364 | Do I have to pay a capital gains tax if I rebuy the same stock within 30 days? | 23217, 537916, 407602, 318321, 526661, 343219, 14732, 295082, 250500, 487256 |
| 10497 | 304284, 34913, 398622, 575729, 71898, 196423 | Why would you elect to apply a refund to next year's tax bill? | 159880, 46902, 445230, 324994, 147765, 246453, 545873, 80994, 187695, 468047 |
| 3500 | 141935, 174019, 565409 | Why invest in becoming a landlord? | 11601, 41356, 572061, 587514, 517218, 480773, 249195, 64984, 436875, 71424 |
| 4205 | 134239 | How and why does the exchange rate of a currency change almost everyday? | 227897, 442503, 163303, 133540, 310748, 353028, 96791, 42951, 135220, 17469 |
| 1306 | 484437, 204075, 167684 | I made an investment with a company that contacted me, was it safe? | 87632, 594206, 471138, 366010, 105671, 327729, 177442, 381259, 110200, 527894 |
| 6262 | 26799, 390877 | Help required on estimating SSA benefit amounts | 118707, 20409, 529927, 11998, 420511, 430407, 34538, 410395, 424981, 412226 |
| 8632 | 213976 | Is it best to exercise options shares when they vest, or wait | 200784, 163396, 237718, 313372, 43497, 237783, 100628, 220147, 76556, 595605 |
| 1948 | 467509 | Which colors can one use to fill out a check in the US? | 90507, 219935, 401294, 95251, 456773, 310992, 594437, 118694, 2456, 552216 |
| 6133 | 415705 | What happens to all of the options when they expire? | 177559, 7733, 581672, 575408, 116436, 199966, 293605, 566111, 358492, 480879 |
| 3771 | 521712, 128471, 488948, 198349, 217683, 49601 | Best way to buy Japanese yen for travel? | 490384, 217715, 24912, 548340, 96211, 274683, 490113, 263364, 59994, 223359 |
| 1736 | 396933, 25543, 443419 | How can people have such high credit card debts? | 391085, 390598, 569056, 315275, 421379, 89622, 264029, 81206, 76248, 277855 |
| 9188 | 265167 | Selling mutual fund and buying equivalent ETF: Can I 1031 exchange? | 580802, 153660, 539263, 472663, 29184, 200360, 500486, 161019, 393090, 367960 |
| 2051 | 558042 | Where to find the 5 or 10 year returns for a mutual fund? | 290757, 496820, 447552, 93882, 346762, 270992, 112223, 549364, 409603, 427300 |
| 6814 | 340214, 223206 | Selling Stock - All or Nothing? | 5948, 154976, 276883, 198583, 369166, 66834, 459494, 91831, 137073, 61518 |
| 6146 | 7403 | Lost credit card replaced with new card and new numbers. Credit score affected? | 116243, 160125, 89888, 478461, 408763, 229875, 271472, 188406, 446909, 584419 |
| 4756 | 340254 | What is the formula for the Tesla Finance calculation? | 40414, 568220, 312893, 205070, 204917, 43964, 472646, 277217, 445348, 149555 |
| 2076 | 184646, 278824 | Can vet / veterinary bills be considered deductions (tax-deductible) for Income Tax purposes [Canada]? | 328073, 527776, 550345, 518624, 541809, 193251, 212783, 71232, 7423, 507408 |
| 1322 | 399418, 114231, 115552, 64138 | Is this follow-up after a car crash a potential scam? | 166313, 465648, 160568, 288936, 252517, 112095, 332916, 285077, 455565, 12350 |
| 5172 | 529418 | does interest payment on loan stay the same if I pay early | 210673, 188672, 6339, 156195, 534493, 390642, 64752, 205542, 38786, 82809 |
| 5185 | 210236, 317354 | Invest in low cost small cap index funds when saving towards retirement? | 196992, 255902, 241202, 272458, 376485, 517391, 247051, 262180, 483268, 541054 |
| 6909 | 127012 | Why do stocks priced above $2.00 on the ASX sometimes move in $0.005 increments? | 358164, 139089, 452175, 118232, 434925, 310837, 436128, 241059, 559157, 117576 |
| 2348 | 211867, 566573, 410166, 211622, 474234, 352271, 543714, 146441, 265874, 134864 | Why can't you just have someone invest for you and split the profits (and losses) with him? | 447619, 186631, 306430, 195398, 14255, 151412, 83698, 573313, 120649, 67415 |
| 687 | 146021, 268992 | Online tool to connect to my bank account and tell me what I spend in different categories? | 273308, 173348, 506831, 258423, 478807, 291278, 107398, 447597, 584175, 328835 |
| 4499 | 76996 | Is investing exlusively in a small-cap index fund a wise investment? | 241202, 68773, 78837, 90858, 520963, 46099, 161445, 323363, 206841, 479420 |
| 42 | 272709, 327263, 331981 | What are the ins/outs of writing equipment purchases off as business expenses in a home based business? | 326261, 288564, 210117, 376446, 28764, 88967, 145313, 374443, 122569, 42831 |
| 3530 | 189190, 239998 | How to exclude stock from mutual fund | 424192, 226967, 569948, 346474, 161445, 472663, 467575, 370244, 539263, 238360 |
| 659 | 439467, 264297, 168796, 120279, 365240, 449079 | Buying from an aggressive salesperson | 125868, 13139, 235925, 230908, 103753, 235772, 250644, 186206, 429899, 232083 |
| 9108 | 272021, 472585 | Starting an investment portfolio with Rs 5,000/- | 183074, 563986, 69012, 138746, 272513, 379900, 199397, 450586, 305579, 46967 |
| 6835 | 102243 | Are bond ETF capital gains taxed similar to stock or stock funds if held for more than 1 year? | 110343, 246221, 581632, 314342, 23217, 159703, 537916, 169240, 153112, 437907 |
| 6803 | 142726 | What are the common moving averages used in a “Golden Cross” stock evaluation? | 35006, 227669, 489933, 565501, 28271, 544053, 488285, 43526, 257185, 42620 |
| 9381 | 384983 | Trade? Buy and hold? Or both? | 75658, 313706, 112659, 7391, 597351, 295277, 8857, 468388, 408123, 110716 |
| 3067 | 406156, 517299 | Should I make extra payments to my under water mortgage or increase my savings? | 131365, 476068, 83543, 423403, 448791, 11791, 160932, 477907, 465801, 294167 |
| 4125 | 344648, 72046 | Alternative means of salary for my employees | 312797, 498284, 345795, 52185, 476980, 312369, 75642, 70357, 529344, 515365 |
| 1150 | 531698, 43603, 353369, 19936 | How are the best way to make and save money at 22 years old | 519346, 10476, 305946, 41732, 580711, 288073, 151774, 458475, 509266, 327240 |
| 7705 | 195191 | Why would I pick a specific ETF over an equivalent Mutual Fund? | 580802, 500486, 161019, 153112, 253971, 153660, 270992, 539263, 183898, 472663 |
| 10039 | 67785 | Do individual investors use Google to obtain stock quotes? | 567531, 587959, 415121, 574777, 187124, 414088, 237323, 201771, 131421, 53263 |
| 9808 | 40702, 431946 | Selling To Close | 229573, 27401, 322891, 521529, 91831, 557582, 355580, 374204, 594614, 292338 |
| 3888 | 307083, 319213, 239632 | Why I can't view my debit card pre-authorized amounts? | 432077, 208169, 316652, 418580, 185434, 533933, 276186, 140809, 292490, 223951 |
| 10109 | 506374, 156029, 406974, 499849, 566591 | Why does Charles Schwab have a Mandatory Settlement Period after selling stocks? | 28314, 113644, 570248, 93231, 266725, 511760, 108671, 563826, 296475, 476721 |
| 957 | 321500 | How can I withdraw money from my LLC? | 144563, 540334, 264075, 237760, 207997, 1066, 88124, 591318, 101748, 151023 |
| 2790 | 423403, 279329, 100483, 469125, 4612 | Should I pay more than 20% down on a home? | 357200, 531750, 480225, 472484, 117509, 56794, 234890, 517356, 321877, 371560 |
| 4999 | 314898, 338803 | Looking for a good source for Financial Statements | 9938, 270269, 82479, 171964, 546115, 269656, 153605, 248349, 24311, 597241 |
| 3189 | 225395 | Diversify my retirement investments with a Roth IRA | 568322, 262322, 482544, 94496, 162592, 561636, 399543, 406239, 471204, 585422 |
| 5134 | 158523 | Why does Yahoo Finance's data for a Vanguard fund's dividend per share not match the info from Vanguard? | 532616, 206727, 286335, 497344, 70072, 408103, 564876, 54225, 477637, 219477 |
| 1321 | 216456, 292065 | Are social media accounts (e.g. YouTube, Twitter, Instagram, etc.) considered assets? | 229640, 203715, 521823, 392379, 265765, 381322, 225389, 429254, 426184, 479130 |
| 6890 | 558703, 240196 | Where does the money go when I buy stocks? | 335631, 166885, 336532, 235772, 107751, 424352, 336722, 461018, 592636, 239714 |
| 3049 | 127974 | How to calculate my estimated taxes. 1099 MISC + Self Employment | 279538, 174025, 408434, 434351, 277812, 477476, 254151, 446117, 406042, 481459 |
| 849 | 557186 | Accounting for reimbursements that exceed actual expenses | 89297, 217748, 400291, 256542, 62869, 18889, 192843, 544947, 211810, 362060 |
| 4539 | 275925 | How should I save money if the real interest rate (after inflation) is negative? | 42475, 32744, 472837, 203926, 61792, 257547, 322816, 583678, 368338, 490648 |
| 715 | 579763, 546538, 187404 | what would you do with $100K saving? | 93936, 54574, 174336, 451307, 3104, 537721, 287501, 573199, 337561, 230215 |
| 2416 | 162612 | Why should a company go public? | 502332, 420476, 171236, 349684, 521566, 467594, 390529, 458071, 471247, 117177 |
| 7509 | 178303 | Investment Portfolio Setup for beginner | 576890, 195156, 282947, 467926, 413856, 406711, 139368, 375671, 495473, 492423 |
| 2204 | 280056, 174363, 50809, 83922 | What's an economic explanation for why greeting cards are so expensive? | 424523, 374030, 4066, 382347, 59758, 278430, 213927, 219313, 334097, 46625 |
| 2334 | 150650 | How do you determine “excess cash” for Enterprise Value calculations from a balance sheet? | 346760, 228403, 192515, 571234, 366249, 105949, 301998, 446652, 473963, 594011 |
| 2903 | 527776 | How should I file my taxes as a contractor? | 305791, 586026, 139501, 223170, 489898, 593694, 25762, 352640, 341220, 59686 |
| 5534 | 423272, 421136 | How does “taking over payments” work? | 297535, 67716, 293501, 316852, 202768, 455044, 351698, 151758, 446932, 4877 |
| 9668 | 42438 | Do stock option prices predicate the underlying stock's movement? | 111768, 511861, 350748, 393496, 181924, 13260, 415281, 196001, 236504, 320787 |
| 1530 | 313361, 219425 | What is the proper way to report additional income for taxes (specifically, Android development)? | 17284, 346374, 261280, 356237, 475596, 28764, 236517, 516548, 541315, 533929 |
| 504 | 500755, 344203, 498751 | Have plenty of cash flow but bad credit | 77248, 78754, 475405, 559718, 376246, 315168, 175691, 440609, 374510, 339648 |
| 2296 | 106424, 83330, 366594, 253563, 130850 | How does a bank make money on an interest free secured loan? | 32880, 395769, 223166, 198007, 28375, 67728, 396853, 354883, 272150, 307767 |
| 3186 | 545421 | United States Treasury Not Endorsing Checks | 555486, 231986, 433817, 411742, 588999, 244115, 316219, 211290, 169921, 200457 |
| 10975 | 61022 | How to contribute to Roth IRA when income is at the maximum limit & you have employer-sponsored 401k plans? | 163865, 140330, 164628, 189989, 360533, 441632, 81148, 447482, 110114, 32009 |
| 3859 | 230261 | Buying an investment property in Australia - what are the advantages and disadvantages of building a house vs buying... | 502242, 96580, 79378, 300229, 4739, 447498, 74393, 427166, 123929, 511515 |
| 1994 | 565157, 51491, 156640 | Does the IRS reprieve those who have to commute for work? | 457615, 231990, 4247, 544381, 486382, 434846, 165494, 300505, 513254, 340551 |
| 6647 | 69790, 428017 | What is meant by “priced in”? | 387515, 214000, 588540, 209863, 349475, 40131, 59652, 462265, 13260, 534734 |
| 9164 | 365298, 263390 | Bonds vs equities: crash theory | 309326, 115648, 296516, 427707, 321941, 418528, 52149, 464824, 287656, 149900 |
| 1812 | 530570 | splitting a joint mortgage - one owner in home | 58652, 159590, 373017, 142966, 395520, 492856, 219274, 138113, 473957, 461042 |
| 10462 | 8266, 11378, 35680, 437879, 204035, 581204 | Is it okay to be married, 30 years old and have no retirement? | 458795, 268023, 198045, 438326, 237923, 151774, 135646, 113660, 467044, 74603 |
| 8855 | 208165, 474296, 312821 | How do i get into investing stocks [duplicate] | 201391, 7814, 322053, 578314, 330694, 545760, 156253, 93397, 98345, 282483 |
| 7594 | 573899 | Converting annual interbank rates into monthly rates | 37954, 17469, 311834, 109454, 349611, 383849, 473477, 597459, 561968, 18939 |
| 7071 | 124230 | ESPP strategy - Sell right away or hold? | 127702, 294573, 35575, 511678, 133644, 361345, 178614, 575213, 71713, 238629 |
| 8974 | 523331, 134931, 356595, 170625, 10476, 478242 | As a 22-year-old, how risky should I be with my 401(k) investments? | 357555, 279570, 437706, 102501, 92941, 551545, 425703, 424766, 397553, 335146 |
| 10912 | 518721 | Forex independent investments | 413015, 387723, 358524, 586207, 597285, 79469, 245931, 80826, 266481, 329161 |
| 2964 | 95116 | Unmarried Couple Splitting up with Joint Ownership of Home | 280177, 58652, 144731, 352883, 153541, 159590, 78409, 77304, 415737, 569179 |
| 5178 | 240261, 561056, 111815, 393833, 39819 | Formula that predicts whether one is better off investing or paying down debt | 139788, 507544, 286567, 262772, 278638, 106215, 554465, 557506, 257016, 33912 |
| 5061 | 23747 | What fiscal scrutiny can be expected from IRS in early retirement? | 449387, 521843, 532781, 558661, 433256, 28662, 422062, 103668, 248536, 526334 |
| 7880 | 85319 | Are there index tracking funds that avoid the “buy high - sell low” problem? | 241202, 148721, 70194, 126151, 416839, 299284, 81652, 410123, 87261, 408524 |
| 4411 | 68239 | How does the importance of a cash emergency fund change when you live in a country with nationalized healthcare? | 570632, 587710, 323686, 70730, 36240, 179702, 361005, 598272, 225235, 347865 |
| 2075 | 170042, 44417, 260983, 359580, 60459, 523393, 301866, 14967 | Are stories of turning a few thousands into millions by trading stocks real? | 562784, 464810, 78224, 77802, 216494, 188129, 33357, 114439, 394357, 423056 |
| 8230 | 319599 | Why would this kind of penny stock increase so much in value? | 347521, 6369, 277305, 247802, 152097, 363204, 588384, 64961, 504243, 291220 |
| 4335 | 357013, 484148 | What is the US Fair Tax? | 318583, 585212, 363178, 322246, 3181, 185973, 516548, 122074, 419466, 328362 |
| 7533 | 93853 | Investing tax (savings) | 316230, 516756, 278638, 262772, 400714, 280530, 534670, 162668, 134332, 523127 |
| 1393 | 352838, 539133 | Which is better when working as a contractor, 1099 or incorporating? | 234436, 586026, 578196, 277812, 106684, 220022, 139501, 489898, 232544, 394871 |
| 9487 | 165544 | Is a public company allowed to issue new shares below market price without consulting shareholders? | 100546, 320472, 481761, 217006, 498752, 162454, 21975, 162612, 465542, 566553 |
| 5264 | 576564 | Does a company's stock price give any indication to or affect their revenue? | 431814, 165544, 52579, 246062, 8643, 115134, 264396, 471611, 10017, 233988 |
| 9733 | 110163, 38655 | Due Diligence - Dilution? | 121262, 450132, 61860, 316321, 308281, 301880, 261280, 19354, 65348, 144431 |
| 7311 | 323768 | Finance, Social Capital IPOA.U | 125355, 547865, 511480, 507841, 583646, 144059, 215232, 463599, 557607, 223113 |
| 744 | 490443, 566480, 78176 | What options are available for a home loan with poor credit but a good rental history? | 310790, 571801, 573276, 495431, 121063, 313896, 124699, 218849, 153088, 421575 |
| 7141 | 132288 | Do investors go long option contracts when they cannot cover the exercise of the options? | 103147, 288289, 401447, 305676, 570046, 166597, 538054, 243714, 220147, 94262 |
| 4071 | 129875 | If our economy crashes, and cash is worthless, should i buy gold or silver | 30305, 264898, 421752, 516596, 448375, 474187, 507038, 291862, 180404, 70575 |
| 4103 | 440270 | What causes US Treasury I bond fixed interest to increase? | 180958, 182249, 372657, 29073, 112659, 543165, 565356, 104160, 416513, 388391 |
| 3254 | 484891 | Why do people buy US dollars on the black market? | 443511, 502358, 531953, 594320, 148454, 313706, 476834, 310489, 356970, 221169 |
| 7512 | 191060 | understanding the process/payment of short sale dividends | 238903, 480949, 202985, 259450, 483305, 107045, 115553, 339875, 19234, 596518 |
| 2598 | 593029 | Is it possible for US retail forex traders to trade exotic currencies? | 358524, 586207, 79469, 245931, 316497, 594655, 394924, 311834, 591230, 153179 |
| 6800 | 35191 | I don't live in America. How can I buy IPO stock of newly listed companies in the United States? | 386278, 147792, 32589, 510163, 590937, 529638, 570046, 110584, 102742, 20359 |
| 1391 | 266229, 562176, 440745 | How is taxation for youtube/twitch etc monetization handled in the UK? | 267067, 367015, 198692, 560548, 527120, 30654, 593250, 149534, 186533, 243939 |
| 7534 | 358125 | Can you explain why it's better to invest now rather than waiting for the market to dip? | 175821, 473025, 426157, 302448, 320600, 301224, 3279, 11224, 86147, 237052 |
| 5356 | 71553, 381362, 312405, 562259 | Historical stock prices: Where to find free / low cost data for offline analysis? | 560213, 391171, 535343, 560108, 105607, 594159, 70185, 225818, 226749, 286335 |
| 2579 | 197052, 432808, 191977, 432020 | What to do when a job offer is made but with a salary less than what was asked for? | 552290, 181213, 260385, 157919, 366282, 340653, 524471, 489554, 229809, 280838 |
| 5790 | 134794 | FX losses on non-UK mortgage for UK property - tax deductable? | 594529, 243939, 129695, 149341, 443536, 528880, 333623, 256395, 1562, 484375 |
| 7823 | 583549 | Retirement Funds: Betterment vs Vanguard Life strategy vs Target Retirement | 451196, 502271, 268731, 175927, 11094, 331492, 411856, 105666, 347825, 169886 |
| 689 | 411044 | Receive credit card payment sending my customer details to a credit card processing company? | 446932, 433993, 131177, 122030, 235808, 410243, 425713, 59023, 438032, 9814 |
| 604 | 231947 | Is there a dollar amount that, when adding Massachusetts Sales Tax, precisely equals $200? | 535673, 547574, 52458, 571430, 258504, 193396, 366444, 204870, 431349, 436128 |
| 9174 | 431652, 535317, 160218, 544576, 405217 | Which U.S. online discount broker is the best value for money? | 451729, 596567, 476873, 365465, 413856, 192910, 55443, 879, 150632, 522798 |
| 6867 | 443804, 466143, 878, 540799, 502607, 445258, 538750 | Will there always be somebody selling/buying in every stock? | 230343, 229573, 573077, 235772, 118633, 250027, 212687, 133760, 527076, 61006 |
| 1670 | 290887 | Investing in hemp producers in advance of possible legalization in Canada? | 493792, 393898, 457882, 250542, 462439, 356212, 390668, 261271, 62966, 512265 |
| 6901 | 388571 | Rules for Broker Behavior with Covered Calls | 344065, 254474, 30557, 282143, 170700, 333674, 138201, 105373, 40447, 561884 |
| 10808 | 487052, 484599 | What are a few sites that make it easy to invest in high interest rate mutual funds? | 324012, 40227, 387722, 119700, 393733, 460402, 402523, 493336, 560783, 434734 |
| 2383 | 232199 | Should I Purchase Health Insurance Through My S-Corp | 17215, 573394, 224406, 46893, 128533, 476085, 154931, 563482, 544869, 207437 |
| 3085 | 248619, 527010 | How long can I convert 401(k) to Roth 401(k)? | 341493, 507892, 175470, 272789, 395840, 349384, 496253, 58103, 344526, 178340 |
| 104 | 575869, 523158 | Investing/business with other people's money: How does it work? | 251347, 386803, 596598, 179042, 24846, 385881, 462113, 348982, 298707, 477716 |
| 5083 | 138845 | Co-signer deceased | 369075, 267182, 18257, 273759, 453263, 157496, 270952, 142876, 437574, 44635 |
| 684 | 441120 | Beyond RRSP deductions, how does a high income earner save on taxes? | 68636, 181330, 388215, 212783, 363026, 376781, 196807, 599739, 342833, 499189 |
| 10674 | 99857, 257241, 543589, 3095 | How to sell a stock in a crashing market? | 72860, 303037, 503912, 328556, 365331, 435963, 99132, 105343, 318185, 273565 |
| 3594 | 246882, 554171, 525247 | If I were to get into a life situation where I would not be able to make regular payments, do lenders typically provi... | 24138, 113822, 140049, 479527, 391819, 592780, 14035, 398132, 347050, 501433 |
| 10526 | 39185 | What extra information might be obtained from the next highest bids in an order book? | 546493, 298551, 260153, 427747, 486692, 138830, 322798, 443804, 485973, 69150 |
| 2181 | 376631, 397329 | What are the risks & rewards of being a self-employed independent contractor / consultant vs. being a permanent emplo... | 277812, 526158, 181652, 197870, 37725, 139501, 223170, 524788, 489898, 145016 |
| 5903 | 231863 | Fees aside, what factors could account for performance differences between U.S. large-cap index ETFs? | 372233, 395842, 562919, 408524, 517391, 246531, 335136, 526346, 97836, 144005 |
| 5620 | 448784, 329552, 548740 | What's the fuss about identity theft? | 91986, 90632, 598801, 252475, 470024, 260580, 97686, 158285, 551747, 581889 |
| 5254 | 392851 | How do I calculate the quarterly returns of a stock index? | 508405, 457873, 122260, 105607, 563169, 531066, 205437, 99708, 471723, 559168 |
| 7221 | 186578 | How Technical Analysts react to non-market hours effects | 140266, 593521, 305117, 297231, 478539, 261331, 487918, 520424, 28590, 55002 |
| 3446 | 236899 | What's the difference between Term and Whole Life insurance? | 109675, 377477, 253202, 117627, 323498, 211839, 509077, 31550, 505027, 405178 |
| 2472 | 401125, 370334, 307315 | How do I deal with a mistaken attempt to collect a debt from me that is owed by someone else? | 62109, 422484, 255887, 337165, 592512, 122592, 584273, 189824, 543607, 568230 |
| 2306 | 315875 | To whom should I report fraud on both of my credit cards? | 556219, 584237, 596284, 506233, 456887, 6341, 578738, 581889, 270449, 249960 |
| 7633 | 197839 | Can a trade happen “in between” the bid and ask price? | 494727, 284235, 137175, 353396, 402482, 505244, 122323, 5018, 450489, 164008 |
| 2400 | 564271 | Will I be paid dividends if I own shares? | 311214, 433210, 158934, 315345, 126479, 186643, 29306, 400497, 584128, 571258 |
| 5549 | 286227, 309361 | Pros / cons of being more involved with IRA investments [duplicate] | 202140, 110465, 471686, 268802, 429106, 315847, 301609, 324273, 171253, 488037 |
| 3801 | 307776 | Can a bunch of wealthy people force Facebook to go public? | 390529, 209242, 216432, 168565, 92014, 69017, 169304, 187815, 391101, 362905 |
| 3735 | 314478 | Shorting Stocks And Margin Account Minimum | 549422, 5573, 370161, 68853, 537153, 547460, 35102, 279185, 527654, 565909 |
| 622 | 369239 | Accidentally opened a year term CD account, then realized I need the money sooner. What to do? | 243571, 373697, 586151, 50394, 515974, 250800, 4044, 397655, 260162, 497993 |
| 9088 | 569461, 561377 | Brokerage account for charity | 390089, 329849, 236186, 266342, 434224, 326167, 174543, 540285, 179408, 546509 |
| 4605 | 313306, 453941, 504661, 229310, 210759 | If the U.S. defaults on its debt, what will happen to my bank money? | 169691, 526384, 400826, 178586, 373717, 127268, 124624, 354896, 131116, 538582 |
| 2330 | 104221 | How can I determine if a debt consolidation offer is real or a scam? | 478514, 120837, 362721, 142582, 441836, 585692, 592325, 464477, 257483, 26263 |
| 6792 | 485973 | Where to find the full book of outstanding bids/asks for a stock? | 546493, 298551, 382067, 459650, 67069, 138830, 317365, 482517, 100188, 313570 |
| 2885 | 367360, 414692, 359579, 85229, 454810 | Merits of buying apartment houses and renting them | 5759, 17651, 159403, 215712, 585247, 309231, 434330, 502291, 410128, 6447 |
| 6110 | 331850, 94117, 259706 | Why does short selling require borrowing? | 188531, 320450, 49794, 314478, 363610, 67107, 206544, 79764, 233379, 107045 |
| 4823 | 104726, 362919 | Close to retirement & we may move within 7 years. Should we re-finance our mortgage, or not? | 561929, 108845, 34389, 426227, 340131, 366841, 283971, 175200, 254152, 72026 |
| 3694 | 282442 | Has anyone created a documentary about folks who fail to save enough for retirement? | 204747, 91911, 91265, 207356, 484683, 184243, 582307, 466552, 237924, 580956 |
| 8 | 566392 | How to deposit a cheque issued to an associate in my business into my business account? | 65404, 261856, 318108, 508754, 543812, 590102, 590837, 25397, 216200, 456636 |
| 1309 | 156162, 489401 | Why does FlagStar Bank harass you about payments within grace period? | 439548, 489368, 438869, 271040, 471630, 309171, 65982, 15824, 152279, 332160 |
| 7109 | 447781 | How do I analyse moving averages? | 42620, 565501, 140804, 489933, 227669, 43526, 488285, 193012, 35006, 257185 |
| 5080 | 256055 | Is there a standard or best practice way to handle money from an expiring UTMA account? | 62079, 451189, 414429, 69841, 445521, 212558, 200912, 279291, 327600, 318338 |
| 4981 | 247894 | Where can I find open source portfolio management software? | 102684, 259463, 303680, 45218, 78436, 419171, 186902, 529790, 38392, 160700 |
| 3932 | 107697, 527443 | How do historically low interest rates affect real estate prices? | 337793, 391831, 24822, 598764, 360199, 319126, 74668, 494992, 503988, 567749 |
| 3934 | 457034, 209604, 201769 | Should market based health insurance premiums be factored into 6 months emergency fund savings? | 402438, 46893, 490223, 431028, 184210, 157990, 122952, 526158, 79903, 52190 |
| 7445 | 153178, 104343, 296231 | IS it the wrong time to get into the equity market immediately after large gains? | 434252, 79111, 400782, 483025, 127401, 345129, 120395, 324210, 125940, 366847 |
| 2264 | 412819 | Personal Tax Return software for Linux? | 286321, 5862, 534454, 304594, 416778, 419245, 554114, 12987, 145650, 542213 |
| 2895 | 521996, 328691 | Where should a young student put their money? | 307162, 135013, 232508, 133487, 426461, 399863, 564206, 573076, 166109, 517313 |
| 6787 | 587120 | Would it make sense to sell a stock, then repurchase it for tax purposes? | 23217, 58290, 390864, 219762, 112195, 263751, 480967, 140835, 88540, 474981 |
| 7096 | 482238 | What's the formula for profits and losses when I delta hedge? | 202432, 272929, 65946, 407941, 157504, 338344, 187401, 197863, 315106, 572579 |
| 2747 | 540571 | What evidence do I need to declare tutoring income on my income tax? | 419768, 442391, 234975, 597328, 261280, 14609, 11654, 511583, 54952, 184698 |
| 1748 | 576295 | How high should I set my KickStarter funding goal in order to have $35,000 left over? | 18001, 451492, 124705, 240351, 554997, 29323, 197562, 349926, 521844, 114443 |
| 5862 | 130209, 269898, 170141 | Can I get a discount on merchandise by paying with cash instead of credit? | 518710, 495751, 122908, 10967, 486419, 21194, 562511, 564180, 104291, 546318 |
| 6041 | 241308 | Most effective Fundamental Analysis indicators for market entry | 81655, 425020, 108579, 224695, 96910, 327737, 204297, 367391, 111091, 194240 |
| 7700 | 273761, 507468, 2653, 179328 | Should I re-allocate my portfolio now or let it balance out over time? | 571913, 117634, 347411, 495473, 51334, 18436, 224392, 123789, 278678, 65532 |
| 10558 | 483268 | Investment strategy for 401k when rolling over soon | 262322, 9861, 509837, 217365, 64459, 63532, 72402, 436930, 396038, 426189 |
| 3014 | 341399, 273282 | What investments are positively related to the housing market decline? | 96118, 415121, 356346, 41884, 85749, 156180, 242321, 93890, 124493, 358586 |
| 3149 | 98112 | Tips for insurance coverage for one-man-teams | 418668, 54400, 472243, 73281, 457034, 518949, 153285, 124142, 153664, 229727 |
| 5343 | 2860 | “International credit report” for French nationals? | 70185, 21872, 486729, 329713, 118383, 179073, 128878, 34609, 327623, 509739 |
| 547 | 6349 | What percentage of my company should I have if I only put money? | 346537, 323355, 477853, 84642, 556421, 335248, 281707, 49614, 140349, 280845 |
| 3394 | 445971, 342258, 129319, 570664 | What is the easiest way to back-test index funds and ETFs? | 224765, 391215, 138383, 565296, 127566, 41176, 437875, 65567, 578530, 412830 |
| 5410 | 507813, 368802 | Dealership made me the secondary owner to my own car | 147530, 454526, 295355, 197138, 70456, 499392, 416606, 197352, 54501, 13975 |
| 3033 | 265866 | Tax consequences of changing state residency? | 159983, 424175, 486965, 394059, 598607, 32610, 571430, 41922, 60281, 360773 |
| 8332 | 445526 | Why do put option prices go higher when the underlying stock tanks (drops)? | 13260, 499811, 350748, 92616, 215118, 147361, 415281, 115652, 480337, 193303 |
| 3512 | 115042 | As an employee, when is it inappropriate to request to see your young/startup company's financial statements? | 269656, 72536, 204847, 9938, 295258, 466835, 258306, 295738, 374258, 129196 |
| 4102 | 241101, 39115, 448699 | How can I determine if my rate of return is “good” for the market I am in? | 162488, 272093, 597437, 528316, 88801, 126274, 453656, 454892, 153274, 369035 |
| 3566 | 99943, 307424 | Where can I buy stocks if I only want to invest a little bit at a time, and not really be involved in trading? | 113644, 311192, 206556, 282483, 519073, 138658, 306561, 45029, 506078, 195191 |
| 2551 | 413832, 450742, 143100 | How to find cheaper alternatives to a traditional home telephone line? | 498728, 66805, 185616, 118909, 549223, 419931, 353689, 124892, 42980, 137182 |

### Lowest MRR@10

| Query ID | MRR@10 | Recall@100 | First Relevant Rank | Query |
| --- | ---: | ---: | ---: | --- |
| 10039 | 0.0000 | 0.0000 | missing | Do individual investors use Google to obtain stock quotes? |
| 10137 | 0.0000 | 0.0000 | missing | F-1 student investing in foreign markets |
| 10213 | 0.0000 | 0.0000 | missing | Looking for good investment vehicle for seasonal work and savings |
| 10246 | 0.0000 | 0.0000 | missing | Understanding the T + 3 settlement days rule |
| 104 | 0.0000 | 0.0000 | missing | Investing/business with other people's money: How does it work? |
| 10462 | 0.0000 | 0.0000 | missing | Is it okay to be married, 30 years old and have no retirement? |
| 10482 | 0.0000 | 0.0000 | missing | Rollover into bond fund to do dollar cost averaging [duplicate] |
| 10674 | 0.0000 | 0.0000 | missing | How to sell a stock in a crashing market? |
| 1074 | 0.0000 | 0.0000 | missing | How common is “pass-through” health insurance? |
| 10827 | 0.0000 | 0.0000 | missing | How much should I be contributing to my 401k given my employer's contribution? |
