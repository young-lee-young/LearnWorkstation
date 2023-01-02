/**
 * @Time:    2022/11/26 22:05 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	LeetCode No040 组合总和Ⅱ

	标签：回溯

	特点：给定数组有重复元素，结果集中需要去重

	去重：树层去重 & 树枝去重
 */
package main

import (
	"fmt"
)

var ret [][]int

func main() {
	ret = [][]int{}

	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	path := []int{}
	used := make(map[int]bool)

	// ⚠️注意：数组去重时需要数组有序
	quickSort(candidates, 0, len(candidates) - 1)

	backtrack(candidates, target, 0, path, used)

	fmt.Println("ret:", ret)
}

func backtrack(candidates []int, target int, start int, path []int, used map[int]bool) {
	// 递归终止条件
	if sum(path) >= target {
		if sum(path) == target {
			temp := make([]int, len(path))
			copy(temp, path)
			ret = append(ret, temp)
		}
		return
	}

	// 单层逻辑处理
	for i := start; i < len(candidates); i ++ {
		// ⚠️去重开始
		// 数层去重，需要在 for 循环里去重
		// 如果当前层和上一层值相同，上一层没有使用，当前层也不使用（如果使用，当前层的效果和上一层一样，就重复了），修改 repeat 值
		if i > 0 && candidates[i] == candidates[i - 1] && !used[i - 1] {
			continue
		}
		// ⚠️去重结束

		// 处理节点
		path = append(path, candidates[i])
		used[i] = true

		// 递归函数
		backtrack(candidates, target, i + 1, path, used)

		// 回溯
		path = path[:len(path)-1]
		used[i]= false
	}
}

func sum(path []int) int {
	s := 0
	for _, v := range path {
		s += v
	}
	return s
}

func quickSort(array []int, L int, R int) {
	if L >= R {
		return
	}

	left := L
	right := R

	// ⚠️ 要取 array[left]，不要取 array[0]
	pivot := array[left]

	for left < right {
		// 第一步，把比 pivot 大的值放到 pivot 右边
		// 如果右指针的值比选定值大，右指针左移
		for left < right && array[right] >= pivot {
			right --
		}
		// 上面的循环退出来，说明右侧的值比 pivot 要小，将小的值换到左指针位置
		if left < right {
			array[left] = array[right]
		}

		// 第二步，把比 pivot 小的值放到 pivot 左边
		// 如果左指针的值比选定值小，左指针右移
		for left < right && array[left] <= pivot {
			left ++
		}
		// 上面的循环退出来，说明左侧的值比 pivot 要大，将大的值换到右指针位置
		if left < right {
			array[right] = array[left]
		}

		// 将 pivot 放到合适的位置
		if left >= right {
			array[left] = pivot
		}
	}

	quickSort(array, L, right-1)
	quickSort(array, right+1, R)
}