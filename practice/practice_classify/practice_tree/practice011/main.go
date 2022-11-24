package main

import (
	tree2 "practice/base/tree"
	"fmt"
)

/**
	判断树路径和是否等于一个数 LeetCode No112
 */
func main() {
	tree := tree2.Tree{}
	tree.GenerateTree()
	root := tree.Root
	result := solution(root, 98)
	fmt.Println(result)
}

func solution(node *tree2.Node, sum int) bool {
	if node == nil {
		return false
	}
	if node.Left == nil && node.Right == nil && node.Data == sum {
		return true
	}
	return solution(node.Left, sum - node.Data) || solution(node.Right, sum - node.Data)
}
