/**
 * @Time:    2022/11/29 15:52
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

	LeetCode No707 设计链表
*/
package main

import "fmt"

func main() {
	ll := Constructor()

	ll.AddAtIndex(1, 0)

	fmt.Println(ll.Get(0))

	//ll.AddAtHead(7)
	//ll.AddAtHead(2)
	//ll.AddAtHead(1)
	//
	//ll.AddAtIndex(3, 0)
	//
	//ll.DeleteAtIndex(2)
	//
	//ll.AddAtHead(6)
	//
	//ll.AddAtTail(4)
	//
	//fmt.Println(ll.Get(4))
	//
	//ll.AddAtHead(4)
	//
	//ll.AddAtIndex(5, 0)
	//
	//ll.AddAtHead(6)

}

type Node struct {
	val  int
	next *Node
}

type MyLinkedList struct {
	Root *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		Root: nil,
	}
}

func (this *MyLinkedList) Get(index int) int {
	node := this.Root

	for i := 0; i < index; i++ {
		if node == nil {
			return -1
		}
		node = node.next
	}

	if node == nil {
		return -1
	}

	return node.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &Node{
		val:  val,
		next: this.Root,
	}
	this.Root = newNode
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := this.Root

	if node == nil {
		newNode := &Node{
			val:  val,
			next: nil,
		}

		this.Root = newNode
		return
	}

	for node.next != nil {
		node = node.next
	}

	newNode := &Node{
		val:  val,
		next: nil,
	}

	node.next = newNode
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	node := this.Root

	if node == nil {
		if index > this.Length() {
			return
		}

		newNode := &Node{
			val:  val,
			next: nil,
		}
		this.Root = newNode
		return
	}

	if index <= 0 {
		this.AddAtHead(val)
	}

	for i := 0; i < index-1; i++ {
		node = node.next
	}

	if node == nil {
		return
	}

	next := node.next

	newNode := &Node{
		val:  val,
		next: next,
	}

	node.next = newNode
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	node := this.Root

	if index == 0 {
		this.Root = node.next
		return
	}

	for i := 0; i < index-1; i++ {
		if node == nil {
			return
		}
		node = node.next
	}

	if node == nil || node.next == nil {
		return
	}

	node.next = node.next.next
}

func (this *MyLinkedList) Length() int {
	i := 0

	node := this.Root

	for node != nil {
		i++
		node = node.next
	}

	return i
}

func (this *MyLinkedList) Print() {
	node := this.Root
	i := 0
	for node != nil {
		fmt.Printf("index: %d, val: %d\n", i, node.val)
		node = node.next
		i++
	}
}
