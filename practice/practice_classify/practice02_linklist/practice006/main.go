/**
LeetCode No160 相交链表
*/
package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	headA := &ListNode{}
	headB := &ListNode{}

	ret := solution(headA, headB)

	fmt.Println("ret:", ret)
}

/**
思路：
1. 计算链表 A 的长度
2. 计算链表 B 的长度
3. 计算链表长度的差值为 n
4. 让长的链表先走 n 步
5. A、B 链表一起走，相等就是相交点

A：9 -> 8 -> 7 -> 6
                   \
					 -> 5 -> 4
			       /
B：          2 -> 3
*/
func solution(headA *ListNode, headB *ListNode) *ListNode {
	// 1. 计算链表 A 的长度
	lenA := 0
	nodeA := headA
	for nodeA != nil {
		lenA += 1
		nodeA = nodeA.Next
	}

	// 2. 计算链表 B 的长度
	lenB := 0
	nodeB := headB
	for nodeB != nil {
		lenB += 1
		nodeB = nodeB.Next
	}

	// 3. 计算链表差值
	n := lenA - lenB

	curA := headA
	curB := headB

	if n < 0 {
		curA, curB = curB, curA
	}

	n = int(math.Abs(float64(n)))

	// 长的链表先走 n 步
	for n != 0 {
		curA = curA.Next
		n--
	}

	// A、B 链表一起走
	for curA != nil {
		if curA == curB {
			return curA
		}

		curA = curA.Next
		curB = curB.Next
	}
	return nil
}
