package main

import (
	"fmt"
)

func main() {
	g := 42 == 38 // equal operator
	h := 42 <= 38 // less than or equal operator
	i := 42 >= 38 // greater than or equal operator
	j := 42 != 38 // not equal operator
	k := 42 < 38  // less than operator
	l := 42 > 38  // greater than operator

	fmt.Println(g, h, i, j, k, l) // Print variables
}
