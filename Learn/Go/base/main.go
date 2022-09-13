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
	nums := []int{1, 2, 3}

	nums = append(nums, []int{4, 5, 6, 7, 8, 9}...)

	fmt.Println("nums", nums, len(nums), cap(nums))
}
