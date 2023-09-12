package main

import "fmt"

func main() {
	var x int = 5 // immutable
	y := x
	y = 7
	fmt.Println(x, y) // => 5 7

	var x1 []int = []int{3, 4, 5} // mutable
	y1 := x1
	y1[0] = 100
	fmt.Println(x1, y1) // => [100 4 5] [100 4 5]

	var x2 map[string]int = map[string]int{"hello": 3} // mutable
	y2 := x2
	y2["y"] = 100
	fmt.Println(x2, y2) // => map[hello:3 y:100] map[hello:3 y:100]

	var x3 [2]int = [2]int{3, 4} // immutable
	y3 := x3
	y3[0] = 100
	fmt.Println(x3, y3) // => [3 4] [100 4]
}
