package main

import "fmt"

/**
	添加于搜索单词 LeetCode No211

	标签：字典树
 */
type Node struct {
	IsWord bool
	Next map[uint8]*Node
}

type Trie struct {
	Root *Node
}

func Constructor() Trie {
	node := &Node{false, map[uint8]*Node{}}
	wordDict := Trie{node}
	return wordDict
}

func (trie *Trie) AddWord(word string) {
	current := trie.Root
	for i := 0; i < len(word); i ++ {
		char := word[i]
		if _, ok := current.Next[char]; !ok {
			newNode := map[uint8]*Node{}
			current.Next[char] = &Node{false, newNode}
		}
		current = current.Next[char]
	}
	if !current.IsWord {
		current.IsWord = true
	}
}

func (trie *Trie) Search(word string) bool {
	return match(trie.Root, word, 0)
}

func match(node *Node, word string, index int) bool {
	if index == len(word) {
		return node.IsWord
	}

	// 字符串当前index的字符
	char := word[index]
	// .是46
	if char != 46 {
		// 判断当前节点孩子节点是否存在需要查询的字符，不存在返回false
		if _, ok := node.Next[char]; !ok {
			return false
		}
		// 如果存在则在下个节点中寻找下个字符
		return match(node.Next[char], word, index + 1)
	// 当前字符是.
	} else {
		// 循环当前节点所有孩子节点
		for _, value := range node.Next {
			// 递归孩子节点是否有存在下个字符
			if match(value, word, index + 1) {
				return true
			}
		}
		return false
	}
}

/**

 */
func main() {
	trie := Constructor()
	fmt.Println(trie)
}
