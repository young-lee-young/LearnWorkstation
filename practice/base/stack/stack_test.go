package stack

import (
	"testing"
	"fmt"
)

func TestStack(t *testing.T) {
	stack := Stack{}
	fmt.Println(stack.GetLength())
	fmt.Println(stack.IsEmpty())
	i := 0
	for i < 10 {
		stack.Push(i)
		i ++
	}

	fmt.Println(stack.GetLength())
	fmt.Println(stack.IsEmpty())

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
