package stack

import "golang.org/x/exp/constraints"

// Your implementation of Stack goes here!

type Stack[T constraints.Ordered] []T

func (s *Stack[T]) Push(vv ...T) {
	*s = append(*s, vv...)
}

func (s Stack[T]) Len() int {
	return len(s)
}

func (s *Stack[T]) Pop() (v T, ok bool) {
	if len(*s) == 0 {
		return v, false
	}
	v = (*s)[s.Len()-1]
	*s = (*s)[:s.Len()-1]
	return v, true
}
