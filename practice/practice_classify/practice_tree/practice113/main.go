package main

import (
	"fmt"
	tree2 "practice/base/tree"
)

/**
判断是否是子树 LeetCode No572
*/
func main() {
	tree := tree2.Tree{}
	tree.GenerateTree()
	root := tree.Root

	subTree := tree2.Tree{}
	subTree.GenerateTree2()
	root2 := subTree.Root

	if root == nil {
		fmt.Println(false)
	}

	result := isSubtree(root, root2)
	fmt.Println(result)
}

func isSubtree(node1 *tree2.Node, node2 *tree2.Node) bool {
	if node1 == nil {
		return false
	}
	return isSubtreeWithRoot(node1, node2) || isSubtree(node1.Left, node2) || isSubtree(node1.Right, node2)
}

func isSubtreeWithRoot(node1 *tree2.Node, node2 *tree2.Node) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	if node1.Data != node2.Data {
		return false
	}
	return isSubtreeWithRoot(node1.Left, node2.Left) && isSubtreeWithRoot(node1.Right, node2.Right)
}
