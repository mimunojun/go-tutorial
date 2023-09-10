package main

import "fmt"

func changeFirst(slice []int) {
	slice[0] = 1000
}

func main() {
	var x []int = []int{3,4,5}
	fmt.Println(x)
	changeFirst(x)
	fmt.Println(x)
}
