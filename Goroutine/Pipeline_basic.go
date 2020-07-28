/* 3단계의 파이프라인 예시 */

package main

import "fmt"

/*
 stage.1: gen
 정수형 리스트를 받아서, 받은 리스트 안의 정수를 내보내는 채널로 변환한다.
 리스트의 모든 값이 전송되고 나면 채널을 닫는다.
*/
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

/*
 stage.2: sq
 채널로부터 정수값을 받고, 받은 정수의 제곱을 내보내는 채널을 리턴한다.
 in 채널이 닫히고 out 채널로 모든 값을 내보낸 후에 out 채널을 닫는다.
*/
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
 stage.3: main
 파이프라인을 설정하고 마지막 단계를 실행한다.
 채널이 닫힐 때까지 두번째 단계(sq) 에서 값을 수신하고, 각 값을 출력한다.
*/
func main() {
	// 파이프라인 설정 및 output 소비
	for n := range sq(sq(gen(2, 3))) {
		fmt.Println(n) // 16, 81
	}
}
