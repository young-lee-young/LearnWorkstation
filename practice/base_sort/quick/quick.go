package quick

import (
	"practice/utils/random"
)

/**
06 快速排序

思路：以第一个数为基准，小于这个数的放在前面，大于这个数的放在后面
*/
func QuickSort(nums []int) []int {
	count := len(nums)
	quickSort(nums, 0, count-1)
	return nums
}

/**
对 nums [left ... right]部分进行快速排序
*/
func quickSort(nums []int, left int, right int) {
	// 递归结束
	if left >= right {
		return
	}

	// 	优化：进行分组后，每组数据量小的时候，可以使用插入排序法
	// if right - left <= 15 {
	//     insertion.InsertionSort(nums, left, right)
	// }

	partition := partition(nums, left, right)
	quickSort(nums, left, partition-1)
	quickSort(nums, partition+1, right)
}

/**
对nums[left ... right]部分进行partition操作
返回p，使得nums[left ... p - 1] < nums[p]，nums[p + 1 ... right] > nums[p]
*/
func partition(nums []int, left int, right int) int {
	/**
	假如数组本来有序，选择第一个值为基准，则快速排序退化为O(n^2)，解决办法为随机选择基准
	*/
	// 取数组第一个元素
	temp := nums[left]
	// 随机选择
	randomIndex := random.GenerateIntFromMinMax(left, right)
	// 交换第一个值和随机选择的值
	nums[left] = nums[randomIndex]
	nums[randomIndex] = temp

	// 数组第一个元素
	first := nums[left]

	// 第一个元素应该在的位置的索引
	partition := left
	// 循环数组，[left ... partition] < first，nums[partition + 1, i) > first
	for i := left + 1; i <= right; i++ {
		// 如果当前值大于基准值，还在原来的位置，如果小于基准值，将当前值和partition后一个位置的值交换，并将partition前移一位
		if nums[i] < first {
			temp := nums[i]
			nums[i] = nums[partition+1]
			nums[partition+1] = temp
			partition++
		}
	}

	// 交换第一个值和partition位置的值
	nums[left] = nums[partition]
	nums[partition] = first

	return partition
}
