package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 { // pass by value
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// methods on non-structs type
type MyFloat float64

func (f MyFloat) Pos() float64 { // method on MyFloat type - pass by value
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// pointer receiver
func (v *Vertex) Scale(f float64) { // pass by reference
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{4, 3}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt(2))
	fmt.Println(f.Pos())

	// pointer reciever
	v.Scale(10)
	fmt.Println(v.Abs())
}
