package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for {
		n := <- c
		fmt.Printf("worker %d received %d\n", id, n)
	}
}

func chanDemo() {
	c := make(chan int)

	go worker(1, c)
	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
}
