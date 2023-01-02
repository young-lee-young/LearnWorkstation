/**
 * @Time:    2022/11/27 00:46 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	LeetCode No131 字符串分割后的回文串

	标签：回溯
 */
package main

import "fmt"

var ret [][]string

func main() {
	ret = [][]string{}

	s := "aab"

	path := []string{}

	backtrack(s, 0, path)

	fmt.Println("ret:", ret)
}

func backtrack(s string, start int, path []string) {
	// 递归结束条件
	if start == len(s) {
		temp := make([]string, len(path))
		copy(temp, path)
		ret = append(ret, temp)
		return
	}

	// 单层搜索逻辑
	for i := start; i < len(s); i ++ {
		// 处理节点
		checkRet := check(s, start, i)
		if !checkRet {
			continue
		}

		path = append(path, s[start:i+1]) // ⚠️注意：这里区间取的 [start : i + 1]

		// 递归函数
		backtrack(s, i+1, path)

		// 回溯
		path = path[:len(path)-1]
	}
}

// 判断回文
func check(s string, start int, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start ++
		end --
	}
	return true
}
