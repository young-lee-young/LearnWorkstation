/**
 * @Time:    2021/3/26 22:45 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content: 二进制矩阵中最短路径
 */
package main

import (
	queue2 "practice/base/queue"
	"fmt"
)

/**
	二进制矩阵中的最短路径 LeetCode 1091

	标签：广度优先搜索
 */
func main() {
	grid := [][]int{{0, 0, 0},{1, 1, 0},{1, 1, 0}}
	result := solution(grid)
	fmt.Println(result)
}

func solution(grid [][]int) int {
	graph := &Graph{
		Row:    len(grid),
		Column: len(grid[0]),
		Grid:   grid,
	}
	visited := make([][]bool, graph.Row)
	for i := 0; i < graph.Row; i ++ {
		visited[i] = make([]bool, graph.Column)
	}
	graph.Visited = visited
	distance := make([][]int, graph.Row)
	for j := 0; j < graph.Row; j ++ {
		distance[j] = make([]int, graph.Column)
	}
	graph.Distance = distance
	direction := [][]int{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}}
	graph.Direction = direction

	// 起始点阻塞
	if graph.Grid[0][0] == 1 {
		return -1
	}
	// 只有一个点
	if graph.Row == 1 && graph.Column == 1 {
		return 1
	}

	// 广度优先搜索
	queue := queue2.Queue{}
	// 添加起始点
	queue.Enqueue(0)
	graph.Visited[0][0] = true
	graph.Distance[0][0] = 1
	for !queue.IsEmpty() {
		current := queue.Dequeue().(int)
		// 二维转一维
		currentX := current / graph.Column
		currentY := current % graph.Column
		// 检查八个方向上的值
		for d := 0; d < 8; d ++ {
			nextX := currentX + graph.Direction[d][0]
			nextY := currentY + graph.Direction[d][1]
			// 如果坐标合法、没有被访问过并且非阻塞
			if graph.inArea(nextX, nextY) && !graph.Visited[nextX][nextY] && graph.Grid[nextX][nextY] == 0 {
				queue.Enqueue(nextX * graph.Column + nextY)
				graph.Visited[nextX][nextY] = true
				graph.Distance[nextX][nextY] = graph.Distance[currentX][currentY] + 1

				// 已经到终点了
				if nextX == graph.Row - 1 && nextY == graph.Column - 1 {
					return graph.Distance[nextX][nextY]
				}
			}
		}
	}
	return -1
}

type Graph struct {
	Row       int
	Column    int
	Grid      [][]int
	Visited   [][]bool
	Distance  [][]int
	Direction [][]int
}

func (g *Graph) inArea(x int, y int) bool {
	return x >= 0 && x < g.Row && y >= 0 && y < g.Column
}
