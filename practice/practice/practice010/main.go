package main

import "fmt"

/**
	不重复数字全排列 LeetCode No046

	标签：回溯法标准方式
 */
func main() {
	nums := []int{1, 2, 3}
	ret := make([][]int, 0)
	tmpRet := make([]int, 0)
	ret = solution(tmpRet, nums, ret)
	fmt.Println(ret)
}

func solution(tmpRet []int, nums []int, ret [][]int) [][]int {
	// 回溯结束条件
	if len(nums) == 0 {
		ret = append(ret, tmpRet)
		return ret
	}

	for index, num := range nums {
		// append完当前数字的结果
		newRet := make([]int, 0)
		newRet = append(newRet, tmpRet...)
		newRet = append(newRet, num)

		// 把当前数字删除
		newNums := make([]int, 0)
		newNums = append(newNums, nums...)
		newNums = append(newNums[0:index], newNums[index+1:]...)
		ret = solution(newRet, newNums, ret)
	}

	return ret
}
