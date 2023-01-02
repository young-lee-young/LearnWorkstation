/**
 * @Time:    2022/11/28 11:24 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	LeetCode No491 递增子序列

	标签：回溯

	特点：不能对原始数组进行排序
 */
package main

import "fmt"

var ret [][]int

func main() {
	ret = [][]int{}

	nums := []int{4, 7, 6, 7}

	path := []int{}

	backtrack(nums, 0, path)

	fmt.Println("ret:", ret)
}

func backtrack(nums []int, start int, path []int) {
	// 收集结果
	if len(path) >= 2 {
		temp := make([]int, len(path))
		copy(temp, path)
		ret = append(ret, temp)
	}

	// 递归结束条件
	if start == len(nums) {
		return
	}

	// ⚠️注意：
	// 去重条件：标记当前层数字是否使用
	used := make(map[int]bool)

	// 单层搜索逻辑
	for i := start; i < len(nums); i ++ {
		// 当前值比结果集最后一个值小
		if len(path) >= 1 && nums[i] < path[len(path)-1] {
			continue
		}

		// ⚠️注意：去重开始
		if used[nums[i]] {
			continue
		}
		used[nums[i]] = true
		// ⚠️注意：去重结束

		// 处理节点
		path = append(path, nums[i])

		// 递归
		backtrack(nums, i+1, path)

		// 回溯
		path = path[:len(path)-1]
	}
}
