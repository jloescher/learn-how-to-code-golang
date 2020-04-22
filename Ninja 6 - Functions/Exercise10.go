package main

import "fmt"

/*
Closure is when we have “enclosed” the scope of a variable in some code block.
For this hands-on exercise, create a func which “encloses” the scope of a variable:
*/

func main() {
	x := 10
	f := func() int {
		x := 30
		return x
	}
	fmt.Println("This is x:", x)
	fmt.Println("This is y:", f())
}
