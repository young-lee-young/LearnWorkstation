package main

import "fmt"

/**
缺失的数字 LeetCode No268

使用哈希表
*/
func main() {
	nums := []int{3, 0, 1}
	ret := solution2(nums)
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

/**
[0, n] 的和 - 实际的和
 */
func solution2(nums []int) int {
	sum := 0

	real := 0

	for i := 0; i < len(nums); i ++ {
		real += nums[i]
		sum += i
	}

	sum += len(nums)

	return sum - real
}
