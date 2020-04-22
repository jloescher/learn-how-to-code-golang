package main

import "fmt"

func main() {
	j := 38 // Initialize j variable with value of 38

	/* Using format print, print the number in decimal, binary, and hex.
	 * %d displays decimal value
	 * %b displays binary value
	 * %#x displays hex value
	 * \t inserts a tab character
	 * \n inserts a new line
	 */
	fmt.Printf("%d\t%b\t%#x\n", j, j, j)
}
