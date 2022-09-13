/**
 * @Time:    2021/4/14 20:57 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content: 不同路径
 */
package main

import "fmt"

/**
	不同路径 LeetCode 980

	标签：哈密尔顿路径
 */
type Graph struct {
	Row       int
	Column    int
	Grid      [][]int
	Direction [][]int
	Memory    [][]int
	Start     int
	End       int
}

func main() {
	grid := [][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 2, -1}}
	result := solution(grid)
	fmt.Println(result)
}

func solution(grid [][]int) int {
	graph := &Graph{
		Row:    len(grid),
		Column: len(grid[0]),
		Grid:   grid,
	}

	memoryLen := 1 << uint(graph.Row*graph.Column)
	memory := make([][]int, memoryLen)
	for i := 0; i < memoryLen; i ++ {
		memory[i] = make([]int, graph.Row*graph.Column)
	}
	for i := 0; i < memoryLen; i ++ {
		for j := 0; j < graph.Row * graph.Column; j ++ {
			memory[i][j] = -1
		}
	}
	graph.Memory = memory

	// 四联通位置
	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	graph.Direction = directions

	// 格子总数
	left := graph.Row * graph.Column
	// 将所有格子进行预处理
	for i := 0; i < graph.Row; i ++ {
		for j := 0; j < graph.Column; j ++ {
			// 处理起始点
			if graph.Grid[i][j] == 1 {
				graph.Start = i*graph.Column + j
				graph.Grid[i][j] = 0
				// 处理终止点
			} else if graph.Grid[i][j] == 2 {
				graph.End = i*graph.Column + j
				graph.Grid[i][j] = 0
				// 处理障碍点
			} else if graph.Grid[i][j] == -1 {
				left --
			}
		}
	}

	visited := 0
	return graph.dfs(visited, graph.Start, left)
}

func (g *Graph) dfs(visited int, v int, left int) int {
	// 如果值不为-1，说明当前的顶点路径已经计算过了，直接返回就好
	if g.Memory[visited][v] != -1 {
		return g.Memory[visited][v]
	}

	x := v / g.Column
	y := v % g.Column
	// 将当前位上的值设置位1，表示已经访问过
	visited += 1 << uint(v)
	left --

	// left为0表名所有的点都已经走过，并且当前v是终点
	if left == 0 && v == g.End {
		// 由于需要找到所有的路径，所以需要进行进行回溯
		visited -= 1 << uint(v)
		// 将memoery赋值为1
		g.Memory[visited][v] = 1
		return 1
	}

	result := 0
	// 此当前格子向四周遍历
	for d := 0; d < 4; d ++ {
		nextX := x + g.Direction[d][0]
		nextY := y + g.Direction[d][1]
		// 格子坐标合法、是通路并且没有访问过
		next := nextX * g.Column + nextY
		if g.inArea(nextX, nextY) && g.Grid[nextX][nextY] == 0 && (visited & (1 << uint(next)) == 0){
			result += g.dfs(visited, next, left)
		}
	}

	visited -= 1 << uint(v)
	g.Memory[visited][v] = result
	return result
}

func (g *Graph) inArea(x int, y int) bool {
	return x >= 0 && x < g.Row && y >= 0 && y < g.Column
}
