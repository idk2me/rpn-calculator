package structs

type Stack[T any] struct {
	Data []T
}

type StackEmptyError struct{}

func (s StackEmptyError) Error() string {
	return "Popping from an empty stack"
}

func (s *Stack[T]) Push(el T) T {
	s.Data = append(s.Data, el)
	return el
}

func (s *Stack[T]) Pop() (T, error) {
	if len(s.Data) == 0 {
		return *new(T), StackEmptyError{}
	}
	el := s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]

	return el, nil
}

func (s *Stack[T]) Peek() T {
	if len(s.Data) == 0 {
		var zero T
		return zero
	}

	return s.Data[len(s.Data)-1]
}
