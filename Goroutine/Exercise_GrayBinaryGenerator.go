package main

import "fmt"

func grayCode(ch chan []int, code []int, n, index, reverse int) {
	if index == n {
		ch <- append(code[:0:0], code...)
		return
	}

	code[index] = reverse
	grayCode(ch, code, n, index+1, 0)

	code[index] = 1 - reverse
	grayCode(ch, code, n, index+1, 1)
}

func GrayBinaryGenerator(n int) <-chan []int {
	grayStream := make(chan []int)
	code := make([]int, n)

	go func() {
		defer close(grayStream)
		grayCode(grayStream, code, n, 0, 0)
	}()
	return grayStream
}

func main() {
	for g := range GrayBinaryGenerator(3) {
		fmt.Println(g)
	}
}
