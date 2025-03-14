package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
	ownerInfo owner
	owner
	int
}

type owner struct {
	name string
}

// Interfaces
// Interfaces are a way to define a set of methods that a type must have
type engine interface {
	milesLeft() uint8
}


func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func canMakeIt(e engine, miles uint8) bool {
	return e.milesLeft() >= miles
}

// struct can have methods
// methods are defined outside the struct like normal functions
func (e gasEngine) drive() {
	fmt.Printf("%v is driving with %d gallons left\n", e.ownerInfo.name, e.gallons)
}

func main() {
	var myEngine gasEngine = gasEngine{25 ,15, owner{"John"}, owner{"Engi"}, 10}
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo.name,  myEngine.name, myEngine.int)

	// Anonymous struct
	// not reusable struct
	var myEngine2 = struct {
		mpg uint8
		gallons uint8
	}{25, 15}

	fmt.Println(myEngine2)
	myEngine.drive()
	var ok bool = canMakeIt(myEngine, 100)
	if ok {
		fmt.Println("You can make it")
	} else {
		fmt.Println("You can't make it")
	}
}