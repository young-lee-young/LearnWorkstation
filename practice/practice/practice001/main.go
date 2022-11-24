package main

import (
	"fmt"
	"math"
)

/**
两个柱子间最大面积

盛水最多的容器 LeetCode No011

标签：双指针
*/
func main() {
	numArray := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	ret := solution(numArray)
	fmt.Println("ret:", ret)
}

func solution(height []int) int {
	var max float64

	i := 0
	j := len(height) - 1

	for i < j {
		var minHeight int

		if height[i] < height[j] {
			minHeight = height[i]
			i++
		} else {
			minHeight = height[j]
			j--
		}

		area := (j - i + 1) * minHeight
		max = math.Max(max, float64(area))
	}

	return int(max)
}
