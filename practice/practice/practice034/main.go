package main

import "fmt"

/**
LeetCode No062 不同路径

标签：动态规划
*/
func main() {
	m := 3
	n := 7
	result := solution(m, n)
	fmt.Println(result)
}

func solution(m int, n int) int {
	// 动态二维数组定义
	numArray := make([][]int, m)
	// 将最左列和最上行的值设置为1，因为到达这些点只有一种走法
	for i := 0; i < m; i++ {
		numArray[i] = make([]int, n)
		numArray[i][0] = 1
	}
	for i := 0; i < n; i++ {
		numArray[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 当前点走法由左面、上面点走法的和
			numArray[i][j] = numArray[i-1][j] + numArray[i][j-1]
			fmt.Println(i, j, numArray[i][j])
		}
	}
	return numArray[m-1][n-1]
}
