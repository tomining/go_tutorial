package main

import "fmt"

func checkValue(v interface{}) {
	//Build fail. non-bool v used as if condition
	if v {
		fmt.Printf("value is %v\n", v)
		return
	}
	fmt.Println("value is nil")
}

func checkValue2(v interface{}) {
	if v != nil {	//if statement should return bool type
		fmt.Printf("value is %v\n", v)
		return
	}
	fmt.Println("value is nil")
}