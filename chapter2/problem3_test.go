package chapter2

import (
	"reflect"
	"testing"
)

func TestDeleteMiddleNode(t *testing.T) {
	vals := []int{1, 2, 3, 4, 5, 6}
	ll := GetLinkedListFromValues(vals)
	node := ll.GetNode(2)

	DeleteNode(node)

	expected := []int{1, 2, 4, 5, 6}
	actual := ll.Slice()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %v, actual: %v", expected, actual)
	}
}
