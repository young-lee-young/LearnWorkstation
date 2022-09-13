package main

import "fmt"

/**
	map

	map key
	1. map使用哈希表，key必须可以比较相等
	2. 除了slice、map、func的内建类型都可以作为key
 */
// map初始化
func mapInit() {
	// 这种定义的map是nil
	var mNil map[string]string

	// 这种定义的map是空map
	m := map[string]string{}

	// 初始化赋值
	mValue := map[string]int{
		"scores1": 60,
		"scores2": 80,
		"scores3": 100,
	}

	mMake := make(map[string]bool)
	fmt.Println(m, mValue, mMake, mNil)
}

// map循环，map是无序的
func mapRange() {
	m := map[string]int{
		"scores1": 60,
		"scores2": 80,
		"scores3": 100,
	}
	for key, value := range m {
		fmt.Println(key, value)
	}
}

// map取值
func mapGetValue() {
	m := map[string]int{
		"scores1": 60,
		"scores2": 80,
		"scores3": 100,
	}

	scores1, exists := m["scores1"]
	fmt.Println(scores1, exists)

	scores4, exists := m["scores4"]
	fmt.Println(scores4, exists)
}

// map删除元素
func mapDelete() {
	m := map[string]int{
		"scores1": 60,
		"scores2": 80,
		"scores3": 100,
	}
	delete(m, "scores2")
	fmt.Println(m)
}

func main() {
	mapInit()
	mapRange()
	mapGetValue()
	mapDelete()
}
