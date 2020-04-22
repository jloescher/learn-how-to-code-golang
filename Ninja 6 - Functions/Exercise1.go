package main

import "fmt"

/*
Hands on exercise
create a func with the identifier foo that returns an int
create a func with the identifier bar that returns an int and a string
call both funcs
print out their results
*/

func main() {
	f := foo()
	fmt.Println(f)
	b, b2 := bar()
	fmt.Println(b, b2)
}

func foo() int {
	return 43
}

func bar() (int, string) {
	return 43, "This is from bar."
}
