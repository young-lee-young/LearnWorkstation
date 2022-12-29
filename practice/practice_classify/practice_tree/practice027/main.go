/**
 * @Time:    2022/12/9 15:00 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	LeetCode No654 最大二叉树
 */
package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	nums := []int{3, 2, 1, 6, 0, 5}

	ret := solution(nums)

	fmt.Println("ret:", ret)
}

func solution(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	max := math.MinInt64
	var maxIndex int

	for i := 0; i < len(nums); i ++ {
		if nums[i] > max {
			maxIndex = i
			max = nums[maxIndex]
		}
	}

	numsLeft := nums[0:maxIndex]
	numsRight := nums[maxIndex+1:]

	left := solution(numsLeft)
	right := solution(numsRight)

	root := &TreeNode{
		Val:   max,
		Left:  left,
		Right: right,
	}

	return root
}
