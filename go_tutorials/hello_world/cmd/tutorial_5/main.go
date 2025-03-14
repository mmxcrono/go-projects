package main

import (
	"fmt"
	"strings"
)

// Go uses utf-8 for strings
// What about emoji's?
// Use more bits
// utf-32 can be a lot of memory
// utf-8 is variable length
func main() {
	// Create a var with a string that has utf-8 characters
	var myString string = "Hello, 世界"

	for i, v := range myString {
		fmt.Printf("%#U starts at byte position %d\n", v, i)
		fmt.Printf("Byte value %v, %T\n\n", myString[i], myString[i])
	}

	// Stringbuilder because go strings are immutable

	var stringBuilder strings.Builder

	var strSlice = []string{"s", "t", "r", "i", "n", "g"}

	for i := range strSlice {
		stringBuilder.WriteString(strSlice[i])
	}

	var catStr = stringBuilder.String()
	fmt.Println(catStr)
}