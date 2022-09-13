/**
 * @Time:    2021/3/26 23:25 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content: 打开转盘锁
 */
package main

import (
	queue2 "practice/base/queue"
	"fmt"
)

/**
	打开转盘锁 LeetCode No752
 */
func main() {
	deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	target := "0202"
	result := solution(deadends, target)
	fmt.Println(result)
}

func solution(deadends []string, target string) int {
	// 初始化map，用来存放deadend的数据，后续查找为O(1)时间复杂度
	deadendMap := make(map[string]string)
	for _, deadend := range deadends {
		deadendMap[deadend] = deadend
	}

	// 判读起始字符串是否合法
	if _, ok := deadendMap["0000"]; ok {
		return -1
	}
	if _, ok := deadendMap[target]; ok {
		return -1
	}

	// 起始点就是目标值
	if target == "0000" {
		return 0
	}

	// 广度优先遍历
	queue := queue2.Queue{}
	// 用于存放是否访问过，同时存储到该点所需步数
	visited := make(map[string]int)
	queue.Enqueue("0000")
	visited["0000"] = 0
	for !queue.IsEmpty() {
		current := queue.Dequeue().(string)

		// 获取当前节点相邻节点
		nextSlice := make([]string, 0)
		currentCharSlice := []byte(current)
		for i := 0; i < 4; i ++ {
			o := currentCharSlice[i]
			// 向前一位
			currentCharSlice[i] = byte((currentCharSlice[i] - '0' + 1) % 10) + byte(48)
			nextSlice = append(nextSlice, string(currentCharSlice))
			currentCharSlice[i] = o
			// 向后一位
			currentCharSlice[i] = byte((currentCharSlice[i] - '0' + 9) % 10) + byte(48)
			nextSlice = append(nextSlice, string(currentCharSlice))
			currentCharSlice[i] = o
		}

		// 循环当前节点相邻节点
		for _, next := range nextSlice {
			// 判断节点是否是死亡数字并且是否被访问过
			if _, ok := deadendMap[next]; !ok {
				if _, ok := visited[next]; !ok {
					queue.Enqueue(next)
					// 步数等于当前节点步数 + 1
					visited[next] = visited[current] + 1
					if next == target {
						return visited[next]
					}
				}
			}
		}
	}
	return -1
}
