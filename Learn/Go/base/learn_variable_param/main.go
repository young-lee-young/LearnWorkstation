package main

import (
	"fmt"
	"reflect"
)

/**
Go可变参数
*/
func main() {
	//variableParam("1", "2", "3")

	params := []int{1, 2, 3, 4, 5}
	// 参数必须是被调用方法参数类型相同的slice，且只能是slice
	//called(params...)

	// 传递任意类型参数
	anyTypeParam(1, "lee", true, params)
}

/**
可变参数，params是一个slice，len 3, cap 3
*/
func variableParam(params ...string) {
	for _, param := range params {
		fmt.Println(param)
	}
}

func called(params ...int) {
	// append后面的参数就是可变参数
	params = append(params, 6, 7)
	fmt.Println(params)
}

/**
传递任意类型参数
*/
func anyTypeParam(params ...interface{}) {
	for _, param := range params {
		switch reflect.Type(param) {
		case param:
		}
	}
}
