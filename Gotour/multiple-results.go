package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func tri_swap(x, y, z string) (string, string, string) {
	return z, x, y
}

func main() {
	a, b := swap("world!", "Hello")
	c, d, e := tri_swap("is", "me.", "This")
	fmt.Println(a, b)
	fmt.Println(c, d, e)
}
