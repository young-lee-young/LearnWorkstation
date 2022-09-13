package main

import (
	"sync"
)

var wg sync.WaitGroup

func loop() {
	for i := 0; i < 10000; i++ {
		i += 1
	}
	wg.Done()
}

/**
goroutine调度问题
*/

func main() {
	// 两个loop是顺序执行的，因为loop中是CPU密集型，并没有让出CPU
	// 更新：上面的结论已经不正确了，goroutine已经是抢占式调度，防止CPU密集型某些goroutine被饿死
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go loop()
	}
	wg.Wait()
}
