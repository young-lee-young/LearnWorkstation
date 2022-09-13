package main

import (
	"fmt"
)

func main() {

	fmt.Println(191 % 10)
}

/**
递归
*/
func recursion(num int) int {
	fmt.Println(num)
	result := recursion(num + 1)
	fmt.Println(result)
	return result
}

/**
字符串转字符串列表
*/
func strToSlice(str string) []string {
	strSlice := make([]string, 0)
	for _, value := range str {
		strSlice = append(strSlice, string(value))
	}
	return strSlice
}

/**
字符串转字符列表
*/
func strToRuneSlice(str string) []rune {
	runeSlice := make([]rune, 0)
	for _, item := range str {
		runeSlice = append(runeSlice, item)
	}
	return runeSlice
}
