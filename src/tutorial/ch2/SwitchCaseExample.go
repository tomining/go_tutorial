package main

import "fmt"

func main() {
	c := 'a'
	switch {
	case '0' <= c && c <= '9':
		fmt.Println("%c은(는) 숫자입니다.", c)
	case 'a' <= c && c <= 'z':
		fmt.Println("%c은(는) 소문자입니다.", c)
	case 'A' <= c && c <= 'Z':
		fmt.Println("%c은(는) 대문자입니다.", c)
	}
}