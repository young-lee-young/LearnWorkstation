package quick_three_ways

import "practice/utils/random"

/**
	三路快速排序法（用于解决多重复元素）

	使用多余双路快速排序
 */
func QuickSortThreeWays(nums []int) []int {
	count := len(nums)
	quickSort(nums, 0, count - 1)
	return nums
}

/**
	三路快速排序nums[left ... right]
	将nums[left ... right]分为 < first; == first; > first三部分
	之后递归对 < first; > first 两部分继续进行三路快速排序
 */
func quickSort(nums []int, left int, right int) {
	// 递归结束
	if left >= right {
		return
	}

	/**
		partition操作
	 */
	// 交换第一个值和随机选择值
	temp := nums[left]
	randomIndex := random.GenerateIntFromMinMax(left,right)
	nums[left] = nums[randomIndex]
	nums[randomIndex] = temp

	// 取第一个值
	first := nums[left]

	lt := left // nums[left + 1 ... lt] < first
	i := left + 1 // nums[lt + 1 ... i] == first
	gt := right + 1 // nums[gt ... right] > first

	for i < gt {
		if nums[i] < first {
			// 交换小于和等于
			temp := nums[i]
			nums[i] = nums[lt + 1]
			nums[lt + 1] = temp

			lt ++
			i ++
		} else if nums[i] > first {
			temp := nums[i]
			nums[i] = nums[gt - 1]
			nums[gt - 1] = temp

			gt --
		} else {
			i ++
		}
	}
	nums[left] = nums[lt]
	nums[lt] = nums[left]
	/**
		partition结束
	 */

	 quickSort(nums, left, lt - 1)
	 quickSort(nums, gt, right)
}
