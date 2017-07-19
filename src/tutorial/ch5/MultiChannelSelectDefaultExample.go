package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <- tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("Boom!")
			return
		default:	// tick, boom 채널에 메시지가 존재하지 않을 때 실행
			fmt.Println("       .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}