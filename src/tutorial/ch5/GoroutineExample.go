package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main 함수 시작", time.Now())

	go long()
	go short()

	/**
	 * goroutine은 메인 스레드가 종료되면 goroutine도 종료된다.
	 * 아래 예제에서 5초를 2초로 변경하면 long() goroutine이 종료되기 전에 프로그램이 종료되는 것을 확인할 수 있다.
	 */
	time.Sleep(5 * time.Second)	//5초 대기
	fmt.Println("main 함수 종료", time.Now())
}

func long() {
	fmt.Println("long 함수 시작", time.Now())
	time.Sleep(3 * time.Second)	//3초 대기
	fmt.Println("long 함수 종료", time.Now())
}

func short() {
	fmt.Println("short 함수 시작", time.Now())
	time.Sleep(1 * time.Second)	//1초 대기
	fmt.Println("short 함수 종료", time.Now())
}