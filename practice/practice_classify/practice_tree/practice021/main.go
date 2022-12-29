/**
 * @Time:    2022/12/7 22:43 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	LeetCode No257 二叉树的所有路径
 */
package main

import (
	"fmt"
	"strings"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ret []string

func main() {
	root := &TreeNode{}

	ret = []string{}

	path := []string{}

	solution(root, path)

	fmt.Println("ret:", ret)
}

func solution(node *TreeNode, nums []string) {
	nums = append(nums, strconv.Itoa(node.Val))

	// 递归终止条件
	if node.Left == nil && node.Right == nil {
		temp := strings.Join(nums, "->")
		ret = append(ret, temp)
		return
	}

	// 单层递归逻辑
	if node.Left != nil {
		solution(node.Left, nums)
	}

	if node.Right != nil {
		solution(node.Right, nums)
	}
}
