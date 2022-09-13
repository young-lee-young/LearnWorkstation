package main

import (
	"fmt"
	queue2 "practice/base/queue"
)

/**
	单词接龙 LeetCode No127

	标签：图广度优先搜索 BFS
 */
func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	result := solution(beginWord, endWord, wordList)
	fmt.Println(result)
}

func solution(beginWord string, endWord string, wordList []string) int {
	// 将所有单词放到map里
	wordMap := make(map[string]string)
	for _, word := range wordList {
		wordMap[word] = word
	}
	// 如果中间集为空，返回0
	if len(wordMap) == 0 {
		return 0
	}

	// 如果中间集没有最后值，返回0
	if _, ok := wordMap[endWord]; !ok {
		return 0
	}

	// 将开始值移除
	delete(wordMap, beginWord)

	// 广度优先遍历需要的队列
	queue := queue2.Queue{}
	queue.Enqueue(beginWord)

	// 图的遍历，得将已经访问过的添加到哈希表中
	visitedMap := make(map[string]string)
	visitedMap[beginWord] = beginWord

	step := 1
	for !queue.IsEmpty() {
		count := queue.GetLength()
		for i := 0; i < count; i ++ {
			currentWord := queue.Dequeue().(string)
			fmt.Println("----------", currentWord)
			changeRet := changeWord(currentWord, endWord, &queue, &visitedMap, wordMap)
			if changeRet {
				return step + 1
			}
		}
		step ++
	}
	return 0
}

func changeWord(currentWord string, endWord string, queue *queue2.Queue, visitedMap *map[string]string, wordMap map[string]string) bool {
	runeSlice := strToSlice(currentWord)
	fmt.Println(len(endWord))
	for i := 0; i < len(endWord); i ++ {
		// 保存原来的字符
		originChar := runeSlice[i]
		for k := 'a'; k <= 'z'; k ++ {
			if k == originChar {
				continue
			}
			// 字符替换
			runeSlice[i] = k
			nextWord := string(runeSlice)
			fmt.Println(nextWord)
			// 判断是否在给定的库中
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
		runeSlice[i] = originChar
	}
	return false
}

func strToSlice(str string) []rune {
	runeSlice := make([]rune, 0)
	for _, item := range str {
		runeSlice = append(runeSlice, item)
	}
	return runeSlice
}
