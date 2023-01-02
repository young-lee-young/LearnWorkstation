/**
 * @Time:    2022/12/9 15:36 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	LeetCode No700 二叉搜索树中的搜索
 */
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{}

	val := 0

	ret := solution(root, val)

	fmt.Println("ret:", ret)
}

func solution(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return nil
	}

	if node.Val > val {
		return solution(node.Left, val)
	}

	if node.Val < val {
		return solution(node.Right, val)
	}

	return node
}
