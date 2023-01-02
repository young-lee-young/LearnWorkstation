package main

import "fmt"

/**
	最长和谐序列 LeetCode No594

	解题思路：
 */
func main() {
	numArray := [...]int{1,3,2,2,5,2,3,7}
	hashMap := make(map[int]int)

	// 把所有元素放到map里，元素是键，个数为值
	for _, num := range numArray {
		hashMap[num] ++
	}

	result := 0
	// 循环map，如果 x + 1也在map里，统计这两个的数量和，取大的
	for key, value := range hashMap {
		if _, ok := hashMap[key + 1]; ok {
			count := value + hashMap[key + 1]
			if count > result {
				result = count
			}
		}
	}
	fmt.Println(result)
}
