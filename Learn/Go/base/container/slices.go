package main

import (
	"fmt"
)

/**
	切片

	slice底层
	ptr：指向切片开始的位置
	len：切片长度
	cap：切片容量
 */
// 创建slice
func createSlice() {
	var slice []int
	fmt.Println(slice)

	// 创建slice并指定长度
	slice2 := make([]int, 10)
	fmt.Printf("createSlice slice2 %v, len %d, cap %d\n", slice2, len(slice2), cap(slice2))

	// 创建slice并指定长度，容量
	slice3 := make([]int, 12, 16)
	fmt.Printf("createSlice slice3 %v, len %d, cap %d\n", slice3, len(slice3), cap(slice3))

	// slice扩容，cap成倍扩容
	for i := 0; i < 10; i++ {
		slice = append(slice, i)
		fmt.Printf("slice %v, len %d, cap %d\n", slice, len(slice), cap(slice))
	}

	// slice拷贝，将slice拷贝进slice3
	copy(slice3, slice)
	fmt.Println(slice3)

	// 删除元素
	slice3 = append(slice3[:4], slice3[5:]...)
	// 删除第一个元素
	slice3 = slice3[1:]
	// 删除最后一个元素
	slice3 = slice3[:len(slice3) - 1]
	fmt.Println(slice3)
}

func updateSlice()  {
	array := [...]int{1, 2, 3, 4, 5, 6, 7}

	slice1 := array[2:]
	slice2 := array[:6]

	slice1[0] = 100
	fmt.Println(slice1, slice2)
}

// slice扩展，不能向前扩展，只能向后扩展，且扩展的数量不能超过容量
func extendSlice() {
	array := [...]int{1, 2, 3, 4, 5, 6, 7}

	slice1 := array[2:6]
	// 二次切片
	slice2 := slice1[3:5]
	fmt.Printf("slice1 %v, len %d, cap %d\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2 %v, len %d, cap %d\n", slice2, len(slice2), cap(slice2))
}

// slice添加值
func appendSlice() {
	array := [...]int{1, 2, 3, 4, 5, 6, 7}
	slice1 := array[5:6]

	slice3 := append(slice1, 10)
	// 添加元素时，如果超过cap，系统会重新分配更大的底层数组
	slice4 := append(slice3, 11)
	slice5 := append(slice4, 12)
	fmt.Println(slice5)
}

func main() {
	createSlice()
	updateSlice()
	extendSlice()
	appendSlice()
}
