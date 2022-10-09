package main

import (
	"fmt"
	"math"
)

/**
存在重复元素Ⅲ LeetCode No220

标签：
*/
func main() {
	nums := []int{1, 2, 3, 1}
	k := 3
	t := 0
	ret := solution(nums, k, t)
	fmt.Println("ret:", ret)
}

func solution(nums []int, indexDiff int, valueDiff int) bool {
	numMap := make(map[int]int)

	for i := 0; i < indexDiff; i++ {
		if i >= len(nums) {
			return false
		}

		if checkExist(numMap, nums[i], valueDiff) {
			return true
		}

		numMap[nums[i]] = nums[i]
	}

	for i := indexDiff; i < len(nums); i++ {
		if checkExist(numMap, nums[i], valueDiff) {
			return true
		}

		// 增加窗口右界限值
		numMap[nums[i]] = nums[i]
		// 删除窗口左界限值
		delete(numMap, nums[i-indexDiff])
	}

	return false
}

func checkExist(numMap map[int]int, num int, valueDiff int) bool {
	for _, v := range numMap {
		if int(math.Abs(float64(v-num))) <= valueDiff {
			return true
		}
	}

	return false
}
