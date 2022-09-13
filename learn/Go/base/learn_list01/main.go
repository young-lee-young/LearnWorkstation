package main

import "fmt"

func main() {
	arrayA := [2]int{0, 1}
	var arrayB [2]int

	// 数组赋值是值复制
	arrayB = arrayA

	fmt.Printf("arrayA : %p , %v\n", &arrayA, arrayA)
	fmt.Printf("arrayB : %p , %v\n", &arrayB, arrayB)

	arrayParam(arrayA)
}

// 数组传惨是值复制
func arrayParam(array [2]int) {
	fmt.Printf("arrayC : %p , %v\n", &array, array)
}
