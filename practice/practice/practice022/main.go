package main

import "fmt"

/**
	重复数字的全排列 LeetCode No047

	标签：回溯法 + 剪枝
 */
func main() {
	nums := []int{1, 1, 2}
	// 最终结果
	ret := make([][]int, 0)
	// 每个子结果
	tempRet := make([]int, 0)
	ret = solution(nums, tempRet, ret)
	fmt.Println(ret)
}

func solution(nums []int, tempRet []int, ret [][]int) [][]int {
	if len(nums) == 0 {
		ret = append(ret, tempRet)
		return ret
	}

	// 用于判断是否是相同的元素，剪枝的过程
	hashMap := make(map[int]int)
	for index, num := range nums {
		if _, ok := hashMap[num]; ok {
			continue
		}
		hashMap[num] = num
		// append当前数字
		newRet := make([]int, 0)
		newRet = append(newRet, tempRet...)
		newRet = append(newRet, num)

		newNums := make([]int, 0)
		newNums = append(newNums, nums...)
		newNums = append(newNums[0:index], newNums[index + 1:]...)
		ret = solution(newNums, newRet, ret)
	}
	return ret
}
