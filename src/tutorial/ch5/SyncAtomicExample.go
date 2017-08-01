package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

type counter struct {
	i int64
}

func (c *counter) increment() {
	atomic.AddInt64(&c.i, 1)	// 1증가에 따른 원자성 보장
}

func (c *counter) display() {
	fmt.Println(c.i)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := counter{i:0}
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.increment()
		}()
	}

	wg.Wait()

	c.display()
}