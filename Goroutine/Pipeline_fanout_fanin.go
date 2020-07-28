/* Fan-out, Fan-in 예시 */

package main

import (
	"fmt"
	"sync"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

/*
 merge 함수는 각 인바운드 채널에 대해 값을 단일 아웃바운드 채널로 복사하는 고루틴을 실행함으로써
 여러 채널을 단일 채널로 변환한다.
 그리고 모든 output 고루틴이 시작되고 나면 아웃바운드 채널을 닫기 위해 한번 더 고루틴을 실행한다.
 닫힌 채널에 전송을 하는 것은 패닉을 유발하므로
 채널을 닫기 전에 모든 전송이 완료되었는지 확인하는 것이 중요하다.
 sync.WaitGroup을 통해서 이와 같은 동기화를 간단히 할 수 있다.
*/

func merge(ch ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// c 채널이 닫힐 때까지 c 채널의 값을 out 채널로 송신한 다음 wg.Done을 호출하는 함수 정의
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done() // 대기중인 고루틴의 수행이 종료됨을 알려줌
	}
	wg.Add(len(ch)) // WaitGroup에 대기중인 고루틴 개수를 인풋채널 갯수만큼 추가

	// 각 인풋 채널(ch)에 대해 output 고루틴을 실행
	for _, c := range ch {
		go output(c)
	}

	// 모든 output 고루틴이 끝나면 out 채널을 닫는 고루틴을 실행
	// 이 고루틴은 반드시 wg.Add 호출이 끝난 후 시작해야 함
	go func() {
		wg.Wait() // 모든 고루틴이 종료될 때까지 대기
		close(out)
	}()
	return out
}

func main() {
	in := gen(2, 3)

	// 동일하게 in 채널을 읽는 sq작업을 두개의 고루틴에 분배함 (fan-out)
	c1 := sq(in)
	c2 := sq(in)

	// merge 함수를 통해 c1, c2 채널을 단일 채널로 병합 (fan-in)
	for n := range merge(c1, c2) {
		fmt.Println(n) // 4 9 or 9 4
	}
}
