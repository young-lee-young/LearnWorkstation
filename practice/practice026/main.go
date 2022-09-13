package main

import "fmt"

/**
	模拟行走机器人 LeetCode No874

	解题思路：
 */
func main() {
	commands := []int{7, -2, -2, 7, 5}

	obstacles := [][]int{
		{-3, 2},
		{-2, 1},
		{0, 1},
		{-2, 4},
		{-1, 0},
		{-2, -3},
		{0, -3},
		{4, 4},
		{-3, 3},
		{2, 2},
	}
	result := solution(commands, obstacles)
	fmt.Println(result)
}

func solution(commands []int, obstacles [][]int) int {
	// 欧式距离最大值
	ret := 0
	x := 0
	y := 0

	// 方向，初始值为北
	direction := "n"
	directionMap := map[string][]int{
		"n": {0, 1},
		"s": {0, -1},
		"e": {1, 0},
		"w": {-1, 0},
	}

	for _, command := range commands {
		xNext := 0
		yNext := 0
		// 前进
		if command >= 0 {
			// 开始往前走command步
			for i := 0; i < command; i ++ {
				// 下一步坐标
				xNext = x + directionMap[direction][0]
				yNext = y + directionMap[direction][1]
				// 判断下一步知否是障碍物
				existObstacles := false
				for j := 0; j < len(obstacles); j ++ {
					if xNext == obstacles[j][0] && yNext == obstacles[j][1] {
						existObstacles = true
						break
					}
				}
				// 如果存在障碍物，直接执行下一个命令
				if existObstacles {
					break
				}
				x = xNext
				y = yNext
				fmt.Println(x, y)
				if x * x + y * y > ret {
					ret = x * x + y * y
				}
			}
		} else {
			// 改变方向
			if command == -1 {
				// 向右转
				switch direction {
				case "n":
					direction = "e"
				case "e":
					direction = "s"
				case "s":
					direction = "w"
				case "w":
					direction = "n"
				}
			} else {
				// 向左转
				switch direction {
				case "n":
					direction = "w"
				case "w":
					direction = "s"
				case "s":
					direction = "e"
				case "e":
					direction = "n"
				}
			}
		}
	}
	return ret
}
