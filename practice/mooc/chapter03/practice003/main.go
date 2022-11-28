package main

import "fmt"

/**
有序数组中删除重复元素 LeetCode No026

标签：快慢指针
*/
func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	result := solution(nums)
	fmt.Println(result)
}

func solution(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
			fmt.Println(nums)
		}
	}
	return slow + 1
}
