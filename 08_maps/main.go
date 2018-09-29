package main

import "fmt"

// func main() {
// 	// Define map (Key value pairs)
// 	emails := make(map[string]string)

// 	// Assigning values
// 	emails["Vijay"] = "vijay@gmail.com"
// 	emails["Deepak"] = "deepak@gmail.com"
// 	emails["mike"] = "mike@gmail.com"
// 	fmt.Println(emails)
// 	fmt.Println(len(emails))
// 	fmt.Println(emails["Vijay"])

// 	// Delete from map
// 	delete(emails, "mike")
// 	fmt.Println(emails)
// }

func main() {
	// Declare map and add key value pair

	emails := map[string]string{"vj": "vj@vvv.co", "vj2": "vj2@vvv.co"}
	emails["Vijay"] = "vijay@gmail.com"
	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["vj"])
}
