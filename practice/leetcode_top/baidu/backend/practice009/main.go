package main

import "fmt"

/**
缺失的数字 LeetCode No268

使用哈希表
*/
func main() {
	nums := []int{3, 0, 1}
	ret := solution(nums)
	fmt.Println("ret:", ret)
}

func solution(nums []int) int {
	hash := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		hash[nums[i]] = true
	}

	for i := 0; i < len(nums)+1; i++ {
		if !hash[i] {
			return i
		}
	}

	return -1
}
