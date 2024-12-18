package main

import (
	"fmt"
	"math"
)

func sqrt(i float64) string {
	if i < 0 {
		return sqrt(-i) + "i"
	}
	return fmt.Sprint(math.Sqrt(i))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim) // printf is used to print the value in the specific format
	}

	return lim
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
