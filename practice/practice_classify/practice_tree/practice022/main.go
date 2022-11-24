package main

import (
	tree2 "practice/base/tree"
	"math"
	"fmt"
)

/**
	验证是否是二叉搜索树 LeetCode No98

	解题思路：递归，左子树最大值为根节点，最小值无限小，右子树最小值为根节点，最大值无限大
 */
func main() {
	tree := tree2.Tree{}
	tree.GenerateTree()
	root := tree.Root
	result := solution(root, math.MinInt64, math.MaxInt64)
	fmt.Println(result)
}

func solution(node *tree2.Node, min int, max int) bool {
	if node == nil {
		return true
	}
	if node.Data <= min || node.Data >= max {
		return false
	}
	return solution(node.Left, min, node.Data) && solution(node.Right, node.Data, max)
}
