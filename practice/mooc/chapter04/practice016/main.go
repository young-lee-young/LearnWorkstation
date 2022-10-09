package main

import (
	"fmt"
)

/**
存在重复元素Ⅱ LeetCode No219

标签：滑动窗口 + 哈希Map
*/
func main() {
	nums := []int{1, 2, 3, 1, 2, 3}
	k := 2
	ret := solution(nums, k)
	fmt.Println("ret:", ret)
}

func solution(nums []int, k int) bool {
	numMap := make(map[int]int)

	for i := 0; i < k; i++ {
		if i >= len(nums) {
			return false
		}

		if _, ok := numMap[nums[i]]; ok {
			return true
		}
		numMap[nums[i]] = nums[i]
	}

	for i := k; i < len(nums); i++ {
		if _, ok := numMap[nums[i]]; ok {
			return true
		}

		// 增加窗口右界限值
		numMap[nums[i]] = nums[i]
		// 删除窗口左界限值
		delete(numMap, nums[i-k])
	}

	return false
}
