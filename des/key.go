package des

var pc1Table = []byte{0,
57, 49, 41, 33, 25, 17, 9,
1, 58, 50, 42, 34, 26, 18,
10, 2, 59, 51, 43, 35, 27,
19, 11, 3, 60, 52, 44, 36,
63, 55, 47, 39, 31, 23, 15,
7, 62, 54, 46, 38, 30, 22,
14, 6, 61, 53, 45, 37, 29,
21, 13, 5, 28, 20, 12, 4}

var pc2Table = []byte{0,
14, 17, 11, 24, 1, 5,
3, 28, 15, 6, 21, 10,
23, 19, 12, 4, 26, 8,
16, 7, 27, 20, 13, 2,
41, 52, 31, 37, 47, 55,
30, 40, 51, 45, 33, 48,
44, 49, 39, 56, 34, 53,
46, 42, 50, 36, 29, 32}

// makeKey 生成密钥 Test OK!
func makeKey(key []byte) [][]byte {
	if len(key) != 8 {
		panic("Error key")
	}
	var subKeys [][]byte
	// 第一个为空
	subKeys = append(subKeys, make([]byte, 1))
	cd := pc1Transform(key)
	for i := 1; i <= 16; i++ {
		if i == 1 || i == 2 || i == 9 || i == 16 {
			// 循环左移一位
			headC := cd[1]
			headD := cd[29]
			for j := 1; j < 28; j++ {
				cd[j] = cd[j + 1]
				cd[j + 28] = cd[j + 29]
			}
			cd[28] = headC
			cd[56] = headD
		} else {
			// 循环左移两位
			headC1, headC2 := cd[1], cd[2]
			headD1, headD2 := cd[29], cd[30]
			for j := 1; j < 27; j++ {
				cd[j] = cd[j + 2]
				cd[j + 28] = cd[j + 30]
			}
			cd[27], cd[28] = headC1, headC2
			cd[55], cd[56] = headD1,headD2
		}
		subKeys = append(subKeys, pc2Transform(cd))
	}
	return subKeys
}

func pc1Transform(key []byte) []byte {
	newKey := make([]byte, 57)
	rawKey := to64(key)
	for i := 1; i < len(pc1Table); i++ {
		newKey[i] = rawKey[pc1Table[i]]
	}
	return newKey
}

// pc2Transform 56 to 48
func pc2Transform(bit []byte) []byte {
	key := make([]byte, 49)
	for i := 1; i <= 48; i++ {
		key[i] = bit[pc2Table[i]]
	}
	return key
}
