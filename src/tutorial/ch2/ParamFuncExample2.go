package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	f := func(c rune) bool {
		return unicode.Is(unicode.Hangul, c)	//c가 한글이면 true를 반환
	}

	fmt.Println(strings.IndexFunc("Hello, 월드", f))
	fmt.Println(strings.IndexFunc("Hello, world", f))
}