package main

import "fmt"

func main() {
	four()
}

type Point struct { // always begin with a capital letter
	x int32
	y int32
}

func one() {
	// combine existing types to make useful types
	var p1 Point = Point{1, 2}        // 1 is assigned to x, 2 is assigned to y
	var p2 Point = Point{x: -3, y: 4} // alternative way to construct a new `Point`
	fmt.Println(p1)
	fmt.Println(p2)
}

func changeX(pt *Point) {
	pt.x = 100

}

func changeX2(pt Point) {
	pt.x = 100
}

func two() {
	p1 := &Point{y: 3}
	fmt.Println(p1) // => {0, 3}
	changeX(p1)     // make change in a function when a pointer is passed
	fmt.Println(p1) // => {100, 3}

	p2 := Point{y: 3}
	fmt.Println(p2) // => {0, 3}
	changeX2(p2)    // doesn't make any changes in a function when a value is passed
	fmt.Println(p2) // => {0, 3}
}

func three() {
	p1 := &Point{y: 3}
	p1.x = 8
	fmt.Println(p1) // custom type (struct) itself is mutable
}

type Circle struct {
	radius float64
	center *Point
}

type Circle2 struct {
	radius float64
	*Point // member doesn't have to need its name, the name of custom type itself become the name of the member
}

type Circle3 struct {
	radius float64
	int    // above rule can be applied to also non-custom type like int, float, etc.
}

func four() {
	c1 := Circle{4.56, &Point{4, 5}}
	fmt.Println(c1)        // => {4.56 0xc0000b4000}
	fmt.Println(c1.center) // => {4, 5}

	var c2 Circle2 = Circle2{4.13, &Point{3, 1}}
	fmt.Println(c2.x, c2.y) // => 3 1	the member name will be omitted
	fmt.Println(c2.Point)

	var c3 Circle3 = Circle3{2.31, 2}
	fmt.Println(c3.int) // => 2		it can be referenced like this
}
