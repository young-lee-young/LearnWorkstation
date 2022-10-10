package main

import (
	"fmt"
	"practice/base/link"
)

/**
反转链表Ⅱ LeetCode No092
*/
func main() {
	linkList := link.LinkedList{}
	for i := 0; i < 10; i++ {
		linkList.Append(i)
	}
	linkList.ShowList()

	ret := solution(linkList.Head, 2, 5)
	for ret != nil {
		fmt.Println(ret.Data)
		ret = ret.Next
	}
}

func solution(head *link.Node, left int, right int) *link.Node {

}
