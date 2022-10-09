package main

import "fmt"

/**
字符串中所有字母的异位词 LeetCode No438

标签：滑动窗口 + 数组哈希
*/
func main() {
	s := "cbaebabacd"
	p := "abc"
	ret := solution(s, p)
	fmt.Println("ret:", ret)
}

func solution(s string, p string) []int {
	ret := make([]int, 0)

	if len(s) < len(p) {
		return ret
	}

	sMap := [26]int{}
	pMap := [26]int{}

	for i := range p {
		pMap[p[i]-'a']++
		sMap[s[i]-'a']++
	}

	if sMap == pMap {
		ret = append(ret, 0)
	}

	// i 是右边界的索引
	for i := len(p); i < len(s); i++ {
		// 右边界字符数量加 1
		sMap[s[i]-'a']++
		// 左边界的索引为 i - len(p)，左边界字符数量减 1
		sMap[s[i-len(p)]-'a']--

		if sMap == pMap {
			ret = append(ret, i-len(p)+1)
		}
	}

	return ret
}
