package main

import (
	"practice/base/ntree"
	"fmt"
)

/**
	N叉树前序遍历 LeetCode No589

	解题思路：
 */
func main() {
	tree := ntree.NTree{}
	tree.Generate(5, 3)
	root := tree.Root
	tree.ShowTree(3)

	result := solution(root)
	fmt.Println(result)
}

func solution(node *ntree.Node) []int {
	if node == nil {
		return []int{}
	}
	var result []int
	result = append(result, node.Data)
	for _, subNode := range node.Children {
		result = append(result, solution(subNode)...)
	}
	return result
}
