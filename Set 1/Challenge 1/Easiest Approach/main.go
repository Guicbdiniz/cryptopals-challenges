package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	fmt.Printf("Hex: %s\n", hexString)

	decodedHexString, _ := hex.DecodeString(hexString)

	fmt.Printf("Decoded Hex: %s\n", decodedHexString)

	encodedBase64String := base64.StdEncoding.EncodeToString([]byte(decodedHexString))

	fmt.Printf("Encoded Base64: %s\n", encodedBase64String)
}
