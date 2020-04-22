package main

import "fmt"

// Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.

func main() {
	favSport := "Bicycling"

	switch favSport {
	case "Football":
		fmt.Println("Your favorite sport is Football.")
	case "Baseball":
		fmt.Println("Your favorite sport is Baseball.")
	case "Basketball":
		fmt.Println("Your favorite sport is Basketball.")
	case "Hockey":
		fmt.Println("Your favorite sport is Hockey.")
	default:
		fmt.Printf("Your favorite sport %v is not popular.", favSport)
	}
}
