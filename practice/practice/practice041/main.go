package main

import (
	"fmt"
	"practice/utils/compare"
)

/**
	打家劫舍 LeetCode No198

	标签：动态规划
 */
func main() {
	numArray := []int{1, 2, 3, 1}
	result := solution(numArray)
	fmt.Println(result)
}

func solution(numArray []int) int {
	// 定义动态规划数组，加两个长度，放到数组开始
	dpArray := make([]int, len(numArray) + 2)
	for i := 0; i < len(numArray); i ++ {
		j := i + 2
		// 动态规划转移方程，动态规划数组从第2个开始 nums[i] = max(nums[i - 1], nums[i - 2] + nums[i])
		dpArray[j] = compare.Max(dpArray[j - 1], dpArray[j - 2] + numArray[i])
	}
	return dpArray[len(dpArray) - 1]
}
