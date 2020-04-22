package main

import "fmt"

// Define variables in package level. These are zero valued.
var x int
var y string
var z bool

func main() {
	// Print defines variables in function
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
