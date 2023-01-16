/**
LeetCode No096 不同的二叉搜索树
*/
package main

import "fmt"

func main() {
	n := 1

	ret := solution(n)

	fmt.Println("ret:", ret)
}

/**
1. 'DP 数组' 和 '下标' 的含义：dp[i] 表示 i 个节点有 dp[i] 个二叉搜索树

2. 递推公式：
	如 n = 3，则有二叉搜索树个数为：以 1 为头节点的二叉树个数 + 以 2 为头节点的二叉树个数 + 以 3 为头节点的二叉树个数

	以 1 为头节点二叉树个数：左子树 0 节点 * 右子树 2 节点

	以 2 为头节点二叉树个数：左子树 1 节点 * 右子树 1 节点

	以 3 为头节点二叉树个数：左子树 2 节点 * 右子树 0 节点

	dp[3] = dp[0] * dp[2] + dp[1] * dp[1] + dp[2] * dp[0]

	dp[n] =
	for i := 2; i < n + 1; i ++ {
		sum = 0
		for j := 0; j < i; j ++ {
			sum = sum + dp[i] * dp[i - 1 - j]
		}
	}

3. 'DP 数组' 初始化：

	dp[0] = 1 // 空二叉树
	dp[1] = 1
	dp[2] = 2

4. 遍历数组：从前往后遍历
*/
func solution(n int) int {
	dp := make([]int, n+1)

	dp[0] = 1
	dp[1] = 1

	for i := 2; i < n+1; i++ {
		sum := 0
		for j := 0; j < i; j++ {
			sum += dp[j] * dp[i-1-j]
		}
		dp[i] = sum
	}

	return dp[n]
}
