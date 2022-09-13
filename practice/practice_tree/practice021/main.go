package main

import (
	tree2 "practice/base/tree"
	"fmt"
)

/**
	二叉树最左后一层最左边值 LeetCode 513
 */
func main() {
	tree := tree2.Tree{}
	tree.GenerateTree()
	root := tree.Root

	currentSlice := []*tree2.Node{root}
	for len(currentSlice) != 0 {
		nextSlice := []*tree2.Node{}
		for i := 0; i < len(currentSlice); i ++ {
			if currentSlice[i].Left != nil {
				nextSlice = append(nextSlice, currentSlice[i].Left)
			}
			if currentSlice[i].Right != nil {
				nextSlice = append(nextSlice, currentSlice[i].Right)
			}
		}
		if len(nextSlice) == 0 {
			break
		}
		currentSlice = nextSlice
		fmt.Println(len(currentSlice))
	}
	fmt.Println(currentSlice[0].Data)
}
