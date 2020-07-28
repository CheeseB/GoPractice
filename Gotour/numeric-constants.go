package main

import (
	"fmt"
	"math"
)

//high-precision value
//takes the type needed by its context
const (
	Big = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x * 10 + 1}
func needFloat(x float64) float64 {
	return x * 0.1
}
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	//can't use v here
	return lim
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
