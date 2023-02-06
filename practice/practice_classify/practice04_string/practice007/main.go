/**
 * @Time:    2023/2/5 22:05
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

LeetCode No459 重复的子字符串
*/
package main

import "fmt"

func main() {
	s := "ababa"

	ret := solution(s)

	fmt.Println("ret:", ret)
}

func solution(s string) bool {

	return false
}

/**
暴力解法：

逐渐加长子串长度，从剩余的字符串里面从前往后截取

*/
func force(s string) bool {
	for i := 1; i < len(s); i++ {
		sub := s[0:i]

		subLen := len(sub)

		j := i

		for j+subLen <= len(s) {
			cut := s[j : j+subLen]

			if cut != sub {
				break
			}

			j = j + subLen
		}

		if j == len(s) {
			return true
		}
	}

	return false
}
