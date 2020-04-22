package main

import "fmt"

/*
Using iota, create 4 constants for the NEXT 4 years. Print the constant values.
*/

const (
	y1 = iota + 2020
	y2
	y3
	y4
)

func main() {
	fmt.Println(y1, y2, y3, y4)
}
