package main

import (
	"fmt"
)

var c, python, java bool

func add(x, y int) int {
	return x + y
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))

	var i int
	fmt.Println(i, c, python, java)
}
