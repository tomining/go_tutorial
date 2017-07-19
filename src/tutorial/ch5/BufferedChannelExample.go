package main

import "fmt"

func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	//c <- 3	// buffer가 full인 상태에서 채널에 전송하므로 오류
	go func() {c <- 3}()	//goroutine은 채널에 메시지를 전송할 수 있을 때까지 대기 후 전송
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}