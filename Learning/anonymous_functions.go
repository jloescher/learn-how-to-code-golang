package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("Anonymous function ran.")
	}()

	func(x int) {
		fmt.Println("The meaning of life:", x)
	}(42)
}

func foo() {
	fmt.Println("Foo ran")
}
