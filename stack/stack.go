package stack

import "slices"

type Stack[T any] []T

func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

func (s *Stack[T]) Pop() T {
	r := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return r
}

func (s *Stack[T]) Drain(n int) {
	*s = (*s)[:n]
}

func (s *Stack[T]) SplitOff(n int) Stack[T] {
	r := (*s)[n:]
	*s = (*s)[:n]
	return slices.Clone(r)
}

func (s *Stack[T]) Len() int {
	return len(*s)
}
