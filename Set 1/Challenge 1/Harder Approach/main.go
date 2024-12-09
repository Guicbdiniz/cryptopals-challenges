package main

import (
	"fmt"
	"strings"
)

func hexToBinary(hexChar byte) string {
	switch hexChar {
	case '0':
		return "0000"
	case '1':
		return "0001"
	case '2':
		return "0010"
	case '3':
		return "0011"
	case '4':
		return "0100"
	case '5':
		return "0101"
	case '6':
		return "0110"
	case '7':
		return "0111"
	case '8':
		return "1000"
	case '9':
		return "1001"
	case 'a', 'A':
		return "1010"
	case 'b', 'B':
		return "1011"
	case 'c', 'C':
		return "1100"
	case 'd', 'D':
		return "1101"
	case 'e', 'E':
		return "1110"
	case 'f', 'F':
		return "1111"
	default:
		return ""
	}
}

func decodeHexString(value string) []byte {
	var binaryStringBuilder strings.Builder

	for i := 0; i < len(value); i++ {
		binaryStringBuilder.WriteString(hexToBinary(value[i]))
	}

	binaryString := binaryStringBuilder.String()
	byteCount := len(binaryString) / 8

	decodedStringAsBytes := make([]byte, byteCount)
	for i := 0; i < byteCount; i++ {
		byteStr := binaryString[i*8 : (i+1)*8]
		var byteValue byte
		for j := 0; j < 8; j++ {
			if byteStr[j] == '1' {
				byteValue |= 1 << (7 - j)
			}
		}
		decodedStringAsBytes[i] = byteValue
	}

	return decodedStringAsBytes
}

func getBitsFromBytes(bytes []byte) []int {
	bits := make([]int, 0, len(bytes)*8)
	for _, b := range bytes {
		for i := 0; i < 8; i++ {
			// Shift the byte right by (7 - i) and mask with 1 to get the bit at position i
			bit := int((b >> (7 - i)) & 1)
			bits = append(bits, bit)
		}
	}
	return bits
}

func encodeToBase64(value []byte) string {
	base64Chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

	bitsFromValue := getBitsFromBytes(value)

	var base64StringBuilder strings.Builder
	for i := 0; i < len(bitsFromValue); i += 6 {
		var base64Char byte
		for j := 0; j < 6; j++ {
			if i+j < len(bitsFromValue) {
				base64Char |= byte(bitsFromValue[i+j] << (5 - j))
			}
		}
		base64StringBuilder.WriteByte(base64Chars[base64Char])
	}

	return base64StringBuilder.String()
}

func main() {
	hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	fmt.Printf("Hex: %s\n", hexString)

	decodedHexStringBytes := decodeHexString(hexString)

	fmt.Printf("Decoded Hex: %s\n", string(decodedHexStringBytes))

	encodedBase64String := encodeToBase64(decodedHexStringBytes)

	fmt.Printf("Encoded Base64: %s\n", encodedBase64String)
}
