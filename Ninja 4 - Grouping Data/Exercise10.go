package main

import "fmt"

/*
Using the code from the previous example, delete a record from your map. Now print the map out using the “range” loop
*/

func main() {
	m := map[string][]string{ // Create a map with a key of Type string, with a value of Type []string which is a string slice
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	m[`loescher_jonathan`] = []string{`Jason`, `Prema`, `Bicycling`} // Adding to the map

	delete(m, `no_dr`) // Delete item by key

	for k, v := range m {
		fmt.Printf("Key: %v, %v\n", k, v)
		for i, v := range v {
			fmt.Printf("\t\tIndex: %v, Value: %v\n", i, v)
		}
	}
}
