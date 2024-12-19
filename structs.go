package main

import "fmt"

type person struct {
	name string
	age  int
}

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p1 = &Vertex{1, 2} // has type *Vertex
)

func main() {
	// declaring a struct
	v := person{"rudra", 15}
	fmt.Println(v)

	// accessing the fields using dot notation
	fmt.Println(v.name)
	fmt.Println(v.age)

	// pointers to structs
	var p *person
	p = &v
	fmt.Println(*p)
	fmt.Println((*p).name)
	fmt.Println(p.name) // same as above, you don't have to dereference and easier to read

	// string literals
	fmt.Println(v1, *p1, v2, v3)
}
