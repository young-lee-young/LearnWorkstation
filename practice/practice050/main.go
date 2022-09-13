/**
 * @Time:    2021/3/8 19:26 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content: 二叉树锯齿型层次遍历
 */
package main

import (
	"fmt"
	tree2 "practice/base/tree"
	queue2 "practice/base/queue"
)

/**
	二叉树锯齿型层次遍历 LeetCode No103

	解题思路：使用队列层次遍历，使用反转标志对偶数行反转
 */
func main() {
	tree := tree2.Tree{}
	tree.GenerateTree()
	root := tree.Root
	result := solution(root)
	fmt.Println(result)
}

func solution(root *tree2.Node) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := queue2.Queue{}
	queue.Enqueue(root)

	direction := "right"
	for !queue.IsEmpty() {
		temp := make([]int, 0)
		queueLen := queue.GetLength()
		for i := 0; i < queueLen; i ++ {
			node := queue.Dequeue().(*tree2.Node)
			temp = append(temp, node.Data)
			if node.Left != nil {
				queue.Enqueue(node.Left)
			}
			if node.Right != nil {
				queue.Enqueue(node.Right)
			}
		}

		if direction == "left" {
			temp = reverse(temp)
			direction = "right"
		} else {
			direction = "left"
		}
		result = append(result, temp)
	}
	return result
}

func reverse(nums []int) []int {
	for i, j := 0, len(nums) - 1; i < j; i, j = i + 1, j - 1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}
