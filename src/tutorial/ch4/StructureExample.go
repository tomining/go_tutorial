package main

import "fmt"

type Option struct {
	name string
	value string
}

type Item struct {
	name string
	price float64
	quantity int
	Option
}

func main() {
	shoes := Item{"Sports Shoes", 30000, 2, Option{"color", "red"}}

	fmt.Println(shoes)

	fmt.Println(shoes.name)
	fmt.Println(shoes.value)	//임베드드 필드인 Option 구조체의 내부 필드 value에 접근
	fmt.Println(shoes.Option.name)
}