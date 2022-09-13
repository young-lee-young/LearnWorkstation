package main

import (
	"fmt"
	"runtime"
	"time"
)

/**
	获取goroutine个数
 */

func receiver() {
	time.Sleep(time.Second * 1)
}

func main() {
	for i := 0; i < 100; i ++ {
		go receiver()
	}
	// 获取goroutine个数
	fmt.Println(runtime.NumGoroutine())
}
