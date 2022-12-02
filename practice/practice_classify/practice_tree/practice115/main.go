package main

import (
	"fmt"
	tree2 "practice/base/tree"
)

/**
树最小深度 LeetCode No111

标签：递归
*/
func main() {
	tree := tree2.Tree{}
	tree.GenerateTree()
	root := tree.Root

	minDepth := solution(root)
	fmt.Println(minDepth)
}

func solution(node *tree2.Node) int {
	if node == nil {
		return 0
	}
	if node.Left == nil {
		return solution(node.Right) + 1
	}
	if node.Right == nil {
		return solution(node.Left) + 1
	}
	leftDepth := solution(node.Left)
	rightDepth := solution(node.Right)
	if leftDepth < rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}
