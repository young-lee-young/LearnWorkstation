/**
 * @Time:    2022/4/21 12:10 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	基本计算器 LeetCode No227

	思路：
		1. 中缀表达式先转换为后缀表达式
		2. 后缀表达式求值

 */
package main

import (
	"fmt"
	stack2 "practice/base/stack"
	"strconv"
)

func main() {
	s := "3+2*2"

	ret := solution(s)
	fmt.Println("ret", ret)
}

func solution(s string) int {
	// 防止最后一个字符为数字，没有放入到tokens中
	s += " "

	// 后缀表达式列表
	tokens := make([]string, 0)

	// 用于存放计算符号
	stack := stack2.Stack{}

	numberStr := ""
	for i := 0; i < len(s); i ++ {
		char := s[i]

		// 如果是数字，将数字拼接
		if char >= '0' && char <= '9' {
			numberStr += string(char)
			continue
			// 如果不是数字，说明遇到了运算符号，把当前数字放入到后缀表达式列表中
		} else {
			if numberStr != "" {
				tokens = append(tokens, numberStr)
				numberStr = ""
			}
		}

		if char == ' ' {
			continue
		}

		// 获取当前操作符号的优先级
		currentRank := getRank(char)

		// 栈顶符号优先级比当前符号优先级高，栈顶的符号应该先算，将栈顶符号加入后缀表达式列表中
		for !stack.IsEmpty() && getRank(stack.Peek().(byte)) >= currentRank {
			tokens = append(tokens, string(stack.Pop().(byte)))
		}

		// 将当前符号入栈
		stack.Push(char)
	}

	// 栈不为空，将栈的符号弹出
	for !stack.IsEmpty() {
		tokens = append(tokens, string(stack.Pop().(byte)))
	}

	fmt.Println(tokens)
	return cal(tokens)
}

/**
计算符号优先级
 */
func getRank(char byte) int {
	switch char {
	case '+':
		return 1
	case '-':
		return 1
	case '*':
		return 2
	case '/':
		return 2
	default:
		return 0
	}
}

/**
	复用后缀表达式求值代码
 */
func cal(tokens []string) int {
	stack := stack2.Stack{}

	for i := 0; i < len(tokens); i ++ {
		char := tokens[i]
		num, err := strconv.Atoi(char)
		// 数字直接入栈
		if err == nil {
			stack.Push(num)
			continue
		}

		num2 := stack.Pop().(int)
		num1 := stack.Pop().(int)

		var ret int
		switch char {
		case "+":
			ret = num1 + num2
		case "-":
			ret = num1 - num2
		case "*":
			ret = num1 * num2
		case "/":
			ret = num1 / num2
		}

		stack.Push(ret)
	}

	return stack.Pop().(int)
}
