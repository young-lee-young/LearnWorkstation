/**
LeetCode No538 把二叉搜索树转换为累加树
*/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var pre int

func main() {
	root := &TreeNode{}

	pre = 0

	ret := solution(root)

	fmt.Println("ret:", ret)
}

/**
解题思路：

双指针
 */
func solution(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	// 右
	solution(node.Right)

	// 中
	node.Val = node.Val + pre

	pre = node.Val

	// 左
	solution(node.Left)

	return node
}
