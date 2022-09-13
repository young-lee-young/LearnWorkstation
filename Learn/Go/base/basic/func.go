package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

/**
	方法
 */
// 多返回值
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operation: " + op)
	}
}

// 函数式编程，函数作为参数
func apply(op func(num1 int, num2 int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	funcName := runtime.FuncForPC(p).Name()
	fmt.Printf("call func %s with args %d %d\n", funcName, a, b)
	return op(a, b)
}

func pow(a , b int) int {
	return int(math.Pow(float64(a) , float64(b)))
}

// 可变参数列表
func sum(numbers ...int) int {
	result := 0
	for i := range numbers {
		result += numbers[i]
	}
	return result
}

// 交换两个数
func swap(a, b *int)  {
	*b, *a = *a, *b
}

func main() {
	a, b := 3, 4
	if result, err := eval(a, b, "="); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	fmt.Println(apply(pow, a, b))
	fmt.Println(sum(1, 2, 3, 4, 5))


	swap(&a, &b)
	fmt.Println(a, b)
}
