package main

import (
	"fmt"
	"runtime"
	"sync"
)

type counter struct {
	i int64
	mu sync.Mutex	// 공유 데이터 i를 보호하기 위한 Mutex
}

func (c *counter) increment() {
	c.mu.Lock()	//i 변경 부분을 Mutax 로 Lock
	c.i += 1
	c.mu.Unlock()
}

func (c *counter) display() {
	fmt.Println(c.i)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := counter{i: 0}
	done := make(chan struct{})

	for i := 0; i < 1000; i++ {
		go func() {
			c.increment()
			done <- struct{}{}
		}()
	}

	for i := 0; i < 1000; i++ {
		<- done
	}

	c.display()
}
