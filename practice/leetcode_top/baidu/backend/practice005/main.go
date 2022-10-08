package main

import (
	"fmt"
	queue2 "practice/base/queue"
	tree2 "practice/base/tree"
)

/**
二叉树的右视图 LeetCode No199

使用队列层序遍历，记录每层最后一个值
*/
func main() {
	tree := tree2.Tree{}
	tree.GenerateTree()
	tree.PreorderTraversal()
	root := tree.Root
	ret := solution(root)
	fmt.Println("ret", ret)
}

func solution(root *tree2.Node) []int {
	ret := make([]int, 0)

	if root == nil {
		return ret
	}

	queue := queue2.Queue{}
	queue.Enqueue(root)

	for !queue.IsEmpty() {
		queueLen := queue.GetLength()
		for i := 0; i < queueLen; i++ {
			node := queue.Dequeue().(*tree2.Node)

			if node.Left != nil {
				queue.Enqueue(node.Left)
			}

			if node.Right != nil {
				queue.Enqueue(node.Right)
			}

			if i == queueLen-1 {
				ret = append(ret, node.Data)
			}
		}
	}

	return ret
}
