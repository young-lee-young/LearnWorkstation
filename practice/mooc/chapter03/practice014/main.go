/**
 * @Time:    2021/6/13 15:52
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:
 */
package main

import (
	"fmt"
	"practice/utils/compare"
)

/**
最长不重复子串 LeetCode No003

标签：滑动窗口
*/
func main() {
	s := "abcabcbb"
	ret := solution(s)
	fmt.Println("ret:", ret)
}

func solution(s string) int {
	// 滑动窗口为[left ... right]
	left := 0
	right := -1
	// 表示一个字符在滑动窗口中出现的次数
	freq := [256]int{0}
	result := 0

	for left < len(s) {
		// 滑动窗口右边界不越界，并且右边扩展的字符没有在滑动窗口中，右边界向右移动，将扩展的字符串频次+1
		if right+1 < len(s) && freq[s[right+1]] == 0 {
			right++
			freq[s[right]]++ // 此时的 right 已经加完 1
			// 左侧边界字符频次-1，并将左侧边界右移
		} else {
			freq[s[left]]--
			left++
		}

		result = compare.Max(result, right-left+1)
	}

	return result
}
