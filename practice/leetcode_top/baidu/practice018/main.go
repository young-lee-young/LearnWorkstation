package main

import (
	"practice/base/tree"
	"strconv"
	"strings"
)

/**
序列化和反序列化二叉搜索树

序列化使用前序遍历
反序列化时，比第一个节点下小的为左子树，比第一个节点大的为右子树
*/
func main() {

}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) serialize(root *tree.Node) string {
	ret := make([]string, 0)

	var preorder func(node *tree.Node)

	preorder = func(node *tree.Node) {
		if node != nil {
			ret = append(ret, strconv.Itoa(node.Data))
			preorder(node.Left)
			preorder(node.Right)
		}
	}

	preorder(root)

	return strings.Join(ret, ",")
}

func (this *Codec) deserialize(data string) *tree.Node {
	if data == "" {
		return nil
	}

	dataList := strings.Split(data, ",")
	numList := make([]int, 0)
	for i := 0; i < len(dataList); i++ {
		num, _ := strconv.Atoi(dataList[i])
		numList = append(numList, num)
	}

	var dfs func(nums []int) *tree.Node
	dfs = func(nums []int) *tree.Node {
		if len(nums) == 0 {
			return nil
		}

		rootNum := nums[0]

		numsLeft := make([]int, 0)
		numsRight := make([]int, 0)
		for i := 1; i < len(nums); i++ {
			if nums[i] < rootNum {
				numsLeft = append(numsLeft, nums[i])
			} else {
				numsRight = append(numsRight, nums[i])
			}
		}

		nodeLeft := dfs(numsLeft)
		nodeRight := dfs(numsRight)

		rootNode := &tree.Node{
			Data:  rootNum,
			Left:  nodeLeft,
			Right: nodeRight,
		}
		return rootNode
	}

	return dfs(numList)
}
