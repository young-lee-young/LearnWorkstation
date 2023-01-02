/**
 * @Time:    2022/12/9 15:47 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	LeetCode No530 二叉搜索树的最小绝对差

	解题思路：最小绝对差一定是相邻的两个数的差，使用中序遍历使二叉树有序
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

var ret int

var pre *TreeNode

func main() {
	node1 := &TreeNode{Val: 1, Left: nil, Right: nil}
	node3 := &TreeNode{Val: 3, Left: nil, Right: nil}
	node2 := &TreeNode{Val: 2, Left: node1, Right: node3}
	node6 := &TreeNode{Val: 6, Left: nil, Right: nil}
	node4 := &TreeNode{Val: 4, Left: node2, Right: node6}

	ret = math.MaxInt64

	pre = nil

	solution(node4)

	fmt.Println("ret:", ret)
}

func solution(cur *TreeNode) {
	// 递归终止条件
	if cur == nil {
		return
	}

	// 单层递归逻辑，中序遍历
	solution(cur.Left) // 左

	if pre != nil { // 中
		ret = min(ret, cur.Val-pre.Val)
	}

	pre = cur

	solution(cur.Right) // 右
}

func min(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}
