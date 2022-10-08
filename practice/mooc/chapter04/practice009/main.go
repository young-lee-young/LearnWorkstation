/**
 * @Time:    2021/7/3 12:08
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
三数之和 LeetCode No15

标签：对撞指针
*/
func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := solution(nums)
	fmt.Println(result)
}

func solution(nums []int) [][]int {
	result := [][]int{}

	if len(nums) < 3 {
		return result
	}

	// 将数组排序
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		// 如果当前值大于0，后面的值都大于0
		if nums[i] > 0 {
			return result
		}

		// 重复值跳过
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		L := i + 1
		R := len(nums) - 1
		for L < R {
			// 等于0
			if nums[i]+nums[L]+nums[R] == 0 {
				result = append(result, []int{nums[L], nums[i], nums[R]})
				// 去除重复值
				for L < R && nums[L] == nums[L+1] {
					L = L + 1
				}
				for L < R && nums[R] == nums[R-1] {
					R = R - 1
				}
				L = L + 1
				R = R - 1
				// 大于0，右指针左移
			} else if nums[i]+nums[L]+nums[R] > 0 {
				R = R - 1
				// 小于0，做指针右移
			} else {
				L = L + 1
			}
		}
	}
	return result
}
