package main

import (
	tree2 "practice/base/tree"
	"practice/base/stack"
	"fmt"
)

/**
	非递归后续遍历二叉树 LeetCode No145

	解题思路：使用双栈
 */
func main() {
	tree := tree2.Tree{}
	tree.GenerateTree()
	root := tree.Root
	tree.PostorderTraversal()

	fmt.Println("-----------")
	stack1 := stack.Stack{}
	stack2 := stack.Stack{}
	current := root
	stack1.Push(current)

	for !stack1.IsEmpty() {
		current = stack1.Pop().(*tree2.Node)
		stack2.Push(current)
		if current.Left != nil {
			stack1.Push(current.Left)
		}
		if current.Right != nil {
			stack1.Push(current.Right)
		}
	}

	for !stack2.IsEmpty() {
		fmt.Println(stack2.Pop().(*tree2.Node).Data)
	}
}
