package oop

// 中序遍历
func (node *TreeNode) Traverse() {
	node.TraverseFunc(func(node *TreeNode) {
		node.Print()
	})
}

func (node *TreeNode) TraverseFunc(f func(*TreeNode)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}
