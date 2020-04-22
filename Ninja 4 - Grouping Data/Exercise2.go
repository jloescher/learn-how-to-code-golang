package main

import "fmt"

/*
Using a COMPOSITE LITERAL:
create a SLICE of TYPE int
assign 10 VALUES
Range over the slice and print the values out.
Using format printing
print out the TYPE of the slice
*/

func main() {
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	for _, v := range s {
		fmt.Println(v)
	}
	fmt.Printf("%T", s)
}
