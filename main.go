package main

import (
	"fmt"
	"github.com/dineshbt/cryptit/decrypt"
	"github.com/dineshbt/cryptit/encrypt"
)

func main() {
	str := "dinesh"
	fmt.Println("original string= ", str)
	encryptedStr := encrypt.Nimbus(str)
	fmt.Println("encrypted String= ", encryptedStr)
	decryptedStr := decrypt.Nimbus(encryptedStr)
	fmt.Println("decrypted string =", decryptedStr)

}
