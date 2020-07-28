package main

import "fmt"

func main() {
	var i int
	j := i //int

	a := 42 //int
	b := 3.142 //float64
	c := 0.867 + 0.5i //complex128
	d := "omg" //string

	fmt.Printf("%T %T %T %T %T", j, a, b, c, d)
}
