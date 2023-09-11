package main

import (
	"fmt"
	"math"
)

func main() {
	one()
}

type shape interface { // define what `shape` has to have as a method, this is interface
	area() float64
}

type circle struct {
	radius float64
}

type rect struct {
	width  float64
	height float64
}

func (c *circle) area() float64 { // `circle` has `area` method which satisfies the requirements of `shape` interface, so it can be implemented as into `shape` object
	return math.Pi * c.radius * c.radius
}

func (r *rect) area() float64 {
	return r.width * r.height
}

func getArea(s shape) float64 {
	return s.area()
}

func one() {
	c1 := circle{4.5}
	r1 := rect{5, 7}
	shapes := []shape{&c1, &r1}

	for _, shape := range shapes {
		fmt.Println(shape.area())
	} // => 63.61725123519331 <br> 35

	for _, shape := range shapes {
		fmt.Println(getArea(shape))
	}
}
