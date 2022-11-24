/**
 * @Time:    2021/3/25 23:20 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content: 岛屿的最大面积
 */
package main

import (
	"fmt"
	"math"
)

/**
	岛屿的最大面积 LeetCode No695

	标签：图深度优先搜索

	优化
 */
func main() {
	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	result := solution(grid)
	fmt.Println(result)
}

func solution(grid [][]int) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	direction := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	graph := &Graph{
		Row:       len(grid),
		Column:    len(grid[0]),
		Grid:      grid,
		Direction: direction,
	}

	var result int
	visited := make([][]bool, graph.Row)
	for item := 0; item < len(visited); item ++ {
		visited[item] = make([]bool, graph.Column)
	}
	graph.Visited = visited
	for i := 0; i < graph.Row; i ++ {
		for j := 0; j < graph.Column; j ++ {
			if !graph.Visited[i][j] && graph.Grid[i][j] == 1 {
				result = int(math.Max(float64(result), float64(graph.dfs(i, j))))
			}
		}
	}
	return result
}

// 图结构体
type Graph struct {
	Row       int
	Column    int
	Grid      [][]int
	Direction [][]int
	Visited   [][]bool
}

// 判断x、y是否在grid中
func (g *Graph) inArea(x int, y int) bool {
	return x >= 0 && x < g.Row && y >= 0 && y < g.Column
}

func (g *Graph) dfs(x int, y int) int {
	g.Visited[x][y] = true
	result := 1
	for d := 0; d < 4; d ++ {
		nextX := x + g.Direction[d][0]
		nextY := y + g.Direction[d][1]
		if g.inArea(nextX, nextY) && !g.Visited[nextX][nextY] && g.Grid[nextX][nextY] == 1 {
			result += g.dfs(nextX, nextY)
		}
	}
	return result
}
