package main

import (
	"practice/base/link"
)

/**
删除链表倒数第 N 个节点 LeetCode No19

快指针先走 N 步
*/
func main() {
	n := 2

	linkList := link.LinkedList{}
	for i := 1; i < 3; i++ {
		linkList.Append(i)
	}
	linkList.ShowList()

	solution(linkList.Head, n)

	linkList.ShowList()
}

func solution(head *link.Node, n int) *link.Node {
	dummy := &link.Node{Next: head}

	fast := dummy
	slow := dummy

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}
