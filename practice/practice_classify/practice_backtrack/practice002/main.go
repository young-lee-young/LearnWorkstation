/**
LeetCode No216

标签：回溯
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
		if Sum(path) == n {
			temp := make([]int, k)
			copy(temp, path)
			ret = append(ret, temp)
		}
		return
	}

	// 剪枝操作
	if Sum(path) >= n {
		return
	}

	// 单层搜索逻辑
	for i := start; i < len(nums); i++ {
		// 处理节点
		path = append(path, nums[i])

		backtrack(nums, n, k, i+1, path)

		path = path[:len(path)-1]
	}
}

func Sum(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
