package main

import "fmt"

func main() {
	r := rect{3, 4}
	describe(r)
	display(r)
}

type shaper interface {
	area() float64
}

func describe(s shaper) {
	fmt.Println("area : ", s.area())
}

// 익명 인터페이스로도 지정할 수 있다.
func display(s interface{ show() float64})  {
	fmt.Println("area : ", s.show())
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) show() float64 {
	return r.width * r.height
}