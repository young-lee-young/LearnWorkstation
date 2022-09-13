package main

import (
	"fmt"
)

func main() {

	// map声明
	mapInit := map[string]string{}

	// map声明并初始化并赋值
	mapInitValue := map[string]int{"scores1": 60, "scores2": 80, "scores3": 100}

	// map声明并初始化并指定容量
	mapInitCap := make(map[int]int, 10)

	// map存值
	mapInit["name"] = "leeyoung"

	// map取值
	scores1 := mapInitValue["scores1"]
	fmt.Println(scores1)

	// 长度
	fmt.Println(len(mapInitCap))

	mapInitCap[0] = 100
	// 判断是否有值，还是初始化的默认值
	if value, isExist := mapInitCap[0]; isExist {
		fmt.Println(isExist)
		fmt.Println(value)
	} else {
		fmt.Println(isExist)
		fmt.Println(value)
	}

	// range
	for key, value := range mapInitValue {
		fmt.Println(key)
		fmt.Println(value)
	}

	mapForSet()
	mapAdvancedUsage()
}

// map高级用法，值为函数
func mapAdvancedUsage() {
	mapAdvanced := map[int]func(param int) int{}
	mapAdvanced[0] = func(param int) int {
		return param
	}
	mapAdvanced[1] = func(param int) int {
		return param * param
	}
	mapAdvanced[2] = func(param int) int {
		return param * param * param
	}

	fmt.Println(mapAdvanced[0](2))
	fmt.Println(mapAdvanced[1](2))
	fmt.Println(mapAdvanced[2](2))
}

// 用map实现set
func mapForSet() {
	mapSet := map[int]bool{}
	mapSet[0] = true

	delete(mapSet, 0)
}
