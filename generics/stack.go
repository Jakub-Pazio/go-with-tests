package generics

import "errors"

type Stack[T any] struct {
	stack []T
}

func (s *Stack[T]) push(item T) {
	s.stack = append(s.stack, item)
}

func (s *Stack[T]) pop() (T, error) {
	if len(s.stack) < 1 {
		var zero T
		return zero, errors.New("Stack is empty")
	}
	index := len(s.stack) - 1
	item := s.stack[index]
	s.stack = s.stack[:index]
	return item, nil
}
