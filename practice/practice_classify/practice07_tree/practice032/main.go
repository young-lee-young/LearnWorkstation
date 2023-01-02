/**
 * @Time:    2023/1/2 18:33 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content: 

	LeetCode No501 二叉搜索树中的众数
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

var pre *TreeNode

var count int

var maxCount int

var ret []int

func main() {
	node2_2 := &TreeNode{Val: 2, Left: nil, Right: nil}
	node2_1 := &TreeNode{Val: 2, Left: node2_2, Right: nil}
	root := &TreeNode{Val: 1, Left: nil, Right: node2_1}

	pre = nil

	count = 0

	maxCount = math.MinInt64

	ret = make([]int, 0)

	solution(root)

	fmt.Println("ret:", ret)
}

func solution(cur *TreeNode) {
	// 递归终止条件
	if cur == nil {
		return
	}

	// 单层递归逻辑，中序遍历
	solution(cur.Left) // 左

	// ⚠️：这里的逻辑需要注意
	if pre == nil { // 中
		count = 1
	} else if cur.Val == pre.Val {
		count ++
	} else {
		count = 1
	}

	pre = cur

	if count == maxCount {
		ret = append(ret, cur.Val)
	}

	if count > maxCount {
		maxCount = count

		ret = ret[:0]
		ret = append(ret, cur.Val)
	}

	solution(cur.Right) // 右
}
