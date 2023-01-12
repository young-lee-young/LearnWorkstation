/**
LeetCode No108 将有序数组转换为二叉搜索树
*/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}

	ret := solution(nums)
	fmt.Println("ret:", ret)
}

func solution(nums []int) *TreeNode {

}
