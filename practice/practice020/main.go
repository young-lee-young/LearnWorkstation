package main

import "fmt"

/**
	岛屿数量 LeetCode No200

	标签：深度优先遍历

	解题思路：沉没陆地
 */
func main() {
	//row1 := []byte{'1', '1', '0', '0', '0'}
	//row2 := []byte{'1', '1', '0', '0', '0'}
	//row3 := []byte{'0', '0', '1', '0', '0'}
	//row4 := []byte{'0', '0', '0', '1', '1'}
	//grid := [][]byte{row1, row2, row3, row4}

	row1 := []byte{'1', '1', '1'}
	row2 := []byte{'0', '1', '0'}
	row3 := []byte{'1', '1', '1'}
	grid := [][]byte{row1, row2, row3}
	result := solution(grid)
	fmt.Println(result)
}

func solution(grid [][]byte) int {
	count := 0
	rowLen := len(grid)
	for i :=0; i < rowLen; i ++ {
		columnLen := len(grid[i])
		for j := 0; j < columnLen; j ++ {
			if grid[i][j] == '1' {
				// 沉没附近为1的陆地
				sinkIslands(i, j, rowLen, columnLen, &grid)
				fmt.Println(grid)
				count ++
			}
		}
	}
	return count
}

/**
	沉没陆地，使用深度优先搜索
 */
func sinkIslands(i int, j int, rowLen int, columnLen int, grid *[][]byte) {
	if (*grid)[i][j] == '1' {
		// 沉没左面陆地
		if j > 0 {
			sinkIslands(i, j - 1, rowLen, columnLen, grid)
		}
		(*grid)[i][j] = '0'
		if i > 0 {
			sinkIslands(i - 1, j, rowLen, columnLen, grid)
		}
		(*grid)[i][j] = '0'
		// 沉没右面陆地
		if j < columnLen - 1 {
			sinkIslands(i, j + 1, rowLen, columnLen, grid)
		}
		(*grid)[i][j] = '0'
		// 沉没下面陆地
		if i < rowLen - 1 {
			sinkIslands(i + 1, j , rowLen, columnLen, grid)
		}
		(*grid)[i][j] = '0'
	}
}
