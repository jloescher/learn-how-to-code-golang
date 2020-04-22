package main

import "fmt"

func main() {
	f := func() { // Assigning a function to a variable.
		fmt.Println("My first func expression")
	}
	f()

	f1 := func(x int) {
		fmt.Println("the year big brother started watching:", x)
	}
	f1(1984)
}
