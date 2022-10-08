package main

import (
	"fmt"
	"practice/utils/char"
	"strings"
)

/**
回文字符串，只考虑字母和数字，忽略大小写 No125

标签：对撞指针
*/
func main() {
	str := "A man, a plan, a canal: Panama"
	fmt.Println(str)
	result := solution(str)
	fmt.Println(result)
}

func solution(str string) bool {
	str = strings.ToUpper(str)
	left := 0
	right := len(str) - 1

	for left < right {
		if !char.IsAlpha(str[left]) && !char.IsDigit(str[left]) {
			left++
			continue
		}
		if !char.IsAlpha(str[right]) && !char.IsDigit(str[right]) {
			right--
			continue
		}

		if str[left] == str[right] {
			left++
			right--
			continue
		} else {
			return false
		}
	}
	return true
}
