/**
	栈
 */
package stack

type Stack struct {
	Data []interface{}
}

func (stack *Stack) Push(data interface{}) {
	stack.Data = append(stack.Data, data)
}

func (stack *Stack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	data := stack.Data[stack.GetLength() - 1]
	stack.Data = stack.Data[:stack.GetLength() - 1]
	return data
}

func (stack *Stack) Peek() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	data := stack.Data[stack.GetLength() - 1]
	return data
}

func (stack *Stack) GetLength() int {
	return len(stack.Data)
}

func (stack *Stack) IsEmpty() bool {
	if stack.GetLength() == 0 {
		return true
	}
	return false
}
