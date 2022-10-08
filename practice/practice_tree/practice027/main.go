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

		/**
		前序：根     左     右
		     0    i + 1
		中序：左     根     右
		     0      i

		中序的根索引为 i，则 左子树有 i 个值，右子树有 i - end 个值

		在前序遍历中：左子树区间 [1 : i + 1]，共 i 个值；右子树区间 [i + 1: ]
		在中序遍历中：左子树区间 [0 : i]，共 i 个值; 右子树区间 [i + 1: ]
		*/

		if inorder[i] == data {
			// 左子树
			// 左子树的前序遍历值
			preorderLeft := preorder[1 : i+1]
			// 左子树的中序遍历值
			inorderLeft := inorder[:i]
			node.Left = solution(preorderLeft, inorderLeft)

			// 右子树
			// 右子树的前序遍历值
			// 右子树的中序遍历值
			preorderRight := preorder[i+1:]
			inorderRight := inorder[i+1:]
			node.Right = solution(preorderRight, inorderRight)
		}
	}

	return node
}
