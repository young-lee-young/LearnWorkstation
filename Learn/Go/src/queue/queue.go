package queue

/**
	Go语言实现队列
 */
type Queue []int

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}
