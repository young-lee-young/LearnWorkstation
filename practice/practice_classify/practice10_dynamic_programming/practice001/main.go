/**
LeetCode No509 斐波那契数
*/
package main

import "fmt"

func main() {
	n := 3

	ret := solution(n)
	fmt.Println("ret:", ret)
}

/**
1. 'DP 数组' 和 '下标' 的含义：dp[i] 表示第 i 个斐波那契数值为 dp[i]
2. 递推公式：dp[i] = dp[i - 1] + dp[i - 2]
3. 'DP 数组初始化'：dp[0] = 1，dp[1] = 1
4. 遍历顺序：从前向后遍历
5. 打印 'DP 数组'
*/
func solution(n int) int {
	if n == 0 {
		return 0
	}

	// 'DP 数组初始化'
	// 从 F(0) 开始，到 F(n)，共 n + 1 个数
	// 如 n = 3，F(0)、F(1)、F(2)、F(3)，共 4 个数
	dp := make([]int, n+1)

	dp[0] = 0
	dp[1] = 1

	// 遍历顺序
	for i := 2; i < n+1; i++ {
		// 递归公式
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
