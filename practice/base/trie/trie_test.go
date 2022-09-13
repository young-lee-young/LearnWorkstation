package trie

import (
	"testing"
	"fmt"
)

func TestTrie(t *testing.T) {
	node := &Node{false, map[uint8]*Node{}}
	customizeTrie := CustomizeTrie{node, 0}

	fmt.Println("---------- add ----------")
	customizeTrie.Add("lee")
	customizeTrie.Add("zhao")

	fmt.Println("---------- contains ----------")
	fmt.Println(customizeTrie.Contains("li"))
	fmt.Println(customizeTrie.Contains("zhao"))

	fmt.Println("---------- prefix ----------")
	fmt.Println(customizeTrie.Prefix("li"))
	fmt.Println(customizeTrie.Prefix("zh"))

	fmt.Println("---------- size ----------")
	fmt.Println(customizeTrie.Size())
}
