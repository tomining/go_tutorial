package main

import "fmt"

const (	//2진수
	Running = 1 << iota	//1
	Waiting				//10 = 2
	Send				//100 = 4
	Receive				//1000 = 8
)

func main() {
	stat := Running | Send	// stat := 101
	display(stat)
}

/**
 * 2진수를 비트연산을 통해서 비교
 */
func display(stat int) {
	if stat&Running == Running {
		fmt.Println("Running")
	}
	if stat&Waiting == Waiting {
		fmt.Println("Waiting")
	}
	if stat&Send == Send {
		fmt.Println("Send")
	}
	if stat&Receive == Receive {
		fmt.Println("Receive")
	}
}
