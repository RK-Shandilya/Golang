package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	// type conversion
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	// type inversion
	var i int = 10
	var f2 float64 = float64(i)
	fmt.Println(i, f2)
}
