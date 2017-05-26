package main

import (
	"fmt"
	_ "math"
)

func main() {
	c1 := 1 + 2i
	c2 := complex64(3 + 4i)
	c3 := complex(5, 6)

	fmt.Println(c1, real(c1), imag(c1))
	fmt.Println(c2, real(c2), imag(c2))
	fmt.Println(c3, real(c3), imag(c3))
}