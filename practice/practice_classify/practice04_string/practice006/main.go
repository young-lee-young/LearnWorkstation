/**
 * @Time:    2023/1/29 12:22 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

LeetCode No028 找出字符串中第一个匹配项的下标
 */
package main

import (
	"fmt"
)

func main() {
	haystack := "mississippi" // 干草堆

	needle := "issip" // 针

	ret := solution(haystack, needle)

	fmt.Println("ret:", ret)
}

func solution(haystack string, needle string) int {
	return -1
}
