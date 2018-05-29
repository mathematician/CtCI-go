package chapter2

import (
	"testing"
)

func TestGetFromTail(t *testing.T) {
	vals := []int{1, 2, 3, 4, 5, 6, 6}
	ll := GetLinkedListFromValues(vals)
	expected := 4
	actual := ll.GetFromTail(4)

	if actual != expected {
		t.Errorf("Expected: %d, actual: %d.", expected, actual)
	}

}
