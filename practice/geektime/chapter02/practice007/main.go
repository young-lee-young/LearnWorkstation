/**
 * @Time:    2022/4/21 19:17 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	接雨水 LeetCode No42

 */
package main

import (
	"fmt"
	stack2 "practice/base/stack"
	"math"
)

func main() {
	nums := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	ret := solution(nums)
	fmt.Println("ret", ret)
}

func solution(nums []int) int {

	sum := 0
	stack := stack2.Stack{}

	for i := 0; i < len(nums); i ++ {

		if stack.IsEmpty() || nums[i] <= nums[stack.Peek().(int)] {
			stack.Push(i)
			continue
		}

		for nums[i] > nums[stack.Peek().(int)] {

			height := nums[stack.Pop().(int)]

			if stack.IsEmpty() {
				break
			}

			distance := i - stack.Peek().(int) - 1

			min := int(math.Min(float64(nums[stack.Peek().(int)]), float64(nums[i])))

			sum = sum + distance*(min-height)
		}

		stack.Push(i)
	}

	return sum
}
