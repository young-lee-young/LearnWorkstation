/**
LeetCode No116 填充每个节点的下一个右侧节点指针

解题思路：二叉树层序遍历
*/
package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func main() {

}

func solution(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := Queue{}
	queue.Enqueue(root)

	for !queue.IsEmpty() {
		queueLen := queue.GetLen()

		var last *Node
		for i := 0; i < queueLen; i ++ {
			node := queue.Dequeue().(*Node)

			if last != nil {
				last.Next = node
			}

			last = node

			if node.Left != nil {
				queue.Enqueue(node.Left)
			}

			if node.Right != nil {
				queue.Enqueue(node.Right)
			}
		}

		last.Next = nil
	}

	return root
}

type Queue struct {
	data []interface{}
}

func (q *Queue) Enqueue(v interface{}) {
	q.data = append(q.data, v)
}

func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	v := q.data[0]
	q.data = q.data[1:]
	return v
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue) GetLen() int {
	return len(q.data)
}
