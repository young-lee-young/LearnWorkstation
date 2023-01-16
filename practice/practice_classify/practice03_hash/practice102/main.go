package main

import "fmt"

/**
判断数组是否有重复元素 LeetCode No217

解决思路：用map存元素
*/
func main() {
	numArray := [...]int{1, 2, 3, 1}
	hashMap := make(map[int]int)

	for _, num := range numArray {
		if _, ok := hashMap[num]; ok {
			fmt.Println(num)
			return
		}
		hashMap[num] = num
	}
}
