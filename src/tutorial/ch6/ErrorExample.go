package main

import (
	_ "errors"
	"fmt"
	"math"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		//return 0, errors.New("음수는 사용할 수 없습니다.")	// errors.New() 사용
		return 0, fmt.Errorf("음수(%g)는 사용할 수 없습니다.", f)	//fmt.Errorf() 사용
	}

	return math.Sqrt(f), nil
}

func main() {
	if f, err := Sqrt(-1); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Println(f)
	}
}