package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for int8
	// rune  - alias for int32
	// float32 float64
	// complex64 complex128

	// Usaing var
	// var name = "Vijay Deepak"
	name := "Vj Deepak"
	size := 2.3
	name, email := "Vijay Deepak", "email@email.com"
	var age int32 = 24
	const isCool = true
	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", size)
}
