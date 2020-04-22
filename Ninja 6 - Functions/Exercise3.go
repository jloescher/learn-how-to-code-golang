package main

import "fmt"

// Use the “defer” keyword to show that a deferred func runs after the func containing it exits.

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	defer foo(xi...) // defer causes foo to run last
	bar(xi)
}

func foo(x ...int) {
	var total int
	for _, v := range x {
		total += v
	}
	fmt.Println("Printing from foo()", total)
}

func bar(x []int) {
	var total int
	for _, v := range x {
		total += v
	}
	fmt.Println("Printing from bar()", total)
}
