package main

import "fmt"

func main() {
	//initPointer()
	assignReferenceAddress()
	assignReferenceAddress2()
	initPointerNew()
}

func initPointer() {
	type rect struct {
		w, h float64
	}

	var pRect *rect
	var pInt *int
	var pFloat *float64
	//var pComplex = *complex128()
	fmt.Println(&pRect, &pInt, &pFloat)	//초기화 하지 않은 포인터는 임의 영역을 할당 받는듯
}

func assignReferenceAddress() {
	var p *int

	i := 1
	p = &i

	fmt.Println(i)
	fmt.Println(&i)	//메모리 주소 값
	fmt.Println(*p)
	fmt.Println(p)	//메모리 주소 값
}

func assignReferenceAddress2() {
	var p *int
	var pp **int

	i := 1
	p = &i
	pp = &p
	fmt.Println(i, *p, **pp) 	// 1 1 1

	i += 1
	fmt.Println(i, *p, **pp)	// 2 2 2

	*p++
	fmt.Println(i, *p, **pp)	// 3 3 3

	**pp++
	fmt.Println(i, *p, **pp)	// 4 4 4
}

func initPointerNew() {
	p := new(int)
	*p = 1
	fmt.Println(p, *p)
}