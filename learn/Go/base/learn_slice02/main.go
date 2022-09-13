package main

import "fmt"

func main() {
	sliceExpansion()
	//copyOrNot()
}

/**
数组扩容

扩容机制

如果扩容大于2倍，直接就是实际容量
如果期望的容量大于两倍原始容量，扩容到期望容量
否则判断底层数组长度是否大于1024，小于1024扩大一倍，否则扩大1.25倍

再进行内存对齐
*/
func sliceExpansion() {
	// append多个元素
	numSlice1 := []int{0}
	println(numSlice1)

	// 新容量大于2倍旧容量，最终是新容量
	appendSlice1 := append(numSlice1, 1, 2, 3)
	println(appendSlice1)

	// 长度小于1024，容量扩大2倍
	numSlice2 := []int{0, 1, 2, 3, 4, 5}
	appendSlice2 := append(numSlice2, 6)
	println(appendSlice2)

	// 长度小于1024，扩大1.25倍
	numSlice3 := make([]int, 0)
	for i := 0; i < 1025; i++ {
		numSlice3 = append(numSlice3, i)
	}
	println(numSlice3)
}

func copyOrNot() {
	array := [4]int{0, 1, 2, 3}

	// 将数组切片，此时长度为2，容量为3
	slice := array[1:3]
	fmt.Printf("slice %v, address %p, len %v, cap %v \n", slice, &slice, len(slice), cap(slice))

	// 修改切片的第2个值，此时会修改底层的数组
	slice[1] = 100
	fmt.Printf("slice %v, address %p, len %v, cap %v \n", slice, &slice, len(slice), cap(slice))
	fmt.Println("array: ", array)

	// append切片，此时长度为3，容量为3，切片容量已满，仍然会修改底层数组
	slice = append(slice, 200)
	fmt.Printf("slice %v, address %p, len %v, cap %v \n", slice, &slice, len(slice), cap(slice))
	fmt.Println("array: ", array)

	// 继续append切片，此时长度为4，容量为6，此时底层数组改变
	newSlice := append(slice, 4)
	//fmt.Printf("slice %v, address %p, len %v, cap %v \n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("newSlice %v, address %p, len %v, cap %v \n", newSlice, &newSlice, len(newSlice), cap(newSlice))
	fmt.Println("array: ", array)

	slice[0] = 300
	fmt.Printf("slice %v, address %p, len %v, cap %v \n", slice, &slice, len(slice), cap(slice))
	fmt.Println("array: ", array)
}
