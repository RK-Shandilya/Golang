// Explains how to create custom error types in Go

package main

// The fmt package is used for formatted I/O, and the time package provides functionality for measuring and displaying time
import (
	"fmt"
	"time"
)

// MyError is a struct that holds information about an error, including the time it occurred (When) and a description of the error (What).
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

// function returns a MyError
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
