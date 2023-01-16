package main

import "fmt"

/**
最长连续序列 LeetCode No128

解题思路：
*/
func main() {
	numArray := []int{100, 4, 200, 1, 3, 2}

	hashMap := make(map[int]int)
	for _, num := range numArray {
		hashMap[num] = 0
	}

	result := 0
	for key, _ := range hashMap {
		count := 0
		for {
			if _, ok := hashMap[key]; ok {
				count++
				key++
				continue
			}
			break
		}
		if count > result {
			result = count
		}
	}
	fmt.Println(result)
}
