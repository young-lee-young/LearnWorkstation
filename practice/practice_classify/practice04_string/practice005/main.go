/**
 * @Time:    2023/1/29 12:11 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

剑指offer No058-Ⅱ 左旋转字符串
 */
package main

import (
	"fmt"
)

func main() {
	s := "abcdefg"

	k := 2

	ret := solution(s, k)

	fmt.Println("ret:", ret)
}

/**
解题思路：
1. 反转区间为前 k 的子串
2. 反转区间为 k 到末尾的子串
3. 反转整个字符串
 */
func solution(s string, k int) string {
	sb := ([]byte)(s)

	reverse(&sb, 0, k-1)
	reverse(&sb, k, len(sb)-1)
	reverse(&sb, 0, len(sb)-1)

	return string(sb)
}

func reverse(sb *[]byte, begin int, end int) {
	for begin < end {
		(*sb)[begin], (*sb)[end] = (*sb)[end], (*sb)[begin]
		begin ++
		end --
	}
}
