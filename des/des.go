package des

import (
	"encoding/hex"
)

func Encrypt(plain, key string) string {
	return hex.EncodeToString(myDES([]byte(plain), []byte(key), true))
}

func Decrypt(cipher, key string) string {
	cipherByte, err := hex.DecodeString(cipher)
	if err != nil {
		panic("Error cipher")
	}
	plain := myDES(cipherByte, []byte(key), false)
	length := len(plain)
	return string(plain[:length - int(plain[length - 1])])
}

func myDES(raw, key []byte, isEncrypt bool) []byte {
	if len(key) != 8 {
		panic("Error key length")
	}
	var res []byte
	// 生成16子密钥 Test OK
	keys := makeKey(key)
	// 补全
	if isEncrypt {
		raw = mPKCS5Padding(raw)
	}
	times := len(raw) / 8
	for i := 0; i < times; i++ {
		cRaw := raw[i*8 : i*8+8]
		// IP变换 Test OK
		cRaw = ipTransform(cRaw)
		cRaw = to64(cRaw)
		// T-迭代 && 交换置换W Test OK
		cRaw = tIteration(cRaw, keys, !isEncrypt)
		cRaw = to8(cRaw)
		// IP逆置换 Test OK
		cRaw = ipInverseTransform(cRaw)
		res = append(res, cRaw...)
	}
	return res
}