/**
 * @Time:    2023/1/28 12:00 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

LeetCode No151 反转字符串中的单词
 */
package main

import (
	"fmt"
)

func main() {
	s := "the sky is blue"

	ret := solution(s)

	fmt.Println("ret:", ret)
}

func solution(s string) string {
	// 先将整个字符串反转
	s = reverse(s)

	sb := ([]byte)(s)

	// 将每个单词反转
	slow := 0
	for fast := 0; fast < len(sb); fast ++ {
		if sb[fast] != ' ' {

		}
	}

	return string(sb)
}

func reverse(s string) string {
	sb := ([]byte)(s)

	i := 0
	j := len(sb) - 1

	for i <= j {
		sb[i], sb[j] = sb[j], sb[i]
		i ++
		j --
	}

	return string(sb)
}
