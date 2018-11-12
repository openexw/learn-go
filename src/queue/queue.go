package queue

// An FIFO queue.
type Queue []int

// Push the int number.
// 	e.g. queue.Push(1)
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

// Pop the int number.
// e.g. queue.Pop()
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// Check the queue is empty.
// e.g. queue.Pop()
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
