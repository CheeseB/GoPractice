package main

import "fmt"

func move(ch chan [2]string, from, to string) {
	var result [2]string
	result[0] = from
	result[1] = to

	ch <- result
}

func hanoiTop(ch chan [2]string, n int, from, to, by string) {
	if n == 1 {
		move(ch, from, to)
	} else {
		hanoiTop(ch, n-1, from, by, to)
		move(ch, from, to)
		hanoiTop(ch, n-1, by, to, from)
	}
}

func Hanoi(n int, from, to, by string) <-chan [2]string {
	hanoiStream := make(chan [2]string)
	go func() {
		defer close(hanoiStream)
		hanoiTop(hanoiStream, n, from, to, by)
	}()
	return hanoiStream
}

func main() {
	for move := range Hanoi(3, "A", "B", "C") {
		fmt.Println(move[0], "->", move[1])
	}
}
