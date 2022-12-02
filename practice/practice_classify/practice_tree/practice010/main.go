package main

import (
	"fmt"
	queue2 "practice/base/queue"
	tree2 "practice/base/tree"
	"strconv"
)

/**
LeetCode No637 二叉树的层平均值
*/
func main() {
	tree := tree2.Tree{}
	tree.GenerateTree()
	root := tree.Root
	solution(root)
}

func solution(root *tree2.Node) {
	result := make([]float64, 0)
	queue := queue2.Queue{}

	queue.Enqueue(root)
	for !queue.IsEmpty() {
		sum := 0
		count := queue.GetLength()
		for i := 0; i < count; i++ {
			node := queue.Dequeue().(*tree2.Node)
			sum += node.Data
			if node.Left != nil {
				queue.Enqueue(node.Left)
			}
			if node.Right != nil {
				queue.Enqueue(node.Right)
			}
		}
		avg, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(sum)/float64(count)), 64)
		result = append(result, avg)
	}
	fmt.Println(result)
}
