/**
 * @Time:    2022/11/27 18:09 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	LeetCode No093 复原IP地址

	标签：回溯
 */
package main

import (
	"fmt"
	"strings"
	"strconv"
)

var ret []string

func main() {
	ret = []string{}

	s := "101023"

	path := []string{}

	backtrack(s, 0, path)

	fmt.Println("ret:", ret)
}

func backtrack(s string, start int, path []string) {
	// 递归函数结束条件
	if start == len(s) {
		if len(path) == 4 {
			ip := strings.Join(path, ".")
			fmt.Println("ip", ip)
			ret = append(ret, ip)
		}
		return
	}

	if len(path) == 4 {
		return
	}

	// 单层搜索逻辑
	for i := start; i < len(s); i ++ {
		// 处理节点
		checkRet := check(s, start, i+1)
		if !checkRet {
			continue
		}

		path = append(path, s[start:i+1])

		// 递归函数
		backtrack(s, i+1, path)

		// 回溯
		path = path[:len(path)-1]
	}
}

func check(s string, start int, end int) bool {
	sub := s[start:end]

	if strings.HasPrefix(sub, "0") && end-start >= 2 { // ⚠️注意：️这里的 end - start >= 2
		return false
	}

	i, err := strconv.Atoi(sub)
	if err != nil || i > 255 {
		return false
	}

	return true
}
