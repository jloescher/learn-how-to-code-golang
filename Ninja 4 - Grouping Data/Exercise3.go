package main

import "fmt"

/*
Using the code from the previous example, use SLICING to create the following new slices which are then printed:
[42 43 44 45 46]
[47 48 49 50 51]
[44 45 46 47 48]
[43 44 45 46 47]
*/

func main() {
	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(s[:5])  // Prints everything to the 5th position not the 5th positions value
	fmt.Println(s[5:])  // Prints the 5th positions values and onward
	fmt.Println(s[2:7]) // Prints the 2nd position up to but not including the 7th position
	fmt.Println(s[1:6]) // Prints the 1st position up to but not including the 6th position
}
