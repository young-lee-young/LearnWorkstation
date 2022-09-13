package main

import (
	"strconv"
	"fmt"
)

func main() {
	mapDisorder()
	fmt.Println("-------------------")
	mapOrder()
	fmt.Println("-------------------")
	remove()
}

/**
	map是无序的
 */
func mapDisorder() {
	mapNew := make(map[string]int)

	for i := 0; i < 100; i ++ {
		// integer to array
		mapNew[strconv.Itoa(i)] = i
	}

	for _, v := range mapNew {
		fmt.Println(v)
	}
}

/**
	实现有序的map，需要借助slice
 */
func mapOrder()  {
	keySlice := make([]string, 0)
	mapNew := make(map[string]int)

	for i := 0; i < 100; i ++ {
		keySlice = append(keySlice, strconv.Itoa(i))
		mapNew[strconv.Itoa(i)] = i
	}

	//sort.Strings(keySlice)
	for _, v := range keySlice {
		fmt.Println(v)
	}
}

/**
	map删除元素
 */
func remove() {
	mapNew := make(map[string]string)
	mapNew["lee"] = "lee"
	delete(mapNew, "lee")
	// 重复删除不会报错
	delete(mapNew, "lee")
	for k, v := range mapNew {
		fmt.Println(k, v)
	}
}
