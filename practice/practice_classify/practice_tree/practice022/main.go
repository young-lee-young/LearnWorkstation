/**
 * @Time:    2022/12/7 23:09 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	LeetCode No404 左叶子之和
 */
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum int

func main() {
	root := &TreeNode{}

	sum = 0

	solution(root)
	fmt.Println("ret:", sum)
}

func solution(node *TreeNode) {
	// 递归终止条件
	if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
		sum += node.Left.Val
	}

	if node.Left != nil {
		solution(node.Left)
	}

	if node.Right != nil {
		solution(node.Right)
	}
}
