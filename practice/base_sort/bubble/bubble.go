package bubble

/**
	冒泡排序

	思路：每次将最大的值移动到未排序数组末尾
 */
func BubbleSort(nums []int) []int {
	count := len(nums)
	for i := 0; i < count; i ++ {
		// 后面的数已经有序，只需要比较 0 - (count - i - 1)这些数就可以了
		for j := 0; j < count - i - 1; j ++ {
			// 如果前一个数比后一个数大 交换两个数
			if nums[j] > nums[j + 1] {
				temp := nums[j]
				nums[j] = nums[j + 1]
				nums[j + 1] = temp
			}
		}
	}
	return nums
}
