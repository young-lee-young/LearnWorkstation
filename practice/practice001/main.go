package main

import (
	"fmt"
	"math"
)

/**
	两个柱子间最大面积

	盛水最多的容器 LeetCode No11

	标签：双指针
 */
func main() {
	numArray := [...]int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	max := 0.0
	i := 0
	j := len(numArray) - 1
	for i < j {
		var minHeight int
		if numArray[i] < numArray[j] {
			minHeight = numArray[i]
			i ++
		} else {
			minHeight = numArray[j]
			j --
		}
		area := (j - i + 1) * minHeight
		max = math.Max(float64(max), float64(area))
		fmt.Println(i, j, minHeight, max)
	}
	fmt.Println(max)
}
