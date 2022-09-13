package main

import "fmt"

func main() {
	// 数组的声明
	var arrayInit [3]int

	// 多维数组初始化
	arrayInitMulDimen := [2][2]int{{1, 2}, {3, 4}}
	arrayInitMulDimen[0][0] = 100
	fmt.Println(arrayInitMulDimen)

	// 数组赋值
	arrayInit[1] = 0

	arraySlice()
	traverseArray()
}

// 遍历数组
func traverseArray() {
	// 数组声明并初始化
	arrayInit := [...]string{"lee", "zhao", "xu", "hu"}

	for i := 0; i < len(arrayInit); i++ {
		fmt.Println(arrayInit[i])
	}

	for _, value := range arrayInit {
		fmt.Println(value)
	}
}

// 数组切片
func arraySlice()  {
	// 声明同时初始化赋值
	arrayInitValue := [5]int{1, 2, 3, 4, 5}

	// 切片
	arraySliced := arrayInitValue[1:3]
	fmt.Println(arraySliced, len(arraySliced), cap(arraySliced))

	// 数组添加元素
	arraySliced = append(arraySliced, 100)
	fmt.Println(arraySliced, len(arraySliced), cap(arraySliced))

	// 原始数组也被改变
	fmt.Println(arrayInitValue)
}
