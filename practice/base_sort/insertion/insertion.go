package insertion

/**
	03 插入排序

	思路：将当前元素插入前面排好序的数组中
 */
func InsertionSort(nums []int) []int {
	count := len(nums)
	for i := 0; i < count; i ++ {
		for j := i; j > 0; j -- {
			// 当前值比前一个值小，进行交换
			if nums[j] < nums[j - 1] {
				temp := nums[j]
				nums[j] = nums[j - 1]
				nums[j - 1] = temp
			}
		}
	}
	return nums
}
