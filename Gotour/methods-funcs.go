package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//func (v Vertex) Abs() float64 // method
func Abs(v Vertex) float64 { //original function
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
	v := Vertex{3, 4}
	//fmt.Println(v.Abs()) //method call
	fmt.Println(Abs(v)) //original function call
}
