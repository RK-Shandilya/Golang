package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) String() string {
	return fmt.Sprintf("%v (%v years)", p.name, p.age)
}

func main() {
	a := person{"Arthur Dent", 42}
	z := person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
