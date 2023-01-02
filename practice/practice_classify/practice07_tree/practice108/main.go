package main

import (
	tree2 "practice/base/tree"
	"math/rand"
	"time"
	"fmt"
)

/**
	树最大路径，Leetcode No543
 */
func main() {
	tree := tree2.Tree{}
	rand.Seed(time.Now().Unix())
	i := 0
	for i < 10 {
		num := rand.Intn(100)
		tree.Insert(num)
		i ++
	}
	root := tree.Root
	maxDepth := getDeep(root)
	tree.PreorderTraversal()
	fmt.Println("max depth : -----------", maxDepth)
}

func getDeep(node *tree2.Node) int {
	if node == nil {
		return 0
	}
	leftDepth := getDeep(node.Left)
	rightDepth := getDeep(node.Right)
	var max int
	if leftDepth > rightDepth {
		max = leftDepth
	} else {
		max = rightDepth
	}
	return max + 1
}
