package overview_test

import (
	"testing"

	"golang.org/x/exp/constraints"
)

type queue[T any] struct {
	data []T
}

func NewQueue[T constraints.Ordered]() *queue[T] {
	return &queue[T]{make([]T, 0)}
}

func (q *queue[T]) Push(v T) {
	q.data = append(q.data, v)
}

func (q *queue[T]) Pop() T {
	v := q.data[0]
	q.data = q.data[1:]
	return v
}

func (q *queue[T]) Size() int {
	return len(q.data)
}

func (q *queue[T]) Empty() bool {
	return len(q.data) == 0
}

func (q *queue[T]) Top() T {
	if q.Empty() {
		var zero T
		return zero
	}

	return q.data[0]
}

func TestQueueElement(t *testing.T) {
	q := NewQueue[int]()
	q.Push(1)
	q.Push(2)
	q.Push(3)

	if q.Size() != 3 {
		t.Errorf("Expected size 3, got %d", q.Size())
	}

	if q.Top() != 1 {
		t.Errorf("Expected 1, got %d", q.Top())
	}

	if q.Pop() != 1 {
		t.Errorf("Expected 1, got %d", q.Pop())
	}

	if q.Pop() != 2 {
		t.Errorf("Expected 2, got %d", q.Pop())
	}

	if q.Pop() != 3 {
		t.Errorf("Expected 3, got %d", q.Pop())
	}

	if !q.Empty() {
		t.Errorf("Expected empty queue, got %d", q.Size())
	}
}
