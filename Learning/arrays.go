package main

import "fmt"

func main() {
	var x [2]string // array with length of 2
	x[0] = "Jonathan"
	x[1] = "Loescher"
	fmt.Println(x)
	fmt.Printf("%T", x)
}
