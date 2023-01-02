/**
 * @Time:    2022/11/27 18:59 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	LeetCode No078 子集

	标签：回溯
 */
package main

import "fmt"

var ret [][]int

func main() {
	ret = [][]int{}

	nums := []int{1, 2, 3}

	path := []int{}

	backtrack(nums, 0, path)

	fmt.Println("ret:", ret)
}

func backtrack(nums []int, start int, path []int) {
	// 收集结果
	temp := make([]int, len(path))
	copy(temp, path)
	ret = append(ret, temp)

	// 递归终止条件
	if start >= len(nums) {
		return
	}

	// 单层逻辑搜索
	for i := start; i < len(nums); i ++ {
		// 处理节点
		path = append(path, nums[i])

		// 递归函数
		backtrack(nums, i+1, path)

		// 回溯
		path = path[:len(path)-1]
	}
}
