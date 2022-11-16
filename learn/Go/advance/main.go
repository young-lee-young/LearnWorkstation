package main

import (
	"fmt"
	"unsafe"
)

type s1 struct {
	num1 int32
	num2 int32
}

type s2 struct {
	num1 int16
	num2 int32
}

func main() {
	a := int64(0)
	fmt.Println(unsafe.Alignof(a))
}
