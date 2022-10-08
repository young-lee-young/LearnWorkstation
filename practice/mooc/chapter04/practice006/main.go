package main

import "fmt"

/**
同构字符串 LeetCode No205

标签：查找表
*/
func main() {
	//str1 := "paper"
	//str2 := "title"
	str1 := "badc"
	str2 := "baba"
	result := solution(str1, str2)
	fmt.Println(result)
}

func solution(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	byte1Map := make(map[byte]byte, 0)
	byte2Map := make(map[byte]byte, 0)

	for i := 0; i < len(str1); i++ {
		if value, ok := byte1Map[str1[i]]; ok {
			if str2[i] != value {
				return false
			}
		} else {
			if value, ok := byte2Map[str2[i]]; ok {
				if str1[i] != value {
					return false
				}
			}

			byte1Map[str1[i]] = str2[i]
			byte2Map[str2[i]] = str1[i]
		}
	}
	return true
}
