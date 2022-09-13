package main

import (
	"practice/base/link"
	"fmt"
)

/**
	两个链表交点 LeetCode No160
 */
func main() {
	i := 0
	linkList1 := link.LinkedList{}

	for i < 10 {
		linkList1.Append(i)
		i ++
	}

	linkList2 := link.LinkedList{}
	j := 10
	for j < 15 {
		linkList2.Append(j)
		j ++
	}

	linkList3 := link.LinkedList{}
	k := 20
	for k < 30 {
		linkList3.Append(k)
		k ++
	}

	current1 := linkList1.Head
	m := 0
	for m < 9 {
		current1 = current1.Next
		m ++
	}

	current2 := linkList2.Head
	n := 0
	for n < 4 {
		current2 = current2.Next
		n ++
	}

	current1.Next = linkList3.Head
	current2.Next = linkList3.Head

	linkList1.ShowList()
	linkList2.ShowList()

	linkListLength1 := linkList1.GetLength()
	linkListLength2 := linkList2.GetLength()

	var lengthDiff int
	if linkListLength1 > linkListLength2 {
		fast := linkList1.Head
		lengthDiff = linkListLength1 - linkListLength2
		p := 0
		for p < lengthDiff {
			fast = fast.Next
			p ++
		}

		slow := linkList2.Head
		for fast != nil && slow != nil {
			if fast == slow {
				fmt.Println(fast.Data)
				return
			} else {
				fast = fast.Next
				slow = slow.Next
			}
		}

	} else if linkListLength2 > linkListLength1 {
		fast := linkList2.Head
		lengthDiff = linkListLength2 - linkListLength1
		q := 0
		for q < lengthDiff {
			fast = fast.Next
			q ++
		}

		slow := linkList1.Head
		for fast != nil && slow != nil {
			if fast == slow {
				fmt.Println(fast.Data)
				return
			} else {
				fast = fast.Next
				slow = slow.Next
			}
		}

	} else {
		fast := linkList1.Head
		slow := linkList2.Head
		for fast != nil && slow != nil {
			if fast == slow {
				fmt.Println(fast.Data)
				return
			} else {
				fast = fast.Next
				slow = slow.Next
			}
		}
	}
	fmt.Println(false)
}
