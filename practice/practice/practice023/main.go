package main

import (
	"fmt"
	"sort"
)

/**
	分发饼干 LeetCode No455

	标签：贪心算法
 */
func main() {
	g := []int{10, 9, 8, 7}
	s := []int{5, 6, 7, 8}
	sort.Slice(g, func(i, j int) bool {
		return g[i] < g[j]
	})
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	result := solution(g, s)
	fmt.Println(result)
}

func solution(g []int, s []int) int {
	count := 0
	next := 0
	for i := 0; i < len(g); i ++ {
		for j := next; j < len(s); j ++ {
			// 胃口比饼干小
			if g[i] <= s[j] {
				count ++
				next = j + 1
				break
			}
		}
	}
	return count
}
