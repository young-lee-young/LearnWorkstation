package main

import (
	"fmt"
	"time"
	"context"
)

/**
	context控制goroutine
 */

func watch(ctx context.Context) {
	for {
		select {
		case <- ctx.Done():
			fmt.Println("监控退出")
			return
		default:
			fmt.Println("监控中...")
			time.Sleep(time.Second * 1)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx)

	time.Sleep(time.Second * 5)
	fmt.Println("通知监控停止")
	// 使用cancel控制goroutine
	cancel()
	time.Sleep(time.Second * 2)
}
