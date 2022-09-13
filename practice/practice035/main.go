package main

import "fmt"

/**
	有障碍物的不同路径 LeetCode No63

	标签：动态规划
 */
func main() {
	numArray := [][]int{
		{0, 0},
		{1, 1},
		{0, 0},
	}
	result := solution(numArray)
	fmt.Println(result)
}

func solution(numArray [][]int) int {
	if len(numArray) == 0 || len(numArray[0]) == 0 || numArray[0][0] == 1 {
		return 0
	}
	m := len(numArray)
	n := len(numArray[0])

	// 实例化二维数组
	newNumArray := make([][]int, m)
	for i := 0; i < m; i ++ {
		newNumArray[i] = make([]int, n)
	}
	// 将最左列和最上值设置为1，到达这些点只有一种走法
	for i := 0; i < m && numArray[i][0] == 0; i ++ {
		// 不是障碍物设置为1，如果是障碍物则是0
		newNumArray[i][0] = 1
	}
	for i := 0; i < n && numArray[0][i] == 0; i ++ {
		newNumArray[0][i] = 1
	}

	for i := 1; i < m; i ++ {
		for j := 1; j < n; j ++ {
			// 如果不是障碍物进行走法计算，是障碍物直接为0
			if numArray[i][j] == 0 {
				newNumArray[i][j] = newNumArray[i - 1][j] + newNumArray[i][j - 1]
			}
		}
	}
	return newNumArray[m - 1][n - 1]
}
