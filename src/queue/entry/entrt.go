package main

import (
	"fmt"
	"queue"
)

func main() {
	/*q := queue.Queue{1}
	q.Push(2)
	q.Push(3)
	q.Pop()*/
	q := queue.Queue{1}
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Pop()
	fmt.Println(q)
}
