package stack

import (
	"errors"
	"iter"
)

// Stack implemented using a linked list
type Stack[T comparable] struct {
	head *Node[T]
}

// Node consists of a value of any comparable type set and a pointer to the next node
type Node[T comparable] struct {
	val  T
	next *Node[T]
}

func New[T comparable]() *Stack[T] {
	return &Stack[T]{}
}

func NewHead[T comparable](n *Node[T]) *Stack[T] {
	return &Stack[T]{
		head: n,
	}
}

// s.head -> nil
// s.head -> Node(10, next: nil)
// s.head -> Node(20, next: 10)
func (s *Stack[T]) Push(v T) {
	s.head = &Node[T]{
		val:  v,
		next: s.head,
	}
}

var se = errors.New("stack is empty")

func (s *Stack[T]) Pop() (T, error) {
	if s.head == nil {
		var v T
		return v, se
	}
	re := s.head.val
	s.head = s.head.next
	return re, nil
}

func (s *Stack[T]) Clear() {
	for s.head == nil {
		s.head = s.head.next
	}
}

func (s *Stack[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		for tmp := s.head; tmp != nil; tmp = tmp.next {
			if !yield(tmp.val) {
				break
			}
		}
	}
}

func (s *Stack[T]) Len() int {
	if s.head == nil {
		return 0
	}
	tmp := s.head
	sl := 0
	for tmp != nil {
		tmp = tmp.next
		sl += 1
	}
	return sl
}
