package trie

/**
	字典树 LeetCode No208
 */

type Node struct {
	IsWord bool
	Next map[uint8]*Node
}

type CustomizeTrie struct {
	Root *Node
	TrieSize int
}

func (customizeTrie *CustomizeTrie) Size() int {
	return customizeTrie.TrieSize
}

func (customizeTrie *CustomizeTrie) Add(word string) {
	current := customizeTrie.Root
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
		customizeTrie.TrieSize ++
	}
}

func (customizeTrie *CustomizeTrie) Contains(word string) bool {
	current := customizeTrie.Root
	for i := 0; i < len(word); i ++ {
		char := word[i]
		if _, ok := current.Next[char]; !ok {
			return false
		}
		current = current.Next[char]
	}
	if !current.IsWord {
		return false
	}
	return true
}

func (customizeTrie *CustomizeTrie) Prefix(prefix string) bool {
	current := customizeTrie.Root
	for i := 0; i < len(prefix); i ++ {
		char := prefix[i]
		if _, ok := current.Next[char]; !ok {
			return false
		}
		current = current.Next[char]
	}
	return true
}
