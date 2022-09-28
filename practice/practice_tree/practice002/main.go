package main

import (
	"fmt"
	stack2 "practice/base/stack"
	tree2 "practice/base/tree"
)

/**
非递归实现二叉树中序遍历 LeetCode No94

解题思路：使用栈
*/
func main() {
	tree := tree2.Tree{}
	tree.GenerateTree()
	root := tree.Root
	tree.InorderTraversal()

	stack := stack2.Stack{}
	current := root
	for !stack.IsEmpty() || current != nil {
		if current != nil {
			stack.Push(current)
			current = current.Left
		} else {
			current = stack.Pop().(*tree2.Node)
			fmt.Println(current.Data)
			current = current.Right
		}
	}
}
