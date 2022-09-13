/**
	二叉树
 */
package tree

import (
	"fmt"
)

type Node struct {
	Left  *Node
	Right *Node
	Data  int
}

/**
	二叉搜索树
 */
type Tree struct {
	Root *Node
}

/**
	插入元素
 */
func (tree *Tree) Insert(data int) {
	newNode := &Node{
		Data: data,
	}
	if tree.Root == nil {
		tree.Root = newNode
	} else {
		addNode(tree.Root, newNode)
	}
}

func addNode(node *Node, newNode *Node) {
	if newNode.Data < node.Data {
		if node.Left == nil {
			node.Left = newNode
		} else {
			addNode(node.Left, newNode)
		}
	} else {
		if node.Right == nil {
			node.Right = newNode
		} else {
			addNode(node.Right, newNode)
		}
	}
}

/**
	前序遍历
 */
func (tree *Tree) PreorderTraversal() {
	root := tree.Root
	if root == nil {
		return
	}
	preorder(root)
}

func preorder(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.Data)
	preorder(node.Left)
	preorder(node.Right)
}

/**
	中序遍历
 */
func (tree *Tree) InorderTraversal() {
	root := tree.Root
	if root == nil {
		return
	}
	inorder(root)
}

func inorder(node *Node) {
	if node == nil {
		return
	}
	inorder(node.Left)
	fmt.Println(node.Data)
	inorder(node.Right)
}

/**
	后序遍历
 */
func (tree *Tree) PostorderTraversal() {
	root := tree.Root
	if root == nil {
		return
	}
	postorder(root)
}

func postorder(node *Node) {
	if node == nil {
		return
	}
	postorder(node.Left)
	postorder(node.Right)
	fmt.Println(node.Data)
}

/**
	生成一棵固定的树，用于测试
						       20
                 10                         30
           5             15           25          35
       2       7     12      18   23      28   31     38
 */
func (tree *Tree) GenerateTree() {
	numArray := [...]int{20, 10, 30, 5, 15, 25, 35, 2, 7, 12, 18, 23, 28, 31, 38}
	for _, value := range numArray {
		tree.Insert(value)
	}
}

func (tree *Tree) GenerateTree2() {
	//numArray := [...]int{10, 5, 15, 2, 7, 12, 18, 20}
	numArray := [...]int{3, 9, 20, 0, 0, 15, 7}
	for _, value := range numArray {
		tree.Insert(value)
	}
}
