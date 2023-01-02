/**
 * @Time:    2022/11/29 19:04
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	LeetCode No107 二叉树的层序遍历Ⅱ

	解题思路：使用队列层序遍历，最后结果集数组翻转
*/
package main

import (
	"fmt"
	queue2 "practice/base/queue"
	tree2 "practice/base/tree"
)

func main() {
	tree := tree2.Tree{}
	tree.GenerateTree()
	root := tree.Root
	ret := solution(root)
	fmt.Println("ret:", ret)
}

func solution(root *tree2.Node) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}

	queue := queue2.Queue{}
	queue.Enqueue(root)

	for !queue.IsEmpty() {
		temp := make([]int, 0)
		queueLen := queue.GetLength()

		for i := 0; i < queueLen; i++ {
			node := queue.Dequeue().(*tree2.Node)

			if node == nil {
				break
			}

			temp = append(temp, node.Data)

			if node.Left != nil {
				queue.Enqueue(node.Left)
			}
			if node.Right != nil {
				queue.Enqueue(node.Right)
			}
		}

		if len(temp) != 0 {
			ret = append(ret, temp)
		}
	}

	// 翻转数组
	reverse(ret)

	return ret
}

func reverse(a [][]int) {
	l, r := 0, len(a)-1
	for l < r {
		a[l], a[r] = a[r], a[l]
		l, r = l+1, r-1
	}
}
