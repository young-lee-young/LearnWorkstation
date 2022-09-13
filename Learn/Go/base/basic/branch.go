package main

import (
	"io/ioutil"
	"fmt"
)

/**
	条件语句
 */
// if 语句
func ifFunc() {
	const filename = "/Users/leeyoung/File/Temp/vim_learn.py"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

// switch 语句
func switchFunc(a int, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator: " + op)
	}
	return result
}

func main() {
	ifFunc()
	result := switchFunc(1, 2, "+")
	fmt.Println(result)
}
