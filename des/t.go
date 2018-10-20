package des

// tIteration T-迭代 raw(L0R0):64 key: 48 * 16 Test OK
func tIteration(raw []byte, key [][]byte, isInv bool) []byte {
	// T-迭代
	for i := 1; i <= 16; i++ {
		newRaw := make([]byte, 65)
		copy(newRaw[1:33], raw[33:65])
		r := make([]byte, 33)
		copy(r[1:33], raw[1:33])
		var f []byte
		if isInv {
			f = feistel(newRaw[:33], key[17-i])
		} else {
			f = feistel(newRaw[:33], key[i])
		}
		for j := 1; j <= 32; j++ {
			r[j] = r[j] ^ f[j]
		}
		copy(newRaw[33:65], r[1:33])
		copy(raw[1:65], newRaw[1:65])
	}
	// W-置换
	res := make([]byte, 65)
	copy(res[1:33], raw[33:65])
	copy(res[33:65], raw[1:33])
	return res
}
