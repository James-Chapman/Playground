package search

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	var s = []int{1578208, -1459287, 149047, 4098454, -4943485, 3035907, 1673804, 744815, 6760469, -3945544, 7224282, 8182859, -6575770, -4661905, 3115289, 1917001, -1084917, -8646996, -3268053, -3699637, 1566614, -3764121, 396738, -618384, -5508573, 2452592, -215375, 5080692, 3619196, -2336927, -468224, -1187775, -596570, -105625, -8940982, 4865729, 254545, -1934594, -357295, -2894681, 2968005, -1102542, 2334004, 6366277, 4881444, 1844413, 584100, -3721557, 667422, 1366847, -310063, -2330755, -2306687, 2133133, -2303495, 3504352, 4382054, -751171, -6052570, 8380109, 1181999, 5151997, 4932070, 2912990, -8636742, -2300576, -422352, -8709059, 3498943, 6287972, -4808191, 1599652, 5840194, -5203156, 5832672, -3705946, -6679517, 2440690, 2975081, 3158425, -7070559, 4281021, 6558798, -5101054, -7223936, -1137093, 858327, -7328718, 1892391, -787228, -5732453, -3440484, -3719625, 4439505, -7616755, 3545698, -7449092, -7934882, 3043929, 1400553, 303599, 3625923, 6847324, 761089, 1851390, -5896723, -8926745, 156374, 1035113, 2038094, -2306413, 643897, -890608, 6195300, -3419560, 5079309, -7539349, -4409129, -3556416, 330954, 2394139, -3495311, -2924562, -2039519, -7644481, 248779, -5387666, 5082563, -4584466, 1150689, 2335887, 6453380, 3615054, -1724601, -8768701, 4702235, -4721034, -8349075, -1242494, -3070428, 1412852, -7461701, -126483, -7236665, 992705, -2429973, 1958773, -4035436, 551387, 6015162, -19240, 4585030, 5746385, -6466555, -81263, 2169671, 4156645, 3872477, 8178019, 6308621, -1139658, 1816073, -2838201, 7375926, -2391037, -1398918, -360671, 989181, -2920085, 1325994, 7830876, -440085, 1723677, 4902068, 39461, -2892829, -3307844, 7676205, -1349335, -1377968, -6918146, -1620085, -1705611, -6435546, -97829, -2186543, 9474019, -8505279, -2264021, 3448386, 5654087, 1551857, 6753790, 964432, -5735653, -5281130, 408351, 3052539, -6616520, -596858, 2615883, -1756041, 7524164, 344169, -7200691, 3459305, 2669362, -1750297, 3077300, 4910895, 3069747, 5238465, -3280284, -310410, 3666781, -807778, -5444866, -280829, 3005388, -2434618, 5601413, 3739446, -459159, -443543, -1852922, -3621556, 6234191, -5668573, -8579353, -384084, 7374628, 5238350, 7537296, -4345887, -9755992, -642806, -3310426, 1065418, 2189932, -141851, -6671007, -5684574, 2194413, 3918074, 947107, -5566527, -3348533, -2814837, -2461508, 689999, 5847565, -2156201, -8053815, 882006, 127573, 2342741, -81332, -2501146, 2092726, -2697358, 5083701, -347376, 1380446, -1894709, -984537, -8518114, -4753750, 661540, -1886897, 26960, 1864380, 4242384, 2358639, -4139139, 5734386, -6545688, -354314, -388276, -6392582, -6750945, 3251518, -4426308, 4596397, 2417102, 234710, -797992, -5714557, 4389226, 3736363, 2852346, -3601768, -499585, 2770221, -3256243, 3836405, 6603488, -4148861, -4614991, 117068, -2316080, -790931, 3529555, 2694232, -2282699, 242024, 8397086, -3607030, 7069813, -3966857, -8586398, -6545943, -2818142, 8848350, 425406, 5526419, 2998398, -4807489, -2066930, -5753068, 8583981, -7844724, -593897, -5859731, 1739528, 881437, -1144473, -6334250, 3325631, 3192681, 420539, 4919549, 4250240, -161738, -321868, -1610853, 280040, 2513726, 3356012, -4493893, -4419862, -1637946, 3027561, -32015, 8080410, -757994, -1621674, -3637500, -4004156, 1449732, 2701987, -913718, 122128, -24511, -5407571, 715598, 1040129, 2372787, 2275143, -4497099, 4097947, -4500460, 3658841, 166426, -2358746, 2791981, 2227909, 495914, -1391938, 4243527, -3277336, -3112135, 5718422, 6069541, 3159171, -3631467, 409250, -4990600, -1336561, 4002118, 5031530, -3108904, -4610777, -4088125, -3141620, 2816015, -961217, 3626569, -130153, -7398238, 1770640, -142970, 5339746, 3085633, -1893606, 226253, 7095788, -5096248, 3005353, 4926144, 3172334, 8564443, 4525814, 9162939, 3481277, -8260972, -609578, -7374531, -6422361, 3945778, -3065884, -4109240, -8026712, 4199917, -5210826, -5748685, 2358447, 1962956, 2154131, -2840429, -5122247, 3309442, 8161728, -2867466, -700496, 7724913, 7411646, 1842574, 1855934, 502862, 1356091, -3597168, -497115, -2998171, -2279103, -3429551, -6184888, -2515531, 7276990, -87396, 2996265, 1892949, -329269, -4161321, -4033470, -1428559, 1287118, -3782609, 5658426, -3581622, -5099684, -1025154, 1033647, 1207674, -2086170, 1348100, 4330367, -4311279, -3191296, -3690060, -4484260, -1243635, -12616, -1112915, 394288, 7298505, 2422986, -3650638, 1871029, 347401, -1264656, -4887911, -2701857, 2471433, -6472793, 16550, 1040124, 3131078, 330154, 5701075, 5524260, -118421, -2982372, 3700779, -9186375, -6220337, -5112976, 2404691, 2796285, -2184606, -5615090, -636394, -2579519, -6946500, 1833802, 1048877, -425490, 4424556, -3182151, 6765807, -5310643}
	SelectionSort(s)
	for i, v := range s {
		r := BinarySearch(v, s)
		assert.Equal(t, i, r, "Position should be %d", i)
	}
}

