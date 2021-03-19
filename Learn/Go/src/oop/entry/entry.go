package main

import (
	"fmt"
	"oop"
)

/**
	面向对象

	Go语言只有封装，没有继承和多态
 */
type myTreeNode struct {
	node *oop.TreeNode
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil{
		return
	}

	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root oop.TreeNode

	root = oop.TreeNode{Value: 3}
	root.Left = &oop.TreeNode{}
	root.Right = &oop.TreeNode{5, nil, nil}
	root.Left.Right = oop.CreateNode(2)
	root.Right.Left = new(oop.TreeNode)

	root.Print()

	// 可以直接用值
	root.Right.Left.SetValue(4)
	fmt.Println(root.Right.Left.Value)

	// 可以用地址
	pointerRoot := &root
	pointerRoot.SetValue(200)
	pointerRoot.Print()

	// nil指针也可以调用方法
	var newRoot *oop.TreeNode
	newRoot.SetValue(99)

	root.Traverse()
	nodeCount := 0
	root.TraverseFunc(func(node *oop.TreeNode) {
		nodeCount ++
	})
	fmt.Println("Node Count", nodeCount)

	myRoot := myTreeNode{&root}
	myRoot.postOrder()

	// 切片
	nodes := []oop.TreeNode {
		{Value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)
}
