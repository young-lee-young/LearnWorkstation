package insertion_optimize

/**
	插入排序优化
 */
func InsertionSortOptimize(nums []int) []int {
	count := len(nums)
	for i := 0; i < count; i ++ {
		// 保存当前值
		temp := nums[i]
		// 应该插入位置的索引
		var j int

		// 将当前元素数组依次后移
		for j = i; j > 0 && nums[j - 1] > temp; j -- {
			nums[j] = nums[j - 1]
		}
		nums[j] = temp
	}
	return nums
}
