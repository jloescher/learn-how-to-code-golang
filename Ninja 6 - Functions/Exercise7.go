package main

import "fmt"

// Assign a func to a variable, then call that func

func main() {
	f := func(age int) {
		fmt.Println("This is a function assigned to a variable.")
		fmt.Println("My age is:", age)
	}

	f(38)
}
