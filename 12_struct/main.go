package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	city      string
	age       int
	gender    string
}

// Greeting Method (value reciever)
func (p Person) greet() string {
	return "Hello, My name is " + p.firstName + " " + p.lastName + " and I'm " + strconv.Itoa(p.age)
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Vijay", lastName: "Deepak", city: "chennai", age: 24, gender: "M"}

	// Alternative
	// person1 := Person{"Vj", "Deepak", "chennai", 23, "M"}

	// fmt.Println(person1)
	// fmt.Println(person1.firstName)
	// person1.age = 24
	// fmt.Println(person1)

	fmt.Println(person1.greet())
}
