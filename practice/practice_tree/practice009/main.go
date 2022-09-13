package main

import (
	tree2 "practice/base/tree"
)

/**
	翻转二叉树 LeetCode No226
 */
func main() {
	tree := tree2.Tree{}
	tree.GenerateTree()
	root := tree.Root
	newRoot := solution(root)
	newTree := tree2.Tree{}
	newTree.Root = newRoot
	newTree.PreorderTraversal()
}

func solution(node *tree2.Node) *tree2.Node {
	if node == nil {
		return nil
	}
	leftTree := node.Left
	node.Left = solution(node.Right)
	node.Right = solution(leftTree)
	return node
}
