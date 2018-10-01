package main

import (
	"fmt"
	"strconv"
)

// Long method

// type Person struct {
// 	firstName string
// 	lastName  string
// 	city      string
// 	age       int
// 	gender    string
// }

// Short method
type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

// Greeting Method (value reciever)
func (p Person) greet() string {
	return "Hello, My name is " + p.firstName + " " + p.lastName + " and I'm " + strconv.Itoa(p.age)
}

// hasBirthDay method (pointer reciever)
func (p *Person) hasBirthDay() {
	p.age++
}

// getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "F" {
		p.lastName = spouseLastName
	} else {
		return
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Vijay", lastName: "Deepak", city: "chennai", age: 24, gender: "M"}

	// Alternative
	person2 := Person{"Samantha", "Johnson", "Coimbatore", "F", 20}

	// fmt.Println(person1)
	// fmt.Println(person1.firstName)
	// person1.age = 24
	// fmt.Println(person1)

	person1.hasBirthDay()
	person1.getMarried("Vj")

	person2.getMarried("Willium")

	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
}
