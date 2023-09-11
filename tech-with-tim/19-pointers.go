package main

import "fmt"

func main() {
	three()
}

func one() {
	x := 7          // type of `x` is `int`
	fmt.Println(&x) // puts the memory location of `x` in hex
	// the reference is called "pointer"

	y := &x           // type of `y` is `*int`, called "pointer"
	fmt.Println(x, y) // => 7 0xc0000b4000
	*y = 8
	fmt.Println(x, y) // => 8 0xc0000b4000
}

func two() {
	toChange := "hello"
	changeValue(&toChange)
	fmt.Println(toChange) // => changed!

	toChange2 := "hello"
	changeValue2(toChange2)
	fmt.Println(toChange2) // => hello
}

func three() {
	var toChange string = "hello"
	var pointer *string = &toChange
	fmt.Println(pointer)  // => 0xc00009e230
	fmt.Println(*pointer) // => hello
}

func changeValue(str *string) { // pass me a pointer, not an actual value
	*str = "changed!" // changing the original string
}

func changeValue2(str string) {
	str = "changed!"
}
