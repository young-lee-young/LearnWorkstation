/**
 * @Time:    2021/6/11 23:16
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:
 */
package main

import (
	"fmt"
	"practice/utils/compare"
)

/**
最短连续子数组和大于等于给定值 LeetCode No209

标签：滑动窗口
*/
func main() {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	ret := solution(target, nums)
	fmt.Println("ret:", ret)
}

func solution(target int, nums []int) int {
	// nums[left ... right]为滑动窗口
	left := 0
	right := -1
	sum := 0
	result := len(nums) + 1

	for left < len(nums) {
		// 当前窗口右边界在范围内并且和小于给定值，将右侧边界右移，当前和加新值
		if right+1 < len(nums) && sum < target {
			right++
			sum += nums[right]
			// 当前和大于给定值，将和减最左边值，将左边界右移
		} else {
			sum -= nums[left]
			left++
		}

		if sum >= target {
			result = compare.Min(result, right-left+1)
		}
	}

	// 没有符合条件的子数组
	if result == len(nums)+1 {
		return 0
	}

	return result
}
