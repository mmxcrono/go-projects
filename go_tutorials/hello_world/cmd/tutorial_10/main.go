package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Generics

type contactInfo struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type purchaseInfo struct {
	Name   string `json:"name"`
	Price  float32 `json:"price"`
	Amount uint `json:"amount"`
}

func main() {
	// var intSlice = []int{1, 2, 3, 4, 5}
	// var float32Slice = []float32{1.1, 2.2, 3.3, 4.4, 5.5}
	// fmt.Println(sumSlice(intSlice))
	// fmt.Println(sumSlice(float32Slice))
	// intSlice = []int{}
	// fmt.Println(isEmpty(intSlice))
	// fmt.Println(isEmpty(float32Slice))


	contacts := loadJson[contactInfo]("contacts.json")
	purchases := loadJson[purchaseInfo]("purchases.json")
	fmt.Println(contacts)
	fmt.Println(purchases)
}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

func loadJson [T contactInfo | purchaseInfo] (filePath string) []T {
	// Load json and return the struct
	data, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	var loaded = []T{}
	err = json.Unmarshal(data, &loaded)

	if err != nil {
		log.Fatal(err)
	}

	return loaded
}