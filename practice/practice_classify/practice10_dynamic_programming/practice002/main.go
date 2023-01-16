/**
LeetCode No070 爬楼梯
*/
package main

import "fmt"

func main() {
	n := 3

	ret := solution(n)
	fmt.Println("ret:", ret)
}

/**
1. 'DP 数组' 和 '下标' 的含义：dp[i] 表示到达第 i 阶楼梯有 dp[i] 种方法
2. 递推公式：dp[i] = dp[i - 1] + dp[i - 2]
3. 'DP 数组初始化'：dp[1] = 1，dp[2] = 2
4. 遍历顺序：从前向后遍历
*/
func solution(n int) int {
	if n < 2 {
		return n
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2

	for i := 3; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
