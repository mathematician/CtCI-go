package chapter2

import (
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	vals := []int{3, 5, 8, 5, 10, 2, 1}
	ll := GetLinkedListFromValues(vals)
	partition := 5

	ll.Partition(partition)

	expected := []int{3, 2, 1, 5, 8, 5, 10}
	actual := ll.Slice()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, actual: %v", expected, actual)
	}
}
