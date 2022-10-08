package main

import "fmt"

/**
删除数组中重复元素，每个元素最多存在2个 LeetCode No80

标签：快慢指针
*/
func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	result := solution(nums)
	fmt.Println(result)
}

func solution(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if slow < 2 || nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
