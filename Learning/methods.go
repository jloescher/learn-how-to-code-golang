package main

import "fmt"

/*
The syntax of a function is the following:
func (r receiver) identifier( "parameter(s)" "example:" x int "or" x ...string  ) ( "return(s)" "example:" bool int string ) { code }
Note: Words in quotes are just for clarity, not part of the syntax.
*/

type Person struct {
	firstName string
	lastName  string
}

type SecretAgent struct {
	Person
	ltk bool
}

func (s SecretAgent) speak() {
	fmt.Println("I am", s.firstName, s.lastName)
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
	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
}
