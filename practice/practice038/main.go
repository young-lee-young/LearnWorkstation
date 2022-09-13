package main

import "fmt"

/**
	最大子序和 LeetCode No53

	标签：动态规划
 */
func main() {
	numArray := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	result := solution(numArray)
	fmt.Println(result)
}

func solution(numArray []int) int {
	if len(numArray) == 0 {
		return 0
	}
	// 定义动态规划数组
	dpArray := make([]int, len(numArray))
	// 将数组第一个数放入动态规划数组
	dpArray[0] = numArray[0]
	for i := 1; i < len(numArray); i ++ {
		// 如果前面的和大于0，加上当前数
		if dpArray[i - 1] > 0 {
			dpArray[i] = numArray[i] + dpArray[i - 1]
		// 如果小于0，和为当前树
		} else {
			dpArray[i] = numArray[i]
		}
	}

	// 取动态规划数组中的最大值
	result := dpArray[0]
	for i := 1; i < len(dpArray); i ++ {
		if dpArray[i] > result {
			result = dpArray[i]
		}
	}
	return result
}
