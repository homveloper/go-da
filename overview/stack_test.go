package overview_test

import "testing"

type stack[T any] struct {
	data []T
}

func NewStack[T any]() *stack[T] {
	return &stack[T]{}
}

func (s *stack[T]) Empty() bool {
	return len(s.data) == 0
}

func (s *stack[T]) Size() int {
	return len(s.data)
}

func (s *stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *stack[T]) Pop() T {

	if s.Empty() {
		panic("stack is empty")
	}

	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v
}

func (s *stack[T]) Top() T {

	if s.Empty() {
		panic("stack is empty")
	}

	return s.data[len(s.data)-1]
}

func TestStackElement(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.Size() != 3 {
		t.Errorf("Expected size 3, got %d", s.Size())
	}

	if s.Top() != 3 {
		t.Errorf("Expected 3, got %d", s.Top())
	}

	if s.Pop() != 3 {
		t.Errorf("Expected 3, got %d", s.Pop())
	}

	if s.Pop() != 2 {
		t.Errorf("Expected 2, got %d", s.Pop())
	}

	if s.Pop() != 1 {
		t.Errorf("Expected 1, got %d", s.Pop())
	}

	if !s.Empty() {
		t.Errorf("Expected empty stack, got %d", s.Size())
	}
}
