package main

import "fmt"

func main() {
	arrayLenChange()
}

// 切片如何实现可变长
func arrayLenChange() {
	// 切片初始化
	arrayInit := []int{}

	for i := 0; i < 10; i++ {
		arrayInit = append(arrayInit, i)
		fmt.Println(len(arrayInit), cap(arrayInit))
	}
}
