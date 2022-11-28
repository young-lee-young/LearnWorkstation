package main

import "fmt"

/**
数组中删除给定的值 LeetCode No027

标签：快慢指针
*/
func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2

	result := solution(nums, val)
	fmt.Println(result)
}

func solution(nums []int, val int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
