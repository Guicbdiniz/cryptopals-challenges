# Convert hex to base64

## hex

Hex is basically a numeral system that uses 16 characters (0-9 and A-F).

But Hex can also refer to Base16, which is a encoding schema that turns bytes to a string.

It works by splitting the binary input value into groups of 4 bits (4 bits can only represent 16 different values) and converting each group into its corresponding hex character.

## base64

Base64 is another encoding schema that turns binary data to a 64 characters (a-z, A-Z, 0-9, / and +) string. It is useful for carrying data into channel that only deals with text data.

It basically works by splitting the binary input value into groups of 6 bits (6 bits can only represent 64 different values) and converting each group into its corresponding Base64 character.

After combining those characters, padding `=` characters can be added so that the length of the final string is a multiple of 4.

## Solution

The easiest way to do it is simply using Go's standard library [base64 package](https://pkg.go.dev/encoding/base64) and (hex package)[https://pkg.go.dev/encoding/hex].

But, of course, the main idea is to understand the encoding schemas better. So I must create my own functions to deal with the encoding and decoding of strings.

### Decoding Hex

To decode a Hex string, I must get each character, get its binary representation in 4 bits and join in with the next 4 bits. 8 bits will represent an ASCII char.

### Encoding base64

To encode the ASCII string to a Base64 string, I must get each 6 bits from the byte slice, turn it into a Base64 character and join all of them.

How to get the 6 bits from the byte slice? I've no idea.

The best way to do it is probably to use 3 bytes. Those 3 bytes have 24 bits. I can separate those 24 bits into 4 groups of 6 bits.

Actually, why not simply get a slice of ints with all bits from the bytes array and them use them 6 by 6?

The solution is definitely related to the bitwise operators.

- To get the bits from one byte, I must first shift the bits to the right (7 for the first bit in the left, 6 for the second, 5 for the third...) and "remove" the rest of the bits by doing & 1 (which is 00000001).
  - So 10101010 >> 7 -> 00000001 & 1 -> 00000001 which is 1.

With this, I can create a int array with the bits from the byte slice.

After that, I must use the values from the slice 6 by 6 to create a byte and use this byte to get the correct character from the base 64 table.
