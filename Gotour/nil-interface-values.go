package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i) // nil interface (<nil>, <nil>)
	// i.M() // runtime error
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
