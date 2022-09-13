package main

import "fmt"

/**
	快乐数：循环各个位的平方和是否位1 LeetCode No202

	标签：查找表
*/
func main() {
	num := 19
	result := solution(num)
	fmt.Println(result)
}

func solution(num int) bool {
	num = countSum(num)
	set := make(map[int]struct{}, 0)

	for num != 1 {
		// 如果不在set中，则将当前值添加到set中
		if _, ok := set[num]; !ok {
			set[num] = struct{}{}
			num = countSum(num)
			// 如果在set中，则存在循环，直接退出
		} else {
			return false
		}
	}
	return true
}

func countSum(num int) int {
	sum := 0
	for num > 0 {
		// 取余求平方和
		remainder := num % 10
		sum += remainder * remainder
		// 取商
		num = num / 10
	}
	return sum
}
