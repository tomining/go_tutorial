package main

import (
	"fmt"
)

func main() {
	fmt.Println("Array initialize...")
	initializeArray()
	fmt.Println("Slice initialize...")
	initializeSlice()
	overCapSlice()
	insertSlice()
	insertSlice2()
}

func initializeArray() {
	var a [5]int
	b := [3]int{1, 2, 3}
	c := [3]int{1, 2}
	d := [...]int{4, 5, 6, 7, 8}
	e := [3][3]int{
		{1, 2, 3},
		{3, 4, 5},	//요소를 여러 줄로 표기할 때, 요소의 마지막에 콤마를 붙여야 함
	}

	fmt.Printf("%-10T %d %v\n", a, len(a), a)
	fmt.Printf("%-10T %d %v\n", b, len(b), b)
	fmt.Printf("%-10T %d %v\n", c, len(c), c)
	fmt.Printf("%-10T %d %v\n", d, len(d), d)
	fmt.Printf("%-10T %d %v\n", e, len(e), e)
}

func initializeSlice() {
	var a []int
	b := []int{}
	c := []int{1, 2, 3}
	d := [][]int {
		{1, 2},
		{3, 4, 5},
	}
	e := make([]int, 0)
	f := make([]int, 3, 5)

	fmt.Printf("%-10T %d %d %v\n", a, len(a), cap(a), a)
	fmt.Printf("%-10T %d %d %v\n", b, len(b), cap(b), b)
	fmt.Printf("%-10T %d %d %v\n", c, len(c), cap(c), c)
	fmt.Printf("%-10T %d %d %v\n", d, len(d), cap(d), d)
	fmt.Printf("%-10T %d %d %v\n", e, len(e), cap(e), e)
	fmt.Printf("%-10T %d %d %v\n", f, len(f), cap(f), f)
}

func overCapSlice() {
	s := make([]int, 0, 3)
	for i := 0; i < 9; i++ {
		s = append(s, i)
		fmt.Printf("len: %d, cap: %d, %v\n", len(s), cap(s), s)
	}
}

func insertSlice() {
	s := []int{1, 2, 3, 4, 5}

	s = insert(s, []int{-3, -2}, 0)
	fmt.Println(s)

	s = insert(s, []int{0}, 2)
	fmt.Println(s)

	s = insert(s, []int{6, 7}, len(s))
	fmt.Println(s)
}

func insertSlice2() {
	s := []int{1, 2, 3, 4, 5}

	s = insert2(s, []int{-3, -2}, 0)
	fmt.Println(s)

	s = insert2(s, []int{0}, 2)
	fmt.Println(s)

	s = insert2(s, []int{6, 7}, len(s))
	fmt.Println(s)
}

func insert(s, new []int, index int) []int {
	return append(s[:index], append(new, s[index:]...)...)
}

func insert2(s, new []int, index int) []int {
	result := make([]int, len(s) + len(new))
	position := copy(result, s[:index])
	position += copy(result[position:], new)
	copy(result[position:], s[index:])
	return result
}