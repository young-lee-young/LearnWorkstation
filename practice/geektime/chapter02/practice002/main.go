/**
 * @Time:    2022/4/21 10:51 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	最小值栈 LeetCode No155

 */
package main

import (
	"practice/base/stack"
	"fmt"
)

func main() {
	nums := []int{1, -2, 0, 3, -5, 4, 2}

	c := Constructor()

	for _, num := range nums {
		c.Push(num)
		fmt.Println("min", c.GetMin())
	}
}

type MinStack struct {
	StackOne stack.Stack
	StackTwo stack.Stack
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.StackOne.Push(val)

	if this.StackTwo.IsEmpty() {
		this.StackTwo.Push(val)
	} else {
		// 将当前最小值和新值进行比较，将较小值放到栈2里面
		currentMin := this.StackTwo.Peek().(int)
		if val < currentMin {
			this.StackTwo.Push(val)
		} else {
			this.StackTwo.Push(currentMin)
		}
	}
}

func (this *MinStack) Pop() {
	this.StackOne.Pop()
	this.StackTwo.Pop()
}

func (this *MinStack) Top() int {
	return this.StackOne.Peek().(int)
}

func (this *MinStack) GetMin() int {
	return this.StackTwo.Peek().(int)
}
