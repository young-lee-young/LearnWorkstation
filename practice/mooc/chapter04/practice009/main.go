/**
 * @Time:    2021/7/3 18:22 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content: 
 */
package main

import "fmt"

/**
	回旋镖的数量 LeetCode 447

	标签：查找表map
 */
func main() {
	nums := [][]int{{0, 0}, {1, 0}, {2, 0}}
	result := solution(nums)
	fmt.Println(result)
}

func solution(nums [][]int) int {
	result := 0

	// 循环每个点
	for i := 0; i < len(nums); i ++ {
		// 存储距离当前点距离相同的点个数
		disMap := make(map[int]int)
		for j := 0; j < len(nums); j ++ {
			// 计算当前点和其他点横纵坐标差
			dx := nums[i][0] - nums[j][0]
			dy := nums[i][1] - nums[j][1]
			dis := dx * dx + dy * dy
			disMap[dis] = disMap[dis] + 1
		}

		for _, value := range disMap {
			result += value * (value - 1)
		}
	}
	return result
}
