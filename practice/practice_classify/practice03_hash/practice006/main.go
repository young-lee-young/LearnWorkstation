/**
LeetCode No383 赎金信
*/
package main

import "fmt"

func main() {
	ransomNote := "aa"

	magazine := "aab"

	ret := solution(ransomNote, magazine)

	fmt.Println("ret:", ret)
}

func solution(ransomNote string, magazine string) bool {
	charMap := make(map[int32]int)

	for _, char := range magazine {
		if _, ok := charMap[char]; ok {
			charMap[char]++
			continue
		}
		charMap[char] = 1
	}

	for _, char := range ransomNote {
		if _, ok := charMap[char]; !ok {
			return false
		}

		count := charMap[char]

		if count <= 0 {
			return false
		}

		charMap[char]--
	}

	return true
}
