package main

import "fmt"

const currentYear = 2020

func main() {
	if currentYear >= 2025 {
		fmt.Println("Your Rich!")
	} else if (currentYear < 2025) && (currentYear > 2023) {
		fmt.Println("In two years or less you will be rich!")
	} else {
		fmt.Println("Your Poor!")
	}
}
