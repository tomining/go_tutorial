package main

import "fmt"

func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d\n", i)
	}
}

func main() {
	f()
}