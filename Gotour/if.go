package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return  fmt.Sprint(math.Sqrt(x)) // turn into string
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
