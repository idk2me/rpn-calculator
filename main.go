package main

import (
	"fmt"
)

type Stack[T any] struct {
	data []T
}

type StackEmptyError struct{}

func (s StackEmptyError) Error() string {
	return "Popping from an empty stack"
}

func (s *Stack[T]) Push(el T) T {
	s.data = append(s.data, el)
	return el
}

func (s *Stack[T]) Pop() (T, error) {
	if len(s.data) == 0 {
		return *new(T), StackEmptyError{}
	}
	el := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return el, nil
}

func (s *Stack[T]) Peek() T {
	if len(s.data) == 0 {
		var zero T
		return zero
	}

	return s.data[len(s.data)-1]
}

func main() {
	var s Stack[int]
	s.data = []int{1, 234, 123, 4, 3, 41, 1234, 12, 3412}
	fmt.Printf("Initial size of stack is: %d\n", len(s.data))
	i := 0
	for {
		fmt.Printf("current index is: %d\n", i)
		fmt.Println(s.data)
		tmp, err := s.Pop()
		if err != nil {
			fmt.Printf("An error has occured: %s", err)
			break
		} else {
			fmt.Println(tmp)
			i++
		}
	}
}
