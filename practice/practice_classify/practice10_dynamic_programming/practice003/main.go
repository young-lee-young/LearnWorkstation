/**
LeetCode No746 使用最小花费爬楼梯
			__
		 __|顶
      __|20
   __|15
__|10
*/
package main

import "fmt"

func main() {
	cost := []int{10, 15, 20}

	ret := solution(cost)
	fmt.Println("ret:", ret)
}

/**
1. 'DP 数组' 和 '下标' 的含义：dp[i] 表示到达 i 楼层花费为 dp[i]
2. 递归公式：dp[i] = min(dp[i - 1] + cost[i - 1], dp[i - 2] + cost[i - 2])
3. 'DP 数组' 初始化：dp[0] = 0，dp[1] = 0
4. 遍历顺序：从前往后遍历
*/
func solution(cost []int) int {
	// ⚠️：这里的 n 是 len(cost)，不用 + 1
	n := len(cost)
	if n < 2 {
		return 0
	}

	dp := make([]int, n+1)

	dp[0] = 0
	dp[1] = 0

	for i := 2; i < n+1; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[n]
}

func min(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}
