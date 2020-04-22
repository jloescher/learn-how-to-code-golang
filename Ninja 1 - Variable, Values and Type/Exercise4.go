package main

import "fmt"

// Create Type with underlying Type of int
type myType int

// Create variable of Type that was created with the underlying Type of int
var x myType

func main() {
	// Print value of x variable
	fmt.Println(x)
	// Print Type of x variable
	fmt.Printf("%T\n", x)
	// Set x variable to 42
	x = 42
	// Print x variable
	fmt.Println(x)
}
