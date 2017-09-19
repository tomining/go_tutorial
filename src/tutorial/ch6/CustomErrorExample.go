package main

import (
	"fmt"
	"math"
	"time"
	"log"
)

type SqrtError struct {
	time time.Time
	value float64
	message string
}

func (e SqrtError) Error() string {
	return fmt.Sprintf("[%v]ERROR - %s(value: %g)", e.time, e.message, e.value)
}

func Sqrt2(f float64) (float64, error) {
	if f < 0 {
		return 0, SqrtError{time: time.Now(), value: f, message: "음수는 사용할 수 없습니다."}
	}

	if math.IsInf(f, 1) {
		return 0, SqrtError{time: time.Now(), value: f, message: "무한대 값은 사용할 수 없습니다."}
	}

	if math.IsNaN(f) {
		return 0, SqrtError{time: time.Now(), value: f, message: "잘못된 수 입니다."}
	}

	return math.Sqrt(f), nil
}

func main() {
	//v, err := Sqrt2(9)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(v)
	//
	//v, err = Sqrt2(-9)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(v)

	test_data := []float64{9, 0, -9, 0}
	for _, data := range(test_data) {
		v, err := Sqrt2(data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(v)
	}
}