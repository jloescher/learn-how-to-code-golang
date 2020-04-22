package main

import "fmt"

/*
Create a for loop using this syntax
for { }
Have it print out the years you have been alive.
*/

var (
	birthYear   = 1981
	currentYear = 2020
)

func main() {
	for {
		if birthYear > currentYear {
			break
		}
		fmt.Println(birthYear)
		birthYear++
	}
}
