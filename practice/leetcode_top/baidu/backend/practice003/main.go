package main

import "fmt"

/**
数组中超过一半的数字 剑指Offer No39

使用哈希表
*/
func main() {
	nums := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	ret := solution(nums)
	fmt.Println("ret", ret)
}

func solution(nums []int) int {
	numMap := make(map[int]int, 0)

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if _, ok := numMap[num]; !ok {
			numMap[num] = 1
		} else {
			numMap[num] += 1
		}
	}

	numLenHalf := len(nums) / 2

	fmt.Println("num map", numMap)

	for k, v := range numMap {
		if v > numLenHalf {
			return k
		}
	}

	return 0
}
