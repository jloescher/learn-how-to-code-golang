package main

import (
	"fmt"
	"time"
)

/*
A “callback” is when we pass a func into a func as an argument. For this exercise,
pass a func into a func as an argument
*/

func main() {
	years := []int{1981, 1993}
	ages := calculateAges(getCurrentYear, years...)
	fmt.Println(ages)
}

func calculateAges(currentYear func() int, birthYears ...int) []int {
	var ages []int
	for _, v := range birthYears {
		ages = append(ages, currentYear()-v)
	}
	return ages
}

func getCurrentYear() int {
	currentTime := time.Now()
	return currentTime.Year()
}
