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
	cd := pcTransform(to64(key), pc1Table)
	for i := 1; i <= 16; i++ {
		if i != 1 && i != 2 && i != 9 && i != 16 {
			cd = leftShift(cd)
		}
		// 循环左移两次
		cd = leftShift(cd)
		subKeys = append(subKeys, pcTransform(cd, pc2Table))
	}
	return subKeys
}

// leftShift 循环左移
func leftShift(raw []byte) []byte {
	headC := raw[1]
	headD := raw[29]
	for j := 1; j < 28; j++ {
		raw[j] = raw[j+1]
		raw[j+28] = raw[j+29]
	}
	raw[28] = headC
	raw[56] = headD
	return raw
}

func pcTransform(bit, table []byte) []byte {
	key := make([]byte, 57)
	for i := 1; i <= 48; i++ {
		key[i] = bit[table[i]]
	}
	return key
}
