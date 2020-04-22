package main

import "fmt"

/*
create a func with the identifier foo that
- takes in a variadic parameter of type int
- pass in a value of type []int into your func (unfurl the []int)
- returns the sum of all values of type int passed in
create a func with the identifier bar that
- takes in a parameter of type []int
- returns the sum of all values of type int passed in
*/

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := foo(xi...)
	fmt.Println("Printing from foo()", sum)

	sum2 := bar(xi)
	fmt.Println("Printing from bar()", sum2)
}

func foo(x ...int) int {
	var total int
	for _, v := range x {
		total += v
	}
	return total
}

func bar(x []int) int {
	var total int
	for _, v := range x {
		total += v
	}
	return total
}
