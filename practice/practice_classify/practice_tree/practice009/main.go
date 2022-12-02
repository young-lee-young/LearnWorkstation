/**
LeetCode No199 二叉树的右视图

解题思路：二叉树的层序遍历，取每层最后一个值
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
	ret := solution(root)
	fmt.Println("ret:", ret)
}

func solution(root *TreeNode) []int {
	ret := make([]int, 0)

	queue := Queue{}
	queue.Enqueue(root)

	for !queue.IsEmpty() {
		queueLen := queue.GetLength()

		for i := 0; i < queueLen; i++ {
			node := queue.Dequeue().(*TreeNode)
			if node == nil {
				break
			}

			if node.Left != nil {
				queue.Enqueue(node.Left)
			}

			if node.Right != nil {
				queue.Enqueue(node.Right)
			}

			if i == queueLen-1 {
				ret = append(ret, node.Val)
			}
		}
	}

	return ret
}
