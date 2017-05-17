package main

import "fmt"

func main() {
	loopWithoutLable()		//Lable 없이 반복문
	loopWithLable()			//Lable을 이용한 반복문
	loopWithLableAndRange()	//Range Keyword를 이용한 반복문
}

func loopWithoutLable() {
	x := 7
	table := [][]int{{1, 5, 9,}, {2, 6, 5, 13}, {5, 3, 7, 4 }}
	found := false
	for row := 0; row < len(table); row++ {
		for col := 0; col < len(table[row]); col++ {
			if table[row][col] == x {
				found = true
				fmt.Println("found %d(row:%d, col:%d)", x, row, col)
				break
			}
		}

		if found {
			break
		}
	}
}

func loopWithLable() {
	x := 7
	table := [][]int{{1, 5, 9,}, {2, 6, 5, 13}, {5, 3, 7, 4 }}

	LOOP:
	for row := 0; row < len(table); row++ {
		for col := 0; col < len(table[row]); col++ {
			if table[row][col] == x {
				fmt.Printf("found %d(row:%d, col:%d)\n", x, row, col)
				break LOOP
			}
		}
	}
}

func loopWithLableAndRange() {
	x := 7
	table := [][]int{{1, 5, 9,}, {2, 6, 5, 13}, {5, 3, 7, 4 }}

	LOOP:
	for row, rowValue := range table {
		for col, colValue := range rowValue {
			if colValue == x {
				fmt.Println("found %d(row:%d, col:%d)", x, row, col)
				break LOOP
			}
		}
	}
}