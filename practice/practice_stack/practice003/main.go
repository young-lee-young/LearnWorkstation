package main

import (
	"fmt"
	stack2 "practice/base/stack"
)

/**
	用栈实现括号匹配 LeetCode No20
 */
func main() {
	testString := "{([{}}])}"
	fmt.Println(solution(testString))
}

func solution(testString string) bool {
	bracketsMap := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	stack := stack2.Stack{}
	for i := 0; i < len(testString); i ++ {
		character := string(testString[i])
		if character == "(" || character == "[" || character == "{" {
			stack.Push(character)
			continue
		}
		top := stack.Peek()
		if top == nil {
			return false
		}
		if bracketsMap[character] == top.(string) {
			stack.Pop()
		} else {
			return false
		}
	}

	if !stack.IsEmpty() {
		return false
	}
	return true
}
