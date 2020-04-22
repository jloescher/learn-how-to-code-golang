package main

import (
	"fmt"
	"time"
)

/*
Create a func which returns a func
assign the returned func to a variable
call the returned func
*/

func main() {
	x := myReturnedFunc()
	age := x(1981)
	fmt.Println(age)
}

func myReturnedFunc() func(int) int {
	return func(yearOfBirth int) int {
		currentTime := time.Now()
		return currentTime.Year() - yearOfBirth
	}
}
