/**
LeetCode No019 删除链表倒数第k个节点

思路：快慢指针，快指针先走k步
*/
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{}
	n := 5

	ret := solution(head, n)
	fmt.Println("ret:", ret)
}

func solution(head *ListNode, n int) *ListNode {
	slow := head
	fast := head

	j := 0
	for j < n {
		fast = fast.Next
		j++
	}

	pre := slow
	for fast != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next
	}

	pre.Next = pre.Next.Next

	return head
}
