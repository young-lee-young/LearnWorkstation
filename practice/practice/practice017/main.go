package main

import (
	"fmt"
	queue2 "practice/base/queue"
)

/**
	最小基因变化 LeetCode No433

	标签：图广度优先遍历
 */
func main() {
	beginWord := "AAAAACCC"
	endWord := "AACCCCCC"
	wordSlice := []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}

	result := solution(beginWord, endWord, wordSlice)
	fmt.Println(result)
}

func solution(beginWord string, endWord string, wordSlice []string) int {
	// 将所有单词放到哈希表中
	wordMap := make(map[string]string)
	for _, word := range wordSlice {
		wordMap[word] = word
	}

	// 如果中间集为空，返回0
	if len(wordSlice) == 0 {
		return -1
	}

	if _, ok := wordMap[endWord]; !ok {
		return -1
	}

	// 将开始值移除
	delete(wordMap, beginWord)

	// 广度优先遍历需要的队列
	queue := queue2.Queue{}
	queue.Enqueue(beginWord)

	// 图广度优先遍历判断是否已经访问过
	visitedMap := make(map[string]string)
	visitedMap[beginWord] = beginWord

	step := 0
	for !queue.IsEmpty() {
		count := queue.GetLength()
		for i := 0; i < count; i ++ {
			currentWord := queue.Dequeue().(string)
			changeRet := changeWord(currentWord, endWord, wordMap, &queue, &visitedMap)
			if changeRet {
				return step + 1
			}
		}
		step ++
	}
	return -1
}

func changeWord(currentWord string, endWord string, wordMap map[string]string, queue *queue2.Queue, visitedMap *map[string]string) bool {
	changeRuneSlice := []rune{'A', 'C', 'G', 'T'}
	runeSlice := strToRuneSlice(currentWord)
	for i := 0; i < len(endWord); i ++ {
		originRune := runeSlice[i]
		for _, changeRune := range changeRuneSlice {
			if changeRune == originRune {
				continue
			}
			// 字符替换
			runeSlice[i] = changeRune
			nextWord := string(runeSlice)
			if _, ok := wordMap[nextWord]; ok {
				// 成功找到结尾元素
				if nextWord == endWord {
					return true
				}
				// 判断是否已经访问过
				if _, ok := (*visitedMap)[nextWord]; !ok {
					queue.Enqueue(nextWord)
					(*visitedMap)[nextWord] = nextWord
				}
			}
		}
		runeSlice[i] = originRune
	}
	return false
}

func strToRuneSlice(str string) []rune {
	runeSlice := make([]rune, 0)
	for _, item := range str {
		runeSlice = append(runeSlice, item)
	}
	return runeSlice
}
