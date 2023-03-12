package overview_test

import (
	"testing"

	"golang.org/x/exp/constraints"
)

type set[T constraints.Ordered] struct {
	inner map[T]bool
}

func (s *set[T]) Add(v T) bool {
	if s.Contains(v) {
		return false
	}

	s.inner[v] = true
	return true
}

func (s *set[T]) Remove(v T) bool {
	if !s.Contains(v) {
		return false
	}

	delete(s.inner, v)
	return true
}

func (s *set[T]) Contains(v T) bool {
	return s.inner[v]
}

func (s *set[T]) Intersect(other *set[T]) *set[T] {
	result := NewSet[T]()
	for k := range s.inner {
		if other.Contains(k) {
			result.Add(k)
		}
	}

	return result
}

func (s *set[T]) Union(other *set[T]) *set[T] {
	result := NewSet[T]()
	for k := range s.inner {
		result.Add(k)
	}

	for k := range other.inner {
		result.Add(k)
	}

	return result
}

func (s *set[T]) Size() int {
	return len(s.inner)
}

func NewSet[T constraints.Ordered]() *set[T] {
	return &set[T]{make(map[T]bool)}
}

func TestSetElement(t *testing.T) {
	s := NewSet[int]()

	s.Add(1)
	s.Add(2)
	s.Add(3)

	if true == s.Add(2) {
		t.Error("Add 2 failed")
	}

	if true == s.Contains(10) {
		t.Error("Contains 10 failed")
	}

	if true == s.Remove(10) {
		t.Error("Remove 10 failed")
	}
}

func TestSetIntersect(t *testing.T) {
	s1 := NewSet[int]()
	s2 := NewSet[int]()

	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	s2.Add(2)
	s2.Add(3)
	s2.Add(4)

	s3 := s1.Intersect(s2)
	if s3.Size() != 2 {
		t.Error("Intersect failed")
	}
}

func TestSetUnion(t *testing.T) {
	s1 := NewSet[int]()
	s2 := NewSet[int]()

	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	s2.Add(2)
	s2.Add(3)
	s2.Add(4)

	s3 := s1.Union(s2)
	if s3.Size() != 4 {
		t.Error("Union failed")
	}
}
