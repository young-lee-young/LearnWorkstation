package main

import (
	"fmt"
	"math"
)

/**
	乘积最大子数组 LeetCode No152

	标签：动态规划

	解题思路：
 */
func main() {
	//numArray := []int{2, 3, -2, 4}
	numArray := []int{-2, 0, -1}
	result := solution(numArray)
	fmt.Println(result)
}

func solution(numArray []int) int {
	// 当前结果存最小值
	result := math.MinInt64
	maxTemp := 1
	minTemp := 1
	for i := 0; i < len(numArray); i ++ {
		// 当前值大于0，将最大值、最小值分别乘当前值
		if numArray[i] > 0 {
			maxTemp = maxTemp * numArray[i]
			minTemp = minTemp * numArray[i]
		// 当前值等于0，将暂时最大值、暂时最小值复位，取最大值和0较大值
		} else if numArray[i] == 0 {
			maxTemp = 1
			minTemp = 1
			result = max(result, 0)
			continue
		// 存在负数，导致最大的变最小的，最小的变最大的，需要反过来乘
		} else {
			temp := maxTemp
			maxTemp = max(minTemp*numArray[i], numArray[i])
			minTemp = min(temp*numArray[i], numArray[i])
		}
		// 取结果值和暂时最大值较大值
		result = max(result, maxTemp)
	}
	return result
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
