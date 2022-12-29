/**
 * @Time:    2022/12/9 15:47 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	LeetCode No530 二叉搜索树的最小绝对差
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

func main() {
	node1 := &TreeNode{Val: 1, Left: nil, Right: nil}
	node3 := &TreeNode{Val: 3, Left: nil, Right: nil}
	node2 := &TreeNode{Val: 2, Left: node1, Right: node3}
	node6 := &TreeNode{Val: 6, Left: nil, Right: nil}
	node4 := &TreeNode{Val: 4, Left: node2, Right: node6}

	ret = math.MaxInt64

	solution(node4.Left, node4.Val)

	solution(node4.Right, node4.Val)

	fmt.Println("ret:", ret)
}

func solution(node *TreeNode, parentValue int) {
	if node == nil {
		return
	}

	v := int(math.Abs(float64(node.Val - parentValue)))

	if v < ret {
		ret = v
	}

	if node.Left != nil {
		solution(node.Left, node.Val)
	}

	if node.Right != nil {
		solution(node.Right, node.Val)
	}
}
