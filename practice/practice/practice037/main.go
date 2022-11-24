package main

import "fmt"

/**
	三角形最小路径和 LeetCode No120

	标签：动态规划
 */
func main() {
	numArray := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	result := solution(numArray)
	fmt.Println(result)
}

func solution(numArray [][]int) int {
	arrayLen := len(numArray)

	// 定义动态规划二维数组，比给定数组三角形行数多1
	dpArray := make([][]int, arrayLen + 1)
	for i := 0; i < arrayLen + 1; i ++ {
		dpArray[i] = make([]int, arrayLen + 1)
	}

	// 从三角形数组最后一行开始递推
	for i := arrayLen - 1; i >= 0; i -- {
		for j := 0; j <= i; j ++ {
			dpArray[i][j] = min(dpArray[i + 1][j], dpArray[i + 1][j + 1]) + numArray[i][j]
			fmt.Println(i, j, dpArray)
		}
	}
	return dpArray[0][0]
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
