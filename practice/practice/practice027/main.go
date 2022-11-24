package main

import "fmt"

/**
	跳跃游戏 LeetCode No55

	解题思路：
 */
func main() {
	nums := []int{2, 3, 1, 1, 4}
	result := solution(nums)
	fmt.Println(result)
}

func solution(nums []int) bool {
	// 可以达到最远的位置
	max := 0
	for index, num := range nums {
		// 如果当前位置大于可达最大位置
		if index > max {
			return false
		}
		// 更新最大可达位置
		if index + num > max {
			max = index + num
		}
	}
	return true
}
