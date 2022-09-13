package main

import "fmt"

/**
	x的n次幂 LeetCode No50

	标签：回溯

	解题思路： pow(x, n) := pow(x * x, n / 2) * (x if n % 2 else 1)
             pow(x, 0) := 1
 */
func main() {
	var num float64
	var count int

	num = 2
	count = -2

	var result float64
	if count >= 0 {
		result = solution(num, count)
	} else {
		result= 1.0 / solution(num, count)
	}
	fmt.Println(result)
}

func solution(num float64, count int) float64 {
	if count == 0 {
		return 1
	}
	if count % 2 == 0 {
		return solution(num * num, count / 2) * 1
	}
	return solution(num * num, count / 2) * num
}
