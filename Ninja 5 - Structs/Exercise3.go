package main

import "fmt"

/*
Create a new type: vehicle.
- The underlying type is a struct.
- The fields:
	- doors
	- color
- Create two new types: truck & sedan.
	- The underlying type of each of these new types is a struct.
	- Embed the “vehicle” type in both truck & sedan.
	- Give truck the field “fourWheel” which will be set to bool.
	- Give sedan the field “luxury” which will be set to bool. solution
Using the vehicle, truck, and sedan structs:
	- using a composite literal, create a value of type truck and assign values to the fields;
	- using a composite literal, create a value of type sedan and assign values to the fields.
Print out each of these values.
Print out a single field from each of these values.
*/

type Vehicle struct {
	doors int
	color string
}

type Truck struct {
	Vehicle
	fourWheel bool
}

type Sedan struct {
	Vehicle
	luxury bool
}

func main() {
	v1 := Truck{
		Vehicle: Vehicle{
			doors: 2,
			color: "Gun Metal",
		},
		fourWheel: true,
	}

	v2 := Sedan{
		Vehicle: Vehicle{
			doors: 4,
			color: "Blue",
		},
		luxury: true,
	}

	fmt.Println(v1)
	fmt.Println(v2)

	fmt.Println(v1.fourWheel)
	fmt.Println(v2.doors)
}
