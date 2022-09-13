package main

import (
	"math/rand"
	"time"
	"fmt"
	"practice/base/stack"
)

/**
	数组中元素与下一个比它大的元素之间的距离 LeetCode No739
 */
func main() {
	rand.Seed(time.Now().Unix())
	var array [10]int
	i := 0
	for i < 10 {
		array[i] = rand.Intn(100)
		i ++
	}
	fmt.Println(array)
	solution(array)
}

func solution(array [10]int) {
	var resultArray [10]int
	indexStack := stack.Stack{}
	for index, value := range array {
		for indexStack.GetLength() > 0 && value > array[indexStack.Peek().(int)] {
			top := indexStack.Pop().(int)
			resultArray[top] = index - top
		}
		indexStack.Push(index)
	}
	fmt.Println(resultArray)
}
