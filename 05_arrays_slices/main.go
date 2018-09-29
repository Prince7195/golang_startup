package main

import "fmt"

func main() {
	// Declare Arrays
	// var fruitsList [2]string

	// // Assign value to array
	// fruitsList[0] = "Apple"
	// fruitsList[1] = "Mango"

	// // Declare and Assign
	// fruitsList := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitsList)
	// fmt.Println(fruitsList[0])

	fruitSlice := []string{"Apple", "Orange", "Mango", "Grape"}
	fmt.Println(fruitSlice)      // [Apple Orange Mango Grape]
	fmt.Println(len(fruitSlice)) // 4
	fmt.Println(fruitSlice[1:2]) // [Orange]
	fmt.Println(fruitSlice[1:3]) // [Orange Mango]
}
