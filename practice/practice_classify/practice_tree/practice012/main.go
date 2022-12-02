package main

import (
	"fmt"
	"math"
	queue2 "practice/base/queue"
	tree2 "practice/base/tree"
)

/**
二叉树每层中的最大值 LeetCode No515

使用队列层次遍历
*/
func main() {
	tree := tree2.Tree{}
	tree.GenerateTree()
	root := tree.Root
	result := solution(root)
	fmt.Println(result)
}

func solution(root *tree2.Node) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	queue := queue2.Queue{}
	queue.Enqueue(root)

	for !queue.IsEmpty() {
		max := math.MinInt64
		queueLen := queue.GetLength()
		for i := 0; i < queueLen; i++ {
			node := queue.Dequeue().(*tree2.Node)
			if node.Data > max {
				max = node.Data
			}
			if node.Left != nil {
				queue.Enqueue(node.Left)
			}
			if node.Right != nil {
				queue.Enqueue(node.Right)
			}
		}
		result = append(result, max)
	}
	return result
}
