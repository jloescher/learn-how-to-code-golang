package main

import (
	"fmt"
)

/*
The syntax of a function is the following:
func (r receiver) identifier( "parameter(s)" "example:" x int "or" x ...string  ) ( "return(s)" "example:" bool int string ) { code }
Note: Words in quotes are just for clarity, not part of the syntax.
*/

func main() {
	s1 := fmt.Sprintf("for the item in index")
	s2 := fmt.Sprintf("we are now adding")
	s3 := fmt.Sprintf("to the total which is now")
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	x := sum(s1, s2, s3, xi...) // Since the function is expecting a series of Type int it is necessary to use the ... to expand the slice to Type int
	fmt.Println("The sum of x is:", x)
	bar(2)
}

func sum(s1, s2, s3 string, x ...int) int { // accepts unlimited items of Type int and returns one int additionally variadic parameters much be the final parameter
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println(s1, i, s2, v, s3, sum)
	}
	return sum
}

func bar(x int) { // accepts only one item of Type int
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
