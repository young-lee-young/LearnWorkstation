package main

import (
	"fmt"
)

/**
	搜索二维矩阵 LeetCode No74

	标签：二分搜索

	解题思路：第一列二分搜索，再在该行二分搜索
 */
func main() {
	//matrix := [][]int{
	//	{1, 3, 5, 7},
	//	{10, 11, 16, 20},
	//	{23, 30, 34, 50},
	//}
	matrix := [][]int{
		{1, 3},
	}
	target := 1
	result := solution(matrix, target)
	fmt.Println(result)
}

func solution(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	rowLen := len(matrix) - 1
	columnLen := len(matrix[0]) - 1

	start := 0
	end := rowLen
	var mid int
	// 对第一列进行二分搜索
	for start < end {
		mid = start + (end - start + 1) / 2
		if matrix[mid][0] == target {
			return true
		}
		if matrix[mid][0] > target {
			end = mid - 1
		} else {
			start = mid
		}
	}
	if start == end {
		mid = start
	}

	// 对该行进行二分搜索
	newStart := 0
	newEnd := columnLen
	var newMid int
	for newStart < newEnd {
		newMid = newStart + (newEnd - newStart + 1) / 2
		if matrix[mid][newMid] == target {
			return true
		}
		if matrix[mid][newMid] > target {
			newEnd = newMid - 1
		} else {
			newStart = newMid
		}
	}
	if newStart == newEnd {
		newMid = newStart
	}

	if start == end && newStart == newEnd && matrix[mid][newMid] == target {
		return true
	}
	return false
}
