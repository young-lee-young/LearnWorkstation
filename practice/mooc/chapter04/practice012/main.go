/**
 * @Time:    2021/7/3 17:38
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:
 */
package main

import "fmt"

/**
四数之和 LeetCode No454

标签：查找表map
*/
func main() {
	nums1 := []int{1, 2}
	nums2 := []int{-2, -1}
	nums3 := []int{-1, 2}
	nums4 := []int{0, 2}
	result := solution(nums1, nums2, nums3, nums4)
	fmt.Println(result)
}

func solution(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	numMap := make(map[int]int)
	for i := 0; i < len(nums3); i++ {
		for j := 0; j < len(nums4); j++ {
			numMap[nums3[i]+nums4[j]]++
		}
	}

	result := 0
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if value, ok := numMap[0-nums1[i]-nums2[j]]; ok {
				result += value
			}
		}
	}
	return result
}