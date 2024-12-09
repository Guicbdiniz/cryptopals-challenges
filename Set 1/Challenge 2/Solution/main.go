package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	hexStringA := "1c0111001f010100061a024b53535009181c"
	hexStringB := "686974207468652062756c6c277320657965"

	fmt.Printf("A: %s\n", hexStringA)
	fmt.Printf("B: %s\n", hexStringB)

	decodedA, _ := hex.DecodeString(hexStringA)
	decodedB, _ := hex.DecodeString(hexStringB)

	fmt.Printf("Decoded A: %v\n", decodedA)
	fmt.Printf("Decoded B: %v\n", decodedB)

	xored := make([]byte, len(decodedA))

	for i := 0; i < len(decodedA); i++ {
		xored[i] = decodedA[i] ^ decodedB[i]
	}

	fmt.Printf("XOR as bytes: %v\n", xored)

	encodedHex := hex.EncodeToString(xored)

	fmt.Printf("XOR as hex: %s\n", encodedHex)
}
