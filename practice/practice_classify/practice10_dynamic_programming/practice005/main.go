/**
LeetCode No063 不用路径Ⅱ
*/
package main

import "fmt"

func main() {
	obstacleGrid := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}

	ret := solution(obstacleGrid)
	fmt.Println("ret:", ret)
}

/**
1. 'DP 数组' 和 '下标' 的含义：dp[i][j] 表示到达坐标 (i, j) 的不同路径条数

2. 递推公式：
	if obstacleGrid[i][j] != 1 {
		dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
	}

3. 'DP 数组' 初始化：
	dp[i][0] = 1 // 第一行初始化为 1，如果有障碍，后面都初始化为 0
	dp[0][j] = 1 // 第一列初始化为 1，如果有障碍，后面都初始化为 0

4. 遍历顺序：从前往后遍历
*/
func solution(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	// 'DP 数组' 初始化
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// 第一行初始化为 1
	obstacle := false
	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			obstacle = true
		}
		if !obstacle {
			dp[0][j] = 1
		}
	}

	// 第一列初始化为 1
	obstacle = false
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			obstacle = true
		}
		if !obstacle {
			dp[i][0] = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] != 1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}
