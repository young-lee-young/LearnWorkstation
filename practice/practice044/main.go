package main

import "fmt"

/**
	摩斯密码 LeetCode No804

	标签：集合
 */
func main() {
	wordArray := []string{"gin", "zen", "gig", "msg"}
	result := solution(wordArray)
	fmt.Println(result)
}

func solution(wordArray []string) int {
	// 构造字母和密码之间映射关系
	wordMap := make(map[int32]string)
	wordSlice := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	for i := 0; i < 26; i ++ {
		wordMap['a' + int32(i)] = wordSlice[i]
	}

	mapSet := make(map[string]string)
	// 遍历单词
	for _, word := range wordArray {
		str := ""
		// 遍历每个单词字符，找到对应的密码，拼接
		for _, char := range word {
			str = str + wordMap[char]
		}
		mapSet[str] = str
	}
	return len(mapSet)
}
