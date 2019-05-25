package des

func mPKCS5Padding(raw []byte) []byte {
	count := 8 - len(raw)%8
	for i := 0; i < count; i++ {
		raw = append(raw, byte(count))
	}
	return raw
}
