package chapter2

import (
	"reflect"
	"testing"
)

func TestGetIntersection(t *testing.T) {
	vals1 := []int{1, 2, 3, 4, 5, 6}
	vals2 := []int{3, 4, 5, 4, 5, 1}

	list1 := GetLinkedListFromValues(vals1)
	list2 := GetLinkedListFromValues(vals2)

	intersection := GetIntersection(list1, list2)
	actual := intersection.Slice()
	expected := []int{4}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %v, actual: %v", expected, actual)
	}
}
