package queue

import (
	"testing"
	"fmt"
)

func TestQueue(t *testing.T) {
	queue := Queue{}
	fmt.Println(queue.GetLength())
	fmt.Println(queue.IsEmpty())

	i := 0
	for i < 10 {
		queue.Enqueue(i)
		i ++
	}

	fmt.Println(queue.GetLength())
	fmt.Println(queue.IsEmpty())

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
}