/**
 * @Time:    2023/1/29 12:22 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

LeetCode No028 找出字符串中第一个匹配项的下标

解题思路：KMP 算法
 */
package main

import (
	"fmt"
)

func main() {
	haystack := "mississippi" // 干草堆

	needle := "issip" // 针

	ret := solution(haystack, needle)

	fmt.Println("ret:", ret)
}

func solution(haystack string, needle string) int {
	next := getNext(needle)

	j := 0

	n := len(needle)

	for i := 0; i < len(haystack); i ++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}

		if haystack[i] == needle[j] {
			j ++
		}

		if j == n {
			return i - n + 1
		}
	}

	return -1
}

/**
构建 next 数组的步骤

1. 初始化：初始化 next 数组和函数里的变量
2. 处理前后缀不相同的情况
3. 处理前后缀相同的情况
4. 更新 next 数组的值

i 指针：后缀末尾位置
j 指针：前缀末尾位置，也代表 [0, i] 最长相等前后缀长度
 */
func getNext(needle string) []int {
	next := make([]int, len(needle))

	j := 0

	next[0] = j

	for i := 1; i < len(needle); i ++ {
		for j > 0 && needle[i] != needle[j] {
			j = next[j-1]
		}

		if needle[i] == needle[j] {
			j ++
		}

		next[i] = j
	}

	return next
}
