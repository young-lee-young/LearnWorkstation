/**
 * @Time:    2022/4/20 13:48 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	移动零 LeetCode No283

 */
package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}

	ret := solution(nums)
	fmt.Println(ret)
}

func solution(nums []int) int {
	n := 0
	for i := 0; i < len(nums); i ++ {
		if nums[i] != 0 {
			nums[n] = nums[i]
			n ++
		}
	}
	// 补零
	for n < len(nums) {
		nums[n] = 0
		n ++
	}
	fmt.Println("ret", nums)
	return n
}
