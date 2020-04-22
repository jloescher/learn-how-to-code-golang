package main

import "fmt"

/*
Write down what these print:
fmt.Println(true && true)
fmt.Println(true && false)
fmt.Println(true || true)
fmt.Println(true || false)
fmt.Println(!true)
*/

func main() {
	fmt.Printf("true && true evaluates to %v\n", true && true)
	fmt.Printf("true && false evaluates to %v\n", true && false)
	fmt.Printf("true || true evaluates to %v\n", true || true)
	fmt.Printf("true || false evaluates to %v\n", true || false)
	fmt.Printf("!true evaluates to %v\n", !true)
}
