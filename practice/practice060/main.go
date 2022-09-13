/**
 * @Time:    2021/3/30 22:44 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content: 倒水问题
 */
package main

import (
	queue2 "practice/base/queue"
	"fmt"
	"practice/utils/compare"
	"practice/utils/reverse"
)

/**
	两个水桶，一个5L，一个3L，利用两个水桶倒出4L

	标签：广度优先遍历
 */
func main() {
	end, pre := solution()
	result := path(end, pre)
	fmt.Println(result)
}

/*
	用一个两位数字表示水桶里的水量
 */
func solution() (int, []int) {
	end := -1

	// 广度优先遍历
	queue := queue2.Queue{}
	visited := make(map[int]bool)
	pre := make([]int, 54)

	// 两个桶均为空入队
	queue.Enqueue(0)
	visited[0] = true

	for !queue.IsEmpty() {
		current := queue.Dequeue().(int)

		// max a = 5；max b = 3
		a := current / 10
		b := current % 10

		// 下一步可以获得的两个桶水量
		nextSlice := make([]int, 0)
		// a灌满水，b不变
		nextSlice = append(nextSlice, 5 * 10 + b)
		// a不变，b灌满水
		nextSlice = append(nextSlice, a * 10 + 3)
		// a倒掉，b不变
		nextSlice = append(nextSlice, 0 * 10 + b)
		// a不变，b倒掉
		nextSlice = append(nextSlice, a * 10 + 0)

		// a中水倒入b中
		// a小，表示a中的水全部倒入了b；(3 - b)小，表示将b倒满后a还剩水
		x := compare.Min(a, 3 - b)
		nextSlice = append(nextSlice, (a - x) * 10 + (b + x))

		// b中水倒入a中
		// (5 - a)小，表示将a倒满后b还剩水，b小，表示b中的水全部倒入了a
		y := compare.Min(5 - a, b)
		nextSlice = append(nextSlice, (a + y) * 10 + (b - y))

		for _, next := range nextSlice {
			// 没有遍历过
			if _, ok := visited[next]; !ok {
				queue.Enqueue(next)
				visited[next] = true
				pre[next] = current

				// 已经倒出4L水
				if next / 10 == 4 {
					end = next
					return end, pre
				}
			}
		}
	}
	return end, pre
}

func path(end int, pre []int) []int {
	fmt.Println(pre)
	result := make([]int, 0)
	if end == -1 {
		return result
	}

	current := end
	for current != 0 {
		result = append(result, current)
		current = pre[current]
	}
	result = append(result, 0)
	// 反转
	result = reverse.ReverseSlice(result)
	return result
}
