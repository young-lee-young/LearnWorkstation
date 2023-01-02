/**
 * @Time:    2022/11/29 17:56 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	LeetCode No145 二叉树的后序遍历

 */
package main

import "fmt"

var ret []int

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	ret = []int{}

	root := &TreeNode{}

	recursion(root)

	fmt.Println("ret:", ret)
}

func recursion(node *TreeNode) {
	// 递归终止条件
	if node == nil {
		return
	}

	recursion(node.Left)
	recursion(node.Right)
	ret = append(ret, node.Val)
}
