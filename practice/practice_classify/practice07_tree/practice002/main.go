package main

import (
	tree2 "practice/base/tree"
	stack2 "practice/base/stack"
	"fmt"
)

/**
	LeetCode No144 二叉树的前序遍历（非递归实现）

	解题思路：使用栈

	前序遍历：中 左 右
 */
func main() {
	tree := tree2.Tree{}
	tree.GenerateTree()
	root := tree.Root
	tree.PreorderTraversal()

	stack := stack2.Stack{}
	stack.Push(root)

	for !stack.IsEmpty() {
		node := stack.Pop().(*tree2.Node)

		fmt.Println(node.Data)

		// 将右子树入栈
		if node.Right != nil {
			stack.Push(node.Right)
		}

		// 将左子树入栈
		if node.Left != nil {
			stack.Push(node.Left)
		}
	}
}
