package main

import (
	tree2 "practice/base/tree"
	"fmt"
)

/**
	LeetCode No104 二叉树的最大深度

	标签：递归

	根节点的最大深度 = 根节点的高度
 */
func main() {
	tree := tree2.Tree{}
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(4)
	root := tree.Root

	ret := solution(root)
	fmt.Println("ret:", ret)
}

func solution(root *tree2.Node) int {
	if root == nil {
		return 1
	}
	return getHeight(root)
}

func getHeight(node *tree2.Node) int {
	// 递归终止条件
	if node == nil {
		return 0
	}

	// 单层递归逻辑，后序遍历
	leftHeight := getHeight(node.Left)		// 左
	rightHeight := getHeight(node.Right)	// 右
	return max(leftHeight, rightHeight) + 1	// 中
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
