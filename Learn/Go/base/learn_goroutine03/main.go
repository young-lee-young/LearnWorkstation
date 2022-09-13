package main

import (
	"fmt"
	"runtime"
	"sync"
)

/**
	goroutine真正并行，规定多核去跑
 */

var wg sync.WaitGroup

func loop() {
	for i := 0; i < 10000; i ++ {
		fmt.Println(i)
	}
	wg.Done()
}

func main() {
	// 最多使用两个核
	runtime.GOMAXPROCS(2)
	go loop()
	go loop()
	wg.Add(2)
	wg.Wait()
}
