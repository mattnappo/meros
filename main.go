package main

import (
	"fmt"

	"github.com/xoreo/meros/crypto"
)

func main() {
	payload := []byte("test hash")
	hash := crypto.Sha3(payload)
	fmt.Println(hash.String())
}
