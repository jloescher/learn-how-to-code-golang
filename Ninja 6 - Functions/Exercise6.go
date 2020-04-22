package main

import "fmt"

// Build and use an anonymous func

func main() {
	func(s string, i int) {
		fmt.Println(s, i)
	}("This is an anonymous function.", 38)
}
