package main

import "practice/base/link"

/**
合并两个有序链表 LeetCode No021
*/
func main() {
	linkList1 := link.LinkedList{}
	linkList2 := link.LinkedList{}

	count := 0
	for count < 20 {
		if count%2 == 0 {
			linkList1.Append(count)
		} else {
			linkList2.Append(count)
		}
		count += 1
	}

	newHead := mergeTwoOrderLinkList(linkList1.Head, linkList2.Head)
	linkListResult := link.LinkedList{}
	linkListResult.Head = newHead
	linkListResult.ShowList()
}

func mergeTwoOrderLinkList(node1 *link.Node, node2 *link.Node) *link.Node {
	if node1 == nil {
		return node2
	}
	if node2 == nil {
		return node1
	}

	if node1.Data.(int) > node2.Data.(int) {
		node2.Next = mergeTwoOrderLinkList(node1, node2.Next)
		return node2
	} else {
		node1.Next = mergeTwoOrderLinkList(node1.Next, node2)
		return node1
	}
}
