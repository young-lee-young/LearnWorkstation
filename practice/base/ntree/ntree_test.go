package ntree

import "testing"

func TestNTree(t *testing.T) {
	tree := NTree{}
	tree.Generate(5, 3)
	tree.ShowTree(3)
}

