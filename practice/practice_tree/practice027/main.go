package main

import (
	"fmt"
	"practice/base/tree"
)

/**
根据前序遍历和后序遍历构建二叉树 LeetCode No105

标签：递归调用
*/
func main() {
	preorder := []int{20, 10, 5, 2, 7, 15, 12, 18, 30, 25, 23, 28, 35, 31, 38}
	inorder := []int{2, 5, 7, 10, 12, 15, 18, 20, 23, 25, 28, 30, 31, 35, 38}
	newRoot := solution(preorder, inorder)
	fmt.Println(newRoot)
	tree := tree.Tree{}
	tree.Root = newRoot
	tree.InorderTraversal()
}

func solution(preorder []int, inorder []int) *tree.Node {
	fmt.Println(preorder)
	fmt.Println(inorder)
	// 递归结束条件
	if len(preorder) == 0 {
		return nil
	}
	// 当前树的根节点
	data := preorder[0]
	node := &tree.Node{Data: data}

	// 遍历中序列表
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == data {
			// 左子树
			preorderLeft := preorder[1 : i+1]
			// 中序遍历，前面是左子树，后面是右子树
			inorderLeft := inorder[:i]
			node.Left = solution(preorderLeft, inorderLeft)

			// 右子树
			preorderRight := preorder[i+1:]
			inorderRight := inorder[i+1:]
			node.Right = solution(preorderRight, inorderRight)
		}
	}
	return node
}
