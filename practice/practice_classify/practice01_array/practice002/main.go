/**
 * @Time:    2022/11/28 22:30 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	LeetCode No027 移除元素

	标签：双指针
 */
package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}

	val := 3

	ret := solution(nums, val)

	fmt.Println(nums)

	fmt.Println("ret:", ret)
}

func solution(nums []int, val int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	i := 0
	for j := 0; j < len(nums); j ++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i ++
		}
	}

	return i
}
