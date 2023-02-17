package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := "lee"

	// 反射出类型
	t := reflect.TypeOf(s)
	fmt.Println(t)

	// 反射出值
	v := reflect.ValueOf(s)
	fmt.Println(v)

	s2 := v.Interface().(string)
	fmt.Println(s2)
}
