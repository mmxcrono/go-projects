package main

import (
	"fmt"
	"time"
)

func main() {
	var intArr [3]int32
	fmt.Println(intArr[0])
	// [0, 0, 0] is default
	// 0-indexed

	// See memory allocation
	fmt.Println(&intArr[0])

	// Initialize array
	var initArr [3]int32 = [3]int32{1, 2, 3}
	fmt.Println(initArr)

	// Infer
	initArr2 := [3]int32{1, 2, 3}
	fmt.Println(initArr2)
	
	// Infer with dot syntax
	initArr3 := [...]int32{1, 2, 3}
	fmt.Println(initArr3)

	// Slices
	var intSlice []int32 = []int32{4,5,6}
	fmt.Println(intSlice)
	fmt.Printf("length is %v capacity %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
	fmt.Printf("length is %v capacity %v\n", len(intSlice), cap(intSlice))

	// Append multiple using spread operator
	intSlice = append(intSlice, []int32{8, 9, 10}...)
	fmt.Println(intSlice)
	fmt.Printf("length is %v capacity %v\n", len(intSlice), cap(intSlice))

	// Create slice with make command, length of slice and capacity of slice
	var intSlice2 []int32 = make([]int32, 3, 5)
	fmt.Println(intSlice2)
	fmt.Printf("length is %v capacity %v\n", len(intSlice2), cap(intSlice2))

	// Maps
	var intMap = map[string]uint8{"one": 1, "two": 2}
	fmt.Println(intMap)

	var age, ok = intMap["one"]
	fmt.Println(age, ok)

	age, ok = intMap["three"]
	fmt.Println(age, ok)

	delete(intMap, "one")
	fmt.Println(intMap)

	// No order is preserved when iterating over map
	for key, value := range intMap {
		fmt.Println(key, value)
	}

	for index, value := range initArr2 {
		fmt.Println(index, value)
	}

	var i int = 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	var k int
	for {
		if k >= 10 {
			break
		}
		fmt.Println(k)
		k++
	}

	for x:=range 10 {
		fmt.Println(x)
	}

	for x:=0; x<10; x++ {
		fmt.Println(x)
	}

	var t0 = time.Now()

	var sum int = 0

	for i:=range 100000000 {
		sum += i
	}

	fmt.Println(sum)

	var t1 = time.Since(t0)
	fmt.Println(t1)
}