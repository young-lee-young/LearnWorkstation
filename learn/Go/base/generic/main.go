package main

import "fmt"

type Generic interface {
	int | int32 | int64 | float32 | float64 | string
}

func sum[T Generic](a T, b T) T {
	return a + b
}

func example[T any, R comparable]() {

}

func main() {
	a := sum(1, 2)
	b := sum(10.0, 20.0)
	c := sum("hello", "world")
	fmt.Println(a, b, c)
}
