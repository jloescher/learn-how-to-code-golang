package main

import (
	"encoding/json"
	"fmt"
)

/*
Starting with this code, unmarshal the JSON into a Go data structure
Hint: https://mholt.github.io/json-to-go/
*/

type Person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

type People interface {
	Says()
	Introduce()
}

func (p Person) Says() { // Method on Person
	p.Introduce()
	for _, v := range p.Sayings {
		fmt.Printf("\t\t%v\n", v)
	}
}

func (p Person) Introduce() { // Method on Person
	fmt.Println("Hello, my name is", p.First, p.Last, "and I am", p.Age, "years old. I like to say these phrases:")
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println("JSON Format:", s)

	bs := []byte(s)

	var people []Person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("In Go Struct:", people)

	fmt.Printf("%T\n", people)
	for i, v := range people {
		fmt.Println("Person:", i)
		v.Says()
	}
}
