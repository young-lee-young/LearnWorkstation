package main

import (
	"fmt"
)

/**
	旋转排序数组中最小值 LeetCode No153

	标签：二分搜索
 */
func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	result := solution(nums)
	fmt.Println(result)
}

func solution(nums []int) int {
	start := 0
	end := len(nums) - 1

	for start < end {
		mid := (start + end) / 2
		if nums[mid] > nums[end] {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return nums[start]
}
