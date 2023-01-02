/**
 * @Time:    2022/11/27 19:36 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	LeetCode No090 子集Ⅱ

	标签：回溯

	⚠️总结：这种结果集需要去重的，需要将数组排序，使用一个数组或哈希表记录索引位置上的的数字是否使用过

	去重时的逻辑：	如果当前数和上一个数相同，并且上一个数没有使用，则当前数也不使用，直接 continue
	if i > 0 && nums[i] == nums[i - 1] && !used[i - 1] {
		continue
	}
 */
package main

import "fmt"

var ret [][]int

func main() {
	ret = [][]int{}

	nums := []int{1, 2, 2}

	path := []int{}

	used := make(map[int]bool)

	quickSort(nums, 0, len(nums) -1)

	backtrack(nums, 0, path, used)

	fmt.Println("ret:", ret)
}

func backtrack(nums []int, start int, path []int, used map[int]bool) {
	// 收集结果
	temp := make([]int, len(path))
	copy(temp, path)
	ret = append(ret, temp)

	// 递归结束条件
	if start == len(nums) {
		return
	}

	// 单层搜索逻辑
	for i := start; i < len(nums); i ++ {
		// ⚠️注意：去重开始
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}
		// ⚠️注意：去重结束

		// 处理节点
		path = append(path, nums[i])
		used[i] = true

		// 递归函数
		backtrack(nums, i+1, path, used)

		// 回溯
		path = path[:len(path)-1]
		used[i] = false
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