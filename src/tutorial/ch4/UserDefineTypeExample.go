package main

import "fmt"

type quantity int
type dozen []quantity
type costCalculator func(quantity, float64) float64

func main() {
	initUserDefineType()
	userDescribe()
	initStruct()
	docTyping()
}

func initUserDefineType() {
	var d dozen
	for i := quantity(1); i <= 12; i++ {
		d = append(d, i)
	}
	fmt.Println(d)
}

func describe(q quantity, price float64, c costCalculator) {
	fmt.Printf("quantity: %d, price: %0.0f, cost: %0.0f\n", q, price, c(q, price))
}

func userDescribe() {
	var offBy10Percent, offBy1000Won costCalculator

	offBy10Percent = func(q quantity, price float64) float64 {
		return float64(q) * price * 0.9
	}
	offBy1000Won = func(q quantity, price float64) float64 {
		return float64(q) * price - 1000
	}

	describe(3, 10000, offBy10Percent)
	describe(3, 10000, offBy1000Won)
}

type rect struct {
	width float64
	height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func initStruct() {
	r := rect{3, 4}
	fmt.Println("area: ", r.area())
}

type shaper interface {
	area() float64
}

func describe2(shaper shaper) {
	fmt.Println("area: ", shaper.area())
}

func docTyping() {
	r := rect{3, 4}
	describe2(r)	//rect 구조체가 area()를 구현하고 있으므로 shaper 타입 인터페이스로 처리 됨
}