package main

import "fmt"

const (
	a     = 42 // unTyped constant
	b int = 38 // Typed constant
)

func main() {
	fmt.Println(a, b) // print constants
}
