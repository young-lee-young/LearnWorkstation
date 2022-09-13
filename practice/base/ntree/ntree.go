package ntree

import (
	"math/rand"
	"time"
	"fmt"
	queue2 "practice/base/queue"
)

/**
	多叉树节点
 */
type Node struct {
	Data int
	Children []*Node
}

/**
	N叉树
 */
type NTree struct {
	Root *Node
}

/**
	生成N叉数
 */
func (ntree *NTree) Generate(maxHeight int, width int) {
	rand.Seed(time.Now().UnixNano())
	root := generate(maxHeight, 1, width)
	ntree.Root = root
}

/**
	层序遍历N叉数
 */
func (ntree *NTree) ShowTree(width int) {
	queue := queue2.Queue{}
	if ntree.Root == nil {
		return
	}
	queue.Enqueue(ntree.Root)
	for !queue.IsEmpty() {
		count := queue.GetLength()
		for i := 0; i < count; i ++ {
			node := queue.Dequeue().(*Node)
			fmt.Printf(" %d ", node.Data)
			for _, value := range node.Children {
				queue.Enqueue(value)
			}
			if (i + 1) % width == 0 {
				fmt.Printf(" | ")
			}
		}
		fmt.Println()
	}
}

/**
	递归生成
 */
func generate(maxHeight, height int, width int) *Node {
	if height > maxHeight {
		return nil
	}
	children := make([]*Node, 0)
	for i := 0; i < width; i ++ {
		child := generate(maxHeight, height + 1, width)
		if child != nil {
			children = append(children, child)
		}
	}
	num := randInt((height - 1) * 10, height * 10)
	newNode := &Node{
		Data: num,
		Children: children,
	}
	return newNode
}

func randInt(min int, max int) int {
	for {
		num := rand.Intn(max)
		if num > min {
			return num
		}
	}
}
