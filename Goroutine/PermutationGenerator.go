package main

import "fmt"

func permgen(ch chan []string, a []string, n int) {
	if n == 1 {
		ch <- append(a[:0:0], a...) // dup slice 'a'
		return
	}
	for i := 0; i < n; i++ {
		a[n-1], a[i] = a[i], a[n-1]
		permgen(ch, a, n-1)
		a[n-1], a[i] = a[i], a[n-1]
	}
}

func Permutations(a []string) <-chan []string {
	permStream := make(chan []string)
	go func() {
		defer close(permStream)
		permgen(permStream, a, len(a))
	}()
	return permStream
}

func main() {
	for p := range Permutations([]string{"사과", "배", "포도"}) {
		fmt.Println(p)
	}
}
