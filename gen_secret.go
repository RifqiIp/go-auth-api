package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func main() {
	bytes := make([]byte, 64) // 64 byte = sangat aman
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}

	fmt.Println(hex.EncodeToString(bytes))
}