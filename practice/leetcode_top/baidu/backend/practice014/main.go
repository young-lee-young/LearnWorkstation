package main

import (
	"fmt"
	"strconv"
)

/**
字符串压缩
*/
func main() {
	str := "aabcccccaaa"
	ret := solution(str)
	fmt.Println("ret:", ret)
}

func solution(str string) string {
	ret := ""

	if str == "" || len(str) == 1 {
		return str
	}

	current := str[0]
	count := 1

	for i := 1; i < len(str); i++ {
		s := str[i]
		if s == current {
			count += 1
		} else {
			ret = ret + string(current) + strconv.Itoa(count)
			count = 1
			current = s
		}
	}

	ret = ret + string(current) + strconv.Itoa(count)

	if len(ret) >= len(str) {
		return str
	}

	return ret
}
