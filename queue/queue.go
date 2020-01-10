package queue

// interface{} 表示任意类型
type Queue []interface{}

func (q *Queue) Push(v interface{})  {
	*q = append(*q, v)
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
