package main

import (
	"fmt"
	"sync"
)

/**
	mutex是不可重入锁
 */
type Counter struct {
	sync.Mutex
	Count int
}

// unlock一个没加锁的mutex
func unlockWithoutLock() {
	var mu sync.Mutex
	defer mu.Unlock()
	fmt.Println("hello world")
	// fatal error: sync: unlock of unlocked mutex
}

// 复制锁
func copyLock() {
	var counter Counter
	counter.Lock()
	defer counter.Unlock()
	// 这里复制了一个使用了的mutex，会导致死锁
	copyMutex(counter)
}

// 这里Counter参数是通过复制方式传入
func copyMutex(counter Counter) {
	counter.Lock()
	defer counter.Unlock()
	fmt.Println("dead lock")
}

func main() {
	copyLock()
}
