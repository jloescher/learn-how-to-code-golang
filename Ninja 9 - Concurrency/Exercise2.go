package main

import "fmt"

/*
This exercise will reinforce our understanding of method sets:
 - create a type person struct
 - attach a method speak to type person using a pointer receiver
  - *person
 - create a type human interface
  - to implicitly implement the interface, a human must have the speak method
 - create func “saySomething”
  - have it take in a human as a parameter
  - have it call the speak method
 - show the following in your code
  - you CAN pass a value of type *person into saySomething
  - you CANNOT pass a value of type person into saySomething
 - here is a hint if you need some help
  - https://play.golang.org/p/FAwcQbNtMG
*/

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p *Person) Speak() {
	fmt.Println("Hello my name is", p.FirstName, p.LastName, "and I am", p.Age, "years old.")
}

type Human interface {
	Speak()
}

func main() {
	p1 := Person{
		"James",
		"Bond",
		32,
	}
	SaySomething(&p1)
}

func SaySomething(h Human) {
	h.Speak()
}
