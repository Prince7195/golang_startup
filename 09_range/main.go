package main

import "fmt"

func main() {
	ids := []int{33, 44, 55, 66, 67, 12}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i+1, id)
	}

	// without index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// range with maps
	emails := map[string]string{"vj": "vj@vvv.co", "vj2": "vj2@vvv.co"}
	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	// get keys
	for k := range emails {
		fmt.Println("Name: " + k)
	}
}
