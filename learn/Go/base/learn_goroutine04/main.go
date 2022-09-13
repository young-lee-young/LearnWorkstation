package main

import (
	"fmt"
	"runtime"
	"time"
)

/**
主goroutine睡眠 线程被让出来
*/
func goroutineNormal() {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			fmt.Println("A: ", i)
		}()
	}

	time.Sleep(time.Second * 1)
}

/**
会发生死锁

main goroutine阻塞住 其他goroutine都退出了

ch永远不会收到值 所以死锁
*/
func goroutineDeadlock() {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			fmt.Println("A: ", i)
		}()
	}

	var ch = make(chan int)
	<-ch
}

func main() {
	//goroutineNormal()
	goroutineDeadlock()
}
