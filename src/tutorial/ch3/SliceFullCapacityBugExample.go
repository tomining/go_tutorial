package main

import "fmt"

func a() {
	x := []int{}
	x = append(x, 0)
	x = append(x, 1)  // commonTags := labelsToTags(app.Labels)
	y := append(x, 2) // Tags: append(commonTags, labelsToTags(d.Labels)...)
	z := append(x, 3) // Tags: append(commonTags, labelsToTags(d.Labels)...)
	fmt.Println(y, z)
}

func b() {
	x := []int{}
	x = append(x, 0)
	x = append(x, 1)
	x = append(x, 2)  // commonTags := labelsToTags(app.Labels)
	y := append(x, 3) // Tags: append(commonTags, labelsToTags(d.Labels)...)
	z := append(x, 4) // Tags: append(commonTags, labelsToTags(d.Labels)...)
	fmt.Println(y, z)
}

/**
 * slice 자료형 append시 버그
 * http://allegro.tech/2017/07/golang-slices-gotcha.html
 */
func main() {
	a()
	b()
}