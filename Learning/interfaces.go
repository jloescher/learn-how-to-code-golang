package main

/*
Values can have more than one Type
*/

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

type SecretAgent struct {
	Person
	ltk bool
}

func (s SecretAgent) speak() {
	fmt.Println("I am", s.firstName, s.lastName, "- the SecretAgent speak")
}

func (p Person) speak() {
	fmt.Println("I am", p.firstName, p.lastName, "- the Person speak")
}

type human interface {
	speak() // Any Type with the method speak() is also of Type human
}

func bar(h human) {
	switch h.(type) {
	case Person:
		fmt.Println("I was passed into barrrr", h.(Person).firstName) // Assertion
	case SecretAgentt:
		fmt.Println("I was passed into barrrr", h.(SecretAgent).firstName)
	}
	fmt.Println("I was passed into bar", h)
}

func main() {
	sa1 := SecretAgent{
		Person: Person{
			firstName: "James",
			lastName:  "Bond",
		},
		ltk: true,
	}

	sa2 := SecretAgent{
		Person: Person{
			firstName: "Miss",
			lastName:  "Moneypenny",
		},
		ltk: false,
	}

	p1 := Person{
		firstName: "Dr.",
		lastName:  "Yes",
	}
	fmt.Println(sa1)
	fmt.Printf("%T\n", sa1)
	sa1.speak()
	sa2.speak()
	fmt.Println(p1)
	p1.speak()
	bar(sa1)
	bar(sa2)
	bar(p1)
}
