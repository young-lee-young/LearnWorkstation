/**
 * @Time:    2022/11/26 21:40 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	LeetCode No039 组合总和

	标签：回溯

	特点：给定数组无重复元素，并且元素可以重复使用
 */
package main

import "fmt"

var ret [][]int

func main() {
	ret = [][]int{}

	candidates := []int{2, 3, 6, 7}
	target := 7
	path := []int{}

	backtrack(candidates, target, path)

	fmt.Println("ret:", ret)
}

func backtrack(candidates []int, target int, path []int) {
	// 递归结束条件
	if sum(path) >= target {
		if sum(path) == target {
			temp := make([]int, len(path))
			copy(temp, path)
			ret = append(ret, temp)
		}
		return
	}

	// 单层逻辑处理
	for i := 0; i < len(candidates); i ++ {
		// 处理单个节点
		path = append(path, candidates[i])

		// 剪枝：和大于目标值，不用再继续递归
		if sum(path) > target {
			path = path[:len(path)-1]
			continue
		}

		// 递归函数
		// ⚠️注意：元素可以重复使用，体现在 i，而不是使用 i + 1
		backtrack(candidates[i:], target, path)

		// 回溯
		path = path[:len(path)-1]
	}
}

func sum(path []int) int {
	s := 0
	for _, v := range path {
		s += v
	}
	return s
}
