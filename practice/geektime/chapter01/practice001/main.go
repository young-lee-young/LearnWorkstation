/**
 * @Time:    2022/4/20 13:31 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	删除有序数组中的重复项 LeetCode No26

	思想：保序操作数组，使用过滤器思想

 */
package main

import "fmt"

func main() {
	nums := []int{1, 1, 2}

	result := solution(nums)
	fmt.Println(result)
}

func solution(nums []int) int {
	n := 0
	for i := 0; i < len(nums); i ++ {
		if i == 0 || nums[i] != nums[i - 1] {
			nums[n] = nums[i]
			n ++
		}
	}
	return n
}
