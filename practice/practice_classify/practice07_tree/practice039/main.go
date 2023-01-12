/**
LeetCode No538 把二叉搜索树转换为累加树
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

func solution(root *TreeNode) *TreeNode {

}
