package main

import (
	"fmt"
	"strings"
)

func slices_are_refrences() {
	names := []string{"John", "Mary", "David", "Ringo"}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	names[0] = "Mohini"

	b[0] = "Rudra"
	fmt.Println(a, b, names)
}

func struct_slice_defaluts() {
	p := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println(p)

	p = p[1:4]
	fmt.Println(p)

	p = p[:2]
	fmt.Println(p)

	p = p[1:]
	fmt.Println(p)
}

func slices_len_cap() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)
	// s = append(s, 10, 12) // to append the elements 10 and 12
	s = append(s[1:3])
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func nil_slices() {
	var sl []int
	fmt.Println(sl, len(sl), cap(sl))
	if sl == nil {
		fmt.Println("nil")
	}
}

func slices_of_slices() {
	// creating a tic-toc board
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var s []int = primes[1:4]
	fmt.Println(s)

	// using make keyword
	highScore := make([]int, 4) // memory allocation happens here and it returns a slice but of exactly 4 elements
	highScore[0] = 100
	highScore[1] = 200
	highScore[2] = 300
	highScore[3] = 400
	// highScore[4] = 500 // we can't do this because we have only 4 slots and the memory is allocated for 4 slots
	highScore = append(highScore, 556, 557, 558) // It will allocate new memory and copy the old data to the new memory and append the values.

	fmt.Println(highScore)

	// slices_are_refrences()
	// struct_slice_defaluts()
	// slices_len_cap()
	// nil_slices()
	slices_of_slices()
}
