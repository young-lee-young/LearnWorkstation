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
	// 递归终止条件
	if node == nil {
		return nil
	}

	// 在树里进行查找，在子树里找到了 p 或 q
	if node == p || node == q {
		return node
	}

	var left *TreeNode
	var right *TreeNode

	// 根据二叉搜索树性质
	if p.Val < node.Val && q.Val < node.Val {
		// p 和 q 都在 node 左边
		left = solution(node.Left, p, q)
	} else if p.Val > node.Val && q.Val > node.Val {
		// p 和 q 都在 node 右边
		right = solution(node.Right, p, q)
	} else {
		// p 和 q 在 node 两边
		left = solution(node.Left, p, q)
		right = solution(node.Right, p, q)
	}

	// 中
	// 左、右子树 分别有 p 和 q，则当前节点即为最近公共祖先
	if left != nil && right != nil {
		return node
	}

	// 左子树里有 p 或 q
	if left != nil {
		return left
	}

	// 右子树里有 p 或 q
	if right != nil {
		return right
	}

	// 左、右子树都没有 p 和 q
	return nil
}

func solution2(node *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {

}
