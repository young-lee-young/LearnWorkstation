package main

import (
	"practice/base/stack"
	"fmt"
)

/**
	用两个栈实现队列 LeetCode No232
 */
func main() {
	queue := CustomizeQueue{}
	i := 0
	for i < 10 {
		queue.EnQueue(i)
		i ++
	}
	fmt.Println(queue.IsEmpty())
	fmt.Println(queue.GetLength())
	for !queue.IsEmpty() {
		fmt.Println(queue.DeQueue())
	}
}

type CustomizeQueue struct {
	StackOne stack.Stack
	StackTwo stack.Stack
}

func (queue *CustomizeQueue) EnQueue(data interface{}) {
	queue.StackOne.Push(data)
}

func (queue *CustomizeQueue) DeQueue() interface{} {
	if queue.StackTwo.IsEmpty() && queue.StackOne.IsEmpty() {
		return nil
	}
	if queue.StackTwo.IsEmpty() {
		for !queue.StackOne.IsEmpty() {
			data := queue.StackOne.Pop()
			queue.StackTwo.Push(data)
		}
	}
	return queue.StackTwo.Pop()
}

func (queue *CustomizeQueue) GetLength() int {
	return queue.StackOne.GetLength() + queue.StackTwo.GetLength()
}

func (queue *CustomizeQueue) IsEmpty() bool {
	if queue.StackOne.IsEmpty() && queue.StackTwo.IsEmpty() {
		return true
	}
	return false
}
