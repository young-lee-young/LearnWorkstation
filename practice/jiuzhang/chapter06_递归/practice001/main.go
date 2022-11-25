package main

import "fmt"

/**
🥇标准实现

递归实现二分搜索
*/
func main() {
	nums := []int{1, 3, 5, 7, 9}
	target := 5
	ret := BinarySearch(nums, 0, len(nums)-1, target)
	fmt.Println("ret:", ret)
}

func BinarySearch(nums []int, start int, end int, target int) int {
	// 递归结束条件
	if start > end {
		return -1
	}

	mid := (start + end) / 2

	// target 在 mid 右边
	if nums[mid] < target {
		return BinarySearch(nums, mid+1, end, target)
	}

	// target 在 mid 左边
	if nums[mid] > target {
		return BinarySearch(nums, start, mid-1, target)
	}

	return mid
}
