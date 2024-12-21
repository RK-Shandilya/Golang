package main

import (
	"fmt"
	"math"
)

type MyFloat float64

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// methods on non-structs type
func (f MyFloat) Pos() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{4, 3}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt(2))
	fmt.Println(f.Pos())
}
