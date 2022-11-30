/**
 * @Time:    2022/11/29 14:12 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	LeetCode No059 螺旋矩阵

 */
package main

import "fmt"

func main() {
	n := 4

	ret := solution(n)

	fmt.Println("ret:", ret)
}

func solution(n int) [][]int {
	// ⚠️注意：二维数组初始化
	ret := make([][]int, n)
	for i := 0; i < n; i ++ {
		ret[i] = make([]int, n)
	}

	num := 1

	turns := n / 2 // 圈数
	for turn := 0; turn < turns; turn ++ {
		startx := turn
		starty := turn

		endx := n - (turn + 1)
		endy := n - (turn + 1)

		// 上横排
		for j := startx; j < endx; j ++ {
			ret[startx][j] = num
			num ++
		}

		// 右竖排
		for i := starty; i < endy; i ++ {
			ret[i][endy] = num
			num ++
		}

		// 下横排
		for j := endx; j > startx; j -- {
			ret[endx][j] = num
			num ++
		}

		// 左竖排
		for i := endy; i > starty; i -- {
			ret[i][starty] = num
			num ++
		}
	}

	// 如果是单列，补充最中间的值
	if n % 2 == 1 {
		ret[turns][turns] = num
	}

	return ret
}
