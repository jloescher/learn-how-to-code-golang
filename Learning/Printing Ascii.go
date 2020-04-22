package main

import "fmt"

func main() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("%v in Hex is %#x and in Ascii is: %#U\n", i, i, i)
	}
}
