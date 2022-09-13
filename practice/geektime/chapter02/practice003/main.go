/**
 * @Time:    2022/4/21 11:38 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	逆波兰式求值 LeetCode No150

	思路：后缀表达式求值
		1. 遇到一个数，则把该数入栈
		2. 遇到运算符，把栈顶两个数取出进行计算，然后把结果入栈
		3. 扫描完成后，栈中剩的数为最终结果

	时间复杂度：O(n)

 */
package main

import (
	stack2 "practice/base/stack"
	"strconv"
	"fmt"
)

func main() {
	tokens := []string{"4", "13", "5", "/", "+"}

	ret := solution(tokens)
	fmt.Println("ret", ret)
}

func solution(tokens []string) int {
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
