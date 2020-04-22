package main

import "fmt"

// Set variables in package level
var x = 42
var y = "James Bond"
var z = true

func main() {
	// Combine variables into one string
	s := fmt.Sprintf("%v %v %v", x, y, z)

	// Print string
	fmt.Println(s)

	// Print Type of s variable
	fmt.Printf("%T\n", s)
}
