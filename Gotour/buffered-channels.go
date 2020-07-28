package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2
	// first-in-first-out
	fmt.Println(<- ch) // 1
	fmt.Println(<- ch) // 2

	/* fatal error: all goroutines are asleep - deadlock!
	ch <- 1
	ch <- 2
	ch <- 3
	//
	fmt.Println(<- ch)
	fmt.Println(<- ch)
	fmt.Println(<- ch)
	*/
}
