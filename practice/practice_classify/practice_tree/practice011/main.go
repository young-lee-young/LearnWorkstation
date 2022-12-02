package main

import "fmt"

/**
N叉树层序遍历 LeetCode No429

标签：使用队列
*/

type Node struct {
	Val      int
	Children []*Node
}

func main() {
	root := &Node{
		Val: 5,
	}

	result := make([][]int, 0)
	if root == nil {
		return
	}
	// 使用slice模仿队列
	queue := make([]*Node, 0)
	queue = append(queue, root)

	for len(queue) != 0 {
		count := len(queue)
		numSlice := make([]int, 0)
		var newQueue []*Node
		for i := 0; i < count; i++ {
			for j := 0; j < len(queue[i].Children); j++ {
				newQueue = append(newQueue, queue[i].Children[j])
			}
			numSlice = append(numSlice, queue[i].Val)
		}
		result = append(result, numSlice)
		queue = newQueue
	}
	fmt.Println(result)
}
