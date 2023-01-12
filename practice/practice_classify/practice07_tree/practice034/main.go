/**
 * @Time:    2023/1/2 20:10
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	LeetCode No235 二叉搜索树的最近公共祖先
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

	p := &TreeNode{}
	q := &TreeNode{}

	ret := solution(root, p, q)
	fmt.Println("ret:", ret)
}

func solution(node *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if node.Val > p.Val && node.Val > q.Val {
		return solution(node.Left, p, q)
	}
	if node.Val < p.Val && node.Val < q.Val {
		return solution(node.Right, p, q)
	}
	return node
}
