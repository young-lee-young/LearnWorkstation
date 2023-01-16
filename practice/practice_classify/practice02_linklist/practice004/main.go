/**
LeetCode No024 两两交换链表中的节点
*/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//node4 := &ListNode{Val: 4}
	node3 := &ListNode{Val: 3, Next: nil}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}

	head := node1

	ret := solution(head)

	fmt.Println("ret:", ret)
}

func solution(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}

	last := dummy

	node := dummy.Next

	for node != nil && node.Next != nil {
		node1 := node
		node2 := node.Next

		next := node.Next.Next

		// 交换
		node2.Next = node1

		last.Next = node2

		last = node1

		node = next
	}

	if node != nil {
		last.Next = node
	} else {
		last.Next = nil
	}

	return dummy.Next
}
