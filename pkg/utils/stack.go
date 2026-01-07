package utils

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(item T) {
	s.elements = append(s.elements, item)
}

func (s *Stack[T]) Pop() T {
	if len(s.elements) == 0 {
		panic("Cannot pop from an empty stack")
	}

	defer func() {
		s.elements = s.elements[:len(s.elements)-1]
	}()

	return s.elements[len(s.elements)-1]
}

func (s *Stack[T]) Top() T {
	if len(s.elements) == 0 {
		panic("Cannot top on an empty stack")
	}

	return s.elements[len(s.elements)-1]
}

func (s *Stack[T]) Size() int {
	return len(s.elements)
}
