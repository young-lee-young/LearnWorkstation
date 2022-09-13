package main

import (
	"math/rand"
	"time"
	"fmt"
	"practice/base/stack"
)

/**
	循环数组中比当前元素大的下一个元素 LeetCode No503
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
	// 左移运算，10 * 2 ^ 1
	arrayLength := 10 << 1
	for i := 0; i < arrayLength; i ++ {
		index := i % 10
		value := array[index]
		for indexStack.GetLength() > 0 && value > array[indexStack.Peek().(int)] {
			top := indexStack.Pop().(int)
			resultArray[top] = array[index]
		}
		indexStack.Push(index)
	}
	fmt.Println(resultArray)
}
