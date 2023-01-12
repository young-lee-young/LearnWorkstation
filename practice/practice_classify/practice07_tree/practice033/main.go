/**
LeetCode No236 二叉树最近公共祖先

标签：深度优先遍历，后序遍历
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

	// 去左子树里面找 p 和 q
	left := solution(node.Left, p, q) // 左

	// 去右子树里面找 p 和 q
	right := solution(node.Right, p, q) // 右

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
