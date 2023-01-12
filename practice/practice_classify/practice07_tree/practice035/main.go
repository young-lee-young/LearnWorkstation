/**
LeetCode No701 二叉搜索树中的插入操作
*/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	v := 1

	ret := solution(nil, v)

	fmt.Println("ret addr", ret)
}

func solution(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{Val: val}
		return root
	}

	if root.Val > val {
		root.Left = solution(root.Left, val)
	} else {
		root.Right = solution(root.Right, val)
	}

	return root
}
