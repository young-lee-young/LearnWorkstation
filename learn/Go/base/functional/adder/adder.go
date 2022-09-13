package main

import "fmt"

/**
	函数式编程

	正统函数式编程
	不可变性：不能有状态，只有常量和函数
	函数只能有一个参数
 */
func adder() func(int) int {
	// 自由变量
	sum := 0
	return func (v int) int {
		sum += v
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + 1 + ... + %d = %d\n", i, a(i))
	}

	b := adder2(0)
	for j := 0; j < 10; j++ {
		var s int
		s, b = b(j)
		fmt.Printf("0 + 1 + ... + %d = %d\n", j, s)
	}
}
