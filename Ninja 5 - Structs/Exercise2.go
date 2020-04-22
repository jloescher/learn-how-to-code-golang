package main

import "fmt"

/*
Take the code from the previous exercise,
then store the values of type person in a map with the key of last name.
Access each value in the map. Print out the values, ranging over the slice.
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

	m := map[string]Person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	fmt.Println(m["Nguyen"])
	fmt.Println(m["Loescher"])

	for _, v := range m {
		fmt.Printf("%v %v likes the following ice cream flavors:", v.firstName, v.lastName)
		for _, v := range v.iceCreamFlavors {
			fmt.Println(v)
		}
	}
}
