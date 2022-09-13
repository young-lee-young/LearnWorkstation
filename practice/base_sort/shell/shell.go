package shell

/**
	04 希尔排序

	思路：插入排序的变种，设定一个间隔值，间隔值逐渐缩小为1
 */
func ShellSort(nums []int) []int {
	count := len(nums)
	gap := count

	for gap > 1 {
		gap = gap / 2
		// 从第gap个元素，逐个对其所在组进行直接插入排序操作
		for i := gap; i < count; i ++ {
			j := i
			// 对每个组进行插入排序，后面的值小于前面值，将两个值交换
			for j - gap >= 0 && nums[j] < nums[j - gap] {
				temp := nums[j]
				nums[j] = nums[j - gap]
				nums[j - gap] = temp
				j = j - gap
			}
		}
	}
	return nums
}
