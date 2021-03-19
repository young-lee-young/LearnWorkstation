package main

import (
	"fmt"
	"time"
)

func main() {
	username, scores := multiReturn()
	fmt.Println(username, scores)

	decoFunc := timeSpent(slowFunc)
	decoFunc(10)

	resultSum := varParam(1, 2, 3, 4, 5)
	fmt.Println(resultSum)

	deferFunc()
}

// 多返回值函数
func multiReturn() (string, int) {
	return "leeyoung", 100
}

// 参数和返回值为函数
func timeSpent(inner func(param int) int) func(param int) int {
	return func(n int) int{
		start := time.Now()
		result := inner(n)

		fmt.Println(time.Since(start).Seconds())
		return result
	}
}

func slowFunc(param int) int {
	time.Sleep(time.Second * 2)
	return param
}

// 可变长参数
func varParam(params ...int) int {
	result := 0

	for _, value := range params {
		result += value
	}
	return result
}

// 延迟执行函数
func deferFunc()  {
	// 延迟执行，程序报错也会执行
	defer func() {
		fmt.Println("defer func running")
	}()

	fmt.Println("outer func running")
	// 报错
	panic("error")
}
