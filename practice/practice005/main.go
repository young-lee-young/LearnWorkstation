package main

import "fmt"

/**
LRU Cache LeetCode No146
*/

type Node struct {
	Pre   *Node
	Next  *Node
	key   interface{}
	Value interface{}
}

type LRUCache struct {
	Capacity int
	Length   int
	Head     *Node
	Tail     *Node
	HashMap  map[interface{}]*Node
}

func (cache *LRUCache) Get(key int) interface{} {
	if node, ok := cache.HashMap[key]; ok {
		cache.MoveToHead(node)
		return cache.HashMap[key].Value
	}
	return nil
}

func (cache *LRUCache) Put(key int, value interface{}) {
	if node, ok := cache.HashMap[key]; ok {
		node.Value = value
		cache.MoveToHead(node)
	} else {
		if cache.Length == cache.Capacity {
			delete(cache.HashMap, cache.Tail.key)
			cache.RemoveTail()
		} else {
			cache.Length++
		}
		nodeNew := &Node{
			Pre:   nil,
			Next:  nil,
			key:   key,
			Value: value,
		}
		cache.HashMap[key] = nodeNew
		cache.InsertToHead(nodeNew)
	}
}

func (cache *LRUCache) MoveToHead(node *Node) {
	switch node {
	case cache.Head:
		return
	case cache.Tail:
		cache.RemoveTail()
	default:
		node.Pre.Next = node.Next
		node.Next.Pre = node.Pre
	}
	cache.InsertToHead(node)
}

func (cache *LRUCache) InsertToHead(node *Node) {
	if cache.Tail == nil {
		cache.Tail = node
	} else {
		node.Next = cache.Head
		cache.Head.Pre = node
	}
	cache.Head = node
}

func (cache *LRUCache) RemoveTail() {
	if cache.Tail != nil {
		cache.Tail = cache.Tail.Pre
	}
}

func main() {
	cache := LRUCache{
		Capacity: 2,
		Length:   0,
		Head:     nil,
		Tail:     nil,
		HashMap:  make(map[interface{}]*Node, 2),
	}
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))
	cache.Put(4, 4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
}
