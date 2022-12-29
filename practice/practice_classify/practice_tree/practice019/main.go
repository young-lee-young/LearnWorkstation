/**
 * @Time:    2022/12/5 23:57 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	LeetCode No222 完全二叉树的节点个数

	利用满二叉树的性质：满二叉树的节点个数为 (2 ^ height) - 1
 */
package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node6 := &TreeNode{Val: 6, Left: nil, Right: nil}
	node5 := &TreeNode{Val: 5, Left: nil, Right: nil}
	node4 := &TreeNode{Val: 4, Left: nil, Right: nil}
	node3 := &TreeNode{Val: 3, Left: node6, Right: nil}
	node2 := &TreeNode{Val: 2, Left: node4, Right: node5}
	node1 := &TreeNode{Val: 1, Left: node2, Right: node3}

	ret := solution2(node1)
	fmt.Println("ret:", ret)
}

// 方法1：遍历所有节点
func solution(node *TreeNode) int {
	// 递归终止条件
	if node == nil {
		return 0
	}

	// 单层递归逻辑，后序遍历
	left := solution(node.Left)   // 左
	right := solution(node.Right) // 右
	return left + right + 1       // 中
}

func solution2(node *TreeNode) int {
	// 递归终止条件
	if node == nil {
		return 0
	}

	// 单层递归逻辑
	leftHeight := getLeftHeight(node.Left)
	rightHeight := getRightHeight(node.Right)

	if leftHeight == rightHeight {
		// ⚠️注意：这里的 leftHeight 要加 1
		return int(math.Pow(2, float64(leftHeight+1))) - 1
	}

	left := solution2(node.Left)
	right := solution2(node.Right)
	return left + right + 1
}

// 获取左侧高度
func getLeftHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftHeight := getLeftHeight(node.Left)

	return leftHeight + 1
}

func getRightHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}

	rightHeight := getRightHeight(node.Right)

	return rightHeight + 1
}
