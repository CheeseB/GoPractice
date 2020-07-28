package main

import "fmt"

func combgen(ch chan []string, arr []string, combArr []int, n, r, index, target int) {
	if r == 0 {
		comb := make([]string, r)
		for i := 0; i < index; i++ {
			comb = append(comb, arr[combArr[i]])
		}
		ch <- comb
	} else if target == n {
		return
	} else {
		combArr[index] = target
		combgen(ch, arr, combArr, n, r-1, index+1, target+1)
		combgen(ch, arr, combArr, n, r, index, target+1)
	}
}

func Combinations(arr []string, r int) <-chan []string {
	combStream := make(chan []string)
	n := len(arr)
	combArr := make([]int, n)
	go func() {
		defer close(combStream)
		combgen(combStream, arr, combArr, n, r, 0, 0)
	}()
	return combStream
}

func main() {
	for c := range Combinations([]string{"사과", "배", "복숭아", "포도", "귤"}, 2) {
		fmt.Println(c)
	}
}
