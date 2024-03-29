package chapter2

import (
	"reflect"
	"testing"
)

func TestLinkedListInsertAndGet(t *testing.T) {
	ll := GetLinkedList()
	vals := []int{1, 2, 3, 4}

	for _, val := range vals {
		ll.Insert(val)
	}

	actual := make([]int, len(vals))
	for i, _ := range vals {
		actual[i] = ll.Get(i)
	}

	if !reflect.DeepEqual(actual, vals) {
		t.Errorf("Expected: %v, actual: %v\n", vals, actual)
	}
}

func TestSlice(t *testing.T) {
	vals := []int{1, 2, 3, 4}
	ll := GetLinkedListFromValues(vals)
	actual := ll.Slice()
	if !reflect.DeepEqual(actual, vals) {
		t.Errorf("Expected: %v, actual: %v\n", vals, actual)
	}
}
