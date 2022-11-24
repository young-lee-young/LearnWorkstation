package main

import "fmt"

/**
	字母异位词分组 LeetCode No49

	解题思路：
 */
func main() {
	array := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	hashMap := map[string][]string{}
	for _, value := range array {
		// 将每个单词放到26位数组里，值为次数
		charSlice := [26]int{}
		for _, char := range value {
			charSlice[char - 'a'] ++
		}
		// 将次数拼接成字符串，作为map的键，单词作为值
		var key string
		for _, num := range charSlice {
			key += string(num)
		}
		hashMap[key] = append(hashMap[key], value)
	}

	// 键map值append到切片中
	result := make([][]string, 0)
	for _, value := range hashMap {
		result = append(result, value)
	}
	fmt.Println(result)
}
