/**
 * @Time:    2021/3/11 15:26 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content: 调整数组顺序使奇数位于偶数后面
 */
package main

import (
	"practice/utils/random"
	"fmt"
)

/**
	调整数组顺序使奇数位于偶数后面 剑指Offer No21

	标签：双指针
 */
func main() {
	numSlice := random.GenerateArray(10, 100)
	fmt.Println(numSlice)
	result := solution(numSlice)
	fmt.Println(result)
}

func solution(numSlice []int) []int {
	start := 0
	end := len(numSlice) - 1

	for start < end {
		numStart := numSlice[start]
		numEnd := numSlice[end]
		if numStart % 2 == 0 {
			// 前后都是偶数，后面指针前移
			if numEnd % 2 == 0 {
				end --
			// 前面的是偶数，后面的是奇数，交换位置
			} else {
				temp := numSlice[start]
				numSlice[start] = numSlice[end]
				numSlice[end] = temp
				start ++
				end --
			}
		} else {
			// 前面是奇数，后面是偶数，只需移动前后指针就好
			if numEnd % 2 == 0 {
				start ++
				end --
			// 前后都是奇数，前面指针后移
			} else {
				start ++
			}
		}
	}
	return numSlice
}
