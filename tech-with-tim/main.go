package main

import "fmt"

func main() {
	var x []int = []int{3, 4}
	y := x
	y[0] = 100

	fmt.Println(x, y)

}
