package main

import "fmt"

type Vertex struct {
	X, Y float64
}

var m map[string]Vertex

func map_literal() {
	var p = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(p)
}

func mutating_map() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	m["Google"] = Vertex{37.42202, -122.08408}
	fmt.Println(m)
	fmt.Println(m["Bell Labs"])

	map_literal()
	mutating_map()
}
