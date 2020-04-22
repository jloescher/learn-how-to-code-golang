package main

import "fmt"

// Create a program that uses a switch statement with no switch expression specified.

func main() {
	switch { // default is true
	case 2 == 4:
		fmt.Println("This is false, won't print")
	case 2 < 4:
		fmt.Println("This is true, it will print")
	}
}
