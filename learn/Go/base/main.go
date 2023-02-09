package main

import (
	"fmt"
)

type Person struct {
	firstName string
	lastName  string
}

func changeName(p Person) {
	p.firstName = "Bob"
}

func main() {
	var s []int

	s = append(s, 1)

	m := make(map[int]int)

	fmt.Printf("s %p, m %p", &s, &m)

	t(s, m)

	fmt.Println(s, m)

	//nums := []int{1, 2, 3}
	//
	//nums = append(nums, []int{4, 5, 6, 7, 8, 9}...)
	//
	//fmt.Println("nums", nums, len(nums), cap(nums))
	//
	//a := &Person{
	//	firstName: "lee",
	//	lastName:  "young",
	//}
	//hello(a)
}

func t(s []int, m map[int]int) {
	m[1] = 1
	s = append(s, 2)
	fmt.Printf("s %p, m %p", &s, &m)
}

func hello(v interface{}) {
	fmt.Println("v:", v)
}
