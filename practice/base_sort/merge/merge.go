package merge

/**
	05 归并排序

	思路：先对半递归分组，对每个组进行排序，排序后再合并两个数组

	Level 0                   8 6 2 3 1 5 7 4
	Level 1          8 6 2 3                 1 5 7 4
	Level 2      8 6         2 3         1 5         7 4
	Level 3    8     6     2     3     1     5     7     4

	一共有n个元素，层数level = logn
	每层都有n个元素

	每层元素排序时间复杂度为O(n)，总的时间复杂度为O(nlogn)
 */
func MergeSort(nums []int) []int {
	count := len(nums)
	mergeSort(nums, 0, count - 1)
	return nums
}

/**
	递归算法，对[l ... r]的范围进行排序
 */
func mergeSort(nums []int, left int, right int) {
	// 递归结束
	if left >= right {
		return
	}

	// 	优化：进行分组后，每组数据量小的时候，可以使用插入排序法
	// if right - left <= 15 {
	//     insertion.InsertionSort(nums, left, right)
	// }

	// 中间值，这里取第一个数组最后一个元素索引
	mid := (left + right) / 2
	mergeSort(nums, left, mid)
	mergeSort(nums, mid + 1, right)
	merge(nums, left, mid, right)
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
