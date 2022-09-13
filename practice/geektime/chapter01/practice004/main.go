/**
 * @Time:    2022/4/20 19:08 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	反转链表 LeetCode No206

 */
package main

import (
	"practice/base/link"
	"fmt"
)

func main() {
	linkList := link.LinkedList{}
	for i := 0; i < 10; i ++ {
		linkList.Append(i)
	}
	linkList.ShowList()

	ret := solution(linkList.Head)

	for ret != nil {
		fmt.Println(ret.Data)
		ret = ret.Next
	}
}

func solution(head *link.Node) *link.Node {
	var lastNode *link.Node
	for head != nil {
		nextNode := head.Next
		head.Next = lastNode
		lastNode = head
		head = nextNode
	}
	return lastNode
}
