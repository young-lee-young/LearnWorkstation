/**
 * @Time:    2022/12/7 23:32 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	LeetCode No513 找树左下角的值
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

var maxDepth int

func main() {
	root := &TreeNode{}

	maxDepth = math.MinInt64

	solution(root, 0)

	fmt.Println("ret:", ret)
}

func solution(node *TreeNode, depth int) {
	// 递归结束条件
	if node.Left == nil && node.Right == nil {
		// 单层递归逻辑
		if depth > maxDepth {
			maxDepth = depth
			ret = node.Val
		}
	}

	if node.Left != nil {
		solution(node.Left, depth+1)
	}

	if node.Right != nil {
		solution(node.Right, depth+1)
	}
}
