package main

import (
	tree2 "practice/base/tree"
	"fmt"
)

/**
	LeetCode No101 对称二叉树

	解题思路：递归，后序遍历
*/
func main() {
	tree := tree2.Tree{}
	tree.GenerateTree()
	root := tree.Root

	ret := solution(root)
	fmt.Println("ret:", ret)
}

func solution(root *tree2.Node) bool {
	if root == nil {
		return true
	}
	return isSymmetric(root.Left, root.Right)
}

func isSymmetric(left *tree2.Node, right *tree2.Node) bool {
	// 递归终止条件
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Data != right.Data {
		return false
	}

	// 单层递归逻辑，后序遍历
	outside := isSymmetric(left.Left, right.Right)	// 左
	inside := isSymmetric(left.Right, right.Left)	// 右
	return outside && inside						// 中
}
