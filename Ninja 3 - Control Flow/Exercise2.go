package main

import "fmt"

// Print every rune code point of the uppercase alphabet three times.

func main() {

	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for x := 0; x < 3; x++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
