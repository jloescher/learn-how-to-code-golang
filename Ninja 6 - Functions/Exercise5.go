package main

import (
	"fmt"
	"math"
)

/*
create a type SQUARE
create a type CIRCLE
attach a method to each that calculates AREA and returns it
- circle area= π r 2
- square area = L * W
create a type SHAPE that defines an interface as anything that has the AREA method
create a func INFO which takes type shape and then prints the area
create a value of type square
create a value of type circle
use func info to print the area of square
use func info to print the area of circle
*/

type Square struct {
	length float64
}

func (s Square) area() float64 {
	return s.length * s.length
}

type Rectangle struct {
	length, width float64
}

func (r Rectangle) area() float64 {
	return r.width * r.length
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}

type Shape interface {
	area() float64
}

func info(shape Shape) {
	fmt.Printf("%T: %v\n", shape, shape.area())
}

func main() {
	c1 := Circle{radius: 30}
	info(c1)

	s1 := Square{
		length: 10,
	}
	info(s1)

	r1 := Rectangle{
		length: 10,
		width:  20,
	}
	info(r1)

}
