package main

/*
Create a for loop using this syntax
for condition { }
Have it print out the years you have been alive.
*/

var (
	birthYear   = 1981
	currentYear = 2020
)

func main() {
	for birthYear <= currentYear {
		println(birthYear)
		birthYear++
	}
}
