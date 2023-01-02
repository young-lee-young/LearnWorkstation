package main

import (
	"math"
	"fmt"
)

/**
	LeetCode No110 判断是否是平衡树

	平衡二叉树：左右子树高度差不超过1
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{}

	ret := solution(root)
	fmt.Println("ret:", ret)
}

func solution(node *TreeNode) int {
	// 递归终止条件
	if node == nil {
		return 0
	}

	// 单层递归逻辑，后序遍历
	leftDeep := solution(node.Left)
	if leftDeep == -1 {
		return -1
	}

	rightDeep := solution(node.Right)
	if rightDeep == -1 {
		return -1
	}

	absDeep := int(math.Abs(float64(leftDeep - rightDeep)))
	// 左右子树的高度差大于1，不是平衡二叉树
	if absDeep > 1 {
		return -1
	}

	return max(leftDeep, rightDeep) + 1
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
