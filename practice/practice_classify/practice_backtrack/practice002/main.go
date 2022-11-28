/**
组合总和Ⅲ LeetCode No216

标签：回溯

特点：不含重复元素，元素只能使用一次
*/
package main

import "fmt"

var ret [][]int

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := 9 // 和
	k := 3 // 个数
	path := []int{}

	ret = [][]int{}

	backtrack(nums, n, k, 0, path)

	fmt.Println("ret:", ret)
}

func backtrack(nums []int, n int, k int, start int, path []int) {
	// 递归函数结束条件
	if len(path) == k {
		if sum(path) == n {
			temp := make([]int, k)
			copy(temp, path)
			ret = append(ret, temp)
		}
		return
	}

	// 单层搜索逻辑
	for i := start; i < len(nums); i++ {
		// 处理节点
		path = append(path, nums[i])

		// TODO 剪枝操作

		// 递归函数
		// ⚠️注意：不含重复元素，元素只能使用一次，体现在 i + 1
		backtrack(nums, n, k, i+1, path)

		// 回溯
		path = path[:len(path)-1]
	}
}

func sum(nums []int) int {
	s := 0
	for _, v := range nums {
		s += v
	}
	return s
}
