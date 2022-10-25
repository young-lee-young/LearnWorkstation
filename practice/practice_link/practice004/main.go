package main

import (
	"practice/base/link"
)

/**
有序链表中删除重复节点 LeetCode No083
*/
func main() {
	i := 0
	linkList := link.LinkedList{}
	for i < 20 {
		if i%2 == 0 {
			linkList.Append(i)
		} else {
			linkList.Append(i - 1)
		}
		i++
	}
	linkList.ShowList()

	// 正式解答
	current := linkList.Head
	for current != nil && current.Next != nil {
		if current.Data == current.Next.Data {
			// 删除下一个节点
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	linkList.ShowList()
}
