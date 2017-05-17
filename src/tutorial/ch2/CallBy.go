package main

import "fmt"

func main() {
	i := 10
	increaseCallByValue(i)	//i는 증가 안 함
	fmt.Println(i)
	increaseCallByReference(&i) 	//i가 증가
	fmt.Println(i)
}

func increaseCallByValue(i int) {
	i = i + 1
}

func increaseCallByReference(i *int)  {
	*i = *i + 1
}