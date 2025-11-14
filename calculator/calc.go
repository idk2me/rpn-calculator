package calculator

import (
	"fmt"

	"example.com/m/structs"
)

func sum(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mult(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

var ops = map[string]func(int, int) int{
	"+": sum,
	"-": sub,
	"/": div,
	"*": mult,
}

func Calc(expr string) (int, error) {
	tokens, err := GetTokens(expr)
	if err != nil {
		return -1, fmt.Errorf("An error has occured for expression %s: %s", expr, err)
	}

	s := structs.Stack[structs.Token]{}

	for _, v := range tokens {
		switch v.Kind {
		case structs.INT:
			s.Push(v)
		case structs.OP:
			a, err := s.Pop()
			if err != nil {
				return -1, fmt.Errorf("An error has occured for expression %s: %s", expr, err)
			}

			b, err := s.Pop()
			if err != nil {
				return -1, fmt.Errorf("An error has occured for expression %s: %s", expr, err)
			}

			tmp := structs.Token{
				Kind:  structs.INT,
				Value: ops[v.Op](b.Value, a.Value),
			}
			s.Push(tmp)
		}
	}

	res, err := s.Pop()
	if err != nil {
		return -1, fmt.Errorf("An error has occured for expression %s: %s", expr, err)
	}

	if len(s.Data) > 0 {
		return -1, fmt.Errorf("Expected empty stack but got: %v", s.Data)
	}

	return res.Value, nil
}
