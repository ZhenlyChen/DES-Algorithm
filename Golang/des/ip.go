package des

// ipTransform 8 to 8 Test OK!
func ipTransform(raw []byte) []byte {
	res := make([]byte, 8)
	table := []uint8{2, 4, 6, 8, 1, 3, 5, 7}
	for j := 0; j < 8; j++ {
		var b byte = 0
		for k := 0; k < 8; k++ {
			b = b << 1
			if raw[7-k]&(1<<(8-table[j])) != 0 {
				b |= 1
			}
		}
		res[j] = b
	}
	return res
}

// ipInverseTransform 8 to 8
func ipInverseTransform(raw []byte) []byte {
	res := make([]byte, 8)
	table := []uint8{5, 1, 6, 2, 7, 3, 8, 4}
	for j := 0; j < 8; j++ {
		var b byte = 0
		for k := 0; k < 8; k++ {
			b = b << 1
			if raw[table[k]-1]&(1<<uint8(j)) != 0 {
				b |= 1
			}
		}
		res[j] = b
	}
	return res
}
