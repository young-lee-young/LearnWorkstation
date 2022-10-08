/**
 * @Time:    2021/7/3 16:38
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:
 */
package main

import (
	"fmt"
	"math"
	"sort"
)

/**
最接近的三数之和 LeetCode No16

标签：对撞指针
*/
func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	result := solution(nums, target)
	fmt.Println(result)
}

func solution(nums []int, target int) int {
	numsLen := len(nums)
	result := nums[0] + nums[1] + nums[2]

	sort.Ints(nums)

	for i := 0; i < numsLen; i++ {
		L := i + 1
		R := numsLen - 1

		for L < R {
			sum := nums[i] + nums[L] + nums[R]
			if math.Abs(float64(target-sum)) < math.Abs(float64(target-result)) {
				result = sum
			}

			if sum > target {
				R--
			} else if sum < target {
				L++
			} else {
				return result
			}
		}
	}
	return result
}
