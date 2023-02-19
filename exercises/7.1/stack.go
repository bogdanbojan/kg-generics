package stack

// Your implementation of Stack goes here!

type Stack[E any] struct {
	data []E
}

func (s Stack[E]) Len() int {
	return len(s.data)
}

func (s *Stack[E]) Push(vals ...E) {
	s.data = append(s.data, vals...)
}

func (s *Stack[E]) Pop() (v E, ok bool) {
	if len(s.data) == 0 {
		return v, false
	}
	v = s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v, true
}

// Implementation with pointer semantics. It's quite confusing
// and another way of doing things should be preferred - like
// having a more elegant approach by using structs.
// type Stack[T constraints.Ordered] []T

// func (s *Stack[T]) Push(vv ...T) {
// 	*s = append(*s, vv...)
// }

// func (s Stack[T]) Len() int {
// 	return len(s)
// }

// func (s *Stack[T]) Pop() (v T, ok bool) {
// 	if len(*s) == 0 {
// 		return v, false
// 	}
// 	v = (*s)[s.Len()-1]
// 	*s = (*s)[:s.Len()-1]
// 	return v, true
// }
