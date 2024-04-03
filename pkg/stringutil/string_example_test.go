package stringutil_test

import (
	"fmt"

	"github.com/toolvox/utilgo/pkg/stringutil"
)

func ExampleString() {
	// A complex string containing ASCII, Japanese Kanji and Hiragana, Greek characters, and emojis
	complexStr := stringutil.String([]rune("Go🚀は素晴らしい! Κόσμος👍"))

	// Display initial string
	fmt.Println("Original string:", complexStr.String())

	// Reverse the string rune-wise and print each character
	// Demonstrating manual rune manipulation
	reversedRunes := make([]rune, 0, complexStr.Len())
	for i := complexStr.Len() - 1; i >= 0; i-- {
		reversedRunes = append(reversedRunes, complexStr[i])
	}
	fmt.Println("Reversed by runes:", string(reversedRunes))

	// Extract "は素晴らしい!" using ByteSlice by calculating byte positions manually
	// Assuming known positions here for simplicity
	extract := complexStr.ByteSlice(6, 25) // Note: Adjust byte positions for your specific string
	fmt.Println("Extracted segment (by bytes):", extract.String())

	// Find and extract the Greek and emoji part, "Κόσμος👍", using rune positions
	// This assumes knowledge of the string structure
	greekPart := complexStr.Slice(-7, -1)
	fmt.Println("Greek part (by runes):", greekPart.String())

	// Demonstrating conversion back to bytes and manipulating bytes
	byteData := greekPart.Bytes()
	fmt.Printf("Greek part in bytes (hex): % x\n", byteData)

	// Output:
	// Original string: Go🚀は素晴らしい! Κόσμος👍
	// Reversed by runes: 👍ςομσόΚ !いしら晴素は🚀oG
	// Extracted segment (by bytes): は素晴らしい!
	// Greek part (by runes): Κόσμος👍
	// Greek part in bytes (hex): ce 9a cf 8c cf 83 ce bc ce bf cf 82 f0 9f 91 8d
}

func ExampleString_String() {
	str := stringutil.String([]rune("Hello, 世👍界👎"))
	fmt.Println(str.String())
	// Output: Hello, 世👍界👎
}

func ExampleString_Len() {
	str := stringutil.String([]rune("Hello, 世👍界👎"))
	fmt.Println(str.Len())
	// Output: 11
}

func ExampleString_Bytes() {
	str := stringutil.String([]rune("Hello, 世👍界👎"))
	fmt.Printf("% x\n", str.Bytes())
	// Output: 48 65 6c 6c 6f 2c 20 e4 b8 96 f0 9f 91 8d e7 95 8c f0 9f 91 8e
}

func ExampleString_ByteLen() {
	str := stringutil.String([]rune("Hello, 世👍界👎"))
	fmt.Println(str.ByteLen())
	// Output: 21
}

func ExampleString_ByteSlice() {
	str := stringutil.String([]rune("Hello, 世👍界👎"))
	sliced := str.ByteSlice(0, 5)
	fmt.Println(sliced.String())
	// Output: Hello
}

func ExampleString_Slice() {
	str := stringutil.String([]rune("Hello, 世👍界👎"))
	sliced := str.Slice(8, 10)
	fmt.Println(sliced.String())
	// Output: 👍界
}
