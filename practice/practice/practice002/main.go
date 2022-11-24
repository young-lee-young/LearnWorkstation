package main

import "fmt"

/**
	走楼梯问题 LeetCode No70

	标签：动态规划

	解决思路：第n级台阶可以从 (n - 1) 级走1步或者 (n - 2) 级走两部，f(n) = f(n - 1) + f(n - 2)
 */
func main() {
	result := solution(4)
	fmt.Println(result)
}

func solution(n int) int {
	if n <= 2 {
		return n
	}
	f1, f2, f3 := 1, 2, 3

	i := 3
	for i <= n {
		f3 = f1 + f2
		f1 = f2
		f2 = f3
		i ++
	}
	return f3
}
