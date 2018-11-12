package queue

import "fmt"

func ExampleQueue_Pop() {
	q := Queue{1}
	q.Push(1)
	q.Push(2)
	q.Push(3)

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())

	fmt.Println(q.IsEmpty())

	// output
	// 3
	// 2
	// 1
	// false
}
