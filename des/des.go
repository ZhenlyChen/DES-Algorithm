package des

func Encrypt(plain, key string) string {
	if len(key) != 8 {
		return ""
	}
	// 补全
	raw := completion([]byte(plain))
	// IP变换
	raw = ipTransform(raw)
	return ""
}

func Decrypt(cipher, key string) string {
	if len(key) != 8 {
		return ""
	}

	return ""
}
