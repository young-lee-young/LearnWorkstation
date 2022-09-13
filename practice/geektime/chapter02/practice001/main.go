/**
 * @Time:    2022/4/21 10:17 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	有效的括号 LeetCode No20

 */
package main

import (
	"fmt"
	stack2 "practice/base/stack"
)

func main() {
	str := "["

	ret := solution(str)
	fmt.Println("ret", ret)
}

func solution(str string) bool {
	charMap := map[string]string{
		"}": "{",
		")": "(",
		"]": "[",
	}

	stack := stack2.Stack{}
	for i := 0; i < len(str); i ++ {
		char := string(str[i])

		if char == "{" || char == "(" || char == "[" {
			stack.Push(char)
		} else {
			if stack.IsEmpty() {
				return false
			}
			if stack.Peek().(string) != charMap[char] {
				return false
			}
			stack.Pop()
		}
	}

	if stack.GetLength() != 0 {
		return false
	}

	return true
}
