package main

import "fmt"

/*
Create a user defined struct with
- the identifier “person”
- the fields:
 - first
 - last
 - age
attach a method to type person with
- the identifier “speak”
- the method should have the person say their name and age
create a value of type person
call the method from the value of type person
*/

type Person struct {
	first string
	last  string
	age   int
}

func (p Person) speak() {
	fmt.Printf("Hello my name is %v %v and my age is %v.", p.first, p.last, p.age)
}

func main() {
	p1 := Person{
		first: "Jonathan",
		last:  "Loescher",
		age:   38,
	}
	p1.speak()
}
