/**
 * @Time:    2023/1/17 17:49 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

LeetCode No541 反转字符串Ⅱ
 */
package main

import (
	"fmt"
)

func main() {
	s := "abcd"

	k := 4

	ret := solution(s, k)

	fmt.Println("ret:", ret)
}

func solution(s string, k int) string {
	sb := ([]byte)(s)

	len := len(sb)

	start := 0

	for start+2*k <= len {
		reverse(sb, start, start+k-1)
		start = start + 2*k
	}

	if start+k <= len {
		reverse(sb, start, start+k-1)
	}

	if start+k > len {
		reverse(sb, start, len-1)
	}

	return string(sb)
}

func reverse(sb []byte, left int, right int) []byte {
	for left < right {
		sb[left], sb[right] = sb[right], sb[left]

		left ++

		right --
	}
	return sb
}
