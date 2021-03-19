package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

/**
	首字母大写：public
	首字母小写：private
 */
// 作用域为包内部
var aa = "hello world"

func helloWorld()  {
	fmt.Println(aa)
}

// 初始化
func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

// 初始化并赋值
func variableInitialValue()  {
	var a, b int = 3, 4
	var s string = "lee"
	fmt.Printf("%d %d %s\n", a, b, s)
}

// 类型推断
func varialbleTypeDeduction()  {
	var a, b, c, d = 3, 3, true, "lee"
	var s = "lee"
	fmt.Println(a, b, c, d, s)
}

// 赋值简写
func variableShorter() {
	a, b, c, d := 3, 4, true, "lee"
	c = false
	fmt.Println(a, b, c, d)
}

// 复数
func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))

	// 欧拉公式
	fmt.Println(cmplx.Pow(math.E, 1i * math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i * math.Pi) + 1)
}

// 强制类型转换
func triangle() {
	var a, b = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))
	return c
}

// 常量定义
func consts() {
	const filename = "lee.txt"
	fmt.Println(filename)
}

// 枚举
func enums() {
	const (
		cpp = iota
		java
		python
		golang
	)
	fmt.Println(cpp, java, python, golang)

	// b, kb, mb, gb, tb, pb
	const (
		// 左移多少位
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	helloWorld()
	variableZeroValue()
	variableInitialValue()
	varialbleTypeDeduction()
	variableShorter()

	euler()
	triangle()
	consts()
	enums()
}
