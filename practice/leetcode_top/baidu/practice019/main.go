package main

/**
从前序遍历和中序遍历构造二叉树
*/
func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
}

func solution(preoder []int, inorder []int) {

	rootNum := preoder[0]

	numsLeft := make([]int, 0)
	numsRight := make([]int, 0)

	findRoot := false
	for i := 0; i < len(inorder); i++ {
		if !findRoot {
			if inorder[i] != rootNum {
				numsLeft = append(numsLeft, inorder[i])
			} else {
				findRoot = true
			}
		} else {
			numsRight = append(numsRight, inorder[i])
		}
	}

}
