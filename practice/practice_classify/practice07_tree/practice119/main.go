/**
LeetCode No671 二叉树中第二小的节点
*/
package main

import (
	"fmt"
	tree2 "practice/base/tree"
)

func main() {
	tree := tree2.Tree{}
	tree.GenerateTree()
	root := tree.Root

	result := solution(root)
	fmt.Println(result)
}

func solution(node *tree2.Node) int {
	if node == nil || (node.Left == nil && node.Right == nil) {
		return -1
	}

	left := node.Left.Data
	right := node.Right.Data
	if left == node.Data {
		left = solution(node.Left)
	}
	if right == node.Data {
		right = solution(node.Right)
	}

	if left != -1 && right != -1 {
		if left < right {
			return left
		}
		return right
	} else if left != -1 {
		return left
	} else {
		return right
	}
}
