package main

import (
	tree2 "practice/base/tree"
	"fmt"
)

/**
	统计左叶子节点的和 LeetCode No404
 */
func main() {
	tree := tree2.Tree{}
	tree.GenerateTree()
	root := tree.Root

	result := solution(root)
	fmt.Println(result)
}

func solution(node *tree2.Node) int {
	if node == nil {
		return 0
	}
	if isLeaf(node.Left) {
		return node.Left.Data + solution(node.Right)
	}
	return solution(node.Left) + solution(node.Right)
}

func isLeaf(node *tree2.Node) bool {
	if node == nil {
		return false
	}
	return node.Left == nil && node.Right == nil
}
