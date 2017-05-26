package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	accessUnicodeStringInvalidWay()
	accessUnicodeStringValidWay()
	checkStringLen()
}

func accessUnicodeStringInvalidWay() {
	s1 := "hello"
	s2 := "안녕하세요"	//unicode 문자열 접근시 주의해야 한다.
	fmt.Printf("s1: %c %c %c %c %c\n", s1[0], s1[1], s1[2], s1[3], s1[4])
	fmt.Printf("s2: %c %c %c %c %c\n", s2[0], s2[1], s2[2], s2[3], s2[4])
}

func accessUnicodeStringValidWay() {
	s1 := "hello"
	s2 := "안녕하세요"
	r1 := []rune(s1)
	r2 := []rune(s2)	//rune array로 변환시 index로 정상접근 가능
	fmt.Printf("s1: %c %c %c %c %c\n", r1[0], r1[1], r1[2], r1[3], r1[4])
	fmt.Printf("s2: %c %c %c %c %c\n", r2[0], r2[1], r2[2], r2[3], r2[4])
}

func checkStringLen() {
	s1 := "hello"
	s2 := "안녕하세요"

	fmt.Println(len(s1))
	fmt.Println(len(s2))
	fmt.Println(utf8.RuneCountInString(s2))
	fmt.Println(len([]rune(s2)))
}