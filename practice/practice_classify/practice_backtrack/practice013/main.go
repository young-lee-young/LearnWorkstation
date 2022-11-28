/**
 * @Time:    2022/11/28 16:40 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	LeetCode No051 N皇后

	标签：回溯
 */
package main

import (
	"strings"
	"fmt"
)

var ret [][]string

func main() {
	n := 4

	ret = [][]string{}

	path := make([][]string, 0)
	for i := 0; i < n; i ++ {
		row := make([]string, 0)
		for j := 0; j < n; j ++ {
			row = append(row, ".")
		}
		path = append(path, row)
	}

	backtrack(n, 0, path)

	fmt.Println("ret:", ret)
}

func backtrack(n int, row int, path [][]string) {
	// 递归结束条件
	if row == n {
		temp := make([]string, 0)
		for i := 0; i < len(path); i ++ {
			str := strings.Join(path[i], "")
			temp = append(temp, str)
		}

		ret = append(ret, temp)
		return
	}

	// 单层逻辑搜索
	for column := 0; column < n; column ++ {
		// 处理节点
		if !check(path, row, column, n) {
			continue
		}

		path[row][column] = "Q"

		// 递归函数
		backtrack(n, row+1, path)

		// 回溯
		path[row][column] = "."
	}
}

func check(board [][]string, i int, j int, n int) bool {
	// 检查行
	rowM := board[i]
	for _, column := range rowM {
		if column == "Q" {
			return false
		}
	}

	// 检查列
	for _, row := range board {
		if row[j] == "Q" {
			return false
		}
	}

	/**
	⚠️注意：这里的方向需要注意
		 —————————————————————> 正方向
	    |
	    |
		|
		|
		|
		|
		|
	  正方向

	45' 方向：p 为负方向，q 为正方向

	135' 方向：p 为负方向，q 为负方向
	 */

	// 检查 45 度角
	for p, q := i-1, j+1; p >= 0 && q < n; p, q = p-1, q+1 {
		if board[p][q] == "Q" {
			return false
		}
	}

	// 检查 135 度角
	for p, q := i-1, j-1; p >= 0 && q >= 0; p, q = p-1, q-1 {
		if board[p][q] == "Q" {
			return false
		}
	}

	return true
}
