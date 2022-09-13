package main

import "fmt"

/**
	无序数组两数之和返回下标 LeetCode No1

	解题思路：利用map，键为值、值索引
 */
func main() {
	numArray := []int{2, 7, 11, 15}
	sum := 9

	hashMap := make(map[int]int)
	for i := 0; i < len(numArray); i ++ {
		num := sum - numArray[i]
		if value, ok := hashMap[num]; ok {
			fmt.Println(value, i)
			break
		}
		hashMap[numArray[i]] = i
	}
}
