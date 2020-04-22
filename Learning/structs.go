package main

import (
	"fmt"
)

type Person struct {
	firstName string
	lastName  string
	age       int
}

type SecretAgent struct {
	Person // Pull in Person fields
	ltk    bool
}

func main() {
	p1 := Person{
		firstName: "James",
		lastName:  "Bond",
		age:       32,
	}

	p2 := Person{
		firstName: "Miss",
		lastName:  "Moneypenny",
		age:       27,
	}

	sa1 := SecretAgent{
		Person: Person{
			firstName: "James",
			lastName:  "Bond",
			age:       32,
		},
		ltk: true,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(sa1)

	fmt.Println(p1.firstName, p1.lastName, p1.age) // Accessing fields by dot notation
	fmt.Println(p2.firstName, p2.lastName, p2.age)
	/*
		Using dot notation you can access the fields using the fully qualified version seen below in the firstName field.
		Or since fields are promoted automatically you can access through dot notation the same way as other as seen in the
		lastName and aqe fields. Below.
	*/
	fmt.Println(sa1.Person.firstName, sa1.lastName, sa1.age, sa1.ltk)

}
