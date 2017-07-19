package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main 함수 시작", time.Now())

	done := make(chan bool)
	/**
	 * short() 과 long() 의 수행 순서는 보장되지 않는다.
	 * 순서가 필요하다면 중간에 채널 수신 대기를 해야 한다.
	 */
	go short(done)
	go long(done)

	<- done
	<- done
	fmt.Println("main 함수 종료", time.Now())
}

func long(done chan bool) {
	fmt.Println("long 함수 시작", time.Now())
	time.Sleep(3 * time.Second)	//3초 대기
	fmt.Println("long 함수 종료", time.Now())
	done <- true
}

func short(done chan bool) {
	fmt.Println("short 함수 시작", time.Now())
	time.Sleep(1 * time.Second)	//1초 대기
	fmt.Println("short 함수 종료", time.Now())
	done <- true
}