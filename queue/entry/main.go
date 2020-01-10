package main

import (
	"awesomeProject1/queue"
	"fmt"
)

func main() {
	q := queue.Queue{1}

	q.Push(2)
	q.Push(3)

	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q)

	q.Push("anytype")
	fmt.Println(q.Pop())
}
