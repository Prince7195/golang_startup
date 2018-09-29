package main

import "fmt"

func main() {
	x := 7
	y := 4

	// IF ELSE
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is greater than %d\n", x, y)
	}

	// ELSE IF
	color := "red"
	if color == "red" {
		fmt.Printf("Color is red")
	} else if color == "blue" {
		fmt.Printf("Color is blue")
	} else {
		fmt.Printf("Color Not mached blue or red")
	}

	switch color {
	case "red":
		fmt.Printf("Color is red")
	case "blue":
		fmt.Printf("Color is blue")
	default:
		fmt.Printf("Color Not mached blue or red")
	}
}
