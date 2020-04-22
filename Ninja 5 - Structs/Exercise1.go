package main

import "fmt"

/*
Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
first name
last name
favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.
*/

type Person struct {
	firstName       string
	lastName        string
	iceCreamFlavors []string
}

func main() {
	p1 := Person{
		firstName:       "Jason",
		lastName:        "Nguyen",
		iceCreamFlavors: []string{"Superman", "Crazy Vanilla"},
	}

	p2 := Person{
		firstName:       "Jonathan",
		lastName:        "Loescher",
		iceCreamFlavors: []string{"Raspberry", "Rocky Road Raspberry", "Blueberry Cheesecake"},
	}

	fmt.Println(p1.firstName, p1.lastName, "likes these ice cream flavors:")
	for i, v := range p1.iceCreamFlavors {
		fmt.Printf("%v: %v\n", i+1, v)
	}

	fmt.Println(p2.firstName, p2.lastName, "likes these ice cream flavors:")
	for i, v := range p2.iceCreamFlavors {
		fmt.Printf("%v: %v\n", i+1, v)
	}
}
