package main

import (
	"encoding/hex"
	"fmt"
)

func XORHexStringToChar(hexStringA string, charToHex byte) string {
	decodedA, _ := hex.DecodeString(hexStringA)

	xored := make([]byte, len(decodedA))

	for i := 0; i < len(decodedA); i++ {
		xored[i] = decodedA[i] ^ charToHex
	}

	return string(xored)
}

func main() {
	hexStringA := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	for i := 0; i < 256; i++ {
		hexChar := byte(i)
		xored := XORHexStringToChar(hexStringA, hexChar)
		fmt.Printf("XOR with %v:\n %s\n", hexChar, xored)
	}

	fmt.Printf("Best solution was to XOR with 'X' (88) to get:\n %s\n", XORHexStringToChar(hexStringA, byte(88)))
}
