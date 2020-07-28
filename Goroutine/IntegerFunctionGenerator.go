package main

import "fmt"

func IntegerFunctionGenerator(f func(x int) int, n int) <-chan int {
	next := 0
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 0; i < n; i++ {
			intStream <- f(next)
			next++
		}
	}()
	return intStream
}

func main() {
	f := func(x int) int {
		return x * x
	}

	for i := range IntegerFunctionGenerator(f, 5) {
		fmt.Println(i)
	}
}
