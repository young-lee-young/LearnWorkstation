/**
 * @Time:    2022/11/24 14:31 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

允许删除一个字符，判断是否是回文字符串 LeetCode No680

标签：双指针
 */
package main

import "fmt"

func main() {
	s := "abc"
	ret := solution(s, false)
	fmt.Println("ret:", ret)
}

func solution(s string, deleted bool) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] == s[right] {
			left ++
			right --
			continue
		}

		// 已经删除过一个字符
		if deleted {
			return false
		}

		// 去除左边字符
		s1 := s[left+1 : right+1]

		// 去除右边字符
		s2 := s[left:right]

		if solution(s1, true) || solution(s2, true) {
			return true
		}

		// 去除左右字符都不是回文串
		return false
	}

	return true
}
