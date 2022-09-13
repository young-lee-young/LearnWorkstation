/**
	链表
 */
package link

import (
	"fmt"
)

type Node struct {
	Data interface{}
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (node *Node) GetData() interface{} {
	return node.Data
}

func (node *Node) GetNext() *Node {
	return node.Next
}

/**
	头部插入数据
 */
func (link *LinkedList) Add(data interface{}) {
	node := &Node{Data:data}
	node.Next = link.Head
	link.Head = node
}

/**
	任意位置插入数据
 */
func (link *LinkedList) Insert(index int, data interface{}) {
	if index < 0 {
		link.Add(data)
		return
	}
	if index >= link.GetLength() {
		link.Append(data)
		return
	}
	current := link.Head
	count := 0
	for count < (index - 1) {
		current = current.Next
		count ++
	}
	node := &Node{Data: data}
	node.Next = current.Next
	current.Next = node
}

/**
	尾部插入数据
 */
func (link *LinkedList) Append(data interface{}) {
	node := &Node{Data:data}
	if link.IsEmpty() {
		link.Head = node
		return
	}
	current := link.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = node
}

/**
	删除头部元素
 */
func (link *LinkedList) DeleteHead() {
	if link.IsEmpty() {
		return
	}
	link.Head = link.Head.Next
}

/**
	删除任意位置元素
 */
func (link *LinkedList) Delete(index int) {
	if index < 0 {
		link.DeleteHead()
		return
	}
	if index >= link.GetLength() {
		link.DeleteTail()
		return
	}
	current := link.Head
	count := 0
	for count < (index - 1) {
		current = current.Next
		count ++
	}
	current.Next = current.Next.Next
}

/**
	删除尾部元素
 */
func (link *LinkedList) DeleteTail() {
	if link.IsEmpty() {
		return
	}
	if link.GetLength() == 1 {
		link.Head = nil
		return
	}
	current := link.Head
	for current.Next.Next != nil {
		current = current.Next
	}
	current.Next = nil
}

/**
	是否包含某元素
 */
func (link *LinkedList) Contains(data interface{}) bool {
	current := link.Head
	for current != nil {
		if data == current.GetData() {
			return true
		}
		current = current.Next
	}
	return false
}

/**
	是否为空
 */
func (link *LinkedList) IsEmpty() bool {
	if link.Head == nil {
		return true
	}
	return false
}

/**
	获取长度
 */
func (link *LinkedList) GetLength() int {
	current := link.Head
	count := 0
	for current != nil {
		count ++
		current = current.Next
	}
	return count
}

/**
	打印链表
 */
func (link *LinkedList) ShowList() {
	current := link.Head
	for current != nil {
		fmt.Printf("\t%v", current.Data)
		current = current.Next
	}
	fmt.Println()
}
