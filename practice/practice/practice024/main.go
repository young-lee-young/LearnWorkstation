package main

import "fmt"

/**
	柠檬水找零 LeetCode No860

	标签：贪心算法
 */
func main() {
	bills := []int{10, 10}
	result := solution(bills)
	fmt.Println(result)
}

func solution(nums []int) bool {
	fiveCount := 0
	tenCount := 0
	for _, num := range nums {
		if num == 5 {
			fiveCount ++
		} else if num == 10 {
			if fiveCount > 0 {
				fiveCount --
				tenCount ++
			} else {
				return false
			}
		} else {
			if tenCount > 0 {
				tenCount --
				if fiveCount > 0 {
					fiveCount --
				} else {
					return false
				}
			} else {
				if fiveCount >= 3 {
					fiveCount = fiveCount - 3
				} else {
					return false
				}
			}

		}
	}
	return true
}
