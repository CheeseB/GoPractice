package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
	/* closing is only necessary
	when the receiver must be told
	there are no more values coming,
	such as to terminate a range loop */
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c { // repeat until c is closed
		fmt.Println(i)
	}
}
