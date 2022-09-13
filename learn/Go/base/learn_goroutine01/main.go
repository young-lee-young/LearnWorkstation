package main

import (
	"fmt"
	"sync"
	"time"
)

/**
	waitGroup使用
 */

var wg sync.WaitGroup

func customize(i int) {
	defer wg.Done()
	fmt.Println(i)
	time.Sleep(time.Second * 1)
}

func main() {
	for i := 0; i < 10000; i ++ {
		wg.Add(1)
		go customize(i)
	}
	wg.Wait()
}
