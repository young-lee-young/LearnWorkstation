/**
 * @Time:    2022/11/28 21:22 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	🥇标准实现

	LeetCode No704 二分查找

 */
package main

import "fmt"

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}

	target := 9

	ret := solution2(nums, target)

	fmt.Println("ret:", ret)
}

// 左闭右闭区间 [left, right]
func solution(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	// ⚠️注意：左闭右闭区间，left 可以 <= right
	for left <= right {
		// ⚠️注意：这里的 + 1
		mid := left + (right-left+1)/2

		if target > nums[mid] {
			// ⚠️注意：左闭右闭区间，target 不等于 nums[mid]，所以边界要 mid + 1
			left = mid + 1
		}

		if target < nums[mid] {
			// ⚠️注意：和上面相同
			right = mid - 1
		}

		if target == nums[mid] {
			return mid
		}
	}

	return -1
}

// 左闭右开区间 [left, right)
func solution2(nums []int, target int) int {
	left := 0
	right := len(nums)

	for left < right {
		// ⚠️注意：这里没有 + 1
		mid := left + (right-left)/2

		if target > nums[mid] {
			left = mid + 1
		}

		if target < nums[mid] {
			right = mid
		}

		if target == nums[mid] {
			return mid
		}
	}

	return -1
}
