package main

import "fmt"

func main() {
	a := 10
	b := &a
	fmt.Println(a, b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	// Use * to read the value from address
	fmt.Println(*b)
	fmt.Println(b)
	fmt.Println(*&a)

	// Change value with pointers
	*b = 22
	fmt.Println(a)
}
