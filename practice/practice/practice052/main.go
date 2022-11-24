/**
 * @Time:    2021/3/8 20:22 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content: 左叶子之和
 */
package main

import (
	tree2 "practice/base/tree"
	"fmt"
)

/**
	左叶子之和 LeetCode No404

	标签：二叉树、递归
 */
func main() {
	tree := tree2.Tree{}
	tree.GenerateTree()
	result := solution(tree.Root)
	fmt.Println(result)
}

func solution(node *tree2.Node) int {
	if node == nil {
		return 0
	}
	sum := 0
	// 是左叶子节点
	if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
		sum += node.Left.Data
	}
	sumLeft := solution(node.Left)
	sumRight := solution(node.Right)
	return sum + sumLeft + sumRight
}
