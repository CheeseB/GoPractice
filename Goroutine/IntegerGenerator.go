package main

import "fmt"

func IntegerGenerator(n int) <-chan int {
	next := 0

	intStream := make(chan int) // local variable of IntegerGenerator
	go func() {                 // goroutine and closure
		defer close(intStream) // close channel after the goroutine ends
		for next < n {
			next++
			intStream <- next
		}
	}()
	return intStream // return channel
}

func main() {
	for i := range IntegerGenerator(10) { // iterates until the channel is close
		fmt.Println(i)
	}
}
