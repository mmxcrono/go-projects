package main

import (
	"fmt"
	"unicode/utf8"
)

func getNumber() int {
	return 3
}

func main() {
	fmt.Println("Hello, World!")

	var intNum uint8 = 255
	fmt.Println(intNum)

	var floatNum float32 = 3.14
	fmt.Println(floatNum)

	var floatAdd = floatNum + float32(intNum)

	fmt.Println(floatAdd)

	var myString string = "Hello Ĉ" + " World!"
	fmt.Println(myString)
	fmt.Println(len(myString))
	fmt.Println(utf8.RuneCountInString(myString))

	var defaultInt int
	fmt.Println(defaultInt)

	myVar := 3.14
	fmt.Println(myVar)

	var1, var2 := 2, 3
	fmt.Println(var1, var2)


	var3 := getNumber()

	fmt.Println(var3)

	var var4 int = getNumber() + 1
	fmt.Println(var4)

	const myConst string = "Dev"
	fmt.Println(myConst)
}