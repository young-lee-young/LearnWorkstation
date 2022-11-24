package main

import "fmt"

/**
	字符串中第一个唯一字符 LeetCode No387

	标签：哈希表
 */
func main() {
	s := "loveleetcode"
	result := solution(s)
	fmt.Print(result)
}

func solution(s string) int {
	charArray := [26]int{}

	for i := 0; i < len(s); i ++ {
		charArray[s[i] - 'a'] ++
	}

	for i := 0; i < len(s); i ++ {
		if charArray[s[i] - 'a'] == 1 {
			return i
		}
	}

	return -1
}
