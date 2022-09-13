package tree

import (
	"testing"
	"time"
	"math/rand"
)

func TestTree(t *testing.T) {
	rand.Seed(time.Now().Unix())
	i := 0
	tree := Tree{}
	for i < 10 {
		num := rand.Intn(100)
		tree.Insert(num)
		i ++
	}
	tree.PreorderTraversal()
	tree.InorderTraversal()
	tree.PostorderTraversal()
}
