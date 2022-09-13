package main

import (
	tree2 "practice/base/tree"
	"fmt"
)

/**
	路径长度等于给定值的路径总数 LeetCode No437
 */
func main() {
	tree := tree2.Tree{}
	tree.GenerateTree()
	root := tree.Root
	sum := 63

	result := pathSum(root, sum)
	fmt.Println(result)
}

func pathSum(node *tree2.Node, sum int) int {
	if node == nil {
		return 0
	}
	// 以当前节点为根节点进行处理
	result := pathSumWithRoot(node, sum) + pathSum(node.Left, sum) + pathSum(node.Right, sum)
	return result
}

//
func pathSumWithRoot(node *tree2.Node, sum int) int {
	if node == nil {
		return 0
	}
	result := 0
	if node.Data == sum {
		result ++
	}
	result += pathSumWithRoot(node.Left, sum - node.Data) + pathSumWithRoot(node.Right, sum - node.Data)
	return result
}
