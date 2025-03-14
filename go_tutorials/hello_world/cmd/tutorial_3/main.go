package main

import (
	"errors"
	"fmt"
)


func main() {
	const numerator int = 11
	const denominator int = 2

	fmt.Printf("Divide %v by %v\n", numerator, denominator)
	
	var result, remainder, err = intDivision(numerator, denominator)

	// if (err != nil) {
	// 	fmt.Println(err.Error())
	// } else if remainder == 0 {
	// 	fmt.Printf("result is %v\n", result)
	// } else {
	// 	fmt.Printf("result is %v with remainder %v\n", result, remainder)
	// }

	switch {
		case err != nil:
			fmt.Println(err.Error())
		case remainder == 0:
			fmt.Printf("result is %v\n", result)
		default:
			fmt.Printf("result is %v with remainder %v\n", result, remainder)
	}

	switch remainder {
		case 0:
			fmt.Println("The division was exact")
		case 1,2:
			fmt.Println("The remainder was close")
		default:
			fmt.Println("The remainder was far")
	}
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error

	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	
	var result int = numerator / denominator
	var remainder int = numerator % denominator

	return result, remainder, err
}