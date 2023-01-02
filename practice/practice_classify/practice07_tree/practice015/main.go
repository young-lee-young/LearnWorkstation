package main

import (
	tree2 "practice/base/tree"
)

/**
	LeetCode No226 翻转二叉树

	解题思路：递归，前序遍历或后序遍历
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
	// 递归终止条件
	if node == nil {
		return nil
	}

	// 单层递归逻辑，前序遍历
	node.Left, node.Right = node.Right, node.Left	// 中
	solution(node.Left) 						  	// 左
	solution(node.Right)						  	// 右

	return node
}
