/**
LeetCode No142 环形链表Ⅱ
*/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{}

	ret := solution(head)

	fmt.Println("ret:", ret)
}

/**
1. 判断环形链表是否有环
2. 环入口


解题思路：快慢指针

1. 判断链表是否有环

快指针每次走 2 个节点，慢指针每次走 1 个节点，当快指针和慢指针相等时，证明链表有环

实际上，假如链表有环，当慢指针进入环中时，快指针每次以 1 个节点的速度追上慢指针，两个指针最后肯定会相遇


2. 环入口

⚠️：慢指针只能走不到一圈就被快指针追上

慢指针只能是在圈里走了不到一圈就被快指针追上，证明见 iPad 笔记

慢指针路程 x + y，时间 (x + y) / 1

快指针路程 x + y + n * (y + z)，时间 (x + y + n * (y + z)) / 2

时间相等：x + y = (x + y + n * (y + z)) / 2
*/
func solution(head *ListNode) *ListNode {
	slow := head
	fast := head

	// ⚠️：这里只用判断快指针为空，不用判断慢指针
	// 如果有环，两个都不会为空，如果没有环，肯定是快指针先为空
	for fast != nil && fast.Next != nil {

		fast = fast.Next.Next
		slow = slow.Next

		// 快慢指针相遇
		if fast == slow {
			// ⚠️：这里有一个证明点，即 头节点 -> 入口点 = 相遇点 -> 入口点
			// 相遇点位置
			meet := fast
			// 开始位置
			start := head

			for meet != start {
				meet = meet.Next
				start = start.Next
			}

			// 这里是入口点的位置
			return meet
		}
	}

	return nil
}
