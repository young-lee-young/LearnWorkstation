package main

import (
	"fmt"
	"sort"
)

/**
	前K个高频元素 LeetCode No347

	标签：最小堆
 */
func main() {
	num1 := []int{1, 1, 2, 2, 2, 3}
	k := 2
	result := solution(num1, k)
	fmt.Println(result)
}

func solution(num1 []int, k int) []int {
	numMap := make(map[int]int)
	numSlice := make([]int, 0)

	for _, num := range num1 {
		if times, ok := numMap[num]; ok {
			numMap[num] = times + 1
		} else {
			numMap[num] = 1
			numSlice = append(numSlice, num)
		}
	}

	// 这里最好使用最小堆实现，Go没有实现最小堆，这里只好使用slice排序
	fmt.Println(numMap)
	fmt.Println(numSlice)
	sort.Slice(numSlice, func(i, j int) bool {
		fmt.Println(i, j)
		return numMap[numSlice[i]] > numMap[numSlice[j]]
	})
	fmt.Println(numSlice)

	return numSlice[:k]
}
