/**
 * @Time:    2022/11/26 20:51 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	LeetCode No017 电话号码的字母组合

	标签：回溯
 */
package main

import "fmt"

var charMap map[byte][]string

var ret []string

func main() {
	charMap = make(map[byte][]string)

	charMap['2'] = []string{"a", "b", "c"}
	charMap['3'] = []string{"d", "e", "f"}
	charMap['4'] = []string{"g", "h", "i"}
	charMap['5'] = []string{"j", "k", "l"}
	charMap['6'] = []string{"m", "n", "o"}
	charMap['7'] = []string{"p", "q", "r", "s"}
	charMap['8'] = []string{"t", "u", "v"}
	charMap['9'] = []string{"w", "x", "y", "z"}

	ret = []string{}

	digits := "23"
	path := ""

	backtrack(digits, path)

	fmt.Println("ret:", ret)
}

func backtrack(digits string, path string) {
	// 递归结束条件
	if digits == "" {
		ret = append(ret, path)
		return
	}

	// 单层搜索逻辑
	char := digits[0]
	strList := charMap[char]
	for _, str := range strList {
		// 处理节点：path + str 为处理节点的步骤

		// 递归函数
		backtrack(digits[1:], path+str)

		// 回溯：digits[1:] 已经回溯为 digits
		// 注意⚠️：这里是隐匿回溯
	}
}
