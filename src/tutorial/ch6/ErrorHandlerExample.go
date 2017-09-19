package main

import (
	"log"
	"fmt"
)

type fType func(int, int) int

func errorHandler(fn fType) fType {
	return func(a int, b int) int {
		defer func() {
			if err, ok := recover().(error); ok {
				log.Printf("run time panic: %v", err)
			}
		}()
		return fn(a, b)	// 오류가 발생한 경우 0을 반환(return type이 int)
	}
}

func main() {
	test_data := [][]int{{4, 2}, {1, 0}, {9, 3}}
	for _, data := range(test_data) {
		fmt.Println(errorHandler(divide)(data[0], data[1]))
	}
}

func divide(a, b int) int {
	return a/b
}