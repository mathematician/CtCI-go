package chapter1

import "testing"

func TestIsRotatedString(t *testing.T) {
	cases := []struct {
		str1     string
		str2     string
		expected bool
	}{
		{"waterbottle", "erbottlewat", true},
		{"waterbottlewaterbottle", "erbottlewat", false},
		{"waterbottle", "erbottlewat", true},
		{"waterbottle", "erbottlewaterbottlewat", false},
	}

	for _, c := range cases {
		actual := isRotatedString(c.str1, c.str2)
		if actual != c.expected {
			t.Errorf("Input: %s, %s. Expected: %t, got: %t.", c.str1, c.str2, c.expected, actual)
		}
	}
}
