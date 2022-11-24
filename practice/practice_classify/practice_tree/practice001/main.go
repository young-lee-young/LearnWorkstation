package main

import (
	tree2 "practice/base/tree"
	stack2 "practice/base/stack"
	"fmt"
)

/**
	非递归实现二叉树前序遍历 LeetCode No144

	解题思路：使用栈
 */
func main() {
	tree := tree2.Tree{}
	tree.GenerateTree()
	root := tree.Root
	tree.PreorderTraversal()

	fmt.Println("-------------------")
	stack := stack2.Stack{}
	stack.Push(root)
	for !stack.IsEmpty() {
		node := stack.Pop().(*tree2.Node)
		fmt.Println(node.Data)
		if node.Right != nil {
			stack.Push(node.Right)
		}
		if node.Left != nil {
			stack.Push(node.Left)
		}
	}
}
