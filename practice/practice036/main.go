package main

import "fmt"

/**
	最长公共子序列 LeetCode No1143

	标签：动态规划
 */
func main() {
	text1 := "abcde"
	text2 := "ace"
	result := solution(text1, text2)
	fmt.Println(result)
}

func solution(str1 string, str2 string) int {
	strArray := make([][]int, len(str1) + 1)
	for i := 0; i < len(str1) + 1; i ++ {
		strArray[i] = make([]int, len(str2) + 1)
	}

	// 从1开始循环到字符串末尾
	for i := 1; i < len(str1) + 1; i ++ {
		for j := 1; j < len(str2) + 1; j ++ {
			// 字符串相等，最长公共子序列加1
			if str1[i - 1] == str2[j - 1] {
				strArray[i][j] = strArray[i - 1][j - 1] + 1
			} else {
				// 不相等，取上方和左侧最大值
				strArray[i][j] = max(strArray[i - 1][j], strArray[i][j - 1])
			}
		}
	}
	return strArray[len(str1)][len(str2)]
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
