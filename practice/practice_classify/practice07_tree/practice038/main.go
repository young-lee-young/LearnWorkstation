/**
LeetCode No108 将有序数组转换为二叉搜索树
*/
package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	nums := []int{}

	ret := solution(nums)

	fmt.Println("ret:", ret)
}

/**
解题思路：

每次分割数组，取中间节点作为根节点
 */
func solution(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2

	numsLeft := nums[:mid]
	numsRight := nums[mid+1:]

	node := &TreeNode{Val: nums[mid]}
	node.Left = solution(numsLeft)
	node.Right = solution(numsRight)

	return node
}
