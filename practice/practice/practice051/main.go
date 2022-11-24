/**
 * @Time:    2021/3/8 19:47 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content: 对称二叉树
 */
package main

import (
	tree2 "practice/base/tree"
	"fmt"
)

/**
	判断树是否为对称二叉树 LeetCode 101

	标签：递归
 */
func main() {
	tree := tree2.Tree{}
	tree.GenerateTree()
	result := solution(tree.Root)
	fmt.Println(result)
}

func solution(root *tree2.Node) bool {
	if root == nil {
		return true
	}
	return recursion(root.Left, root.Right)
}

func recursion(left *tree2.Node, right *tree2.Node) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Data == right.Data && recursion(left.Left, right.Right) && recursion(left.Right, right.Left)
}
