package main

import (
	"fmt"
	"math"
)

/**
	单词接龙 LeetCode No126

	标签：图广度优先遍历 BFS
 */
func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	result := solution(beginWord, endWord, wordList)
	fmt.Println(result)
}

func solution(beginWord string, endWord string, wordList []string) [][]string {
	// 将所有单词放到map里，后续O(1)时间复杂度查询
	wordMap := make(map[string]string)
	for _, word := range wordList {
		wordMap[word] = word
	}

	// 如果中间集为空，返回0
	if len(wordMap) == 0 {
		return [][]string{}
	}

	// 将开始值移除
	delete(wordMap, beginWord)

	//
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	ids := map[string]int{}
	for i, word := range wordList {
		ids[word] = i
	}
	if _, ok := ids[beginWord]; !ok {
		wordList = append(wordList, beginWord)
		ids[beginWord] = len(wordList) - 1
	}
	if _, ok := ids[endWord]; !ok {
		return [][]string{}
	}

	n := len(wordList)
	edges := make([][]int, len(wordList))
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if transformCheck(wordList[i], wordList[j]) {
				edges[i] = append(edges[i], j)
				edges[j] = append(edges[j], i)
			}
		}
	}

	res := [][]string{}
	cost := make([]int, n)
	queue := [][]int{[]int{ids[beginWord]}}

	for i := 0; i < n; i++ {
		cost[i] = math.MaxInt32
	}
	cost[ids[beginWord]] = 0

	for i := 0; i < len(queue); i++ {
		now := queue[i]
		last := now[len(now) - 1]
		if last == ids[endWord] {
			tmp := []string{}
			for _, index := range now {
				tmp = append(tmp, wordList[index])
			}
			res = append(res, tmp)
		} else {
			for _, to := range edges[last] {
				if cost[last] + 1 <= cost[to] {
					cost[to] = cost[last] + 1
					tmp := make([]int, len(now))
					copy(tmp, now)
					tmp = append(tmp, to)
					queue = append(queue, tmp)
				}
			}
		}
	}
	return res
}

func transformCheck(from, to string) bool {
	for i := 0; i < len(from); i++ {
		if from[i] != to[i] {
			return from[i + 1:] == to[i + 1:]
		}
	}
	return false
}
