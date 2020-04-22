package main

import "fmt"

func main() {
	defer foo() // defer delays the execution to the end of the calling function
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
