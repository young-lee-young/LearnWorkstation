package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan string, 0)

	go func(ctx context.Context) {
		select {
		case <- ctx.Done():
			fmt.Println("done")
			return
		default:
			message := <- ch
			fmt.Println(message)
		}
	}(ctx)
}
