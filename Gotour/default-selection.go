package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	time.Tick(시간): 시간 간격으로 반복 실행
	time.After(시간): 시간이 지나면 ready 되는 채널을 리턴
	time.Sleep(시간): 시간동인 일시정지
	*/
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <- tick:
			fmt.Println("tick.")
		case <- boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
