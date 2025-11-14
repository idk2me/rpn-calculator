package calculator

import (
	"fmt"
	"strconv"
	"strings"

	"example.com/m/structs"
)

var Operators = map[string]bool{
	"+": true,
	"-": true,
	"/": true,
	"*": true,
}

func GetTokens(expr string) ([]structs.Token, error) {
	parts := strings.Fields(expr)
	tokens := make([]structs.Token, 0, len(parts))

	for _, p := range parts {
		if n, err := strconv.Atoi(p); err == nil {
			tokens = append(tokens, structs.Token{
				Kind:  structs.INT,
				Value: n,
			})
			continue
		}

		if Operators[p] {
			tokens = append(tokens, structs.Token{
				Kind: structs.OP,
				Op:   p,
			})
			continue
		}

		return nil, fmt.Errorf("Invalid token: %q", p)
	}

	return tokens, nil
}
