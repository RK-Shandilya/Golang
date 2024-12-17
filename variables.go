package main

import "fmt"

// package level variable
var i, j int = 1, 2

func main() {
	// function level variable
	var c, python, java bool

	// variable declaration : type can be omitted and compiler will infer it from the initializer
	var x, y = 4, 5

	// shorthand variable declaration : type can be omitted and compiler will infer it from the initializer
	a, b := 6, 7
	fmt.Println(i, j, c, python, java, x, y, a, b)
}
