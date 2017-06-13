package main

import "fmt"

type quantity int
type numberMap map[string]int
type rect struct {
	width float64
	height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func main() {
	receiverCallBy()
	receiverCallBy2()
	omitReceiver()
	methodFuntionalExpression()
}

func (q quantity) greaterThan(i int) bool {
	return int(q) > i
}

func (q *quantity) increment() {
	*q++
}

func (q *quantity) decrement() {
	*q--
}

func receiverCallBy() {
	q := quantity(3)
	q.increment()
	fmt.Printf("Is q(%d) greater than %d? %t\n", q, 3, q.greaterThan(3))
	q.decrement()
	fmt.Printf("Is q(%d) greater than %d? %t\n", q, 3, q.greaterThan(3))
}

func (m numberMap) add(key string, value int) {
	m[key] = value
}

func (m numberMap) remove(key string) {
	delete(m, key)
}

//Map은 참조 타입으로 *(포인터)를 활용하지 않는다.
func receiverCallBy2() {
	m := make(numberMap)
	m["one"] = 1
	m["two"] = 2
	m.add("three", 3)
	fmt.Println(m)
	m.remove("two")
	fmt.Println(m)
}

func (rect) new() rect {	//리시버 변수 생략
	return rect{}
}

func omitReceiver() {
	r := rect{}.new()
	fmt.Println(r)
}

func (r *rect) resize(w, h float64) {
	r.width += w
	r.height += h
}

func methodFuntionalExpression() {
	r := rect{3, 4}
	fmt.Println("area: ", r.area())
	r.resize(10, 10)
	fmt.Println("area: ", r.area())

	// areaFn == func(rect) float64
	areaFn := rect.area
	// resizeFn == func(*rect, float64, float64)
	resizeFn := (*rect).resize

	fmt.Println("area: ", areaFn(r))
	resizeFn(&r, -10, -10)
	fmt.Println("area: ", areaFn(r))
}