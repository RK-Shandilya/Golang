package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type Myfloat float64

type Abser interface {
	Abs() float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// This method means type Myfloat implements the interface Abser,
// but we don't need to explicitly declare that it does so.
func (f1 Myfloat) Abs() float64 {
	return math.Abs(float64(f1))
}

func main() {
	var a Abser
	v := Vertex{3, 4}
	f := Myfloat(-5)

	a = f
	a = &v

	fmt.Println(a)
	fmt.Println(a.Abs())

}
