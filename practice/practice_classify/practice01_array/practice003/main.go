/**
 * @Time:    2022/11/28 22:37
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	LeetCode No977 有序数组的平方

	标签：双指针
 */
package main

import "fmt"

func main() {
	nums := []int{-4, -1, 0, 3, 10}

	ret := solution(nums)

	fmt.Println("ret:", ret)
}

func solution(nums []int) []int {
	ret := make([]int, 0)

	left := 0
	right := len(nums) - 1

	for left <= right {
		if nums[left]*nums[left] >= nums[right]*nums[right] {
			ret = append([]int{nums[left] * nums[left]}, ret...)
			left ++
		} else {
			ret = append([]int{nums[right]*nums[right]}, ret...)
			right --
		}
	}

	return ret
}
