package main

import "fmt"

/**
	搜索旋转排序数组，存在返回索引，不存在返回-1 LeetCode No33

	标签：二分搜索
 */
func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	result := solution(nums, target)
	fmt.Println(result)
}

func solution(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for start < end {
		mid := (start + end) / 2
		// 断点在中间
		if nums[mid] >= nums[0] && target < nums[0] {
			start = mid + 1
			continue
		}
		// 断点在后
		if nums[mid] >= nums[0] && (target > nums[mid] || target < nums[0]) {
			start = mid + 1
			continue
		}
		// 断点在前
		if target < nums[0] && target > nums[mid] {
			start = mid + 1
			continue
		}
		end = mid
	}
	if start == end && nums[start] == target {
		return start
	}
	return -1
}
