package main

import (
	"fmt"
	tree2 "practice/base/tree"
)

/**
LeetCode No111 二叉树的最小深度

标签：递归
*/
func main() {
	tree := tree2.Tree{}
	tree.GenerateTree()
	root := tree.Root

	ret := solution(root)
	fmt.Println("ret:", ret)
}

func solution(node *tree2.Node) int {
	// 递归终止条件
	if node == nil {
		return 0
	}

	if node.Left == nil {
		return solution(node.Right) + 1
	}

	if node.Right == nil {
		return solution(node.Left) + 1
	}

	leftHeight := solution(node.Left)		// 左
	rightHeight := solution(node.Right)		// 右
	return min(leftHeight, rightHeight) + 1	// 中
}

func min(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}
