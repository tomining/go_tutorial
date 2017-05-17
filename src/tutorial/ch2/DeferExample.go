package main

import "fmt"

func main() {
	f1()
}

func f1() {
	fmt.Println("f1 - start")
	defer f2()	//Java의 final 같은 개념
	fmt.Println("f1 - end")
}

func f2() {
	fmt.Println("f2 - deferred")
}