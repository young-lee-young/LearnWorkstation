package main

import (
	tree2 "practice/base/tree"
	"fmt"
)

/**
	二叉树最近公共祖先 LeetCode No236

	标签：深度优先遍历
 */
func main() {
	tree := tree2.Tree{}
	tree.GenerateTree()
	root := tree.Root
	nodeP := &tree2.Node{
		Data: 2,
	}
	nodeQ := &tree2.Node{
		Data: 5,
	}
	node := solution(root, nodeP, nodeQ)
	fmt.Println(node.Data)
}

func solution(node *tree2.Node, p *tree2.Node, q *tree2.Node) *tree2.Node {
	// 在树里进行查找，在子树里找到了p或q
	if node == nil || node.Data == p.Data || node.Data == q.Data {
		return node
	}

	// 去左子树里面找P和Q
	left := solution(node.Left, p, q)
	// 去右子树里面找P和Q
	right := solution(node.Right, p, q)

	// 如果左子树里面没有P或Q
	if left == nil {
		return right
	}
	// 如果右子树里没有P或Q
	if right == nil {
		return left
	}
	// 左右子树都里有P和Q，说明当前节点就是最近祖先节点
	return node
}
