package main

import "fmt"

/**
	两个数组交集 LeetCode No350

	标签：映射
 */
func main() {
	num1 := []int{}
	num2 := []int{}
	result := solution(num1, num2)
	fmt.Println(result)
}

func solution(num1 []int, num2 []int) []int {
	numMap := make(map[int]int)

	// 遍历数组，将每个数组出现的频次记录在map中
	for _, num := range num1 {
		if _, ok := numMap[num]; !ok {
			numMap[num] = 1
		} else {
			numMap[num] = numMap[num] + 1
		}
	}

	result := make([]int, 0)
	for _, num := range num2 {
		// 循环数组2，如果在数组1中，加入到结果集中
		if _, ok := numMap[num]; ok {
			result = append(result, num)
			// 并将数组1中的频次减1
			numMap[num] = numMap[num] - 1
			// 如果频次等于0，直接把键删除
			if numMap[num] == 0 {
				delete(numMap, num)
			}
		}

	}
	return result
}
