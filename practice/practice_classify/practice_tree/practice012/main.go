package main

import (
	"fmt"
	"math"
	tree2 "practice/base/tree"
)

/**
二叉树每层中的最大值 LeetCode No515

解题思路：使用队列层序遍历
*/
func main() {
	tree := tree2.Tree{}
	tree.GenerateTree()
	root := tree.Root
	result := solution(root)
	fmt.Println(result)
}

func solution(root *tree2.Node) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	queue := Queue{}
	queue.Enqueue(root)

	for !queue.IsEmpty() {
		max := math.MinInt64
		queueLen := queue.GetLen()

		for i := 0; i < queueLen; i++ {
			node := queue.Dequeue().(*tree2.Node)
			if node.Data > max {
				max = node.Data
			}

			if node.Left != nil {
				queue.Enqueue(node.Left)
			}

			if node.Right != nil {
				queue.Enqueue(node.Right)
			}
		}

		result = append(result, max)
	}

	return result
}

type Queue struct {
	data []interface{}
}

func (q *Queue) Enqueue(v interface{}) {
	q.data = append(q.data, v)
}

func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	v := q.data[0]
	q.data = q.data[1:]
	return v
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue) GetLen() int {
	return len(q.data)
}
