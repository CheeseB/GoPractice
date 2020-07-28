package main

import (
	"fmt"
	"time"
)

// in the "fmt" package
/*
type error interface {
	Error() string
}
*/

type MyError struct {
	When time.Time
	What string
}

// *MyError type implements fmt.error
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	// err == nil -> success
	// err != nil -> fail
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
