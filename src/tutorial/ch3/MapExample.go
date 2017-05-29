package main

import "fmt"

func main() {
	makeMap()
	makeMap2()
	makeMap3()
	accessMapItem()
	addAndDeleteItemIntoMap()
}

func makeMap() {
	numberMap := map[string]int{}
	numberMap["one"] = 1
	numberMap["two"] = 2
	numberMap["three"] = 3
	fmt.Println(numberMap)
}

func makeMap2() {
	numberMap := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,	//요소를 여러 줄로 표기할 때 요소의 끝에 콤마(,)를 붙여야 함
	}
	fmt.Println(numberMap)
}

func makeMap3() {
	numberMap := make(map[string]int, 3)
	numberMap["one"] = 1
	numberMap["two"] = 2
	numberMap["three"] = 3
	fmt.Println(numberMap)
}

func accessMapItem() {
	numberMap := make(map[string]int, 3)
	numberMap["one"] = 1
	numberMap["two"] = 2
	numberMap["three"] = 3

	for k, v := range numberMap {
		fmt.Println(k, v)
	}
}

func addAndDeleteItemIntoMap() {
	numberMap := make(map[string]int, 3)
	numberMap["one"] = 1
	numberMap["two"] = 2
	fmt.Println(numberMap)

	numberMap["three"] = 3
	fmt.Println(numberMap)

	delete(numberMap, "three")
	fmt.Println(numberMap)
}