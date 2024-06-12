package queue

import (
	"Container/list"
)

// 큐
type Queue struct {
	list *list.List
}

// 생성자
func NewQueue() *Queue {
	return &Queue{list.New()}
}

// 삽입
func (q *Queue) Push(v interface{}) {
	q.list.PushBack(v)
}

// 삭제
func (q *Queue) Pop() interface{} {
	front := q.list.Front()
	if front != nil {
		return q.list.Remove(front)
	}
	return nil
}

// 맨 앞 원소
func (q *Queue) Front() interface{} {
	front := q.list.Front()
	if front != nil {
		return front.Value
	}
	return nil
}
