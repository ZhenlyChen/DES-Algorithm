package main

import (
	"fmt"
	"github.com/ZhenlyChen/goDES/des"
)

func main() {

	key := "megashow"
	plain := "Hello, world!"
	cipher := des.Encrypt(plain, key)
	mPlain := des.Decrypt(cipher, key)
	fmt.Println("plain: "+plain)
	fmt.Println("cipher: "+ cipher)
	if mPlain != plain {
		panic("")
	}
}
