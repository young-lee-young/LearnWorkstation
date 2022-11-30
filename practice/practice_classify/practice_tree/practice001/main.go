/**
 * @Time:    2022/11/29 17:43 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	LeetCode No144 二叉树的前序遍历

 */
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ret []int

func main() {
	ret = []int{}

	root := &TreeNode{}

	recursion(root)

	fmt.Println("ret:", ret)
}

func recursion(node *TreeNode) {
	// 递归结束条件
	if node == nil {
		return
	}

	ret = append(ret, node.Val)
	recursion(node.Left)
	recursion(node.Right)
}
