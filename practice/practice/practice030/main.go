package main

import "fmt"

/**
	有效的完全平方 LeetCode No367

	标签：二分搜索
 */
func main() {
	num := 9
	result := solution(num)
	fmt.Println(result)
}

func solution(num int) bool {
	start := 0
	end := num
	for start < end {
		mid := start + (end - start + 1) / 2
		if (mid * mid) == num {
			return true
		}
		if mid * mid > num {
			end = mid - 1
		} else {
			start = mid
		}
	}
	return false
}
