package main

import (
	"fmt"
	"strings"
)

/**
	单词规律 LeetCode No290

	标签：查找表
*/
func main() {
	pattern := "abba"
	str := "dog cat cat dog"
	//str := "dog dog dog dog"
	result := solution(pattern, str)
	fmt.Println(result)
}

func solution(pattern string, str string) bool {
	strs := strings.Split(str, " ")
	if len(pattern) != len(strs) {
		return false
	}

	patternStrMap := make(map[byte]string, 0)
	strPatternMap := make(map[string]byte, 0)

	for i := 0; i < len(pattern); i++ {
		patternItem := pattern[i]
		if value, ok := patternStrMap[patternItem]; ok {
			if strs[i] != value {
				return false
			}
		} else {
			if value, ok := strPatternMap[strs[i]]; ok {
				if value != patternItem {
					return false
				}
			}

			strPatternMap[strs[i]] = patternItem
			patternStrMap[patternItem] = strs[i]
		}
	}

	return true
}
