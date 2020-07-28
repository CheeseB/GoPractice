package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func add_str(str string) (n, m string) {
	n = str + " world"
	m = str + " neighbor"
	return
}

func main() {
	fmt.Println(split(17))
	a, b := add_str("hello")
	fmt.Println(a, b)
}
