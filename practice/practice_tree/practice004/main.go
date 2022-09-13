package main

import (
	tree2 "practice/base/tree"
	"fmt"
)

/**
	树的高度 LeetCode No104

	标签：递归
 */
func main() {
	//tree := tree2.Tree{}
	//rand.Seed(time.Now().Unix())
	//i := 0
	//for i < 10 {
	//	num := rand.Intn(100)
	//	tree.Insert(num)
	//	i ++
	//}
	//root := tree.Root

	tree := tree2.Tree{}
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(4)
	root := tree.Root

	var deep int
	if root == nil {
		deep = 1
	} else {
		deep = maxDeep(tree.Root)
	}
	fmt.Println(deep)
}

func maxDeep(node *tree2.Node) int {
	if node == nil {
		return 0
	}
	leftDeep := maxDeep(node.Left)
	rightDeep := maxDeep(node.Right)
	if leftDeep > rightDeep {
		return leftDeep + 1
	} else {
		return rightDeep + 1
	}
}
