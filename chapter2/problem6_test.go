package chapter2

import "testing"

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		input    []int
		expected bool
	}{
		{[]int{1, 1, 2, 1, 1}, true},
	}

	for _, c := range cases {
		ll := GetLinkedListFromValues(c.input)
		actual := IsPalindrome(ll)
		if actual != c.expected {
			t.Errorf("Input: %v, expected: %t, actual: %t", c.input, c.expected, actual)
		}
	}
}
