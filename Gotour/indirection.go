package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) ScaleMethod(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunction(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main(){
	v := Vertex{1, 2}
	v.ScaleMethod(2)
	ScaleFunction(&v, 10)

	p := &Vertex{2, 1}
	p.ScaleMethod(2)
	ScaleFunction(p, 10)

	fmt.Println(v, p)
}
