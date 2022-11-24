/**
 * @Time:    2021/3/11 19:50 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content: 回文链表
 */
package main

import (
	"practice/base/link"
	stack2 "practice/base/stack"
	"fmt"
)

/**
	回文链表 LeetCode No234

	标签：双指针
 */
func main() {
	nums := []int{1, 2, 3, 3, 2, 1}
	linkedList := link.LinkedList{}
	for _, num := range nums {
		linkedList.Append(num)
	}
	result := solution(linkedList.Head)
	fmt.Println(result)
}

func solution(head *link.Node) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast.Next != nil {
		slow = slow.Next
	}

	// 创建栈，将链表后半部分入栈
	stack := stack2.Stack{}
	stack.Push(slow.Data)
	for slow.Next != nil {
		stack.Push(slow.Next.Data)
		slow = slow.Next
	}

	for !stack.IsEmpty() {
		if stack.Pop() != head.Data {
			return false
		}
		head = head.Next
	}
	return true
}