func TestLinearSearch(t *testing.T) {
	var s = []int{1578208, -1459287, 149047, 4098454, -4943485, 3035907, 1673804, 744815, 6760469, -3945544, 7224282, 8182859, -6575770, -4661905, 3115289, 1917001, -1084917, -8646996, -3268053, -3699637, 1566614, -3764121, 396738, -618384, -5508573, 2452592, -215375, 5080692, 3619196, -2336927, -468224, -1187775, -596570, -105625, -8940982, 4865729, 254545, -1934594, -357295, -2894681, 2968005, -1102542, 2334004, 6366277, 4881444, 1844413, 584100, -3721557, 667422, 1366847, -310063, -2330755, -2306687, 2133133, -2303495, 3504352, 4382054, -751171, -6052570, 8380109, 1181999, 5151997, 4932070, 2912990, -8636742, -2300576, -422352, -8709059, 3498943, 6287972, -4808191, 1599652, 5840194, -5203156, 5832672, -3705946, -6679517, 2440690, 2975081, 3158425, -7070559, 4281021, 6558798, -5101054, -7223936, -1137093, 858327, -7328718, 1892391, -787228, -5732453, -3440484, -3719625, 4439505, -7616755, 3545698, -7449092, -7934882, 3043929, 1400553, 303599, 3625923, 6847324, 761089, 1851390, -5896723, -8926745, 156374, 1035113, 2038094, -2306413, 643897, -890608, 6195300, -3419560, 5079309, -7539349, -4409129, -3556416, 330954, 2394139, -3495311, -2924562, -2039519, -7644481, 248779, -5387666, 5082563, -4584466, 1150689, 2335887, 6453380, 3615054, -1724601, -8768701, 4702235, -4721034, -8349075, -1242494, -3070428, 1412852, -7461701, -126483, -7236665, 992705, -2429973, 1958773, -4035436, 551387, 6015162, -19240, 4585030, 5746385, -6466555, -81263, 2169671, 4156645, 3872477, 8178019, 6308621, -1139658, 1816073, -2838201, 7375926, -2391037, -1398918, -360671, 989181, -2920085, 1325994, 7830876, -440085, 1723677, 4902068, 39461, -2892829, -3307844, 7676205, -1349335, -1377968, -6918146, -1620085, -1705611, -6435546, -97829, -2186543, 9474019, -8505279, -2264021, 3448386, 5654087, 1551857, 6753790, 964432, -5735653, -5281130, 408351, 3052539, -6616520, -596858, 2615883, -1756041, 7524164, 344169, -7200691, 3459305, 2669362, -1750297, 3077300, 4910895, 3069747, 5238465, -3280284, -310410, 3666781, -807778, -5444866, -280829, 3005388, -2434618, 5601413, 3739446, -459159, -443543, -1852922, -3621556, 6234191, -5668573, -8579353, -384084, 7374628, 5238350, 7537296, -4345887, -9755992, -642806, -3310426, 1065418, 2189932, -141851, -6671007, -5684574, 2194413, 3918074, 947107, -5566527, -3348533, -2814837, -2461508, 689999, 5847565, -2156201, -8053815, 882006, 127573, 2342741, -81332, -2501146, 2092726, -2697358, 5083701, -347376, 1380446, -1894709, -984537, -8518114, -4753750, 661540, -1886897, 26960, 1864380, 4242384, 2358639, -4139139, 5734386, -6545688, -354314, -388276, -6392582, -6750945, 3251518, -4426308, 4596397, 2417102, 234710, -797992, -5714557, 4389226, 3736363, 2852346, -3601768, -499585, 2770221, -3256243, 3836405, 6603488, -4148861, -4614991, 117068, -2316080, -790931, 3529555, 2694232, -2282699, 242024, 8397086, -3607030, 7069813, -3966857, -8586398, -6545943, -2818142, 8848350, 425406, 5526419, 2998398, -4807489, -2066930, -5753068, 8583981, -7844724, -593897, -5859731, 1739528, 881437, -1144473, -6334250, 3325631, 3192681, 420539, 4919549, 4250240, -161738, -321868, -1610853, 280040, 2513726, 3356012, -4493893, -4419862, -1637946, 3027561, -32015, 8080410, -757994, -1621674, -3637500, -4004156, 1449732, 2701987, -913718, 122128, -24511, -5407571, 715598, 1040129, 2372787, 2275143, -4497099, 4097947, -4500460, 3658841, 166426, -2358746, 2791981, 2227909, 495914, -1391938, 4243527, -3277336, -3112135, 5718422, 6069541, 3159171, -3631467, 409250, -4990600, -1336561, 4002118, 5031530, -3108904, -4610777, -4088125, -3141620, 2816015, -961217, 3626569, -130153, -7398238, 1770640, -142970, 5339746, 3085633, -1893606, 226253, 7095788, -5096248, 3005353, 4926144, 3172334, 8564443, 4525814, 9162939, 3481277, -8260972, -609578, -7374531, -6422361, 3945778, -3065884, -4109240, -8026712, 4199917, -5210826, -5748685, 2358447, 1962956, 2154131, -2840429, -5122247, 3309442, 8161728, -2867466, -700496, 7724913, 7411646, 1842574, 1855934, 502862, 1356091, -3597168, -497115, -2998171, -2279103, -3429551, -6184888, -2515531, 7276990, -87396, 2996265, 1892949, -329269, -4161321, -4033470, -1428559, 1287118, -3782609, 5658426, -3581622, -5099684, -1025154, 1033647, 1207674, -2086170, 1348100, 4330367, -4311279, -3191296, -3690060, -4484260, -1243635, -12616, -1112915, 394288, 7298505, 2422986, -3650638, 1871029, 347401, -1264656, -4887911, -2701857, 2471433, -6472793, 16550, 1040124, 3131078, 330154, 5701075, 5524260, -118421, -2982372, 3700779, -9186375, -6220337, -5112976, 2404691, 2796285, -2184606, -5615090, -636394, -2579519, -6946500, 1833802, 1048877, -425490, 4424556, -3182151, 6765807, -5310643}
	SelectionSort(s)
	for i, v := range s {
		r := LinearSearch(v, s)
		assert.Equal(t, i, r, "Position should be %d", i)
	}
}