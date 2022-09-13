package main

import "fmt"

/**
	求子集 LeetCode No78

	标签：回溯，类似全排列
 */
func main() {
	nums := []int{1, 2, 3}
	finalResult := make([][]int, 0)

	sliceLen := len(nums)
	for i := 0; i < sliceLen + 1; i ++ {
		ret := make([][]int, 0)
		tempRet := make([]int, 0)
		backResult := solution(tempRet, nums, ret, i)
		finalResult = append(finalResult, backResult...)
	}
	fmt.Println(finalResult)
}

func solution(tempRet []int, nums []int, ret [][]int, count int) [][]int {
	if len(tempRet) == count {
		ret = append(ret, tempRet)
		return ret
	}

	for index, num := range nums {
		newRet := make([]int, 0)
		newRet = append(newRet, tempRet...)
		newRet = append(newRet, num)

		newNums := nums[index + 1:]
		ret = solution(newRet, newNums, ret, count)
	}
	return ret
}
