/**
LeetCode No349 两个数组的交集

标签：集合
*/
package main

import "fmt"

func main() {
	num1 := []int{}
	num2 := []int{}

	ret := solution(num1, num2)

	fmt.Println("ret:", ret)
}

func solution(num1 []int, num2 []int) []int {
	// 用map实现set
	mapSet := make(map[int]int)

	// num1去重
	for _, num := range num1 {
		mapSet[num] = num
	}

	result := make([]int, 0)
	for _, num := range num2 {
		// 当前数字在num1里存在，放到结果集里
		if _, ok := mapSet[num]; ok {
			result = append(result, num)
			// 将当前值从set里删除，防止添加到result里重复元素
			delete(mapSet, num)
		}
	}
	return result
}
