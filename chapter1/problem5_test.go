package chapter1

import "testing"

func TestIsOneAway(t *testing.T) {
	cases := []struct {
		input1   string
		input2   string
		expected bool
	}{
		{"paleee", "bake", false},
		{"pale", "ppaallee", false},
		{"pale", "ple", true},
		{"pales", "pale", true},
		{"pale", "pales", true},
		{"pale", "bale", true},
		{"pale", "bake", false},
		{"pale", "leap", false},
		{"leap", "pale", false},
		{" ", "", true},
		{"", " ", true},
		{"", "", true},
	}

	for _, c := range cases {
		actual := isOneAway(c.input1, c.input2)
		if actual != c.expected {
			t.Errorf("Inputs: %s, %s. Expected: %t, actual: %t\n", c.input1, c.input2, c.expected, actual)
		}
	}
}
