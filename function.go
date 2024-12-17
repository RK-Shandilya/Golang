package main

import "fmt"

// sub func with int return type
func sub(x int, y int) int {
	return x - y
}

// swap function return two values
func swap(x, y string) (string, string) {
	return y, x
}

// if type is same we can use type only in last parameter
func add(x, y int) int {
	return x + y
}

// naked return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("sub", sub(10, 9))
	fmt.Println("sum", add(6, 7))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
}
