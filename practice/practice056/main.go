/**
 * @Time:    2021/3/23 22:19 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content: 岛屿的最大面积
 */
package main

import (
	"math"
	"fmt"
)

/**
	岛屿最大面积 LeetCode No695

	标签：图深度优先搜索
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
	// 判断grid的合法性
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
	matrix := graph.constructGraph()
	graph.Matrix = matrix

	var result int
	visited := make([]bool, len(graph.Matrix))
	graph.Visited = visited
	for v := 0; v < len(graph.Matrix); v ++ {
		x := v / graph.Column
		y := v % graph.Column
		if !graph.Visited[v] && graph.Grid[x][y] == 1 {
			result = int(math.Max(float64(result), float64(graph.dfs(v))))
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
	Matrix    [][]int
	Visited   []bool
}

// 构建图
func (g *Graph) constructGraph() [][]int {
	// 初始化的临接表，一共有 行 * 列 个顶点
	matrix := make([][]int, g.Row * g.Column)
	for i := 0; i < len(matrix); i ++ {
		matrix[i] = make([]int, 0)
	}

	// 遍历所有顶点
	for v := 0; v < len(matrix); v ++ {
		// 和列取商得到横坐标，和列取余的到纵坐标
		x := v / g.Column
		y := v % g.Column
		// 如果是陆地，需要遍历四周
		if g.Grid[x][y] == 1 {
			// 遍历四周
			for d := 0; d < 4; d ++ {
				nextX := x + g.Direction[d][0]
				nextY := y + g.Direction[d][1]
				// 如果相邻的顶点是陆地
				if g.inArea(nextX, nextY) && g.Grid[nextX][nextY] == 1 {
					next := nextX * g.Column + nextY
					// 构建临接表
					matrix[v] = append(matrix[v], next)
					matrix[next] = append(matrix[next], v)
				}
			}
		}
	}
	return matrix
}

// 判断x、y是否在grid中
func (g *Graph) inArea(x int, y int) bool {
	return x >= 0 && x < g.Row && y >= 0 && y < g.Column
}

func (g *Graph) dfs(v int) int {
	g.Visited[v] = true
	result := 1
	for _, w := range g.Matrix[v] {
		if !g.Visited[w] {
			result += g.dfs(w)
		}
	}
	return result
}
