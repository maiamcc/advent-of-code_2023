package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	inputStr := "seeds: 565778304 341771914 1736484943 907429186 3928647431 87620927 311881326 149873504 1588660730 119852039 1422681143 13548942 1095049712 216743334 3671387621 186617344 3055786218 213191880 2783359478 44001797\n\nseed-to-soil map:\n1136439539 28187015 34421000\n4130684560 3591141854 62928737\n2493176649 2843539493 216586902\n4035246184 3979580848 40675839\n784987951 2449883248 10512167\n1230114095 458474273 89127842\n3591141854 4278550666 16416630\n795500118 1007741104 49669915\n4075922023 4020256687 54762537\n1170860539 385724159 59253556\n1754134353 1447758461 710855281\n2464989634 0 28187015\n3811089926 3654070591 224156258\n367106182 564462768 34737691\n0 3060126395 64826339\n1438999297 87449756 298274403\n1319241937 901480302 106260802\n1425502739 444977715 13496558\n906091129 2158613742 230348410\n401843873 2460395415 383144078\n1737273700 547602115 16860653\n64826339 599200459 302279843\n2709763551 1057411019 390347442\n845170033 2388962152 60921096\n3607558484 4075019224 203531442\n3100110993 3124952734 265337006\n4193613297 3878226849 101353999\n3365447999 62608015 24841741\n\nsoil-to-fertilizer map:\n2997768542 2385088490 141138894\n2483957796 2361581050 23507440\n98641524 1346083581 385280737\n3138907436 2256873732 8670947\n0 2158232208 98641524\n3147578383 2265544679 96036371\n1035235183 2879344429 108036359\n2567031012 2526227384 63435416\n740156227 2589662800 180702628\n2630466428 1790930094 367302114\n1029837856 0 5397327\n1143271542 5397327 1340686254\n483922261 2987380788 256233966\n2507465236 1731364318 59565776\n920858855 2770365428 108979001\n\nfertilizer-to-water map:\n1539871014 1431400479 38399903\n4189242304 3947275099 105724992\n2012473116 0 61612686\n3673653298 3769966020 177309079\n25380533 833117788 21807501\n143369400 1411638591 19761888\n2698209531 61612686 40666379\n401367210 2888296065 27849039\n3850962377 4057978463 170640183\n1076364770 854925289 39443942\n0 2048878915 25380533\n2682826842 1483785677 15382689\n4026580932 4228618646 66348650\n790899137 2074259448 70405647\n2738875910 2609016218 230235412\n2090748148 1854132037 12185242\n163131288 1499168366 238235922\n1115808712 3002097202 63461270\n545943998 1215727887 195910704\n4092929582 3673653298 96312722\n1000530579 3065558472 75834191\n2074085802 2916145104 16662346\n429216249 1737404288 116727749\n1578270917 2174814019 434202199\n2969111322 102279065 44844471\n1179269982 669063687 164054101\n2463111278 507301424 161762263\n741854702 2839251630 49044435\n3013955793 894369231 236513741\n861304784 2144665095 30148924\n1525885719 1469800382 13985295\n132032949 2932807450 11336451\n2102933390 147123536 360177888\n47188034 1130882972 84844915\n4021602560 4053000091 4978372\n2624873541 2944143901 57953301\n891453708 3141392663 109076871\n1343324083 1866317279 182561636\n\nwater-to-light map:\n1509583382 1639808290 20361832\n3841220400 2799952377 116887408\n1472887638 3349716751 36695744\n1375316591 4197396249 97571047\n1030032900 38536653 44339012\n3776233310 1557050237 64987090\n1857053855 3386412495 71799907\n2963593546 2694182899 38493443\n3758462347 1622037327 17770963\n1018869652 82875665 11163248\n1308040556 2732676342 67276035\n1928853762 3953749938 243646311\n2488961036 3503789336 239267964\n3562290347 3458212402 6096433\n3568386780 1308040556 190075567\n2728229000 1997029610 235364546\n215668494 0 38536653\n1646361217 3743057300 109790942\n1529945214 3233300748 116416003\n1756152159 3852848242 100901696\n3958107808 1660170122 336859488\n3503356233 1498116123 58934114\n254205147 552157437 522214475\n3002086989 3464308835 39480501\n776419622 94038913 242450030\n2172500073 2916839785 316460963\n0 336488943 215668494\n3041567490 2232394156 461788743\n\nlight-to-temperature map:\n3498288578 2645051323 42074132\n608593503 673232568 65024140\n0 1287033796 108723708\n3979313387 3634135302 315653909\n2652759587 3018896130 103365881\n1093544955 942695289 7961217\n2756125468 3501628238 132507064\n419683046 1625547778 126722533\n683243352 349510049 26140330\n1101506172 511314709 142580382\n1283347805 375650379 135664330\n673617643 24016268 9625709\n709383682 1238532395 48501401\n3763762608 2402322728 184668443\n3948431051 2621479653 23571670\n3660408219 2687125455 102463666\n3561232935 2586991171 34488482\n2459767392 3122262011 192992195\n1244086554 1754618583 31645052\n1073753155 1497748002 19791800\n3972002721 3949789211 7310666\n3540362710 4077337854 20870225\n1419012135 738256708 79984976\n1746461061 1457945428 39802574\n3595721417 4190079809 64686802\n1616767336 653895091 19337477\n1275731606 823578131 7616199\n108723708 818241684 5336447\n546405579 1395757504 62187924\n1498997111 950656506 117770225\n2888632532 3957099877 102103275\n1744112789 1752270311 2348272\n2233819388 2077946837 225948004\n3762871885 2401432005 890723\n2136282224 2303894841 97537164\n3311914546 3315254206 186374032\n2118147522 4059203152 18134702\n308182087 831194330 111500959\n114060155 0 24016268\n2077946837 4254766611 40200685\n971525741 33641977 102227414\n1636104813 1517539802 108007976\n138076423 1068426731 170105664\n3220042816 4098208079 91871730\n2990735807 2789589121 229307009\n757885083 135869391 213640658\n\ntemperature-to-humidity map:\n1130946446 972737563 146373650\n1277320096 1760175559 41760032\n4151385320 4147641404 143581976\n2634337722 0 466605084\n1992884166 956487184 16250379\n4147641404 4291223380 3743916\n641064346 466605084 489882100\n1319080128 1801935591 673804038\n2009134545 2475739629 625203177\n0 1119111213 641064346\n\nhumidity-to-location map:\n3903940466 3635148971 125939893\n1458128760 2186815403 67353660\n3125319983 1458128760 728686643\n2261072201 3994982121 66012689\n3854006626 2992363154 49933840\n1525482420 3780550419 183145699\n2233668127 3967578047 27404074\n2442260515 3138011064 466023456\n740912129 0 327948845\n2422798960 3761088864 19461555\n1708628119 3963696118 3881929\n2327084890 3042296994 95714070\n4029880359 3604034520 31114451\n1712510048 2471205075 521158079\n367508399 695457244 373403730\n2908283971 2254169063 217036012\n0 327948845 367508399"
	inputLns := strings.Split(inputStr, "\n")

	actual := partOne(inputLns)
	assert.Equal(t, 35, actual)
}

func TestPartTwo(t *testing.T) {
	inputStr := "input\ngoes\nhere"
	inputLns := strings.Split(inputStr, "\n")

	actual := partTwo(inputLns)
	assert.Equal(t, 123, actual)
}