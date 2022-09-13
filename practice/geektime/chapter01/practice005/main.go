/**
 * @Time:    2022/4/20 19:26 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	K个一组反转链表 LeetCode No25

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

	ret := solution(linkList.Head, 3)

	for ret != nil {
		fmt.Println(ret.Data)
		ret = ret.Next
	}
}

func solution(head *link.Node, k int) *link.Node {

	// 保护节点
	protect := &link.Node{0, head}
	last := protect

	for head != nil {
		// 1. 分组，往后走 k - 1 步，找到一组
		// 一组的开头为 head，结尾为 end
		end := findEnd(head, k)
		if end == nil {
			break
		}

		// 下一组的头节点
		nextGroupHead := end.Next

		// 2. 一组内部 (head - end 之间) 要反转
		reverse(head, nextGroupHead)

		// 3. 更新每组跟前一组、后一组的边
		last.Next = end
		head.Next = nextGroupHead


		last = head
		head = nextGroupHead
	}

	return protect.Next
}

/**
	返回走 k - 1 步之后的节点
	返回 nil 表示不够 k 个
 */
func findEnd(head *link.Node, k int) *link.Node {
	for head != nil {
		k --
		if k == 0 {
			return head
		}
		head = head.Next
	}
	return nil
}

/**
	反转链表，在节点 stop 停止
 */
func reverse(head *link.Node, stop *link.Node) {
	lastNode := head
	head = head.Next
	for head != stop {
		nextNode := head.Next
		head.Next = lastNode
		lastNode = head
		head = nextNode
	}
}
