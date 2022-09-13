/**
 * @Time:    2022/4/20 22:37 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	环形链表起始点 LeetCode No142

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

	j := 0
	current := linkList.Head
	ringStartNode := &link.Node{}
	for current.Next != nil {
		current = current.Next
		if j == 5 {
			ringStartNode = current
		}
		j ++
	}
	current.Next = ringStartNode
	ret := solution(linkList.Head)
	fmt.Println("ret", ret)
}

func solution(head *link.Node) *link.Node {

}
