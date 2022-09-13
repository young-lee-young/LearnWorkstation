package main

import (
	"practice/base/stack"
	"fmt"
)

/**
	最小值栈 LeetCode No155
 */
func main() {
	stack := MinStack{}
	i := 10
	for i > 0 {
		stack.Push(i)
		i --
	}
	fmt.Println(stack.IsEmpty())
	fmt.Println(stack.GetLength())
	fmt.Println(stack.GetMin())
	fmt.Println(stack.Pop())
	fmt.Println(stack.GetMin())
}

type MinStack struct {
	StackOne stack.Stack
	StackTwo stack.Stack
}

func (stack *MinStack) Push(data interface{}) {
	stack.StackOne.Push(data)
	if stack.StackTwo.IsEmpty() {
		stack.StackTwo.Push(data)
	} else {
		min := stack.StackTwo.Peek()
		if data.(int) < min.(int) {
			stack.StackTwo.Push(data)
		} else {
			stack.StackTwo.Push(min)
		}
	}
}

func (stack *MinStack) Pop() interface{} {
	stack.StackTwo.Pop()
	return stack.StackOne.Pop()
}

func (stack *MinStack) GetMin() interface{} {
	return stack.StackTwo.Peek()
}

func (stack *MinStack) GetLength() int {
	return stack.StackOne.GetLength()
}

func (stack *MinStack) IsEmpty() bool {
	return stack.StackOne.IsEmpty()
}
