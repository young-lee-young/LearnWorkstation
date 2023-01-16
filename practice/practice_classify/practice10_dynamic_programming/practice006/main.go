/**
LeetCode No343 整数拆分
*/
package main

import (
	"fmt"
)

func main() {
	n := 10

	ret := solution(n)
	fmt.Println("ret:", ret)
}

/**
1. 'DP 数组' 和 '下标' 的含义：dp[i] 表示对 i 进行拆分，得到的最大乘积为 dp[i]

2. 递推公式：
	将 i 拆分成 2 个数相乘：j * (i - j)
	将 i 拆分成 3 个数及以上相乘：j * dp[i - j]

3.'DP 数组' 初始化：
	dp[2] = 1 = 1 * 1

4. 遍历数序：顺序遍历
*/
func solution(n int) int {
	dp := make([]int, n+1)

	dp[2] = 1

	for i := 3; i < n+1; i++ {
		for j := 0; j < i; j++ {
			// max(j * (i - j), j * dp[i - j], dp[i])
			dp[i] = max(max(j*(i-j), j*dp[i-j]), dp[i])
		}
	}

	return dp[n]
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
