/**
 * @Time:    2022/11/28 22:30 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	LeetCode No027 移除元素

	标签：快慢指针
 */
package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}

	val := 3

	ret := solution(nums, val)

	fmt.Println("ret:", ret)
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
