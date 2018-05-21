package chapter1

import "testing"

func TestCompressString(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"aabcccccaaa", "a2b1c5a3"},
		{"aebc", "a1e1b1c1"},
	}

	for _, c := range cases {
		actual := compressString(c.input)
		if actual != c.expected {
			t.Errorf("Input: %s. Expected: %s, actual: %s\n", c.input, c.expected, actual)
		}
	}
}
