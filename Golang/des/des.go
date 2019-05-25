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
	// 删除PKCS5填充
	length := len(plain)
	return string(plain[:length-int(plain[length-1])])
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
		// IP变换 Test OK
		raw := ipTransform(raw[i*8 : i*8+8])
		// T-迭代 && 交换置换W Test OK
		raw = to8(tIteration(to64(raw), keys, !isEncrypt))
		// IP逆置换 Test OK
		raw = ipInverseTransform(raw)
		res = append(res, raw...)
	}
	return res
}
