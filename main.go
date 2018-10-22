package main

import (
	"fmt"

	"github.com/ZhenlyChen/DES-Golang/des"
)

func main() {
	key := "megashow"
	plain := "Hello,world! Megashow is very good."
	cipher := des.Encrypt(plain, key)
	mPlain := des.Decrypt(cipher, key)
	fmt.Println("plain: ", plain)
	fmt.Println("cipher: ", cipher)
	fmt.Println("mPlain: ", mPlain)
	if mPlain != plain {
		panic("")
	}
}
