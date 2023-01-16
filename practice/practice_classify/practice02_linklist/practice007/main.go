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
*/
func solution(head *ListNode) *ListNode {

}
