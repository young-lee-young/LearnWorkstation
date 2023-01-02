package main

import (
	tree2 "practice/base/tree"
	"fmt"
	"strconv"
	"strings"
)

/**
	二叉树序列化和反序列化 LeetCode No297
 */
func main() {
	tree := tree2.Tree{}
	tree.GenerateTree()
	root := tree.Root
	// 序列化
	result := serialize(root)
	fmt.Println(result)
	nodeSlice := strings.Split(result, ",")
	// 反序列化
	newRoot := deserialize(&nodeSlice)
	tree.Root = newRoot
	tree.InorderTraversal()
}

func serialize(node *tree2.Node) string {
	if node == nil {
		return "#"
	}
	return strconv.Itoa(node.Data) + ", " + serialize(node.Left) + ", " + serialize(node.Right)
}

func deserialize(nodeSlice *[]string) *tree2.Node {
	if len(*nodeSlice) == 0 {
		return nil
	}
	node := (*nodeSlice)[0]
	*nodeSlice = (*nodeSlice)[1:]
	node = strings.Replace(node," ", "", -1)
	if node == "#" {
		return nil
	}
	data, _ := strconv.Atoi(node)
	root := &tree2.Node{Data: data}
	root.Left = deserialize(nodeSlice)
	root.Right = deserialize(nodeSlice)
	return root
}
