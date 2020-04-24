package main

import (
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	Age       int
}

type byAge []Person

func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type byName []Person

func (bn byName) Len() int           { return len(bn) }
func (bn byName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn byName) Less(i, j int) bool { return bn[i].FirstName < bn[j].FirstName }

func main() {
	p1 := Person{"James", 32}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4}

	fmt.Println("Unsorted:", people)
	sort.Sort(byAge(people)) // sort by age
	fmt.Println("Sorted by Age:", people)
	sort.Sort(byName(people)) // sort by name
	fmt.Println("Sorted by Name:", people)

}
