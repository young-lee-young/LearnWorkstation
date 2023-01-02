package main

import (
	"practice/base/link"
	"fmt"
)

/**
	删除链表倒数第k个节点 LeetCode No019

	思路：快慢指针，快指针先走k步
 */
func main() {
	i := 0
	linkList := link.LinkedList{}
	for i < 20 {
		linkList.Append(i)
		i ++
	}
	linkList.ShowList()

	slow := linkList.Head
	fast := linkList.Head
	k := 5

	j := 0
	for j < k {
		fast = fast.Next
		j ++
	}

	pre := slow
	for fast != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next
	}

	fmt.Println(slow.Data)
	pre.Next = pre.Next.Next
	linkList.ShowList()
}
