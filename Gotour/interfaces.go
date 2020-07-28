package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64 // MyFloat, *Vertex
}

func main() {
	var a Abser

	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f // MyFloat receiver
	a = &v // *Vertex receiver
	// a = v // Vertex does NOT implement Abser

	fmt.Println(a.Abs())
}

type MyFloat float64

type Vertex struct {
	X, Y float64
}

func (f MyFloat) Abs() float64 { // receiver of Abs: MyFloat
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs() float64 { // receiver of Abs: *Vertex
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}
