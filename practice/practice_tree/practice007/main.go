package main

import (
	"math/rand"
	"time"
	tree2 "practice/base/tree"
	"math"
	"fmt"
)

var result bool

/**
	判断是否是平衡树，左右子树高度差小于1 LeetCode No110


 */
func main() {
	rand.Seed(time.Now().Unix())

	tree := &tree2.Tree{}
	i := 0
	for i < 10 {
		num := rand.Intn(100)
		tree.Insert(num)
		i ++
	}

	result = true
	getDeep(tree.Root)
	fmt.Println(result)
}

func getDeep(node *tree2.Node) int {
	if node == nil {
		return 0
	}
	leftDeep := getDeep(node.Left)
	rightDeep := getDeep(node.Right)
	absDeep := int(math.Abs(float64(leftDeep - rightDeep)))
	if absDeep > 1 {
		result = false
	}

	if leftDeep > rightDeep {
		return leftDeep + 1
	} else {
		return rightDeep + 1
	}
}
