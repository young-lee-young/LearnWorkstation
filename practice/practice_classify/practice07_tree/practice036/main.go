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

func solution(root *TreeNode, val int) *TreeNode {

}
