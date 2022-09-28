package main

import "fmt"

/**
x平方根 LeetCode No69

标签：二分搜索
*/
func main() {
	num := 8
	result := solution(num)
	fmt.Println(result)
}

func solution(num int) int {
	start := 0
	end := num

	for start < end {
		mid := start + (end-start+1)/2
		if mid*mid > num {
			end = mid - 1
		} else {
			start = mid
		}
	}
	return start
}
