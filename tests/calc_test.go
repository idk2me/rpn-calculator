package tests

import (
	"testing"

	"example.com/m/calculator"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		input     string
		shouldErr bool
		output    int
	}{
		// addition
		{input: "4 5 +", shouldErr: false, output: 9},
		{input: "4 5 6 + +", shouldErr: false, output: 15},

		// subtraction
		{input: "10 3 -", shouldErr: false, output: 7},
		{input: "20 5 3 - -", shouldErr: false, output: 18},

		// multiplication
		{input: "3 4 *", shouldErr: false, output: 12},
		{input: "2 3 4 * *", shouldErr: false, output: 24},

		// division
		{input: "20 4 /", shouldErr: false, output: 5},
		{input: "100 5 2 / /", shouldErr: false, output: 50},

		// nesting
		{input: "10 6 2 - /", shouldErr: false, output: 2},
		{input: "3 4 + 5 *", shouldErr: false, output: 35},

		// errors
		{input: "4 +", shouldErr: true},
		{input: "4 4 4", shouldErr: true},
	}

	for _, v := range tests {
		c, err := calculator.Calc(v.input)

		if v.shouldErr {
			if err == nil {
				t.Errorf("expected error for %q, got result %v", v.input, c)
			}
			continue
		}

		if err != nil {
			t.Errorf("unexpected error for %q: %v", v.input, err)
			continue
		}

		if c != v.output {
			t.Errorf("for %q expected %v, got %v", v.input, v.output, c)
		}
	}
}
