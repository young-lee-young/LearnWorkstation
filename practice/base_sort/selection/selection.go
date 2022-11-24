package selection

/**
	01 选择排序

	思路：从当前未排序数组中选择最小的值
 */
func SelectionSort(nums []int) []int {
	count := len(nums)

	for i := 0; i < count - 1; i ++ {
		// 找到 [i, n) 区间最小值索引
		minIndex := i
		for j := i + 1; j < count; j ++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}

		// 交换当前位置值和最小值
		temp := nums[i]
		nums[i] = nums[minIndex]
		nums[minIndex] = temp
	}
	return nums
}
