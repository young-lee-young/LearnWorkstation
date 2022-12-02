package main

import (
	"fmt"
	tree2 "practice/base/tree"
)

/**
  判断对称 LeetCode No101
*/
func main() {
	tree := tree2.Tree{}
	tree.GenerateTree()
	root := tree.Root
	if root == nil {
		fmt.Println(true)
		return
	}
	result := isSymmetric(root.Left, root.Right)
	fmt.Println(result)
}

func isSymmetric(node1 *tree2.Node, node2 *tree2.Node) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	if node1.Data != node2.Data {
		return false
	}
	return isSymmetric(node1.Left, node2.Right) && isSymmetric(node1.Right, node2.Left)
}
