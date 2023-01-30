/**
LeetCode No450 删除二叉搜索树中的节点
*/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	val := 1

	root := &TreeNode{}

	ret := solution(root, val)

	fmt.Println("ret:", ret)
}

/**
解题思路见 iPad
 */
func solution(node *TreeNode, val int) *TreeNode {
	/**
	递归函数终止条件
	 */
	// 1. 没有发现目标节点
	if node == nil {
		return nil
	}

	// 找到删除的节点
	if node.Val == val {
		// 2. 删除节点是叶子节点
		if node.Left == nil && node.Right == nil {
			return nil
		}

		// 3. 左不为空，右为空
		if node.Left != nil && node.Right == nil {
			return node.Left
		}

		// 4. 左为空，右不为空
		if node.Left == nil && node.Right != nil {
			return node.Right
		}

		// 5. 左不为空，右不为空
		cur := node.Right
		// 找到右子树最左节点
		for cur.Left != nil {
			cur = cur.Left
		}
		// 将原节点的左子树挂到右子树最左节点
		cur.Left = node.Left
		// ⚠️：这里返回的是 node.Right，不是 cur
		return node.Right
	}
	/**
	终止条件
	 */

	if node.Val > val {
		node.Left = solution(node.Left, val)
	}

	if node.Val < val {
		node.Right = solution(node.Right, val)
	}

	return node
}
