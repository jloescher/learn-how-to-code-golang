package main

import "fmt"

// Print out the remainder (modulus) which is found for each number between 10 and 100 when it is divided by 4.

func main() {
	for i := 10; i <= 100; i++ {
		m := i % 4
		fmt.Printf("When %v is devided by 4 the remaining value is %v\n", i, m)
	}
}
