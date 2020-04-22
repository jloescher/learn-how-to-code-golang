package main

import "fmt"

// Create and use an anonymous struct.

func main() {
	p1 := struct {
		firstName string
		lastName  string
		age       int
	}{
		firstName: "Joanthan",
		lastName:  "Loescher",
		age:       38,
	}

	fmt.Printf("%v %v is %v", p1.firstName, p1.lastName, p1.age)
}
