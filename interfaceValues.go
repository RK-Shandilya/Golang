package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

type F float64

// type T implements interface I
func (t T) M() {
	fmt.Println(t.S)
}

// type F implements interface I
func (f F) M() {
	fmt.Println(f)
}

func main() {
	var v I = T{"Hello"}
	describe(v)
	v.M()

	var f I = F(math.Pi)
	describe(f)
	f.M()
}

// interface returns a tuple(value, type)
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
