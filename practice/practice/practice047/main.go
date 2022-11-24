package main

import "fmt"

/**
	键值映射 LeetCode No677

	标签：字典树
 */
type Node struct {
	Value int
	Next map[uint8]*Node
}

type MapSum struct {
	Root *Node
}

func Constructor() MapSum {
	node := &Node{0, map[uint8]*Node{}}
	m := MapSum{node}
	return m
}

func (m *MapSum) Insert(word string, value int) {
	current := m.Root
	for i := 0; i < len(word); i ++ {
		char := word[i]
		if _, ok := current.Next[char]; !ok {
			newNode := map[uint8]*Node{}
			current.Next[char] = &Node{0, newNode}
		}
		current = current.Next[char]
	}
	current.Value = value
}

func (m *MapSum) Sum(prefix string) int {
	current := m.Root
	for i := 0; i < len(prefix); i ++ {
		char := prefix[i]
		// 没有以prefix为开头的单词
		if _, ok := current.Next[char]; !ok {
			return 0
		}
		current = current.Next[char]
	}
	return recursionSum(current)
}

func recursionSum(node *Node) int {
	// 当前节点为叶子结点，递归结束
	if len(node.Next) == 0 {
		return node.Value
	}
	// 获取当前节点的值
	result := node.Value
	// 遍历当前节点孩子节点，并将所有孩子节点和相加
	for _, child := range node.Next {
		result += recursionSum(child)
	}
	return result
}

func main() {
	trie := Constructor()
	fmt.Println(trie)
}
