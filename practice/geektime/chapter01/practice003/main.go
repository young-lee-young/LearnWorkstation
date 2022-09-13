/**
 * @Time:    2022/4/20 13:58 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	合并两个有序数组 LeetCode No88

 */
package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}

	solution(nums1, 3, nums2, 3)
}

func solution(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1

	for k := m + n - 1; k >= 0; k -- {
		// 什么时候要num[i]
		// 1. j出界
		// 2. i、j都没出界，要大的
		if j < 0 || (i >= 0 && nums1[i] >= nums2[j]) {
			nums1[k] = nums1[i]
			i --
		} else {
			nums1[k] = nums2[j]
			j --
		}
	}
	fmt.Println("ret", nums1)
}
