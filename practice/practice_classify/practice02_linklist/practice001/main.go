/**
 * @Time:    2022/11/29 15:12 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	LeetCode No203 移除链表元素

 */
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node5 := ListNode{Val: 5, Next: nil}
	node4 := ListNode{Val: 4, Next: &node5}
	node3 := ListNode{Val: 3, Next: &node4}
	node2 := ListNode{Val: 2, Next: &node3}
	node1 := ListNode{Val: 1, Next: &node2}
	head := &node1

	ret := solution(head, 5)

	for ret != nil {
		fmt.Println(ret.Val)
		ret = ret.Next
	}
}

/**
	不使用虚拟头节点
 */
func solution(head *ListNode, val int) *ListNode {
	// 头节点的值为目标值
	for head != nil && head.Val == val {
		head = head.Next
	}

	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return head
}

/**
	使用虚拟头节点
 */
func solution2(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head

	cur := dummyHead
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return dummyHead.Next
}
