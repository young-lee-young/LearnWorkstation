package main

import (
	"fmt"
	queue2 "practice/base/queue"
	tree2 "practice/base/tree"
)

/**
	LeetCode No102 二叉树的层序遍历

	解题思路：使用队列层序遍历
*/
func main() {
	tree := tree2.Tree{}
	tree.GenerateTree()
	root := tree.Root
	result := solution(root)
	fmt.Println(result)
}

func solution(root *tree2.Node) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := queue2.Queue{}
	queue.Enqueue(root)

	for !queue.IsEmpty() {
		temp := make([]int, 0)
		queueLen := queue.GetLength()
		for i := 0; i < queueLen; i++ {
			node := queue.Dequeue().(*tree2.Node)

			// ⚠️ 这里要注意，需要判断空
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

		// ⚠️ 空的地址不要加进来
		if len(temp) != 0 {
			result = append(result, temp)
		}
	}

	return result
}
