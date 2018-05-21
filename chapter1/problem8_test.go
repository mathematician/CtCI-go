package chapter1

import (
	"reflect"
	"testing"
)

func TestZeroMatrix(t *testing.T) {
	cases := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			[][]int{
				[]int{0, 1},
				[]int{2, 3},
			},
			[][]int{
				[]int{0, 0},
				[]int{0, 3},
			},
		},
		{
			[][]int{
				[]int{0, 1, 2},
				[]int{3, 4, 5},
				[]int{6, 7, 8},
			},
			[][]int{
				[]int{0, 0, 0},
				[]int{0, 4, 5},
				[]int{0, 7, 8},
			},
		},
		{
			[][]int{
				[]int{0, 1, 2},
				[]int{3, 0, 5},
				[]int{6, 7, 0},
			},
			[][]int{
				[]int{0, 0, 0},
				[]int{0, 0, 0},
				[]int{0, 0, 0},
			},
		},
	}

	for _, c := range cases {
		zeroMatrix(c.input)
		if !reflect.DeepEqual(c.input, c.expected) {
			t.Errorf("Expected: %v, actual: %v\n", c.expected, c.input)
		}
	}
}
