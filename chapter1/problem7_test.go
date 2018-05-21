package chapter1

import (
	"reflect"
	"testing"
)

func TestRotateMatrix(t *testing.T) {
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
				[]int{2, 0},
				[]int{3, 1}},
		},
		{
			[][]int{
				[]int{0, 1, 2},
				[]int{3, 4, 5},
				[]int{6, 7, 8},
			},
			[][]int{
				[]int{6, 3, 0},
				[]int{7, 4, 1},
				[]int{8, 5, 2},
			},
		},
		{
			[][]int{
				[]int{0, 1, 2, 3},
				[]int{4, 5, 6, 7},
				[]int{8, 9, 10, 11},
				[]int{12, 13, 14, 15},
			},
			[][]int{
				[]int{12, 8, 4, 0},
				[]int{13, 9, 5, 1},
				[]int{14, 10, 6, 2},
				[]int{15, 11, 7, 3},
			},
		},
	}

	for _, c := range cases {
		rotateMatrix(c.input)
		if !reflect.DeepEqual(c.input, c.expected) {
			t.Errorf("Expected: %v, actual: %v\n", c.expected, c.input)
		}
	}
}
