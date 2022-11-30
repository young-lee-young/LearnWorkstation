package main

import (
	"practice/base/link"
)

/**
链表反转 LeetCode No206
*/
func main() {
	linkList := link.LinkedList{}
	for i := 0; i < 10; i++ {
		linkList.Append(i)
	}
	linkList.ShowList()

	// 暴力解法
	current := reverseLinkList(linkList.Head)
	linkList.Head = current
	linkList.ShowList()

	// 递归解法
	currentTwo := recursionReverseLinkList(linkList.Head)
	linkList.Head = currentTwo
	linkList.ShowList()

	// 头插法
	currentThree := HeadInsert(linkList.Head)
	linkList.Head = currentThree
	linkList.ShowList()
}

func reverseLinkList(head *link.Node) *link.Node {
	if head == nil || head.Next == nil {
		return head
	}

	pre := head
	current := head.Next
	temp := &link.Node{}

	for current != nil {
		temp = current.Next
		current.Next = pre
		pre = current
		current = temp
	}

	// ⚠️ 整理需要注意
	head.Next = nil

	head = pre
	return head
}

/**
递归反转链表
*/
func recursionReverseLinkList(node *link.Node) *link.Node {
	if node == nil || node.Next == nil {
		return node
	}
	newHead := recursionReverseLinkList(node.Next)
	node.Next.Next = node
	node.Next = nil
	return newHead
}

/**
头插法反转链表
*/
func HeadInsert(node *link.Node) *link.Node {
	newHead := &link.Node{}
	for node != nil {
		next := node.Next
		node.Next = newHead.Next
		newHead.Next = node
		node = next
	}
	return newHead.Next
}
