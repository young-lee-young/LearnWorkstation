/**
LeetCode No669 修剪二叉搜索树
*/
package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{}

	low := 2

	high := 7

	ret := solution(root, low, high)

	fmt.Println("ret:", ret)
}

func solution(node *TreeNode, low int, high int) *TreeNode {
	if node == nil {
		return nil
	}

	// 当前节点值比 low 小，当前节点和左子树都不要了，只处理右子树，留下来符合条件的
	if node.Val < low {
		right := solution(node.Right, low, high)
		return right
	}

	// 当前节点值比 high 大，当前节点和右子树都不要了，只处理左子树，留下来符合条件的
	if node.Val > high {
		left := solution(node.Left, low, high)
		return left
	}

	// 处理左子树
	node.Left = solution(node.Left, low, high)

	// 处理右子树
	node.Right = solution(node.Right, low, high)

	return node
}
