package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

/**
实现自旋锁
*/
type SpinLock int32

func (sl *SpinLock) Lock() {
	// 自旋锁，如果获取锁失败，则一直重试
	for !atomic.CompareAndSwapInt32((*int32)(sl), 0, 1) {
		runtime.Gosched()
	}
}

func (sl *SpinLock) Unlock() {
	atomic.StoreInt32((*int32)(sl), 0)
}

func main() {
	lock := new(SpinLock)
	var num int
	for i := 0; i < 2; i++ {
		go callFunc(i, &num, lock, 500*time.Millisecond)
	}
	select {}
}

func callFunc(i int, num *int, lock sync.Locker, d time.Duration) {
	for {
		func() {
			lock.Lock()
			defer lock.Unlock()
			*num ++
			fmt.Printf("goroutine %v set num %v\n", i, *num)
			time.Sleep(d)
		}()
	}
}
