package main

import (
	"fmt"
	"practice/base/link"
)

/**
相交链表 LeetCode No160

双指针
*/
func main() {
	i := 0
	linkList1 := link.LinkedList{}

	for i < 10 {
		linkList1.Append(i)
		i++
	}

	linkList2 := link.LinkedList{}
	j := 10
	for j < 15 {
		linkList2.Append(j)
		j++
	}

	linkList3 := link.LinkedList{}
	k := 20
	for k < 30 {
		linkList3.Append(k)
		k++
	}

	current1 := linkList1.Head
	m := 0
	for m < 9 {
		current1 = current1.Next
		m++
	}

	current2 := linkList2.Head
	n := 0
	for n < 4 {
		current2 = current2.Next
		n++
	}

	current1.Next = linkList3.Head
	current2.Next = linkList3.Head

	linkList1.ShowList()
	linkList2.ShowList()

	ret := solution(linkList1.Head, linkList2.Head)
	fmt.Println("ret", ret.Data)
}

func solution(headA *link.Node, headB *link.Node) *link.Node {
	a := headA
	b := headB

	for a != b {
		if a != nil {
			a = a.Next
		} else {
			a = headB
		}

		if b != nil {
			b = b.Next
		} else {
			b = headA
		}
	}

	return a
}
