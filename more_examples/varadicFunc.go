package main

import "fmt"

var groosery []string

func addGroosery(a ...string) {
	for _, d := range a {
		groosery = append(groosery, d)
	}
}

func main() {
	fmt.Println("List of Grooseries: ")
	addGroosery("Glass")
	addGroosery("Bread")
	addGroosery("Fruits")
	addGroosery("Biscuts")
	for i, g := range groosery {
		fmt.Println(i, "", g)
	}
	fmt.Println("Capacity", cap(groosery))
}
