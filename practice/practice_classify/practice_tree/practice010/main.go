package main

import (
	tree2 "practice/base/tree"
)

/**
	归并两棵树 LeetCode No617
 */
func main() {
	treeOne := tree2.Tree{}
	treeOne.GenerateTree()
	treeTwo := tree2.Tree{}
	treeTwo.GenerateTree()

	nodeOne := treeOne.Root
	nodeTwo := treeTwo.Root

	newTree := tree2.Tree{}
	newRoot := solution(nodeOne, nodeTwo)
	newTree.Root = newRoot
	newTree.PreorderTraversal()
}

func solution(nodeOne *tree2.Node, nodeTwo *tree2.Node) *tree2.Node {
	if nodeOne == nil && nodeTwo == nil {
		return  nil
	}
	if nodeOne == nil {
		return nodeTwo
	}
	if nodeTwo == nil {
		return nodeOne
	}
	newNode := &tree2.Node{
		Data: nodeOne.Data + nodeTwo.Data,
	}
	newNode.Left = solution(nodeOne.Left, nodeTwo.Left)
	newNode.Right = solution(nodeOne.Right, nodeTwo.Right)
	return newNode
}
