package main

import "fmt"

/**
	最小次数跳到最后 LeetCode No45

	解题思路：
 */
func main() {
	nums := []int{2, 3, 1, 1, 4}
	result := solution(nums)
	fmt.Println(result)
}

func solution(nums []int) int {
	// 跳跃次数
	count := 0

	start := 0
	end := 1
	for end < len(nums) {
		max := 0
		for i := start; i < end; i ++ {
			if i + nums[i] > max {
				max = i + nums[i]
			}
		}
		start = end
		end = max + 1
		count ++
	}
	return count
}
