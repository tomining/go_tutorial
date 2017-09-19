package main

import (
	"log"
	"fmt"
)

func protect(g func()) {
	defer func() {
		log.Println("done")

		if err := recover(); err != nil {
			log.Println("run time panic: %v", err)
		}
	}()
	log.Println("start")
	g()
}

func main() {
	test_data := [][]int{{4, 2}, {1, 0}, {9, 3}}
	for _, data := range(test_data) {
		protect(func() {
			v := divide(data[0], data[1])
			fmt.Println(v)
		})
	}
}

func divide(a, b int) int {
	return a/b
}