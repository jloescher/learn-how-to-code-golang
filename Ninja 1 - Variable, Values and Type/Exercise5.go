package main

import "fmt"

// Create Type with underlying Type of int
type myType int

// Create variable of Type that was created with the underlying Type of int
var x myType

// Declare a variable y of Type int
var y int

func main() {
	fmt.Println(x)        // Print value of x variable
	fmt.Printf("%T\n", x) // Print Type of x variable
	x = 42                // Set x variable to 42
	fmt.Println(x)        // Print value of x variable
	y = int(x)            // Convert x to Type int and set value to y variable which is Declared as int
	fmt.Println(y)        // Print value of y variable
	fmt.Printf("%T\n", y) // Print Type of y variable
}
