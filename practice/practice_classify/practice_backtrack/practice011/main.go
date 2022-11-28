/**
 * @Time:    2022/11/28 15:14 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	LeetCode No046 全排列

	标签：回溯
 */
package main

import "fmt"

var ret [][]int

func main() {
	nums := []int{1, 2, 3}

	ret = [][]int{}

	path := []int{}

	// ⚠️注意：标记数组中每个数使用情况
	used := make(map[int]bool)

	backtrack(nums, path, used)

	fmt.Println("ret:", ret)
}

func backtrack(nums []int, path []int, used map[int]bool) {
	// 递归结束条件
	if len(path) == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		ret = append(ret, temp)
		return
	}

	// 单层搜索逻辑
	for i := 0; i < len(nums); i ++ {
		if used[i] {
			continue
		}

		// 处理节点
		used[i] = true
		path = append(path, nums[i])

		// 递归函数
		backtrack(nums, path, used)

		// 回溯
		used[i] = false
		path = path[:len(path)-1]
	}
}
