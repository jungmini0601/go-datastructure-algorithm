package stack

import (
	"container/list"
)

// 스택 코드 생성
type Stack struct {
	list *list.List
}

// 스택 생성
func NewStack() *Stack {
	return &Stack{
		list: list.New(),
	}
}

// 스택에 값 추가
func (s *Stack) Push(value interface{}) {
	s.list.PushBack(value)
}

// 스택에서 값 제거
func (s *Stack) Pop() interface{} {
	if s.list.Len() == 0 {
		return nil
	}
	e := s.list.Back()
	s.list.Remove(e)
	return e.Value
}

// 스택 최상단 값 반환
func (s *Stack) Top() interface{} {
	if s.list.Len() == 0 {
		return nil
	}
	return s.list.Back().Value
}
