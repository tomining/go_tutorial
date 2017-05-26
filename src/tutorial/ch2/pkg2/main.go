package main

import (
	"fmt"
	"tutorial/ch2/lib"
)

var v rune

func init() {
	v = '1'
}

func main() {
	fmt.Println(lib.IsDigit(v))	//lib에 있는 IsDigit() 호출
}

func IsDigit(c int32) bool {	//사용안됨
	return '0' <= c && c <= '9'
}