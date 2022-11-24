package main

import (
	"practice/base/queue"
	"fmt"
)

/**
	两个队列实现栈 LeetCode No232
 */
func main() {
	stack := CustomizeStack{}
	i := 0
	for i < 10 {
		stack.Push(i)
		i ++
	}
	fmt.Println(stack.IsEmpty())
	fmt.Println(stack.GetLength())
	for !stack.IsEmpty() {
		fmt.Println(stack.Pop())
	}
}

type CustomizeStack struct {
	QueueOne queue.Queue
	QueueTwo queue.Queue
}

func (stack *CustomizeStack) Push(data interface{}) {
	if stack.QueueOne.IsEmpty() && stack.QueueTwo.IsEmpty() {
		stack.QueueOne.Enqueue(data)
		return
	}
	if stack.QueueOne.IsEmpty() {
		stack.QueueOne.Enqueue(data)
		for !stack.QueueTwo.IsEmpty() {
			stack.QueueOne.Enqueue(stack.QueueTwo.Dequeue())
		}
	} else {
		stack.QueueTwo.Enqueue(data)
		for !stack.QueueOne.IsEmpty() {
			stack.QueueTwo.Enqueue(stack.QueueOne.Dequeue())
		}
	}
}

func (stack *CustomizeStack) Pop() interface{} {
	if stack.QueueOne.IsEmpty() && stack.QueueTwo.IsEmpty() {
		return nil
	}
	if stack.QueueOne.IsEmpty() {
		return stack.QueueTwo.Dequeue()
	} else {
		return stack.QueueOne.Dequeue()
	}
}

func (stack *CustomizeStack) GetLength() int {
	return stack.QueueOne.GetLength() + stack.QueueTwo.GetLength()
}

func (stack *CustomizeStack) IsEmpty() bool {
	if stack.QueueOne.IsEmpty() && stack.QueueTwo.IsEmpty() {
		return true
	}
	return false
}
