package main

import (
	"fmt"
	"time"
	"runtime"
)
/**
	协程

	* 轻量级线程
	* 非抢占式多任务处理，由协程主动交出控制权
	* 编译器、解释器、虚拟机层面多任务，非操作系统层面多任务
	* 多个协程可能在一个或多个线程上运行

	gotoutine定义
	* 任务函数加上go就能送给调度器执行
	* 不需要在定义时区分是否是异步函数
	* 调度器在合适点进行切换（IO、select、channel、等待锁、函数调用、runtime.Gosched()）
	* 使用 -race 来检测数据访问冲突

	go run -race go_file.go
 */
func main() {
	var a [10]int
	for i := 0; i < 10; i ++ {
		// 匿名函数
		go func(j int) {
			for {
				// fmt.Printf("Hello from goroutine %d\n", j)
				// Printf是IO操作，IO操作过程中会有协程切换

				a[j] ++
				runtime.Gosched() // 手动交出控制权
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
