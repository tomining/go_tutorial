package main

import (
	"fmt"
	"tutorial/ch2/lib"
)

func main() {
	fmt.Println("Hello")
	fmt.Println(lib.IsDigit('1'))	//lib 패키지의 IsDigit 함수 사용
	fmt.Println(lib.IsDigit('a'))
	fmt.Println(lib.IsSpace('\t'))
}