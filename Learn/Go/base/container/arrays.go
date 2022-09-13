package main

import (
	"fmt"
)

// 数组初始化
func arrayInit() {
	var array1 [5]int
	array2 := [3]int{1, 3, 5}
	array3 := [...]int{2, 4, 6, 8, 10}
	fmt.Println(array1, array2, array3)

	// 二维数组
	var grid [4][5] int
	fmt.Println(grid)
}

// 遍历数组
func traverse() {
	array := [...]int{2, 4, 6, 8, 10}
	for index, value := range array {
		fmt.Println(index, value)
	}
}

// 数组是值类型
func arrayPass() {
	array := [...]int{2, 4, 6, 8, 10}
	printArray(array)
	fmt.Println(array)

	printArrayPointer(&array)
	fmt.Println(array)
}

func printArray(array [5]int) {
	array[0] = 100
	for index, value := range array {
		fmt.Println(index, value)
	}
}

func printArrayPointer(array *[5]int) {
	array[0] = 100
	for index, value := range array {
		fmt.Println(index, value)
	}
}

func main() {
	arrayInit()
	traverse()
	arrayPass()
}
