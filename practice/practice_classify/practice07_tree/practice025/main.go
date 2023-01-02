/**
 * @Time:    2022/12/8 13:20 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	LeetCode No106 从中序与后序遍历构造二叉树

	解题步骤：
		1. 后序数组为空，空节点
		2. 后序数组中最后一个元素为节点元素
		3. 寻找中序数组位置作为切割点
		4. 切中序数组
		5. 切后序数组
		6. 递归处理左区间和右区间
 */
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	inorder := []int{9, 3, 15, 20, 7}

	postorder := []int{9, 15, 7, 20, 3}

	ret := solution(inorder, postorder)

	fmt.Println("ret:", ret)
}

func solution(inorder []int, postorder []int) *TreeNode {
	// 递归终止条件：后序数组为空，根据'解题步骤' 1，说明是空节点
	if len(postorder) == 0 {
		return nil
	}

	// '解题步骤' 2
	rootValue := postorder[len(postorder)-1]

	root := &TreeNode{
		Val: rootValue,
	}

	// '解题步骤' 3
	var inorderRootIndex int
	for i := 0; i < len(inorder); i ++ {
		if inorder[i] == rootValue {
			inorderRootIndex = i
			break
		}
	}

	// '解题步骤' 4
	inorderLeft := inorder[:inorderRootIndex]
	inorderRight := inorder[inorderRootIndex+1:]

	// '解题步骤' 5：切后序遍历时，根据中序数组切割后长度切割后序数组
	postorderLeft := postorder[:len(inorderLeft)]
	postorderRight := postorder[len(inorderLeft) : len(postorder)-1]

	// '解题步骤' 6
	left := solution(inorderLeft, postorderLeft)
	right := solution(inorderRight, postorderRight)

	root.Left = left
	root.Right = right

	return root
}
