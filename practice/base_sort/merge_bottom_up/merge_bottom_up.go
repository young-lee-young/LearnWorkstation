package merge_bottom_up

import "practice/utils/compare"

/**
	自底向上的归并排序
 */
func MergeSortBottomUp(nums []int) []int {
	count := len(nums)
	for size := 1; size <= count; size += size {
		for i := 0; i + size < count; i += size + size {
			merge(nums, i, i + size - 1, compare.Min(i + size + size - 1, count - 1))
		}
	}
	return nums
}

/**
	将 [left ... mid] 和 [mid + 1 ... right] 两部分进行合并
 */
func merge(nums []int, left int, mid int, right int) {
	// 将旧元素拷贝一份
	numsCopy := make([]int, right - left + 1)
	for i := left; i <= right; i ++ {
		numsCopy[i - left] = nums[i]
	}

	// i是第一个数组索引，j是第二个数组中的索引
	i := left
	j := mid + 1
	for k := left; k <= right; k ++ {
		// 第一个数组已经遍历完
		if i > mid {
			nums[k] = numsCopy[j - left]
			j ++
			// 第二个数组已经遍历完
		} else if j > right {
			nums[k] = numsCopy[i - left]
			i ++
			// 第一个数组中的元素比第二个元素小
		} else if numsCopy[i - left] < numsCopy[j - left] {
			nums[k] = numsCopy[i - left]
			i ++
			// 第一个数组中的元素比第二个元素大
		} else {
			nums[k] = numsCopy[j - left]
			j ++
		}
	}
}
