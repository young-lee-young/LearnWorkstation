package main

import "fmt"

/**
可排序：Int、float、string
可比较：bool、complex、pointer、channel、interface、array
不可比较：slice、map、func

struct比较
不包含不可比较成员变量：==操作符可以比较指针和实例
包含不可比较成员变量：==操作符可以表叫指针，不能直接比较实例
*/

type Person struct {
	Name string
	Age  int
}

type T1 struct {
	Name string
	Age  int
	Arr  [2]bool
	ptr  *int
	//slice []int
	//map1  map[string]string
}

func main() {
	person1 := Person{
		Name: "lee",
		Age:  1,
	}

	person2 := Person{
		Name: "lee",
		Age:  1,
	}

	fmt.Println(person1 == person2)

	t1 := T1{
		Name: "yxc",
		Age:  1,
		Arr:  [2]bool{true, false},
		ptr:  new(int),
		//slice: []int{1, 2, 3},
		//map1:  make(map[string]string, 0),
	}

	t2 := T1{
		Name: "yxc",
		Age:  1,
		Arr:  [2]bool{true, false},
		ptr:  new(int),
		//slice: []int{1, 2, 3},
		//map1:  make(map[string]string, 0),
	}

	// 报错 实例不能比较 Invalid operation: t1 == t2 (operator == not defined on T1)
	fmt.Println(t1 == t2)
	// 指针可以比较
	fmt.Println(&t1 == &t2) // false
}
