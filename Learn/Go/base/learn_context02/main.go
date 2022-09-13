package main

import (
	"context"
	"fmt"
	"time"
)

/**
	context控制多个goroutine
 */

func main() {
	ctx, cancelMulti := context.WithCancel(context.Background())
	go watch(ctx, "【监控1】")
	go watch(ctx, "【监控2】")
	go watch(ctx, "【监控3】")

	time.Sleep(time.Second * 5)
	fmt.Println("通知监控停止")
	// 控制多个goroutine
	cancelMulti()
	time.Sleep(time.Second * 2)
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <- ctx.Done():
			fmt.Println(name, "监控退出")
			return
		default:
			fmt.Println(name, "监控中...")
			time.Sleep(time.Second * 1)
		}
	}
}
