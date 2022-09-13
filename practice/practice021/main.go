package main

import (
	"fmt"
	queue2 "practice/base/queue"
)

/**
	扫雷 LeetCode No529

	标签：广度优先遍历
 */
func main() {
	board := [][]byte{
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'M', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
	}
	click :=[]int{3, 0}
	result := solution(board, click)
	fmt.Println(result)
}

func solution(board [][]byte, click []int) [][]byte {
	x := click[0]
	y := click[1]
	rowLen := len(board)
	columnLen := len(board[0])

	// 挖出地雷
	if board[x][y] == 'M' {
		board[x][y] = 'X'
		return board
	}

	// 挖出空方块
	if board[x][y] == 'E' {
		count := checkAround(x, y, rowLen, columnLen, board)
		// 如果周围有雷，直接改为周围的雷数返回
		if count != 0 {
			board[x][y] = byte(count + '0')
			return board
		}
		// 周围没有雷，使用广度优先搜索
		queue := queue2.Queue{}
		// 将当前及周围方块坐标加入到队列
		board[x][y] = 'B'
		coordinate := []int{x, y}
		queue.Enqueue(coordinate)

		// 标准广度优先
		for !queue.IsEmpty() {
			queueLen := queue.GetLength()
			fmt.Println(queueLen)
			for k := 0; k < queueLen; k ++ {
				// 当前坐标
				coordinateTemp := queue.Dequeue().([]int)
				// 周围的值
				for i := -1; i <= 1; i ++ {
					xTemp := coordinateTemp[0] + i
					if xTemp < 0 || xTemp > rowLen - 1 {
						continue
					}
					for j := -1; j <= 1; j ++ {
						yTemp := coordinateTemp[1] + j
						if yTemp < 0 || yTemp > columnLen - 1 {
							continue
						}
						// 把自己排除
						if i == 0 && j == 0 {
							continue
						}
						count := checkAround(xTemp, yTemp, rowLen, columnLen, board)
						// 周围的雷为0，置为B，并添加到队列中
						if count == 0 && board[xTemp][yTemp] != 'B' {
							board[xTemp][yTemp] = 'B'
							coordinateTemp := []int{xTemp, yTemp}
							queue.Enqueue(coordinateTemp)
						} else if board[xTemp][yTemp] == 'B' {
							continue
						} else {
							board[xTemp][yTemp] = byte(count + '0')
						}
					}
				}
			}
		}
	}

	// 其他情况，返回原数组
	return board
}

/**
	检查周围雷数
 */
func checkAround(x int, y int, rowLen int, columnLen int, board [][]byte) int {
	count := 0
	for i := -1; i <= 1; i ++ {
		// 横向越界
		if x + i < 0 || x + i > rowLen - 1 {
			continue
		}
		for j := -1; j <= 1; j ++ {
			// 竖向越界
			if y + j < 0 || y + j > columnLen - 1 {
				continue
			}

			if i == 0 && j == 0 {
				continue
			}

			if board[x + i][y + j] == 'M' {
				count ++
			}
		}
	}
	return count
}
