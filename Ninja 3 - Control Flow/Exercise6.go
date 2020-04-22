package main

import "fmt"

const currentYear = 2020

func main() {
	if currentYear >= 2025 {
		fmt.Println("Your Rich!")
	} else {
		fmt.Println("Your Poor!")
	}
}
