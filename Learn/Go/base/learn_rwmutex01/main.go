package main

import (
	"fmt"
	"sync"
	"time"
)

func read(rwmu *sync.RWMutex, wg *sync.WaitGroup, i int) {
	defer wg.Done()
	fmt.Println("start read", i)
	// 读锁
	rwmu.RLock()
	fmt.Println("reading", i)
	time.Sleep(time.Second * time.Duration(1))
	rwmu.RUnlock()
	fmt.Println("read over", i)
}

func main() {
	var rwmu sync.RWMutex
	var wg sync.WaitGroup
	for i := 0; i < 3; i ++ {
		wg.Add(1)
		// 这里是并发读，并不会等前一个结束再读
		go read(&rwmu, &wg, i)
	}
	wg.Wait()
}
