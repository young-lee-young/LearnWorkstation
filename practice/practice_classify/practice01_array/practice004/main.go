/**
 * @Time:    2022/11/29 10:33 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	LeetCode No209 长度最小的子数组

	标签：滑动窗口
 */
package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}

	target := 7

	ret := solution(nums, target)

	fmt.Println("ret:", ret)
}

func solution(nums []int, target int) int {
	ret := math.MaxInt64

	sum := 0

	i := 0

	// 窗口右侧滑动
	for j := 0; j < len(nums); j ++ {
		sum = sum + nums[j]
		for sum >= target { // ⚠️注意：这里是 for 循环，而不是 if 判断
			ret = min(ret, j-i+1)
			// 窗口左侧滑动
			// ⚠️注意：注意这里的顺序，先减再移动 i
			sum = sum - nums[i]
			i ++
		}
	}

	if ret == math.MaxInt64 {
		return 0
	}

	return ret
}

func min(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}
