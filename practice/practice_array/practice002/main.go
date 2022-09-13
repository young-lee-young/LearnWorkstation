package main

import "fmt"

/**
	有序数组的两个数的和等于给定元素 LeetCode No167
 */
func main() {
	sum := 9
	//numList := [7]int{1, 2, 5, 7, 9, 11, 13}
	numList := []int{2, 7, 11, 15}

	i := 0
	j := len(numList) - 1

	for i < j {
		sumTemp := numList[i] + numList[j]
		if sumTemp == sum {
			fmt.Println(i, j)
			return
		}
		if sumTemp > sum {
			j --
		} else {
			i ++
		}
	}
}
