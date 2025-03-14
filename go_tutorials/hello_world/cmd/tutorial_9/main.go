package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

const MAX_CHICKEN_PRICE float32 = 12.0
const MAX_TOFU_PRICE float32 = 11.0

func main() {
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"wallmart.com", "costco.com", "wholefoods.com"}

	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)
}

func checkTofuPrices(website string, tofuChannel chan string) {
	for {
		time.Sleep(time.Second)
		var tofuPrice = rand.Float32() * 20
		if tofuPrice <= MAX_TOFU_PRICE {
			tofuChannel <- website
			break
		}
	}
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	select {
		case website := <-chickenChannel:
			fmt.Printf("Text sent: Chicken is available at %v\n", website)
		case website := <-tofuChannel:
			fmt.Printf("Email sent: Tofu is available at %v\n", website)
	}
}


// Channels are way to enable goroutines to pass data back and forth with each other
// Channels are typed, so you can only pass data of the same type
// Listen for data on a channel with <-
// Send data to a channel with ->
// Channels are blocking, so if you try to send data to a channel that is full, the program will block until the channel is empty

// Buffered channels can be created with make(chan int, 10)
// This will allow the channel to hold 10 values before blocking

// func main() {
// 	var c = make(chan int, 5)
// 	go process(c)
// 	// Iterate over channel with range
// 	for i := range c {
// 		fmt.Println(i)
// 		time.Sleep(time.Second)
// 	}
// }

// func process(c chan int) {
// 	// Can defer close with this 
// 	defer close(c)
// 	for i := range 10 {
// 		c <- i
// 	}
// 	// Close the channel when done to avoid deadlock
// 	// close(c)
// 	fmt.Println("Exiting process")
// }