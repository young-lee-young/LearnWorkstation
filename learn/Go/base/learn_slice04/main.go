package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	nums := make([]int, 0)

	count := 10
	wg.Add(count)
	for i := 0; i < count; i ++ {
		// 多个goroutine操作一个slice会有数据丢失问题，解决办法是加锁
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			nums = append(nums, i)
		}()
	}
	wg.Wait()

	fmt.Println(nums)
	fmt.Println("nums len", len(nums))
}
