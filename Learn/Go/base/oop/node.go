package oop

import "fmt"

type TreeNode struct {
	Value int
	Left, Right *TreeNode
}

// 结构方法，有接受者node，相当于其他语言的this或self
// 值接收者
func (node TreeNode) Print() {
	fmt.Println(node.Value)
}

// 指针接收者
func (node *TreeNode) SetValue(value int) {
	if node == nil {
		fmt.Println("setting value to nil node")
		return
	}
	node.Value = value
}

// 工厂函数
func CreateNode(value int) *TreeNode {
	// 返回局部变量的地址，C++是错误的
	/**
		C++局部变量分配在栈上，函数退出后，变量销毁
	 */
	return &TreeNode{Value: value}
}
