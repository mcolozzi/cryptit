package main

import (
	"fmt"

	"github.com/mcolozzi/cryptit/decrypt"
	"github.com/mcolozzi/cryptit/encrypt"
)

func main() {
	encryptedStr := encrypt.Nimbus("Kodekloud")
	fmt.Println(encryptedStr)
	fmt.Println(decrypt.Nimbus(encryptedStr))
}
