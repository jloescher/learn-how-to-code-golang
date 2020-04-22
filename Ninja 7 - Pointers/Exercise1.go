package main

import "fmt"

/*
Create a value and assign it to a variable.
Print the address of that value.
*/

func main() {
	x := 10
	fmt.Println(x, &x) // using the & before the variable displays the address of the value
}
