package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(intToUint8(100))
	fmt.Println(intToUint8(10000))
}

func intToUint8(i int) (uint8, error) {
	if 0 <= i && i <= math.MaxUint8 {
		return uint8(i), nil
	}
	return 0, fmt.Errorf("%d cannot convert to uint8 type", i)
}