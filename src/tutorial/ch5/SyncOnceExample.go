package main

import (
	"fmt"
	"runtime"
	"sync"
)

const initialValue = -500

type counter struct {
	i int64
	mu sync.Mutex	// 공유 데이터 i를 보호하기 위한 Mutex
	once sync.Once	// 한 번만 수행할 함수를 지정하기 위한 Once 구조체
}

func (c *counter) increment() {
	c.once.Do(func() {
		fmt.Println("init value")	// 실행 결과에서 한 번만 수행됨을 확인할 수 있음
		c.i = initialValue
	})
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
