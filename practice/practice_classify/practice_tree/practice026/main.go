package main

import (
	"fmt"
)

/**
LeetCode No105 根据前序遍历和中序遍历构建二叉树

标签：递归调用
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}

	inorder := []int{9, 3, 15, 20, 7}

	ret := solution(preorder, inorder)

	fmt.Println("ret:", ret)
}

func solution(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	rootValue := preorder[0]

	root := &TreeNode{
		Val: rootValue,
	}

	var inorderRootIndex int
	for i := 0; i < len(inorder); i ++ {
		if inorder[i] == rootValue {
			inorderRootIndex = i
			break
		}
	}

	inorderLeft := inorder[:inorderRootIndex]
	inorderRight := inorder[inorderRootIndex+1:]

	preorderLeft := preorder[1 : len(inorderLeft)+1]
	preorderRight := preorder[len(inorderLeft)+1:]

	left := solution(preorderLeft, inorderLeft)
	right := solution(preorderRight, inorderRight)

	root.Left = left
	root.Right = right

	return root
}
