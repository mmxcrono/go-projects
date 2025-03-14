package main

import "fmt"

// Pointers
func main() {
	// stores memory address
	// bits depends on os
	// pointer does not have memory address assigned
	// var p *int32
	
	var p *int32 = new(int32)
	// now it has a memory location assigned

	// we can dereference the pointer to get the memory address
	fmt.Printf("The value p points to is: %v\n", *p)
	fmt.Printf("The memory address p points to is: %v\n", p)

	var i int32
	fmt.Printf("The value of i is: %v\n", i)

	// If our pointer was not assigned a memory address, we would get a runtime error when we try to assign a value to it
	*p = 10
	fmt.Printf("The value p points to is now: %v\n", *p)

	// we can also assign the memory address of a variable to a pointer
	var otherPointer = &i

	fmt.Printf("The value of otherPointer is: %v\n", *otherPointer)
	
	// Slices point to arrays, so second variable pointing to the same array will modify that same array
	var slice = []int{1, 2, 3, 4, 5}
	var slice2 = slice
	slice2[0] = 100
	fmt.Printf("The value of slice is: %v\n", slice)

	// Pointers are good so we don't deal with copies of data
	// We can pass the memory address of a variable to a function to modify the original variable
}