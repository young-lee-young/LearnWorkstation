package main

import (
	"fmt"
	"sort"
)

/**
	零钱兑换 LeetCode No322

	标签：动态规划
 */
func main() {
	numArray := []int{1, 2, 5}
	sort.Slice(numArray, func(i, j int) bool {
		return numArray[i] < numArray[j]
	})
	sum := 11
	result := solution(numArray, sum)
	fmt.Println(result)
}

func solution(numArray []int, sum int) int {
	// 构造二维动态规划数组，横坐标是零钱数，数组的含义是当前纵坐标的钱数使用所给的硬币最少组成
	dpArray := make([][]int, len(numArray) + 1)
	// 纵坐标是目标钱数
	for i := 0; i < len(numArray) + 1; i ++ {
		dpArray[i] = make([]int, sum + 1)
	}
	// 初始化第一行，将第一行除第一个位置全部置为 sum + 1
	for j := 1; j < sum + 1; j ++ {
		dpArray[0][j] = sum + 1
	}

	printArray(dpArray)

	// 循环DP数组的行
	for i := 1; i < len(numArray) + 1; i ++ {
		// 循环DP数组的列
		for j := 0; j < sum + 1; j ++ {
			// 状态转移方程，如果需要兑换的钱比给定的钱大
			if j >= numArray[i - 1] {
				// dpArray[i - 1][j]表示还是使用上一个硬币
				// dpArray[i][j - numArray[i - 1]] + 1)表示使用这次硬币
				dpArray[i][j] = min(dpArray[i - 1][j], dpArray[i][j - numArray[i - 1]] + 1)
			} else {
				dpArray[i][j] = dpArray[i - 1][j]
			}
			fmt.Println(i, j, numArray[i - 1])
			printArray(dpArray)
		}
	}

	result := dpArray[len(numArray)][sum]
	if result > sum {
		return -1
	} else {
		return result
	}
}

func min(x int, y int) int {
	if x <= y {
		return x
	}
	return y
}

/**
	打印二维数组
 */
func printArray(numArray [][]int) {
	for i := 0; i < len(numArray); i ++ {
		fmt.Println(numArray[i])
	}
	fmt.Println("------------------------")
}
