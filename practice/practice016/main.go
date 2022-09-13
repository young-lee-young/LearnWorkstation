package main

import "fmt"

/**
	给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合 LeetCode No77

	标签：回溯法
 */
func main() {
	n := 5
	k := 4
	temp := make([]int, 0)
	ret := make([][]int, 0)
	finalRet := solution(1, n, k, temp, ret)
	fmt.Println(finalRet)
}

func solution(c int, n int, k int, temp []int, ret[][]int) [][]int {
	if len(temp) == k {
		ret = append(ret, temp)
		return ret
	}

	finalRet := make([][]int, 0)
	for i := c; i <= n; i ++ {
		newTemp := make([]int, 0)
		newTemp = append(newTemp, temp...)
		newTemp = append(newTemp, i)
		newRet := solution(i + 1, n, k, newTemp, ret)
		finalRet = append(finalRet, newRet...)
	}
	return finalRet
}
