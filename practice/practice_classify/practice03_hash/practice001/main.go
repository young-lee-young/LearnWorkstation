package main

import (
	"fmt"
	"reflect"
)

/**
LeetCode No242 有效的字母异位词（字母个数相同，位置不同）

解题思路：键存字符 值为个数
*/
func main() {
	s1 := "leeyoung"
	s2 := "younglee"

	ret := solution(s1, s2)

	fmt.Println("ret:", ret)
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
	// ⚠️这里的判断
	return reflect.DeepEqual(hashMap1, hashMap2)
}
