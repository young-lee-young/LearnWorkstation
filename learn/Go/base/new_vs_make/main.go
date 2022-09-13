package main

import (
	"fmt"
	"go/types"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := new(Person)
	fmt.Println("p:", p)
	checkType(p)
	fmt.Println("name:", p.Name)
	fmt.Println("age:", p.Age)

	nums := make([]int, 4, 8)
	fmt.Println("nums", nums)
	checkType(nums)
}

func checkType(i interface{}) {
	switch t := i.(type) {
	case *Person:
		fmt.Println("type person")
	case types.Pointer:
		fmt.Println("type pointer")
	case types.Slice:
		fmt.Println("type slice")
	default:
		fmt.Printf("type unknown %T", t)
	}
}
