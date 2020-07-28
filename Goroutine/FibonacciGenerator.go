package main

import "fmt"

func Fib(n int) chan int {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		x, y := 1, 1
		for k := 0; k < n; k++ {
			intStream <- x
			x, y = y, x+y
		}
	}()
	return intStream
}

func main() {
	for x := range Fib(10) {
		fmt.Println(x)
	}
}
