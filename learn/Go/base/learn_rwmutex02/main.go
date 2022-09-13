package main

import (
	"fmt"
	"sync"
	"time"
)

var count = 0

func read(rwmu *sync.RWMutex, wg * sync.WaitGroup, i int) {
	defer wg.Done()
	fmt.Println("start read", i)
	rwmu.RLock()
	fmt.Println("reading count", i, count)
	time.Sleep(time.Duration(1) * time.Millisecond)
	rwmu.RUnlock()
	fmt.Println("read over", i)
}

func write(rwmu *sync.RWMutex, wg *sync.WaitGroup, i int) {
	defer wg.Done()
	fmt.Println("start write", i)
	rwmu.Lock()
	count ++
	fmt.Println("writing count", i, count)
	time.Sleep(time.Duration(1) * time.Millisecond)
	rwmu.Unlock()
	fmt.Println("write over", i)
}

func main() {
	var rwmu sync.RWMutex
	var wg sync.WaitGroup

	for i := 0; i < 3; i ++ {
		wg.Add(1)
		go write(&rwmu, &wg, i)
	}
	time.Sleep(time.Millisecond * time.Duration(1))

	for i := 0; i < 3; i ++ {
		wg.Add(1)
		go read(&rwmu, &wg, i)
	}
	wg.Wait()
	fmt.Println(count)
}
