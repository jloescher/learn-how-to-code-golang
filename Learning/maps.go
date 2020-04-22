package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	} // Create map with a key with a Type of string and a value with a Type of int
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Barnabas"])
	v, ok := m["Barnabas"] // Getting value which is stored in v and if the value exists which is stored as a boolean in ok
	fmt.Println(v)
	fmt.Println(ok) // Will be false as Barnabas is not a valid key in the map

	if v, ok := m["Barnabas"]; ok { // Only will run if ok is equal to true
		fmt.Println("This is the if print:", v)
	}

	m["Todd"] = 33 // Add Key of Todd with a value of 33

	// Looping over a map
	for k, v := range m {
		fmt.Printf("Key: %v and Value: %v\n", k, v)
	}

	delete(m, "Todd") // delete from map using key, it will delete a key that does not exist.

	fmt.Println("Deleted Todd from map.")
	for k, v := range m {
		fmt.Printf("Key: %v and Value: %v\n", k, v)
	}

	// correct way.
	if _, ok := m["James"]; ok { // underscore "_" tosses the value. Can't have unused variables in golang
		delete(m, "James")
	}

	fmt.Println("Deleted James from map.")
	for k, v := range m {
		fmt.Printf("Key: %v and Value: %v\n", k, v)
	}
}
