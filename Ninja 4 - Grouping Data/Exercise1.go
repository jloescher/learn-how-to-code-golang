package main

import "fmt"

/*
Using a COMPOSITE LITERAL:
create an ARRAY which holds 5 VALUES of TYPE int
assign VALUES to each index position.
Range over the array and print the values out.
Using format printing
print out the TYPE of the array
*/

func main() {
	arr := [5]int{10, 20, 30}

	for i, v := range arr {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}

	arr[3] = 40 // Add or Change values in an array.
	arr[4] = 50

	for _, v := range arr {
		fmt.Println(v)
	}
	fmt.Printf("%T", arr)
}
