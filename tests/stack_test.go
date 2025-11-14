package tests

import (
	"testing"

	"example.com/m/structs"
)

func TestStackPush(t *testing.T) {
	tests := []struct {
		input []int
	}{
		{[]int{}},
		{[]int{1}},
		{[]int{1, 2}},
		{[]int{1, 2, 3, 20}},
		{[]int{23, 4123, 4123, 4, 2314}},
	}

	for _, tt := range tests {
		var s structs.Stack[int]

		for _, v := range tt.input {
			s.Push(v)
		}

		if len(s.Data) != len(tt.input) {
			t.Errorf("Expected len %d, got %d", len(tt.input), len(s.Data))
			continue
		}

		for i := range tt.input {
			if s.Data[i] != tt.input[i] {
				t.Errorf("Expected element %v, but got %v", tt.input[i], s.Data[i])
				break
			}
		}
	}
}

func TestStackPop(t *testing.T) {
	tests := []struct {
		push    []int
		pops    int
		want    int
		wantErr bool
	}{
		{[]int{1}, 1, 1, false},
		{[]int{5, 6}, 1, 6, false},
		{[]int{10}, 2, 0, true},
	}

	for _, tt := range tests {
		var s structs.Stack[int]

		for _, v := range tt.push {
			s.Push(v)
		}

		var got int
		var err error

		for i := 0; i < tt.pops; i++ {
			got, err = s.Pop()
		}

		if tt.wantErr && err == nil {
			t.Error("Expected error but got nil")
		}

		if !tt.wantErr && err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if got != tt.want {
			t.Errorf("Expected %v, but got %v", tt.want, got)
		}
	}
}
