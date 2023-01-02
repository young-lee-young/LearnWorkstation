/**
组合 LeetCode No077

标签：回溯

特点：给定数组无重复元素，元素不可以重复使用
*/
package main

import "fmt"

var ret [][]int

func main() {
	n := 4
	k := 2
	path := []int{}

	ret = [][]int{}

	backtrack(n, k, 1, path)
	fmt.Println("ret:", ret)
}

func backtrack(n int, k int, start int, path []int) {
	// 终止条件，结果集的长度为 k
	if len(path) == k {
		// 注意⚠️：需要 copy 临时变量，将临时变量放进切片里
		temp := make([]int, k)
		copy(temp, path)
		ret = append(ret, temp)
		return
	}

	// 剪枝
	if n-(k-len(path))+1 < start {
		return
	}

	// 单层搜索逻辑
	for i := start; i <= n; i++ {
		// 处理节点
		path = append(path, i)

		// 递归函数
		backtrack(n, k, i+1, path)

		// 回溯
		path = path[:len(path)-1]
	}
}
