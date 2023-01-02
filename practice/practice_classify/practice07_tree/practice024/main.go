package main

import "fmt"

/**
判断树路径和是否等于一个数 LeetCode No112
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ret bool

func main() {
	root := &TreeNode{}

	target := 0

	ret = false

	path := []int{}

	solution(root, target, path)

	fmt.Println("ret:", ret)
}

func solution(root *TreeNode, target int, path []int) {
	if root == nil {
		return
	}

	solution2(root, target, path)
}

func solution2(node *TreeNode, target int, path []int) {
	path = append(path, node.Val)

	if node.Left == nil && node.Right == nil {
		sum := sum(path)
		if sum == target {
			ret = true
			return
		}
	}

	if node.Left != nil {
		solution(node.Left, target, path)
	}

	if node.Right != nil {
		solution(node.Right, target, path)
	}
}

func sum(nums []int) int {
	s := 0
	for i := 0; i < len(nums); i ++ {
		s += nums[i]
	}
	return s
}
