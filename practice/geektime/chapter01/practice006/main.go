/**
 * @Time:    2022/4/20 21:50 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	环形链表 LeetCode No141

 */
package main

import (
	"practice/base/link"
	"fmt"
)

func main() {
	linkList := link.LinkedList{}
	for i := 0; i < 10; i ++ {
		linkList.Append(i)
	}
	ret1 := solution(linkList.Head)
	fmt.Println("ret 1", ret1)

	// 构造环形链表
	j := 0
	current := linkList.Head
	ringStartNode := &link.Node{}
	for current.Next != nil {
		current = current.Next
		if j == 5 {
			ringStartNode = current
		}
		j ++
	}
	current.Next = ringStartNode
	ret2 := solution(linkList.Head)
	fmt.Println("ret 2", ret2)
}

func solution(head *link.Node) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	for slow != nil && fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}
