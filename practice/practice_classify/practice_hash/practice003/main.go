package main

import (
	"fmt"
	"reflect"
)

/**
	是否字母异位（字母个数相同，位置不同） LeetCode No242

	解题思路：键存字符 值为个数
 */
func main() {
	s1 := "leeyoung"
	s2 := "younglee"
	result := solution(s1, s2)
	fmt.Println(result)
}

func solution(s1 string, s2 string) bool {
	hashMap1 := make(map[string]int)
	hashMap2 := make(map[string]int)
	for _, value := range s1 {
		hashMap1[string(value)]++
	}
	for _, value := range s2 {
		hashMap2[string(value)]++
	}
	return reflect.DeepEqual(hashMap1, hashMap2)
}
