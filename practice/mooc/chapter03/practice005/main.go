package main

import "fmt"

/**
只有0，1，2三个数 进行排序 LeetCode No75

标签：三路快排
*/
func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	//result := solution(nums)
	result := solution2(nums)
	fmt.Println(result)
}

/**
计数排序

遍历一遍，统计各个数出现的次数

时间复杂度：O(n)
空间复杂度：O(1)

缺点：需要扫描2遍数组
*/
func solution(nums []int) []int {
	count := [3]int{}

	// 统计0，1，2的个数
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
	}

	// 将个数将0，1，2填入数组中
	base := 0
	for i := 0; i < len(count); i++ {
		for j := 0; j < count[i]; j++ {
			nums[base+j] = i
		}
		base += count[i]
	}

	return nums
}

/**
三路快排

时间复杂度：O(n)
空间负载度：O(1)

优点：只遍历数组一遍
*/
func solution2(nums []int) []int {
	// 指定3个指针
	zeroIndex := -1       // nums[0 ... zeroIndex] == 0
	twoIndex := len(nums) // nums[twoIndex ... n - 1] == 2

	for i := 0; i < twoIndex; {
		if nums[i] == 1 {
			i++
		} else if nums[i] == 2 {
			twoIndex--
			// 交换值
			temp := nums[twoIndex]
			nums[twoIndex] = nums[i]
			nums[i] = temp

			// 这里i不需要++，下一次继续处理i这个值
		} else {
			zeroIndex++
			// 交换值
			temp := nums[zeroIndex]
			nums[zeroIndex] = nums[i]
			nums[i] = temp

			i++
		}
	}

	return nums
}
