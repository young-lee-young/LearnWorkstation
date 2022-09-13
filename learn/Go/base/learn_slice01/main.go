package main

import (
	"fmt"
	"unsafe"
)

/**
slice，是有序的，底层使用数组
*/
func main() {
	//sliceInit()
	//sliceNil()
	//sliceEmpty()
	sliceCut()
}

/**
slice初始化
*/
func sliceInit() {
	// 第二个参数是len，第三个参数是cap，如果实例化时 len > cap 会报 len larger than cap in make([]int) 错误
	initSlice := make([]int, 3, 5)
	fmt.Println(len(initSlice))
	fmt.Println(cap(initSlice))
	println(initSlice)
	fmt.Println(initSlice)

	// 获取地址
	ptr := unsafe.Pointer(&initSlice[0])
	fmt.Println(&ptr)

	initSlice[1] = 1
	initSlice = append(initSlice, 3)
	fmt.Println(initSlice)
}

/**
nil slice
指针为nil，也就是slice不会指向底层数组
|--------|---------|----------|
|  nil   |   0     |     0    |
|  ptr   | Length  | Capacity |
|--------|---------|----------|
*/
func sliceNil() {
	var nilSlice []int
	println(nilSlice)
	fmt.Println(nilSlice)
}

/**
empty slice
指针指向底层数组长度为0的空数组
|--------|---------|----------|
|  ADDR  |   0     |     0    |
|  ptr   | Length  | Capacity |
|--------|---------|----------|
*/
func sliceEmpty() {
	emptySlice := make([]int, 0)
	println(emptySlice)
	fmt.Println(emptySlice)
}

/**
	slice切片
*/
func sliceCut() {
	numSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	newSlice1 := numSlice[2:5]
	fmt.Printf("new slice1: %v, len %v, cap %v, \n", newSlice1, len(newSlice1), cap(newSlice1))

	// 从1切片到5，容量为 8 - 1
	newSlice2 := numSlice[1:5:8]
	fmt.Printf("new slice2: %v, len %v, cap %v, \n", newSlice2, len(newSlice2), cap(newSlice2))
}
