package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x * x + y * y))
	z := uint(f)
	//var h uint = f //error

	fmt.Println(x, y, z)
}
