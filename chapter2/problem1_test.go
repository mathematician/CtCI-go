package chapter2

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	vals := []int{1, 1, 1, 2, 2, 2, 3, 4, 5, 5, 5}
	ll := GetLinkedListFromValues(vals)
	ll.RemoveDuplicates()
	expected := []int{1, 2, 3, 4, 5}
	actual := ll.Slice()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, actual: %v\n", expected, actual)
	}
}
