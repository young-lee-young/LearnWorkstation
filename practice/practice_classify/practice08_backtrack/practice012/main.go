/**
 * @Time:    2022/11/28 15:32 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	LeetCode No047 全排列Ⅱ

	标签：回溯
 */
package main

import "fmt"

var ret [][]int

func main() {
	nums := []int{1, 1, 2}

	quickSort(nums, 0, len(nums)-1)

	ret = [][]int{}

	path := []int{}

	// ⚠️注意：标记数组中每个数使用情况
	used := make(map[int]bool)

	backtrack(nums, path, used)

	fmt.Println("ret:", ret)
}

func backtrack(nums []int, path []int, used map[int]bool) {
	// 递归终止条件
	if len(nums) == len(path) {
		temp := make([]int, len(path))
		copy(temp, path)
		ret = append(ret, temp)
		return
	}

	// 单层搜索逻辑
	for i := 0; i < len(nums); i ++ {
		// ⚠️注意：去重开始
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}
		// ⚠️注意：去重结束

		// 当前位置使用过
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
