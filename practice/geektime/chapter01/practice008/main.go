/**
 * @Time:    2022/4/20 22:49 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	加一 LeetCode No66

 */
package main

import "fmt"

func main() {
	nums := []int{9, 9, 9}

	ret := solution(nums)
	fmt.Println("ret", ret)
}

func solution(nums []int) []int {
	j := len(nums) - 1

	// 末尾无进位
	if nums[j] < 9 {
		nums[j] = nums[j] + 1
		return nums
	}

	// 末尾有进位
	for j >= 0 && nums[j] == 9 {
		nums[j] = 0
		j --
	}

	// 最高位没有发生进位
	if j >= 0 {
		nums[j] = nums[j] + 1
		return nums
	}

	// 最高位发生进位
	nums = append([]int{1}, nums...)

	return nums
}
