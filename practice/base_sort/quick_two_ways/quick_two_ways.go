package quick_two_ways

import "practice/utils/random"

/**
	双路快速排序法（用于解决多重复元素）
 */
func QuickSortTwoWays(nums []int) []int {
	count := len(nums)
	quickSort(nums, 0, count - 1)
	return nums
}

func quickSort(nums []int, left int, right int) {
	// 递归结束
	if left >= right {
		return
	}

	partition := partition(nums, left, right)
	quickSort(nums, left, partition - 1)
	quickSort(nums, partition + 1, right)
}

func partition(nums []int, left int, right int) int {
	// 交换第一个值和随机选择的值
	temp := nums[left]
	randomIndex := random.GenerateIntFromMinMax(left, right)
	nums[left] = nums[randomIndex]
	nums[randomIndex] = temp

	// 取第一个值
	first := nums[left]

	// nums[left + 1 ... i] <= first; nums[j ... right] >= first
	i := left + 1
	j := right
	for {
		// 总左边找到第一个大于first的值
		for i <= right && nums[i] < first {
			i ++
		}
		// 从右边找到第一个小于first的值
		for j >= left + 1 && nums[j] > first {
			j --
		}
		// 循环结束（有可能i和j中间的值都相等）
		if i > j {
			break
		}

		// 交换i和j所在位置的值
		temp := nums[i]
		nums[i] = nums[j]
		nums[j] = temp

		i ++
		j --
	}

	// 将开始的值和j位置的值交换
	nums[left] = nums[j]
	nums[j] = first

	return j
}
