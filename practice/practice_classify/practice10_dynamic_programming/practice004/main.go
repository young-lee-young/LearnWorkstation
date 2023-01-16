/**
LeetCode No062 不同路径
*/
package main

import "fmt"

func main() {
	m := 3
	n := 7

	ret := solution(m, n)
	fmt.Println("ret:", ret)
}

/**
1. 'DP 数组' 和 '下标' 的含义：dp[i][j] 表示走到坐标 (i, j) 的不同路径条数

2. 递推公式：
	if i == 0 {
		dp[i][j] = dp[i][j - 1]
	}
	if j == 0 {
		dp[i][j] = dp[i - 1][j]
	}
	if i != 0 && j != 0 {
		dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
	}

3. 'DP 数组' 初始化：
	dp[0][j] = 1 // 第一行初始化为 1
	dp[i][0] = 1 // 第一列初始化为 1

4. 遍历顺序：从前往后遍历
*/
func solution(m int, n int) int {
	// 'DP 数组' 初始化
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// 第一行初始化
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	// 第一列初始化
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
