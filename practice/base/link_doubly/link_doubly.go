/**
	双端链表
 */
package link_doubly

import (
	"fmt"
)

type Node struct {
	Data interface{}
	Pre *Node
	Next *Node
}

type LinkDoublyList struct {
	Head *Node
	Tail *Node
}

func (node *Node) GetData() interface{} {
	return node.Data
}

func (node *Node) GetPre() *Node {
	return node.Pre
}

func (node *Node) GetNext() *Node {
	return node.Next
}

/**
	右测插入元素
 */
func (doublyLink *LinkDoublyList) RightPush(data interface{}) {
	head := doublyLink.Head
	tail := doublyLink.Tail
	node := &Node{
		Data:data,
	}
	if head == nil && tail == nil {
		doublyLink.Head = node
		doublyLink.Tail = node
	} else {
		tail := doublyLink.Tail
		tail.Next = node
		node.Pre = tail
		doublyLink.Tail = node
	}
}

/**
	左侧删除元素
 */
func (doublyLink *LinkDoublyList) LeftPop() interface{} {
	head := doublyLink.Head
	tail := doublyLink.Tail
	if head == nil && tail == nil {
		return nil
	}
	next := head.Next
	doublyLink.Head = next
	if next == nil {
		doublyLink.Tail = next
	}
	return head.Data
}


/**
	获取双向链表长度
 */
func (doublyLink *LinkDoublyList) GetLength() int {
	count := 0
	head := doublyLink.Head
	tail := doublyLink.Tail
	if head == nil && tail == nil {
		return count
	}
	count = 1
	for head != tail {
		count += 1
		head = head.Next
	}
	return count
}

/**
	打印链表
 */
func (doublyLink *LinkDoublyList) ShowList() {
	current := doublyLink.Head
	for current != nil {
		fmt.Printf("\t%v", current.Data)
		current = current.Next
	}
	fmt.Println()
}
