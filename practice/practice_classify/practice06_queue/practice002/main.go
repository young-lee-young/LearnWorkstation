package main

import "fmt"

/**
	滑动窗口最大值 LeetCode No239

	标签：单调双端队列
 */
func main() {
	numArray := []int{7, 2, 4}
	windowSize := 2

	deque := make([]int, 0)
	result := make([]int, 0)
	for i := 0; i < len(numArray); i ++ {
		// 将队列中小于当前元素出队
		for len(deque) != 0 && numArray[i] >= numArray[deque[len(deque) - 1]] {
			deque = deque[:len(deque) - 1]
		}
		deque = append(deque, i)
		// 如果当前元素和队首元素差大于窗口，将队首移除
		if i - deque[0] + 1 > windowSize {
			deque = deque[1:]
		}
		// 将队首元素添加入结果列表
		if i >= windowSize - 1 {
			result = append(result, numArray[deque[0]])
		}
	}
	fmt.Println(result)
}
