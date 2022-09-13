/**
 * @Time:    2021/3/31 22:58 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content: 滑动谜题
 */
package main

import (
	queue2 "practice/base/queue"
	"strconv"
	"fmt"
)

/**
	滑动谜题 LeetCode No773

	标签 广度优先遍历

	这种问题的时间复杂度为O(n!)，n表示图可能达到的状态，时间复杂度比O(n^2）要大
 */
func main() {
	board := [][]int{{4, 1, 2}, {5, 0, 3}}
	//board := [][]int{{1, 2, 3}, {5, 4, 0}}
	result := solution(board)
	fmt.Println(result)
}

func solution(board [][]int) int {
	initState := boardToStr(board)
	if initState == "123450" {
		return 0
	}

	// 广度优先遍历
	queue := queue2.Queue{}
	queue.Enqueue(initState)
	visited := make(map[string]int)
	visited[initState] = 0

	for !queue.IsEmpty() {
		current := queue.Dequeue().(string)

		// 获取下一步所有状态
		nextSlice := getNextSlice(current)

		// 循环当前状态所有可能下一步状态
		for _, next := range nextSlice {
			// 没有被访问过
			if _, ok := visited[next]; !ok {
				queue.Enqueue(next)
				// 步数等于当前节点步数 + 1
				visited[next] = visited[current] + 1
				if next == "123450" {
					return visited[next]
				}
			}
		}
	}
	return -1
}

func boardToStr(board [][]int) string {
	result := ""
	for i := 0; i < 2; i ++ {
		for j := 0; j < 3; j ++ {
			result += strconv.Itoa(board[i][j])
		}
	}
	return result
}

func strToBoard(s string) [][]int {
	board := make([][]int, 2)
	for i := 0; i < 2; i ++ {
		board[i] = make([]int, 3)
	}

	for n := 0; n < 6; n ++ {
		board[n / 3][n % 3] = int(s[n] - '0')
	}
	return board
}

func getNextSlice(s string) []string {
	// 将字符串转为二维数组
	board := strToBoard(s)
	// 查找0所在位置
	var zeroIndex int
	for zeroIndex = 0; zeroIndex < 6; zeroIndex ++ {
		if board[zeroIndex / 3][zeroIndex % 3] == 0 {
			break
		}
	}

	// 0的坐标
	zeroX := zeroIndex / 3
	zeroY := zeroIndex % 3

	result := make([]string, 0)
	// 四联通位置
	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for d := 0; d < 4; d ++ {
		nextX := zeroX + directions[d][0]
		nextY := zeroY + directions[d][1]

		// 检查x和y的合法性
		if inArea(nextX, nextY) {
			swap(board, zeroX, zeroY, nextX, nextY)
			result = append(result, boardToStr(board))
			swap(board, zeroX, zeroY, nextX, nextY)
		}
	}

	return result
}

func swap(board [][]int, x int, y int, nextX int, nextY int) {
	temp := board[x][y]
	board[x][y] = board[nextX][nextY]
	board[nextX][nextY] = temp
}

func inArea(x int, y int) bool {
	return x >= 0 && x < 2 && y >= 0 && y < 3
}
