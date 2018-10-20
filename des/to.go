package des


// to64 raw: [0-7]8 to [1-64]64
func to64(raw []byte) []byte {
	res := make([]byte, 65)
	for i := 0; i < 8; i++ {
		cur := raw[i]
		for j := 0; j < 8; j++ {
			res[(8-j)+i*8] = byte(cur % 2)
			cur >>= 1
		}
	}
	return res
}

// to8 raw: [1-64]64 to [0-7]8
func to8(raw []byte) []byte {
	res := make([]byte, 8)
	for i := 1; i <= 8; i++ {
		var b byte
		for j := 1; j <= 8; j++ {
			b <<= 1
			if raw[(i - 1) * 8 + j] == 1 {
				b |= 1
			}
		}
		res[i-1] = b
	}
	return res
}
