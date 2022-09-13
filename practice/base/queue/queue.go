/**
	队列
 */
package queue

type Queue struct {
	Data []interface{}
}

func (queue *Queue) Enqueue(data interface{}) {
	queue.Data = append(queue.Data, data)
}

func (queue *Queue) Dequeue() interface{} {
	if queue.IsEmpty() {
		return nil
	}
	data := queue.Data[0]
	queue.Data = queue.Data[1:]
	return data
}

func (queue *Queue) GetLength() int {
	return len(queue.Data)
}

func (queue *Queue) IsEmpty() bool {
	if queue.GetLength() == 0{
		return true
	}
	return false
}
