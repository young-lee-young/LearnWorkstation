package main

import (
	"fmt"
	stack2 "practice/base/stack"
)

/**
	柱状图中最大矩形 LeetCode No84
 */
func main() {
	violence()
	useStack()
}

/**
	暴力法，从左到右一次求最大面积
 */
func violence() {
	heights := []int{2,1,5,6,2,3}

	max := 0
	for i := 0; i < len(heights); i ++ {
		j := i
		minHeight := heights[i]
		for j < len(heights) {
			if heights[j] < minHeight {
				minHeight = heights[j]
			}
			area := (j - i + 1) * minHeight
			if area > max {
				max = area
			}
			j ++
		}
	}
	fmt.Println(max)
}

/**
	使用递增栈
 */
func useStack() {
	//heights := []int{2, 1, 2}
	heights := []int{2,1,5,6,2,3}
	stack := stack2.Stack{}

	max := 0
	for i := 0; i < len(heights); i ++ {
		if stack.IsEmpty() || heights[i] > heights[stack.Peek().(int)] {
			stack.Push(i)
			continue
		}

		for !stack.IsEmpty() {
			top := stack.Pop().(int)
			// 一直到当前值比栈顶元素大
			if heights[i] > heights[top] {
				stack.Push(top)
				break
			}
			// 如果栈为空，当前值就是宽度
			var width int
			if stack.IsEmpty() {
				width = i
			} else {
				width = i - top
			}
			area := width * heights[top]
			if area > max {
				//fmt.Println(i, top, width, heights[top], area)
				max = area
			}
		}
		stack.Push(i)
	}

	for !stack.IsEmpty() {
		top := stack.Pop().(int)
		var width int
		if stack.IsEmpty() {
			width = len(heights)
		} else {
			width = len(heights) - top
		}
		area := width * heights[top]
		if area > max {
			//fmt.Println(len(heights), top, width, heights[top], area)
			max = area
		}
	}
	fmt.Println(max)
}
