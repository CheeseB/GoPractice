/* Explicit cancellation 예시 */

package main

import (
	"fmt"
	"sync"
)

func gen(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}()
	return out
}

func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out) // 함수 리턴 후에 out 채널이 닫히도록 defer 사용
		for n := range in {
			select {
			case out <- n * n:
				// in 채널로부터 값이 들어오면 값의 제곱을 out 채널에 송신
			case <-done:
				return // done 채널이 닫히면 바로 리턴
			}
		}
	}()
	return out
}

func merge(done <-chan struct{}, ch ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// c나 done 채널이 닫힐때 까지만 값을 송신함
	output := func(c <-chan int) {
		defer wg.Done() // 함수 리턴 후에 wg.Done이 호출되도록 defer 사용
		for n := range c {
			select {
			case out <- n:
				// c 채널로부터 값이 들어오면 그대로 out 채널에 송신
			case <-done:
				return // done 채널이 닫히면 바로 리턴
			}
		}
	}
	wg.Add(len(ch))

	for _, c := range ch {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	done := make(chan struct{})
	defer close(done) // main함수 리턴 후에 done 채널이 닫힘

	in := gen(done, 2, 3)

	c1 := sq(done, in)
	c2 := sq(done, in)

	out := merge(done, c1, c2)
	fmt.Println(<-out)
}
