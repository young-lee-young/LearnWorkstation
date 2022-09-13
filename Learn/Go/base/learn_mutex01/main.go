package main

import (
	"fmt"
	"sync"
)

func lockFunc() {
	var count int
	var wg sync.WaitGroup

	// 10个goroutine去操作，count是不安全的
	for i := 0; i < 10; i ++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j ++ {
				count ++
			}
		}()
	}
	wg.Wait()

	fmt.Println("lock func count: ", count)
}

func unlockFunc() {
	var count int
	var wg sync.WaitGroup
	var mu sync.Mutex

	// 这里将count操作加锁，是安全的
	for i := 0; i < 10; i ++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j ++ {
				mu.Lock()
				count ++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()

	fmt.Println("unlock func count: ", count)
}

func main() {
	lockFunc()
	unlockFunc()
}
