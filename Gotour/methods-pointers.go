package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//value receiver
//operates on a copy of the original value
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//pointer receiver
//can modify the original value
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
