/**
 * @Time:    2022/4/21 13:57 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	柱状图中的最大矩形 LeetCode No84

 */
package main

import (
	"fmt"
	stack2 "practice/base/stack"
)

func main() {
	nums := []int{2, 1, 5, 6, 2, 3}

	ret := solution(nums)
	fmt.Println("ret", ret)
}

func solution(nums []int) int {
	nums = append(nums, 0)
	nums = append([]int{0}, nums...)

	max := 0

	stack := stack2.Stack{}
	for i := 0; i < len(nums); i ++ {
		for !stack.IsEmpty() && nums[i] < nums[stack.Peek().(int)] {
			height := nums[stack.Pop().(int)]
			area := (i - stack.Peek().(int) - 1) * height
			if area > max {
				max = area
			}
		}
		stack.Push(i)
	}

	return max
}
