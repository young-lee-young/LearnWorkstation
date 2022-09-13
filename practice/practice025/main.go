package main

import "fmt"

/**
	买卖股票的最佳时机 LeetCode No122

	标签：贪心算法
 */
func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	result := solution(prices)
	fmt.Println(result)
}

func solution(nums []int) int {
	ret := 0
	// 昨天股票价格
	last := nums[0]
	for i := 0; i < len(nums); i ++ {
		// 如果今天高于昨天，计算收益
		if nums[i] > last {
			ret += nums[i] - last
			last = nums[i]
		} else {
			last = nums[i]
		}
	}
	return ret
}
