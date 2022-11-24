package main

import (
	"practice/base/link"
	"fmt"
)

/**
	判断链表是否有环 LeetCode No141

	标签：双指针

	解题思路：快慢指针，快指针每次2步，慢指针每次一步
 */
func main() {
	// 生成有环链表
	i := 0
	linkList := link.LinkedList{}
	for i < 10 {
		linkList.Append(i)
		i ++
	}
	fmt.Println(hasRing(linkList.Head))

	// 构造环状链表
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
	fmt.Println(hasRing(linkList.Head))
}

func hasRing(head *link.Node) bool {
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
