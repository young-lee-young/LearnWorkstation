/**
 * @Time:    2021/7/3 15:16
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:
 */
package main

import (
	"fmt"
	"sort"
)

/**
四数之和 LeetCode No18

标签：对撞指针
*/
func main() {
	nums := []int{-5, -4, -3, -2, -1, 0, 0, 1, 2, 3, 4, 5}
	target := 0
	result := solution(nums, target)
	fmt.Println(result)
}

func solution(nums []int, target int) [][]int {
	result := [][]int{}
	numsLen := len(nums)

	if numsLen < 4 {
		return result
	}

	sort.Ints(nums)

	for i := 0; i < numsLen-3; i++ {
		// 跳过重复值
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < numsLen-2; j++ {
			// 跳过重复值
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			L := j + 1
			R := numsLen - 1
			for L < R {
				sum := nums[i] + nums[j] + nums[L] + nums[R]
				// 等于目标值
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[L], nums[R]})
					//去除重复值
					for L < R && nums[L] == nums[L+1] {
						L = L + 1
					}
					for L < R && nums[R] == nums[R-1] {
						R = R - 1
					}
					L = L + 1
					R = R - 1
				} else if sum > target {
					R = R - 1
				} else {
					L = L + 1
				}
			}
		}
	}
	return result
}
